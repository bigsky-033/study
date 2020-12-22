#include <stdio.h>

void print_pointer_size()
{
    char ch = 'c';
    int number = 934563;
    float pi = 3.1415f;

    char *char_ptr = &ch;
    int *int_ptr = &number;
    float *float_ptr = &pi;

    printf("char size: %d, char* size: %d\n", sizeof(*char_ptr), sizeof(char_ptr));
    printf("int size: %d, int* size: %d\n", sizeof(*int_ptr), sizeof(int_ptr));
    printf("float size: %d, float* size: %d\n", sizeof(*float_ptr), sizeof(float_ptr));
}

int main(void)
{
    print_pointer_size();

    return 0;
}
