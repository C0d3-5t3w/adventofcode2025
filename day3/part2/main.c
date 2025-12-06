#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 1024
#define NUM_DIGITS 12

long long max_joltage(const char *line, int len) {
  if (len < NUM_DIGITS) {
    return 0;
  }

  char result[NUM_DIGITS + 1];
  int result_idx = 0;
  int start = 0;

  for (int i = 0; i < NUM_DIGITS; i++) {

    int max_pos = len - (NUM_DIGITS - i);

    int best_pos = start;
    char best_digit = line[start];

    for (int p = start; p <= max_pos; p++) {
      if (line[p] > best_digit) {
        best_digit = line[p];
        best_pos = p;
      }
    }

    result[result_idx++] = best_digit;
    start = best_pos + 1;
  }

  result[result_idx] = '\0';

  return atoll(result);
}

int main(void) {
  FILE *file = fopen("list.txt", "r");
  if (!file) {
    perror("Failed to open list.txt");
    return 1;
  }

  char line[MAX_LINE_LENGTH];
  long long total = 0;

  while (fgets(line, sizeof(line), file)) {

    int len = strlen(line);
    while (len > 0 && (line[len - 1] == '\n' || line[len - 1] == '\r')) {
      line[--len] = '\0';
    }

    if (len >= NUM_DIGITS) {
      long long joltage = max_joltage(line, len);
      total += joltage;
    }
  }

  fclose(file);

  printf("Total output joltage: %lld\n", total);

  return 0;
}

