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

  bool *beams = calloc(cols, sizeof(bool));
  bool *next_beams = calloc(cols, sizeof(bool));

  beams[start_col] = true;

  int total_splits = 0;

  for (int r = 1; r < rows; r++) {

    memset(next_beams, 0, cols * sizeof(bool));

    for (int c = 0; c < cols; c++) {
      if (beams[c]) {

        if (grid[r][c] == '^') {

          total_splits++;

          if (c - 1 >= 0) {
            next_beams[c - 1] = true;
          }
          if (c + 1 < cols) {
            next_beams[c + 1] = true;
          }
        } else {

          next_beams[c] = true;
        }
      }
    }

    bool *temp = beams;
    beams = next_beams;
    next_beams = temp;
  }

  printf("Total beam splits: %d\n", total_splits);

  free(beams);
  free(next_beams);

  return 0;
}
