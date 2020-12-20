#include <stdio.h>

#include "minion.h"

extern unsigned int g_hp;
extern unsigned int g_strength;

int main(void)
{
    printf("Normal minion:\n");
    printf("hp: %u\n", g_hp);
    printf("strength: %u\n", g_strength);
    /* this does not compile */
    /* printf("gold: %u\n", g_gold); */

    printf("\n");

    go_berserk();
    add_gold(10u);

    printf("Berserk minion:\n");
    printf("hp: %u\n", g_hp);
    printf("strength: %u\n", g_strength);
    /* this does not compile */
    /* printf("gold: %u\n", g_gold); */

    return 0;
}
