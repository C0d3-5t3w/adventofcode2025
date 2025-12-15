#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_NODES 10000
#define MAX_NAME_LEN 64
#define MAX_OUTPUTS 100
#define HASH_SIZE 20011

typedef struct {
  char name[MAX_NAME_LEN];
  char outputs[MAX_OUTPUTS][MAX_NAME_LEN];
  int output_count;
} Node;

typedef struct {
  char node[MAX_NAME_LEN];
  bool visited_dac;
  bool visited_fft;
  long long value;
  bool valid;
} MemoEntry;

static Node graph[MAX_NODES];
static int graph_size = 0;
static MemoEntry memo[HASH_SIZE];
static bool on_path[MAX_NODES];
static long long call_count = 0;

static unsigned int hash_name(const char *name) {
  unsigned int hash = 5381;
  int c;
  while ((c = *name++)) {
    hash = ((hash << 5) + hash) + c;
  }
  return hash;
}

static int find_node(const char *name) {
  for (int i = 0; i < graph_size; i++) {
    if (strcmp(graph[i].name, name) == 0) {
      return i;
    }
  }
  return -1;
}

static unsigned int memo_hash(const char *node, bool visited_dac,
                              bool visited_fft) {
  unsigned int h = hash_name(node);
  h = h * 31 + (visited_dac ? 1 : 0);
  h = h * 31 + (visited_fft ? 1 : 0);
  return h % HASH_SIZE;
}

static bool memo_get(const char *node, bool visited_dac, bool visited_fft,
                     long long *value) {
  unsigned int h = memo_hash(node, visited_dac, visited_fft);
  unsigned int start = h;

  while (memo[h].valid) {
    if (strcmp(memo[h].node, node) == 0 && memo[h].visited_dac == visited_dac &&
        memo[h].visited_fft == visited_fft) {
      *value = memo[h].value;
      return true;
    }
    h = (h + 1) % HASH_SIZE;
    if (h == start)
      break;
  }
  return false;
}

static void memo_set(const char *node, bool visited_dac, bool visited_fft,
                     long long value) {
  unsigned int h = memo_hash(node, visited_dac, visited_fft);

  while (memo[h].valid) {
    if (strcmp(memo[h].node, node) == 0 && memo[h].visited_dac == visited_dac &&
        memo[h].visited_fft == visited_fft) {
      memo[h].value = value;
      return;
    }
    h = (h + 1) % HASH_SIZE;
  }

  strcpy(memo[h].node, node);
  memo[h].visited_dac = visited_dac;
  memo[h].visited_fft = visited_fft;
  memo[h].value = value;
  memo[h].valid = true;
}

static long long count_paths(const char *node, bool visited_dac,
                             bool visited_fft) {
  call_count++;
  if (call_count % 10000000 == 0) {
    printf("[Progress] Calls: %lld\n", call_count);
  }

  if (strcmp(node, "dac") == 0) {
    visited_dac = true;
  }
  if (strcmp(node, "fft") == 0) {
    visited_fft = true;
  }

  if (strcmp(node, "out") == 0) {
    if (visited_dac && visited_fft) {
      return 1;
    }
    return 0;
  }

  int node_idx = find_node(node);
  if (node_idx < 0) {
    return 0;
  }

  if (on_path[node_idx]) {
    return 0;
  }

  long long cached;
  if (memo_get(node, visited_dac, visited_fft, &cached)) {
    return cached;
  }

  if (graph[node_idx].output_count == 0) {
    return 0;
  }

  on_path[node_idx] = true;

  long long total = 0;
  for (int i = 0; i < graph[node_idx].output_count; i++) {
    total += count_paths(graph[node_idx].outputs[i], visited_dac, visited_fft);
  }

  on_path[node_idx] = false;

  memo_set(node, visited_dac, visited_fft, total);

  return total;
}

int main(void) {
  FILE *file = fopen("../list.txt", "r");
  if (!file) {
    printf("Error opening file\n");
    return 1;
  }

  char line[1024];
  while (fgets(line, sizeof(line), file)) {

    line[strcspn(line, "\n")] = 0;

    if (strlen(line) == 0) {
      continue;
    }

    char *colon = strchr(line, ':');
    if (!colon) {
      continue;
    }

    *colon = '\0';
    char *name = line;

    while (*name == ' ')
      name++;
    char *end = name + strlen(name) - 1;
    while (end > name && *end == ' ')
      *end-- = '\0';

    strcpy(graph[graph_size].name, name);
    graph[graph_size].output_count = 0;

    char *outputs_str = colon + 1;
    char *token = strtok(outputs_str, " \t");
    while (token) {
      strcpy(graph[graph_size].outputs[graph[graph_size].output_count], token);
      graph[graph_size].output_count++;
      token = strtok(NULL, " \t");
    }

    graph_size++;
  }

  fclose(file);

  printf("Total devices: %d\n", graph_size);

  long long total = count_paths("svr", false, false);

  printf("Total DFS calls: %lld\n", call_count);
  printf("Answer: %lld\n", total);

  return 0;
}
