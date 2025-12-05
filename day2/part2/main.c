#include <stdio.h>
#include <string.h>

int is_invalid_id(unsigned long long num) {
  char str[30];
  sprintf(str, "%llu", num);
  int len = strlen(str);

  for (int pattern_len = 1; pattern_len <= len / 2; pattern_len++) {

    if (len % pattern_len != 0) {
      continue;
    }

    int repeats = len / pattern_len;
    if (repeats < 2)
      continue;

    int valid = 1;
    for (int i = pattern_len; i < len && valid; i++) {
      if (str[i] != str[i % pattern_len]) {
        valid = 0;
      }
    }

    if (valid) {
      return 1;
    }
  }

  return 0;
}

unsigned long long build_repeated(unsigned long long pattern, int pattern_len,
                                  int repeats) {
  unsigned long long result = 0;
  unsigned long long multiplier = 1;

  for (int i = 0; i < repeats; i++) {
    result = result * multiplier + pattern;
    if (i == 0) {

      for (int j = 0; j < pattern_len; j++) {
        multiplier *= 10;
      }
    }
  }

  return result;
}

#define MAX_CANDIDATES 1000
unsigned long long candidates[MAX_CANDIDATES];
int num_candidates;

void add_candidate(unsigned long long val) {
  if (num_candidates < MAX_CANDIDATES) {
    candidates[num_candidates++] = val;
  }
}

unsigned long long next_invalid_id(unsigned long long start) {
  num_candidates = 0;

  char str[30];
  sprintf(str, "%llu", start);
  int start_len = strlen(str);

  for (int total_len = start_len; total_len <= start_len + 2 && total_len <= 20;
       total_len++) {

    for (int pattern_len = 1; pattern_len <= total_len / 2; pattern_len++) {
      if (total_len % pattern_len != 0)
        continue;

      int repeats = total_len / pattern_len;
      if (repeats < 2)
        continue;

      unsigned long long min_pattern = 1;
      for (int i = 1; i < pattern_len; i++)
        min_pattern *= 10;
      unsigned long long max_pattern = min_pattern * 10 - 1;

      unsigned long long lo = min_pattern, hi = max_pattern;
      unsigned long long best = 0;

      while (lo <= hi) {
        unsigned long long mid = lo + (hi - lo) / 2;
        unsigned long long num = build_repeated(mid, pattern_len, repeats);

        if (num >= start) {
          best = num;
          hi = mid - 1;
        } else {
          lo = mid + 1;
        }
      }

      if (best > 0) {
        add_candidate(best);
      }
    }
  }

  unsigned long long result = 0;
  for (int i = 0; i < num_candidates; i++) {
    if (result == 0 || candidates[i] < result) {
      result = candidates[i];
    }
  }

  return result;
}

unsigned long long sum_invalid_ids_in_range(unsigned long long start,
                                            unsigned long long end) {
  unsigned long long sum = 0;

  unsigned long long candidate = next_invalid_id(start);

  while (candidate != 0 && candidate <= end) {
    sum += candidate;
    candidate = next_invalid_id(candidate + 1);
  }

  return sum;
}

int main() {
  FILE *fp = fopen("list.txt", "r");
  if (!fp) {
    perror("Failed to open list.txt");
    return 1;
  }

  char content[10000];
  int pos = 0;
  int c;
  while ((c = fgetc(fp)) != EOF && pos < 9999) {
    content[pos++] = c;
  }
  content[pos] = '\0';
  fclose(fp);

  unsigned long long total_sum = 0;

  char *ptr = content;
  while (*ptr) {

    while (*ptr && (*ptr < '0' || *ptr > '9'))
      ptr++;
    if (!*ptr)
      break;

    unsigned long long range_start = 0;
    while (*ptr >= '0' && *ptr <= '9') {
      range_start = range_start * 10 + (*ptr - '0');
      ptr++;
    }

    while (*ptr && *ptr != '-' && (*ptr < '0' || *ptr > '9'))
      ptr++;
    if (*ptr == '-')
      ptr++;

    unsigned long long range_end = 0;
    while (*ptr >= '0' && *ptr <= '9') {
      range_end = range_end * 10 + (*ptr - '0');
      ptr++;
    }

    if (range_end > 0) {
      unsigned long long range_sum =
          sum_invalid_ids_in_range(range_start, range_end);
      total_sum += range_sum;
    }
  }

  printf("Sum of all invalid IDs: %llu\n", total_sum);

  return 0;
}
