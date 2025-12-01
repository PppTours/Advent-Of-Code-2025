// Compile commands :
// Launch these from the repository root
// > mkdir build
// > clang Jour1/Vergazon_C_partie1.c -o build/vergazonJour1
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

    int rotation = delta % 100; // Modulate the delta by 100.
    rotation = left * rotation; // Only negate the rotation if left is 1
    int next_position = current_position + rotation; // Apply the rotation
    if (next_position >= 100) {                      // Overflow
      next_position -= 100;
    } else if (next_position < 0) {
      next_position += 100; // Underflow
    }
    current_position = next_position;
    if (current_position == 0) {
      password++;
    }
  }

  fclose(fp);
  free(line);

  printf("The password is : %d\n", password);
  // Password is 1021
  return 0;
}
