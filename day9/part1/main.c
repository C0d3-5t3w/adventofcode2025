#include <stdio.h>
#include <stdlib.h>

#define MAX_POINTS 1000

typedef struct {
  int x;
  int y;
} Point;

int main() {
  FILE *file = fopen("list.txt", "r");
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

  long long max_area = 0;

  for (int i = 0; i < count; i++) {
    for (int j = i + 1; j < count; j++) {
      long long width =
          llabs((long long)points[j].x - (long long)points[i].x) + 1;
      long long height =
          llabs((long long)points[j].y - (long long)points[i].y) + 1;
      long long area = width * height;

      if (area > max_area) {
        max_area = area;
      }
    }
  }

  printf("Largest rectangle area: %lld\n", max_area);

  return 0;
}
