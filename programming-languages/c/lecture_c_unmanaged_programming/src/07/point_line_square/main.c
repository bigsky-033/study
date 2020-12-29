#include <stdio.h>

#include "shape.h"

int main(void)
{
    point_t p0;
    point_t p1;
    line_t line0;
    line_t line1;
    rectangle_t rect;

    /* init points */
    p0.x = 5;
    p0.y = 5;

    p1.x = -3;
    p1.y = 5;

    /* these 2 lines are reversed */
    line0.start = p0;
    line0.end = p1;

    line1.start = p1;
    line1.end = p0;

    printf("line0 length^2: %d\n", get_line_length_sq(line0));
    printf("line1 length^2: %d\n", get_line_length_sq(line1));

    rect = build_rectangle(p0, p1);
    printf("rect0: (%d, %d) (%d, %d)\n",
           rect.top_left.x, rect.top_left.y,
           rect.bottom_right.x, rect.bottom_right.y);

    rect = build_rectangle(p1, p0);
    printf("rect1: (%d, %d) (%d, %d)\n",
           rect.top_left.x, rect.top_left.y,
           rect.bottom_right.x, rect.bottom_right.y);

    printf("rect area: %d\n", get_rectangle_area(rect));

    /* improper rect build */
    rect.top_left = p0;
    rect.bottom_right = p1;

    printf("rect area: %d\n", get_rectangle_area(rect));

    return 0;
}
