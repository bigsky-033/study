#ifndef MATCH_HISTORY_IO_H
#define MATCH_HISTORY_IO_H

#include <stdio.h>

void write_match_history(char *buffer,
                         size_t buffer_size,
                         const char *names[],
                         const int wins[],
                         const int losses[],
                         const float kills[],
                         const float deaths[],
                         const float assists[],
                         size_t count);

void read_match_histroy(char *buffer);

#endif
