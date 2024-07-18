#include <stdio.h>

void sort(int array[], int size) {
    for (int i = 0; i < size - 1; i++) {
        for (int j = 0; j < size - i - 1; j++) {
            if (array[j] > array[j+1]) {
                int temp = array[j];
                array[j] = array[j+1];
                array[j+1] = temp; 
            }
        }
    }
}

void printArray(int array[], int size) {
    for (int i = 0; i < size - 1; i++) {
        printf("%d\n", array[i]);
    }
}

int main() {
    int numbers[] = {1, 2, 5, 6, 8, 34, 35, 4365, 5};
    int size = sizeof(numbers)/sizeof(numbers[0]);

    sort(numbers, size);
    printArray(numbers, size);

    return 0;
}
