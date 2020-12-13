#include <stdio.h>

int add(int op1, int op2)
{
    printf("add()\n");
    return op1 + op2;
}

int subtract(int op1, int op2)
{
    printf("subtract()\n");
    return op1 - op2;
}

int main(void)
{
    int num1 = 128;
    int num2 = 256;

    printf("%d, %d\n", add(num1, num2), subtract(num1, num2));

    return 0;
}
