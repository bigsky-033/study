#include <stdio.h>

#include "buffered_print.h"

int main(void)
{
    buffered_print("Hello, ");
    printf("--- buffered_print(\"Hello, \") is called \n");

    buffered_print("World. ");
    printf("--- buffered_print(\"World, \") is called \n");

    buffered_print("Java is awesome! ");
    printf("--- buffered_print(\"Java is awesome! \") is called \n");

    buffered_print("What about C#? ");
    printf("--- buffered_print(\"What about C#? \") is called \n");

    return 0;
}
