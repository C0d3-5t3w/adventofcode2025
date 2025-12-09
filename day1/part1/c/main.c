#include <stdio.h>

int main() {
  FILE *file = fopen("../list.txt", "r");
  if (!file) {
    perror("Failed to open list.txt");
    return 1;
  }

  int dial = 50;
  int zero_count = 0;
  char direction;
  int distance;

  while (fscanf(file, " %c%d", &direction, &distance) == 2) {
    if (direction == 'L') {
      dial = (dial - distance) % 100;
      if (dial < 0)
        dial += 100;
    } else if (direction == 'R') {
      dial = (dial + distance) % 100;
    }

    if (dial == 0) {
      zero_count++;
    }
  }

  fclose(file);

  printf("Answer: %d\n", zero_count);

  return 0;
}

