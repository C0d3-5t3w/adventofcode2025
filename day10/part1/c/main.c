#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE 4096
#define MAX_LIGHTS 16
#define MAX_BUTTONS 20

int parse_lights(const char *start, int *num_lights) {
  int target = 0;
  *num_lights = 0;
  const char *p = start + 1;
  while (*p && *p != ']') {
    if (*p == '#') {
      target |= (1 << *num_lights);
    }
    (*num_lights)++;
    p++;
  }
  return target;
}

int parse_button(const char *start, const char **end) {
  int mask = 0;
  const char *p = start + 1;
  while (*p && *p != ')') {
    if (*p >= '0' && *p <= '9') {
      int num = 0;
      while (*p >= '0' && *p <= '9') {
        num = num * 10 + (*p - '0');
        p++;
      }
      mask |= (1 << num);
    } else {
      p++;
    }
  }
  if (*p == ')')
    p++;
  *end = p;
  return mask;
}

int min_presses(int target, int *buttons, int num_buttons, int num_lights) {

  int max_state = 1 << num_lights;
  int *dist = malloc(max_state * sizeof(int));
  for (int i = 0; i < max_state; i++)
    dist[i] = INT_MAX;

  int *queue = malloc(max_state * sizeof(int));
  int front = 0, back = 0;

  dist[0] = 0;
  queue[back++] = 0;

  while (front < back) {
    int state = queue[front++];
    if (state == target) {
      int result = dist[target];
      free(dist);
      free(queue);
      return result;
    }

    for (int i = 0; i < num_buttons; i++) {
      int next_state = state ^ buttons[i];
      if (dist[next_state] == INT_MAX) {
        dist[next_state] = dist[state] + 1;
        queue[back++] = next_state;
      }
    }
  }

  int result = dist[target];
  free(dist);
  free(queue);
  return result == INT_MAX ? -1 : result;
}

int main(void) {
  FILE *fp = fopen("../list.txt", "r");
  if (!fp) {
    perror("Failed to open list.txt");
    return 1;
  }

  char line[MAX_LINE];
  long long total = 0;

  while (fgets(line, sizeof(line), fp)) {

    if (line[0] == '\n' || line[0] == '\0')
      continue;

    char *bracket = strchr(line, '[');
    if (!bracket)
      continue;

    int num_lights;
    int target = parse_lights(bracket, &num_lights);

    int buttons[MAX_BUTTONS];
    int num_buttons = 0;

    const char *p = strchr(bracket, ']');
    if (!p)
      continue;
    p++;

    while (*p) {

      while (*p && *p != '(' && *p != '{')
        p++;
      if (*p == '{' || *p == '\0')
        break;

      const char *end;
      buttons[num_buttons++] = parse_button(p, &end);
      p = end;
    }

    int min = min_presses(target, buttons, num_buttons, num_lights);
    if (min >= 0) {
      total += min;
    }
  }

  fclose(fp);
  printf("%lld\n", total);

  return 0;
}
