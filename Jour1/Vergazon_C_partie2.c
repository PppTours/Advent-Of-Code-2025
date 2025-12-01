// Compile commands :
// Launch these from the repository root
// > mkdir build
// > clang Jour1/Vergazon_C_partie2.c -o build/vergazonJour1
// > ./build/vergazonJour1

#include <stdio.h>
#include <stdlib.h>

int main() {

  // File-reading stuff
  FILE *fp;
  char *line = NULL;
  unsigned long len = 0;
  long read;

  // Problem stuff
  int current_position = 50;
  int password = 0;

  fp = fopen("Jour1/Vergazon_inputData.txt", "r");
  if (fp == NULL)
    exit(EXIT_FAILURE);

  while ((read = getline(&line, &len, fp)) != -1) {
    char direction = line[0];
    int delta = atoi(line + 1); // Contains the amount to rotate
    char left = 1;              // -1 if rotating left, 1 if rotating right
    if (direction == 'L') {
      left = -1;
    }

    // This time, we need to be a bit more careful about how we rotate.
    // the delta variable contains the amount to rotate.
    // the left variable contains the direction to rotate : -1 if left and 1 if
    // right.
    // We need to apply each rotation "click" individually and check for zero
    // here.
    printf("Delta : %d\n", delta);

    while (delta > 0) { // Warning : strict inequality

      // Apply the click
      current_position += left; // Substract 1 if left, add 1 if right
      // Check for overflow / underflow
      if (current_position < 0) {
        current_position += 100;
      } else if (current_position > 99) {
        current_position -= 100;
      }
      // Check for equality with 0
      if (current_position == 0) {
        password++;
      }
      delta--;
    }
  }

  fclose(fp);
  free(line);

  printf("The password is : %d\n", password);
  // Password is 5933
  return 0;
}
