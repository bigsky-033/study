#include "string_utils.h"

int is_alpha(int c)
{
    return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z');
}

int to_upper(int c)
{
    if (is_alpha(c))
    {
        /* c로부터 32를 빼는 것과 마찬가지다. 32를 가리키는 비트를 없애 버리겠다는 것이다. */
        /* a(110 0001)과 z(111 1010) 사이의 모든 알파벳은 32를 가리키는 6번째 비트가 1로 되어있다. */
        /* 이 비트를 0으로 만들어 버리겠다는 의미이다. */
        /* ~는 비트를 뒤집는 연산이다. 그래서, 01 1111 을 이용해 c와 & 연산을 하겠다는 것이다. */
        /* A와 Z 사이의 문자에는 아무런 영향이 없다. 이 값들의 6번째 비트는 이미 0이기 때문이다. */
        return c & ~0x20;
    }
    return c;
}

int to_lower(int c)
{
    if (is_alpha(c))
    {
        return c | 0x20;
    }
    return c;
}

void string_toupper(char *str)
{
    while (*str != '\0')
    {
        *str = to_upper(*str);
        ++str;
    }
}

void string_tolower(char *str)
{
    while (*str != '\0')
    {
        *str = to_lower(*str);
        ++str;
    }
}
