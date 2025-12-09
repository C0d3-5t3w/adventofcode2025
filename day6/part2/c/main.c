#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE 16384
#define MAX_ROWS 10

int main() {
  FILE *file = fopen("../list.txt", "r");
  if (file == NULL) {
    perror("Failed to open file");
    return EXIT_FAILURE;
  }

  char lines[MAX_ROWS][MAX_LINE];
  int num_lines = 0;
  size_t max_len = 0;

  while (fgets(lines[num_lines], MAX_LINE, file) != NULL &&
         num_lines < MAX_ROWS) {
    size_t len = strlen(lines[num_lines]);

    if (len > 0 && lines[num_lines][len - 1] == '\n') {
      lines[num_lines][len - 1] = '\0';
      len--;
    }
    if (len > max_len)
      max_len = len;
    num_lines++;
  }
  fclose(file);

  int op_row = num_lines - 1;
  int num_number_rows = op_row;

  for (int i = 0; i < num_lines; i++) {
    size_t len = strlen(lines[i]);
    for (size_t j = len; j < max_len; j++) {
      lines[i][j] = ' ';
    }
    lines[i][max_len] = '\0';
  }

  long long grand_total = 0;
  size_t col = 0;

  while (col < max_len) {

    bool all_space = true;
    for (int r = 0; r < num_number_rows; r++) {
      if (lines[r][col] != ' ') {
        all_space = false;
        break;
      }
    }
    if (all_space) {
      col++;
      continue;
    }

    size_t start_col = col;

    while (col < max_len) {
      all_space = true;
      for (int r = 0; r < num_number_rows; r++) {
        if (lines[r][col] != ' ') {
          all_space = false;
          break;
        }
      }
      if (all_space)
        break;
      col++;
    }
    size_t end_col = col;

    long long numbers[256];
    int num_count = 0;
    char op = '+';

    size_t c = end_col;
    while (c > start_col) {
      c--;

      bool has_digit = false;
      for (int r = 0; r < num_number_rows; r++) {
        if (lines[r][c] >= '0' && lines[r][c] <= '9') {
          has_digit = true;
          break;
        }
      }

      if (!has_digit) {

        continue;
      }

      long long num = 0;
      for (int r = 0; r < num_number_rows; r++) {
        char ch = lines[r][c];
        if (ch >= '0' && ch <= '9') {
          num = num * 10 + (ch - '0');
        }
      }
      numbers[num_count++] = num;
    }

    for (size_t c = start_col; c < end_col; c++) {
      if (lines[op_row][c] == '+' || lines[op_row][c] == '*') {
        op = lines[op_row][c];
        break;
      }
    }

    if (num_count > 0) {
      long long result;
      if (op == '+') {
        result = 0;
        for (int i = 0; i < num_count; i++) {
          result += numbers[i];
        }
      } else {
        result = 1;
        for (int i = 0; i < num_count; i++) {
          result *= numbers[i];
        }
      }
      grand_total += result;
    }
  }

  printf("Grand Total: %lld\n", grand_total);

  return EXIT_SUCCESS;
}
