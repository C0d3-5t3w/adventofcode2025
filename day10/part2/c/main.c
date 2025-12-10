#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE 4096
#define MAX_COUNTERS 16
#define MAX_BUTTONS 20

int parse_button(const char *start, const char **end, int *indices) {
  int count = 0;
  const char *p = start + 1;
  while (*p && *p != ')') {
    if (*p >= '0' && *p <= '9') {
      int num = 0;
      while (*p >= '0' && *p <= '9') {
        num = num * 10 + (*p - '0');
        p++;
      }
      indices[count++] = num;
    } else {
      p++;
    }
  }
  if (*p == ')')
    p++;
  *end = p;
  return count;
}

int parse_targets(const char *start, int *targets) {
  int count = 0;
  const char *p = start + 1;
  while (*p && *p != '}') {
    if (*p >= '0' && *p <= '9') {
      int num = 0;
      while (*p >= '0' && *p <= '9') {
        num = num * 10 + (*p - '0');
        p++;
      }
      targets[count++] = num;
    } else {
      p++;
    }
  }
  return count;
}

int g_num_counters;
int g_num_buttons;
int g_targets[MAX_COUNTERS];
int g_buttons[MAX_BUTTONS][MAX_COUNTERS];
int g_button_counts[MAX_BUTTONS];
int g_coeff[MAX_COUNTERS][MAX_BUTTONS];

bool check_solution(int presses[]) {
  for (int c = 0; c < g_num_counters; c++) {
    int sum = 0;
    for (int b = 0; b < g_num_buttons; b++) {
      sum += g_coeff[c][b] * presses[b];
    }
    if (sum != g_targets[c])
      return false;
  }
  return true;
}

long long solve_recursive(int presses[], int button_idx, int current_sum,
                          long long best_so_far) {

  if (current_sum >= best_so_far)
    return LLONG_MAX;

  if (button_idx == g_num_buttons) {

    if (check_solution(presses)) {
      return current_sum;
    }
    return LLONG_MAX;
  }

  int max_presses = INT_MAX;
  for (int i = 0; i < g_button_counts[button_idx]; i++) {
    int counter = g_buttons[button_idx][i];
    if (counter < g_num_counters && g_targets[counter] < max_presses) {
      max_presses = g_targets[counter];
    }
  }

  if (best_so_far != LLONG_MAX && best_so_far - current_sum < max_presses) {
    max_presses = (int)(best_so_far - current_sum);
  }

  long long result = best_so_far;

  for (int p = 0; p <= max_presses; p++) {
    presses[button_idx] = p;
    long long sub_result =
        solve_recursive(presses, button_idx + 1, current_sum + p, result);
    if (sub_result < result) {
      result = sub_result;
    }
  }

  presses[button_idx] = 0;
  return result;
}

long long solve_gauss_with_optimization(void) {

  double matrix[MAX_COUNTERS][MAX_BUTTONS + 1];
  for (int i = 0; i < g_num_counters; i++) {
    for (int j = 0; j < g_num_buttons; j++) {
      matrix[i][j] = g_coeff[i][j];
    }
    matrix[i][g_num_buttons] = g_targets[i];
  }

  int pivot_row = 0;
  int pivot_col[MAX_COUNTERS];
  int num_pivots = 0;
  bool is_pivot_col[MAX_BUTTONS] = {false};

  for (int col = 0; col < g_num_buttons && pivot_row < g_num_counters; col++) {

    int max_row = pivot_row;
    for (int row = pivot_row + 1; row < g_num_counters; row++) {
      if (fabs(matrix[row][col]) > fabs(matrix[max_row][col])) {
        max_row = row;
      }
    }

    if (fabs(matrix[max_row][col]) < 1e-9)
      continue;

    for (int j = 0; j <= g_num_buttons; j++) {
      double tmp = matrix[pivot_row][j];
      matrix[pivot_row][j] = matrix[max_row][j];
      matrix[max_row][j] = tmp;
    }

    double pivot_val = matrix[pivot_row][col];
    for (int j = 0; j <= g_num_buttons; j++) {
      matrix[pivot_row][j] /= pivot_val;
    }

    for (int row = 0; row < g_num_counters; row++) {
      if (row != pivot_row && fabs(matrix[row][col]) > 1e-9) {
        double factor = matrix[row][col];
        for (int j = 0; j <= g_num_buttons; j++) {
          matrix[row][j] -= factor * matrix[pivot_row][j];
        }
      }
    }

    pivot_col[num_pivots++] = col;
    is_pivot_col[col] = true;
    pivot_row++;
  }

  int free_vars[MAX_BUTTONS];
  int num_free = 0;
  for (int col = 0; col < g_num_buttons; col++) {
    if (!is_pivot_col[col]) {
      free_vars[num_free++] = col;
    }
  }

  if (num_free == 0) {
    double solution[MAX_BUTTONS] = {0};
    for (int i = num_pivots - 1; i >= 0; i--) {
      int col = pivot_col[i];
      solution[col] = matrix[i][g_num_buttons];
      for (int j = col + 1; j < g_num_buttons; j++) {
        solution[col] -= matrix[i][j] * solution[j];
      }
    }

    long long total = 0;
    int int_sol[MAX_BUTTONS];
    for (int i = 0; i < g_num_buttons; i++) {
      int_sol[i] = (int)round(solution[i]);
      if (int_sol[i] < 0)
        return LLONG_MAX;
      total += int_sol[i];
    }

    if (check_solution(int_sol))
      return total;
    return LLONG_MAX;
  }

  int max_free_val = 0;
  for (int i = 0; i < g_num_counters; i++) {
    if (g_targets[i] > max_free_val)
      max_free_val = g_targets[i];
  }

  long long best = LLONG_MAX;
  int free_vals[MAX_BUTTONS] = {0};

  while (true) {

    double solution[MAX_BUTTONS] = {0};
    for (int i = 0; i < num_free; i++) {
      solution[free_vars[i]] = free_vals[i];
    }

    bool valid = true;
    for (int i = 0; i < num_pivots; i++) {
      int col = pivot_col[i];
      double val = matrix[i][g_num_buttons];
      for (int j = col + 1; j < g_num_buttons; j++) {
        val -= matrix[i][j] * solution[j];
      }
      solution[col] = val;

      if (val < -0.5) {
        valid = false;
        break;
      }
    }

    if (valid) {

      int int_sol[MAX_BUTTONS];
      long long total = 0;
      bool all_nonneg = true;
      for (int i = 0; i < g_num_buttons; i++) {
        int_sol[i] = (int)round(solution[i]);
        if (int_sol[i] < 0) {
          all_nonneg = false;
          break;
        }
        total += int_sol[i];
      }

      if (all_nonneg && total < best && check_solution(int_sol)) {
        best = total;
      }
    }

    int idx = 0;
    while (idx < num_free) {
      free_vals[idx]++;
      if (free_vals[idx] <= max_free_val)
        break;
      free_vals[idx] = 0;
      idx++;
    }
    if (idx == num_free)
      break;
  }

  return best;
}

long long solve_min_presses(int targets[], int num_counters,
                            int buttons[][MAX_COUNTERS], int button_counts[],
                            int num_buttons) {

  g_num_counters = num_counters;
  g_num_buttons = num_buttons;
  memcpy(g_targets, targets, num_counters * sizeof(int));
  memcpy(g_buttons, buttons, num_buttons * MAX_COUNTERS * sizeof(int));
  memcpy(g_button_counts, button_counts, num_buttons * sizeof(int));

  memset(g_coeff, 0, sizeof(g_coeff));
  for (int b = 0; b < num_buttons; b++) {
    for (int i = 0; i < button_counts[b]; i++) {
      int counter = buttons[b][i];
      if (counter < num_counters) {
        g_coeff[counter][b] = 1;
      }
    }
  }

  long long result = solve_gauss_with_optimization();

  if (result == LLONG_MAX) {

    if (num_buttons <= 8) {
      int presses[MAX_BUTTONS] = {0};
      result = solve_recursive(presses, 0, 0, LLONG_MAX);
    }
  }

  return result == LLONG_MAX ? -1 : result;
}

int main(void) {
  FILE *fp = fopen("../list.txt", "r");
  if (!fp) {
    perror("Failed to open list.txt");
    return 1;
  }

  char line[MAX_LINE];
  long long total = 0;

  while (fgets(line, sizeof(line), fp)) {
    if (line[0] == '\n' || line[0] == '\0')
      continue;

    char *brace = strchr(line, '{');
    if (!brace)
      continue;

    int targets[MAX_COUNTERS];
    int num_counters = parse_targets(brace, targets);

    int buttons[MAX_BUTTONS][MAX_COUNTERS];
    int button_counts[MAX_BUTTONS];
    int num_buttons = 0;

    const char *p = line;
    while (*p) {
      while (*p && *p != '(' && *p != '{')
        p++;
      if (*p == '{' || *p == '\0')
        break;

      const char *end;
      button_counts[num_buttons] = parse_button(p, &end, buttons[num_buttons]);
      num_buttons++;
      p = end;
    }

    long long min = solve_min_presses(targets, num_counters, buttons,
                                      button_counts, num_buttons);
    if (min >= 0) {
      total += min;
    }
  }

  fclose(fp);
  printf("%lld\n", total);

  return 0;
}
