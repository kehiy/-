#include <stdio.h>

int main() {
    FILE *pF = fopen("lies.txt", "a");
    fprintf(pF, "C is awful.\n");
    fprintf(pF, "Javascript is awesome!!!\n");
    fclose(pF);

    FILE *pRF = fopen("lies.txt", "r");
    char buffer[255];

    if (pRF == NULL) {
        printf("Can't open the file.");
        return 1;
    }

    while (fgets(buffer, 255, pRF)) {
        printf("%s", buffer);
    }

    return 0;
}
