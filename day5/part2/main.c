#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_RANGES 10000
#define MAX_LINE 100

typedef struct {
  long long start;
  long long end;
} Range;

int compare_ranges(const void *a, const void *b) {
  const Range *ra = (const Range *)a;
  const Range *rb = (const Range *)b;
  if (ra->start < rb->start)
    return -1;
  if (ra->start > rb->start)
    return 1;
  return 0;
}

int main(void) {
  FILE *file = fopen("list.txt", "r");
  if (!file) {
    perror("Failed to open list.txt");
    return 1;
  }

  Range *ranges = malloc(MAX_RANGES * sizeof(Range));
  if (!ranges) {
    perror("Failed to allocate memory");
    fclose(file);
    return 1;
  }

  int range_count = 0;
  char line[MAX_LINE];

  while (fgets(line, sizeof(line), file)) {
    line[strcspn(line, "\n")] = '\0';

    if (strlen(line) == 0) {
      break;
    }

    long long start, end;
    if (sscanf(line, "%lld-%lld", &start, &end) == 2) {
      if (range_count < MAX_RANGES) {
        ranges[range_count].start = start;
        ranges[range_count].end = end;
        range_count++;
      }
    }
  }

  fclose(file);

  qsort(ranges, range_count, sizeof(Range), compare_ranges);

  long long total_fresh = 0;
  long long current_start = ranges[0].start;
  long long current_end = ranges[0].end;

  for (int i = 1; i < range_count; i++) {
    if (ranges[i].start <= current_end + 1) {

      if (ranges[i].end > current_end) {
        current_end = ranges[i].end;
      }
    } else {

      total_fresh += (current_end - current_start + 1);
      current_start = ranges[i].start;
      current_end = ranges[i].end;
    }
  }

  total_fresh += (current_end - current_start + 1);

  free(ranges);

  printf("Number of fresh ingredient IDs: %lld\n", total_fresh);

  return 0;
}
