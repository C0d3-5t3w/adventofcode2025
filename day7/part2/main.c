#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_ROWS 256
#define MAX_COLS 256

int main(void) {
  FILE *file = fopen("list.txt", "r");
  if (!file) {
    perror("Failed to open list.txt");
    return 1;
  }

  char grid[MAX_ROWS][MAX_COLS];
  int rows = 0;
  int cols = 0;

  while (fgets(grid[rows], MAX_COLS, file)) {

    int len = strlen(grid[rows]);
    if (len > 0 && grid[rows][len - 1] == '\n') {
      grid[rows][len - 1] = '\0';
      len--;
    }
    if (rows == 0) {
      cols = len;
    }
    rows++;
  }
  fclose(file);

  int start_col = -1;
  for (int c = 0; c < cols; c++) {
    if (grid[0][c] == 'S') {
      start_col = c;
      break;
    }
  }

  if (start_col == -1) {
    fprintf(stderr, "Could not find starting position 'S'\n");
    return 1;
  }

  long long *timelines = calloc(cols, sizeof(long long));
  long long *next_timelines = calloc(cols, sizeof(long long));

  timelines[start_col] = 1;

  for (int r = 1; r < rows; r++) {

    memset(next_timelines, 0, cols * sizeof(long long));

    for (int c = 0; c < cols; c++) {
      if (timelines[c] > 0) {

        if (grid[r][c] == '^') {

          if (c - 1 >= 0) {
            next_timelines[c - 1] += timelines[c];
          }
          if (c + 1 < cols) {
            next_timelines[c + 1] += timelines[c];
          }
        } else {

          next_timelines[c] += timelines[c];
        }
      }
    }

    long long *temp = timelines;
    timelines = next_timelines;
    next_timelines = temp;
  }

  long long total_timelines = 0;
  for (int c = 0; c < cols; c++) {
    total_timelines += timelines[c];
  }

  printf("Total timelines: %lld\n", total_timelines);

  free(timelines);
  free(next_timelines);

  return 0;
}
