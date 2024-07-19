#include <assert.h>
#include <stdio.h>
#include <stdbool.h>

#define CAPACITY 640000
size_t heap_size = 0;

char heap[CAPACITY] = {0};

void *halloc(size_t size)
{
    assert(heap_size + size <= CAPACITY);

    void *result = heap + heap_size;
    heap_size += size;

    return result;	
}

void hfree(void *ptr)
{
    (void) ptr;
    assert(false && "Not implemented");
}

void hcollect()
{
    assert(false && "Not implemented");
}

int main(void)
{
    char *root = halloc(26);
    for (int i = 0; i < 26; i++) {
        root[i] = i + 'A';
	printf("%c\n", root[i]);
    }

    return 0;
}
