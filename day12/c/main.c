#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_SHAPES 6
#define SHAPE_SIZE 3
#define MAX_ROTATIONS 8
#define MAX_WIDTH 60
#define MAX_HEIGHT 60

typedef struct {
  int cells[SHAPE_SIZE][SHAPE_SIZE];
  int cell_count;
} Shape;

typedef struct {
  Shape variants[MAX_ROTATIONS];
  int num_variants;
  int cell_count;
} ShapeVariants;

ShapeVariants shapes[MAX_SHAPES];
int grid[MAX_HEIGHT][MAX_WIDTH];
int width, height;

Shape rotate_shape(Shape s) {
  Shape result;
  result.cell_count = s.cell_count;
  for (int r = 0; r < SHAPE_SIZE; r++) {
    for (int c = 0; c < SHAPE_SIZE; c++) {
      result.cells[c][SHAPE_SIZE - 1 - r] = s.cells[r][c];
    }
  }
  return result;
}

Shape flip_shape(Shape s) {
  Shape result;
  result.cell_count = s.cell_count;
  for (int r = 0; r < SHAPE_SIZE; r++) {
    for (int c = 0; c < SHAPE_SIZE; c++) {
      result.cells[r][SHAPE_SIZE - 1 - c] = s.cells[r][c];
    }
  }
  return result;
}

bool shapes_equal(Shape a, Shape b) {
  for (int r = 0; r < SHAPE_SIZE; r++) {
    for (int c = 0; c < SHAPE_SIZE; c++) {
      if (a.cells[r][c] != b.cells[r][c])
        return false;
    }
  }
  return true;
}

bool variant_exists(ShapeVariants *sv, Shape s) {
  for (int i = 0; i < sv->num_variants; i++) {
    if (shapes_equal(sv->variants[i], s))
      return true;
  }
  return false;
}

void generate_variants(Shape base, ShapeVariants *sv) {
  sv->num_variants = 0;
  sv->cell_count = base.cell_count;

  Shape current = base;
  for (int flip = 0; flip < 2; flip++) {
    for (int rot = 0; rot < 4; rot++) {
      if (!variant_exists(sv, current)) {
        sv->variants[sv->num_variants++] = current;
      }
      current = rotate_shape(current);
    }
    current = flip_shape(base);
  }
}

bool can_place(Shape *s, int row, int col) {
  for (int r = 0; r < SHAPE_SIZE; r++) {
    for (int c = 0; c < SHAPE_SIZE; c++) {
      if (s->cells[r][c]) {
        int nr = row + r;
        int nc = col + c;
        if (nr < 0 || nr >= height || nc < 0 || nc >= width)
          return false;
        if (grid[nr][nc])
          return false;
      }
    }
  }
  return true;
}

void place_shape(Shape *s, int row, int col, int mark) {
  for (int r = 0; r < SHAPE_SIZE; r++) {
    for (int c = 0; c < SHAPE_SIZE; c++) {
      if (s->cells[r][c]) {
        grid[row + r][col + c] = mark;
      }
    }
  }
}

void remove_shape(Shape *s, int row, int col) {
  for (int r = 0; r < SHAPE_SIZE; r++) {
    for (int c = 0; c < SHAPE_SIZE; c++) {
      if (s->cells[r][c]) {
        grid[row + r][col + c] = 0;
      }
    }
  }
}

int count_remaining(int counts[MAX_SHAPES]) {
  int total = 0;
  for (int i = 0; i < MAX_SHAPES; i++) {
    total += counts[i];
  }
  return total;
}

bool solve(int counts[MAX_SHAPES], int placed) {
  if (count_remaining(counts) == 0)
    return true;

  int shape_idx = -1;
  for (int i = 0; i < MAX_SHAPES; i++) {
    if (counts[i] > 0) {
      shape_idx = i;
      break;
    }
  }
  if (shape_idx < 0)
    return true;

  for (int row = 0; row <= height - 1; row++) {
    for (int col = 0; col <= width - 1; col++) {
      for (int var = 0; var < shapes[shape_idx].num_variants; var++) {
        Shape *s = &shapes[shape_idx].variants[var];

        if (can_place(s, row, col)) {
          place_shape(s, row, col, placed + 1);
          counts[shape_idx]--;

          if (solve(counts, placed + 1)) {
            counts[shape_idx]++;
            return true;
          }

          remove_shape(s, row, col);
          counts[shape_idx]++;
        }
      }
    }
  }

  return false;
}

int count_total_cells(int counts[MAX_SHAPES]) {
  int total = 0;
  for (int i = 0; i < MAX_SHAPES; i++) {
    total += counts[i] * shapes[i].cell_count;
  }
  return total;
}

bool can_fit_region(int w, int h, int counts[MAX_SHAPES]) {
  int total_cells = count_total_cells(counts);
  if (total_cells > w * h) {
    return false;
  }

  width = w;
  height = h;
  memset(grid, 0, sizeof(grid));

  int counts_copy[MAX_SHAPES];
  memcpy(counts_copy, counts, sizeof(counts_copy));

  return solve(counts_copy, 0);
}

int main() {
  FILE *f = fopen("../list.txt", "r");
  if (!f) {
    f = fopen("list.txt", "r");
    if (!f) {
      perror("Cannot open list.txt");
      return 1;
    }
  }

  char line[256];

  for (int i = 0; i < MAX_SHAPES; i++) {
    if (!fgets(line, sizeof(line), f))
      break;

    Shape base;
    memset(&base, 0, sizeof(base));
    base.cell_count = 0;

    for (int r = 0; r < SHAPE_SIZE; r++) {
      if (!fgets(line, sizeof(line), f))
        break;
      for (int c = 0; c < SHAPE_SIZE && line[c] && line[c] != '\n'; c++) {
        base.cells[r][c] = (line[c] == '#') ? 1 : 0;
        if (line[c] == '#')
          base.cell_count++;
      }
    }

    fgets(line, sizeof(line), f);
    generate_variants(base, &shapes[i]);
    printf("Shape %d: %d cells, %d variants\n", i, shapes[i].cell_count,
           shapes[i].num_variants);
  }

  int fit_count = 0;
  int region_num = 0;

  while (fgets(line, sizeof(line), f)) {
    if (line[0] == '\n' || line[0] == '\0')
      continue;

    int w, h;
    int counts[MAX_SHAPES];

    if (sscanf(line, "%dx%d: %d %d %d %d %d %d", &w, &h, &counts[0], &counts[1],
               &counts[2], &counts[3], &counts[4], &counts[5]) != 8) {
      continue;
    }

    region_num++;

    int total_cells = count_total_cells(counts);
    printf("Region %d: %dx%d (area=%d, cells=%d) ", region_num, w, h, w * h,
           total_cells);
    fflush(stdout);

    if (can_fit_region(w, h, counts)) {
      fit_count++;
      printf("FIT\n");
    } else {
      printf("NO\n");
    }
  }

  fclose(f);
  printf("\nAnswer: %d\n", fit_count);

  return 0;
}
