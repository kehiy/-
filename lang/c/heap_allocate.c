#include <assert.h>
#include <stdio.h>
#include <stdbool.h>

typedef struct {
    void *start;
    size_t size;
} Heap_Chunk;

#define HEAP_CAP 640000
#define HEAP_ALLOCED_CAP 1024
#define HEAP_FREED_CAP 1024

Heap_Chunk heap_alloced[HEAP_ALLOCED_CAP] = {0};

size_t heap_size = 0;
size_t heap_alloced_size = 0;

Heap_Chunk heap_freed[HEAP_FREED_CAP] = {0};
size_t heap_freed_size = 0;

char heap[HEAP_CAP] = {0};

void *halloc(size_t size)
{
	if (size == 0) {
		return NULL;
	}

    assert(heap_size + size <= HEAP_CAP);

    void *result = heap + heap_size;
    heap_size += size;

    const Heap_Chunk chunk = {
        .start = result,
        .size = size,
    };
    assert(heap_alloced_size < HEAP_ALLOCED_CAP);

    heap_alloced[heap_alloced_size++] = chunk;

    return result;
}

void hfree(void *ptr)
{
    for (size_t i = 0; i < heap_alloced_size; i++) {
		if (heap_alloced[i].start == ptr) {
			printf("free the memory!");
		}
	}
}

void hcollect()
{
    assert(false && "Not implemented");
}

void dump_alloced_chunks(void)
{
    for (size_t i = 0; i < heap_alloced_size; i++) {
        printf("\nstart: %p, size: %zu\n",
	    heap_alloced[i].start,
	    heap_alloced[i].size
	);
    }
}

int main(void)
{
    for (int i = 0; i < 100; i++) {
		void *p = halloc(i);
		if ((i & 0x1) == 0) {
			hfree(p);
		}
	}

    dump_alloced_chunks();

    return 0;
}
