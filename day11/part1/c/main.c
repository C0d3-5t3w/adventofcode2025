#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_DEVICES 1000
#define MAX_OUTPUTS 20
#define MAX_NAME_LEN 32

typedef struct {
  char name[MAX_NAME_LEN];
  char outputs[MAX_OUTPUTS][MAX_NAME_LEN];
  int output_count;
} Device;

Device devices[MAX_DEVICES];
int device_count = 0;
bool visiting[MAX_DEVICES];

int find_device(const char *name) {
  for (int i = 0; i < device_count; i++) {
    if (strcmp(devices[i].name, name) == 0) {
      return i;
    }
  }
  return -1;
}

int add_device(const char *name) {
  int idx = find_device(name);
  if (idx >= 0)
    return idx;

  strcpy(devices[device_count].name, name);
  devices[device_count].output_count = 0;
  return device_count++;
}

long long count_paths(const char *current) {
  if (strcmp(current, "out") == 0) {
    return 1;
  }

  int idx = find_device(current);
  if (idx < 0 || devices[idx].output_count == 0) {
    return 0;
  }

  if (visiting[idx]) {
    return 0;
  }

  visiting[idx] = true;
  long long total = 0;
  for (int i = 0; i < devices[idx].output_count; i++) {
    total += count_paths(devices[idx].outputs[i]);
  }
  visiting[idx] = false;

  return total;
}

int main(void) {
  FILE *fp = fopen("../list.txt", "r");
  if (!fp) {
    perror("Failed to open list.txt");
    return 1;
  }

  char line[1024];
  while (fgets(line, sizeof(line), fp)) {
    line[strcspn(line, "\n")] = 0;
    if (strlen(line) == 0)
      continue;

    char *colon = strchr(line, ':');
    if (!colon)
      continue;

    *colon = 0;
    char *name = line;
    char *rest = colon + 1;

    while (*rest == ' ')
      rest++;

    int idx = add_device(name);

    char *token = strtok(rest, " ");
    while (token && devices[idx].output_count < MAX_OUTPUTS) {
      strcpy(devices[idx].outputs[devices[idx].output_count++], token);
      token = strtok(NULL, " ");
    }
  }
  fclose(fp);

  memset(visiting, 0, sizeof(visiting));

  long long result = count_paths("you");
  printf("%lld\n", result);

  return 0;
}
