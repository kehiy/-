#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main() {

    srand(time(0));

    for (int i = 0; i < 10; i++) {
        int randnum = rand();
        printf("%d\n", randnum);
    }

    return 0;
}