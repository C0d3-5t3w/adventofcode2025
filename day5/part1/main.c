#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_RANGES 1000
#define MAX_LINE 100

typedef struct {
  long long start;
  long long end;
} Range;

int main(void) {
  FILE *file = fopen("list.txt", "r");
  if (!file) {
    perror("Failed to open list.txt");
    return 1;
  }

  Range ranges[MAX_RANGES];
  int range_count = 0;
  char line[MAX_LINE];
  bool parsing_ranges = true;
  int fresh_count = 0;

  while (fgets(line, sizeof(line), file)) {

    line[strcspn(line, "\n")] = '\0';

    if (strlen(line) == 0) {
      parsing_ranges = false;
      continue;
    }

    if (parsing_ranges) {

      long long start, end;
      if (sscanf(line, "%lld-%lld", &start, &end) == 2) {
        if (range_count < MAX_RANGES) {
          ranges[range_count].start = start;
          ranges[range_count].end = end;
          range_count++;
        }
      }
    } else {

      long long id;
      if (sscanf(line, "%lld", &id) == 1) {

        bool is_fresh = false;
        for (int i = 0; i < range_count; i++) {
          if (id >= ranges[i].start && id <= ranges[i].end) {
            is_fresh = true;
            break;
          }
        }
        if (is_fresh) {
          fresh_count++;
        }
      }
    }
  }

  fclose(file);

  printf("Number of fresh ingredient IDs: %d\n", fresh_count);

  return 0;
}
