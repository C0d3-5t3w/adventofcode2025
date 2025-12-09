#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_POINTS 1000

typedef struct {
  int x;
  int y;
} Point;

typedef struct {
  int y;
  int x_min;
  int x_max;
} HSegment;

typedef struct {
  int x;
  int y_min;
  int y_max;
} VSegment;

HSegment h_segments[MAX_POINTS];
VSegment v_segments[MAX_POINTS];
int h_count = 0;
int v_count = 0;

int cmp_hseg(const void *a, const void *b) {
  const HSegment *sa = (const HSegment *)a;
  const HSegment *sb = (const HSegment *)b;
  if (sa->y != sb->y)
    return sa->y - sb->y;
  return sa->x_min - sb->x_min;
}

int cmp_vseg(const void *a, const void *b) {
  const VSegment *sa = (const VSegment *)a;
  const VSegment *sb = (const VSegment *)b;
  if (sa->x != sb->x)
    return sa->x - sb->x;
  return sa->y_min - sb->y_min;
}

bool is_inside_or_on_boundary(int x, int y) {
  int crossings = 0;

  for (int i = 0; i < v_count; i++) {
    if (v_segments[i].x >= x && y >= v_segments[i].y_min &&
        y <= v_segments[i].y_max) {

      if (v_segments[i].x == x) {
        return true;
      }

      if (y > v_segments[i].y_min && y <= v_segments[i].y_max) {
        crossings++;
      }
    }
  }

  for (int i = 0; i < h_count; i++) {
    if (h_segments[i].y == y && x >= h_segments[i].x_min &&
        x <= h_segments[i].x_max) {
      return true;
    }
  }

  return (crossings % 2) == 1;
}

bool is_rectangle_valid(int x1, int y1, int x2, int y2) {
  int min_x = x1 < x2 ? x1 : x2;
  int max_x = x1 > x2 ? x1 : x2;
  int min_y = y1 < y2 ? y1 : y2;
  int max_y = y1 > y2 ? y1 : y2;

  if (!is_inside_or_on_boundary(min_x, min_y))
    return false;
  if (!is_inside_or_on_boundary(min_x, max_y))
    return false;
  if (!is_inside_or_on_boundary(max_x, min_y))
    return false;
  if (!is_inside_or_on_boundary(max_x, max_y))
    return false;

  for (int i = 0; i < v_count; i++) {
    int seg_x = v_segments[i].x;
    if (seg_x > min_x && seg_x < max_x) {

      if (v_segments[i].y_min <= max_y && v_segments[i].y_max >= max_y) {

        if (v_segments[i].y_max > max_y) {

          if (!is_inside_or_on_boundary(seg_x, max_y + 1) ||
              !is_inside_or_on_boundary(seg_x - 1, max_y) ||
              !is_inside_or_on_boundary(seg_x + 1, max_y)) {
          }
        }
      }
    }
  }

  for (int x = min_x; x <= max_x; x++) {
    if (!is_inside_or_on_boundary(x, max_y))
      return false;
  }

  for (int x = min_x; x <= max_x; x++) {
    if (!is_inside_or_on_boundary(x, min_y))
      return false;
  }

  for (int y = min_y; y <= max_y; y++) {
    if (!is_inside_or_on_boundary(min_x, y))
      return false;
  }

  for (int y = min_y; y <= max_y; y++) {
    if (!is_inside_or_on_boundary(max_x, y))
      return false;
  }

  return true;
}

int main() {
  FILE *file = fopen("../list.txt", "r");
  if (!file) {
    fprintf(stderr, "Error opening list.txt\n");
    return 1;
  }

  Point points[MAX_POINTS];
  int count = 0;

  while (fscanf(file, "%d,%d", &points[count].x, &points[count].y) == 2) {
    count++;
    if (count >= MAX_POINTS) {
      fprintf(stderr, "Too many points\n");
      fclose(file);
      return 1;
    }
  }
  fclose(file);

  printf("Read %d red tile coordinates\n", count);

  for (int i = 0; i < count; i++) {
    int next = (i + 1) % count;
    int x1 = points[i].x, y1 = points[i].y;
    int x2 = points[next].x, y2 = points[next].y;

    if (y1 == y2) {

      h_segments[h_count].y = y1;
      h_segments[h_count].x_min = x1 < x2 ? x1 : x2;
      h_segments[h_count].x_max = x1 > x2 ? x1 : x2;
      h_count++;
    } else if (x1 == x2) {

      v_segments[v_count].x = x1;
      v_segments[v_count].y_min = y1 < y2 ? y1 : y2;
      v_segments[v_count].y_max = y1 > y2 ? y1 : y2;
      v_count++;
    } else {
      fprintf(stderr,
              "Warning: non-axis-aligned segment between (%d,%d) and (%d,%d)\n",
              x1, y1, x2, y2);
    }
  }

  printf("Built %d horizontal and %d vertical segments\n", h_count, v_count);

  qsort(h_segments, h_count, sizeof(HSegment), cmp_hseg);
  qsort(v_segments, v_count, sizeof(VSegment), cmp_vseg);

  long long max_area = 0;
  int best_i = -1, best_j = -1;

  for (int i = 0; i < count; i++) {
    for (int j = i + 1; j < count; j++) {
      long long width =
          llabs((long long)points[j].x - (long long)points[i].x) + 1;
      long long height =
          llabs((long long)points[j].y - (long long)points[i].y) + 1;
      long long area = width * height;

      if (area > max_area) {
        if (is_rectangle_valid(points[i].x, points[i].y, points[j].x,
                               points[j].y)) {
          max_area = area;
          best_i = i;
          best_j = j;
        }
      }
    }
  }

  if (best_i >= 0) {
    printf("Best rectangle: (%d,%d) to (%d,%d)\n", points[best_i].x,
           points[best_i].y, points[best_j].x, points[best_j].y);
  }
  printf("Largest valid rectangle area: %lld\n", max_area);

  return 0;
}