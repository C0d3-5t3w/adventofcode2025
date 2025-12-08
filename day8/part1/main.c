#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_BOXES 1000
#define CONNECTIONS 1000

typedef struct {
  int x, y, z;
} Point;

typedef struct {
  int box1, box2;
  double distance;
} Edge;

int parent[MAX_BOXES];
int rank_arr[MAX_BOXES];

void init_union_find(int n) {
  for (int i = 0; i < n; i++) {
    parent[i] = i;
    rank_arr[i] = 0;
  }
}

int find(int x) {
  if (parent[x] != x) {
    parent[x] = find(parent[x]);
  }
  return parent[x];
}

bool union_sets(int x, int y) {
  int px = find(x);
  int py = find(y);

  if (px == py)
    return false;

  if (rank_arr[px] < rank_arr[py]) {
    parent[px] = py;
  } else if (rank_arr[px] > rank_arr[py]) {
    parent[py] = px;
  } else {
    parent[py] = px;
    rank_arr[px]++;
  }
  return true;
}

double distance(Point a, Point b) {
  long long dx = (long long)a.x - b.x;
  long long dy = (long long)a.y - b.y;
  long long dz = (long long)a.z - b.z;
  return sqrt(dx * dx + dy * dy + dz * dz);
}

int compare_edges(const void *a, const void *b) {
  Edge *ea = (Edge *)a;
  Edge *eb = (Edge *)b;
  if (ea->distance < eb->distance)
    return -1;
  if (ea->distance > eb->distance)
    return 1;
  return 0;
}

int compare_int_desc(const void *a, const void *b) {
  return (*(int *)b - *(int *)a);
}

int main() {
  FILE *file = fopen("../list.txt", "r");
  if (!file) {
    fprintf(stderr, "Error opening list.txt\n");
    return 1;
  }

  Point boxes[MAX_BOXES];
  int num_boxes = 0;

  while (fscanf(file, "%d,%d,%d", &boxes[num_boxes].x, &boxes[num_boxes].y,
                &boxes[num_boxes].z) == 3) {
    num_boxes++;
    if (num_boxes >= MAX_BOXES)
      break;
  }
  fclose(file);

  printf("Read %d junction boxes\n", num_boxes);

  int num_edges = (num_boxes * (num_boxes - 1)) / 2;
  Edge *edges = malloc(num_edges * sizeof(Edge));
  if (!edges) {
    fprintf(stderr, "Memory allocation failed\n");
    return 1;
  }

  int edge_count = 0;
  for (int i = 0; i < num_boxes; i++) {
    for (int j = i + 1; j < num_boxes; j++) {
      edges[edge_count].box1 = i;
      edges[edge_count].box2 = j;
      edges[edge_count].distance = distance(boxes[i], boxes[j]);
      edge_count++;
    }
  }

  printf("Calculated %d distances\n", edge_count);

  qsort(edges, edge_count, sizeof(Edge), compare_edges);

  init_union_find(num_boxes);

  int actual_connections = 0;
  for (int i = 0; i < CONNECTIONS && i < edge_count; i++) {
    if (union_sets(edges[i].box1, edges[i].box2)) {
      actual_connections++;
    }
  }

  printf("Processed %d pairs, made %d actual connections\n", CONNECTIONS,
         actual_connections);

  int circuit_size[MAX_BOXES] = {0};
  for (int i = 0; i < num_boxes; i++) {
    int root = find(i);
    circuit_size[root]++;
  }

  qsort(circuit_size, num_boxes, sizeof(int), compare_int_desc);

  int num_circuits = 0;
  for (int i = 0; i < num_boxes; i++) {
    if (circuit_size[i] > 0)
      num_circuits++;
  }

  printf("Number of circuits: %d\n", num_circuits);
  printf("Three largest circuits: %d, %d, %d\n", circuit_size[0],
         circuit_size[1], circuit_size[2]);

  long long answer =
      (long long)circuit_size[0] * circuit_size[1] * circuit_size[2];
  printf("\nAnswer: %lld\n", answer);

  free(edges);
  return 0;
}
