#include <stdio.h>

int main()
{
    for (int i = 1; i <= 9; i++)
    {
        if(i % 3 == 0)
        {
            printf("%d\n", i);
        } else {
            printf("%d\t", i);
        }
    }
    
    return 0;
}
