#include <stdio.h>

#include "string_utils.h"

int main(void)
{
    char str[15] = "Welcome to C";

    printf("Is space alphabet?: %s\n", is_alpha(' ') ? "Yes" : "No");

    printf("m is uppercase: %c\n", to_upper('m'));
    printf("W in lowercase: %c\n", to_lower('W'));

    string_toupper(str);
    printf("Uppercase: %s\n", str);

    string_tolower(str);
    printf("Lowercase: %s\n", str);

    return 0;
}
