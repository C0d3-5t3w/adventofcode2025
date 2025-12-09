#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 1024

int max_joltage(const char *line, int len) {
  int max = 0;

  for (int i = 0; i < len - 1; i++) {
    for (int j = i + 1; j < len; j++) {
      int first = line[i] - '0';
      int second = line[j] - '0';
      int value = first * 10 + second;
      if (value > max) {
        max = value;
      }
    }
  }

  return max;
}

int main(void) {
  FILE *file = fopen("../list.txt", "r");
  if (!file) {
    perror("Failed to open list.txt");
    return 1;
  }

  char line[MAX_LINE_LENGTH];
  long total = 0;

  while (fgets(line, sizeof(line), file)) {

    int len = strlen(line);
    while (len > 0 && (line[len - 1] == '\n' || line[len - 1] == '\r')) {
      line[--len] = '\0';
    }

    if (len > 0) {
      int joltage = max_joltage(line, len);
      total += joltage;
    }
  }

  fclose(file);

  printf("Total output joltage: %ld\n", total);

  return 0;
}

