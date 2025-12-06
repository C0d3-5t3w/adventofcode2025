#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_SIZE 256

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

  int dx[] = {-1, -1, -1, 0, 0, 1, 1, 1};
  int dy[] = {-1, 0, 1, -1, 1, -1, 0, 1};

  int accessible_count = 0;

  for (int r = 0; r < rows; r++) {
    for (int c = 0; c < cols; c++) {

      if (grid[r][c] != '@') {
        continue;
      }

      int adjacent_rolls = 0;
      for (int d = 0; d < 8; d++) {
        int nr = r + dx[d];
        int nc = c + dy[d];

        if (nr >= 0 && nr < rows && nc >= 0 && nc < cols) {
          if (grid[nr][nc] == '@') {
            adjacent_rolls++;
          }
        }
      }

      if (adjacent_rolls < 4) {
        accessible_count++;
      }
    }
  }

  printf("Accessible paper rolls: %d\n", accessible_count);

  return 0;
}

