#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_SIZE 256

int dx[] = {-1, -1, -1, 0, 0, 1, 1, 1};
int dy[] = {-1, 0, 1, -1, 1, -1, 0, 1};

int count_adjacent_rolls(char grid[MAX_SIZE][MAX_SIZE], int rows, int cols,
                         int r, int c) {
  int count = 0;
  for (int d = 0; d < 8; d++) {
    int nr = r + dx[d];
    int nc = c + dy[d];
    if (nr >= 0 && nr < rows && nc >= 0 && nc < cols) {
      if (grid[nr][nc] == '@') {
        count++;
      }
    }
  }
  return count;
}

int main(void) {
  FILE *file = fopen("list.txt", "r");
  if (!file) {
    perror("Failed to open list.txt");
    return 1;
  }

  char grid[MAX_SIZE][MAX_SIZE];
  int rows = 0;
  int cols = 0;

  while (fgets(grid[rows], MAX_SIZE, file) != NULL) {

    int len = strlen(grid[rows]);
    if (len > 0 && grid[rows][len - 1] == '\n') {
      grid[rows][len - 1] = '\0';
      len--;
    }
    if (len > cols) {
      cols = len;
    }
    rows++;
  }
  fclose(file);

  printf("Grid size: %d rows x %d cols\n", rows, cols);

  int total_removed = 0;
  bool removed_any = true;

  while (removed_any) {
    removed_any = false;
    int removed_this_round = 0;

    bool to_remove[MAX_SIZE][MAX_SIZE] = {false};

    for (int r = 0; r < rows; r++) {
      for (int c = 0; c < cols; c++) {
        if (grid[r][c] != '@') {
          continue;
        }

        int adjacent_rolls = count_adjacent_rolls(grid, rows, cols, r, c);

        if (adjacent_rolls < 4) {
          to_remove[r][c] = true;
          removed_this_round++;
          removed_any = true;
        }
      }
    }

    for (int r = 0; r < rows; r++) {
      for (int c = 0; c < cols; c++) {
        if (to_remove[r][c]) {
          grid[r][c] = '.';
        }
      }
    }

    if (removed_this_round > 0) {
      printf("Removed %d rolls this round\n", removed_this_round);
    }
    total_removed += removed_this_round;
  }

  printf("Total paper rolls removed: %d\n", total_removed);

  return 0;
}
