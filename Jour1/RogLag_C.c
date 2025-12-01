#include <stdio.h>
#include <stdlib.h>

int main(int argc, char* argv[]) {
    FILE* file = fopen(argv[1], "r");

    int pos = 50;
    int zeros = 0;
    char dir;
    int steps;

    while (fscanf(file, " %c%d", &dir, &steps) == 2) {
        for (int i = 0; i < steps; i++) {
            if (dir == 'L') {
                pos = (pos - 1 + 100) % 100;
            } else {
                pos = (pos + 1) % 100;
            }
            if (pos == 0) zeros++;
        }
    }

    fclose(file);
    printf("Password: %d\n", zeros);
    return EXIT_SUCCESS;
}