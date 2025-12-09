#include <stdio.h>
#include <string.h>

int is_invalid_id(unsigned long long num) {
  char str[30];
  sprintf(str, "%llu", num);
  int len = strlen(str);

  if (len % 2 != 0) {
    return 0;
  }

  int half = len / 2;

  for (int i = 0; i < half; i++) {
    if (str[i] != str[half + i]) {
      return 0;
    }
  }

  return 1;
}

unsigned long long next_invalid_id(unsigned long long start) {

  char str[30];
  sprintf(str, "%llu", start);
  int len = strlen(str);

  int target_len = (len % 2 == 0) ? len : len + 1;

  while (target_len <= 20) {
    int half = target_len / 2;

    unsigned long long min_half = 1;
    for (int i = 1; i < half; i++)
      min_half *= 10;

    unsigned long long max_half = min_half * 10 - 1;

    unsigned long long start_half = min_half;

    if (target_len == len || (len % 2 != 0 && target_len == len + 1)) {

      unsigned long long multiplier = 1;
      for (int i = 0; i < half; i++)
        multiplier *= 10;
      multiplier += 1;

      unsigned long long needed = (start + multiplier - 1) / multiplier;
      if (needed > start_half) {
        start_half = needed;
      }
    }

    if (start_half <= max_half) {

      unsigned long long multiplier = 1;
      for (int i = 0; i < half; i++)
        multiplier *= 10;

      return start_half * multiplier + start_half;
    }

    target_len += 2;
  }

  return 0;
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
  FILE *fp = fopen("../list.txt", "r");
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

