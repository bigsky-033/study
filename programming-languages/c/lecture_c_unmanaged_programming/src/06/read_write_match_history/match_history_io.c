#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <string.h>

#include "match_history_io.h"

#define LINE_LENGTH (4096)

void write_match_history(char *buffer,
                         size_t buffer_size,
                         const char *names[],
                         const int wins[],
                         const int losses[],
                         const float kills[],
                         const float deaths[],
                         const float assists[],
                         size_t count)
{
    char line[LINE_LENGTH];
    size_t i;
    size_t remaining_buffer_size;

    remaining_buffer_size = buffer_size - 1; /* -1은 널 문자 */
    buffer[0] = '\0';                        /* 이 버퍼는 지금 빈 문자열이다 라는 의미이다. */

    for (i = 0; i < count; ++i)
    {
        size_t num_written;

        sprintf(line, "%8s %.2f %.2f %.2f %d %d\n",
                names[i],
                kills[i], deaths[i], assists[i],
                wins[i], losses[i]);

        num_written = strlen(line);
        if (num_written > remaining_buffer_size)
        {
            break;
        }

        strcpy(buffer, line);
        buffer += num_written; /* strcpy를 한 후 버퍼의 위치(포인터)를 옮겨 주었다. */
        remaining_buffer_size -= num_written;
    }
}

void read_match_histroy(char *buffer)
{
    const char *DELIM = "\n";
    char *tokenizer;

    printf("%8s %7s %7s %7s %7s %6s %6s %9s\n",
           "Champ", "Kills", "Deaths", "Assists",
           "KDA", "Wins", "Losses", "Win Ratio");

    tokenizer = strtok(buffer, DELIM);
    while (tokenizer != NULL)
    {
        char name[LINE_LENGTH];
        float kills;
        float deaths;
        float assists;
        int wins;
        int losses;

        float kda;
        float win_ratio;

        if (sscanf(tokenizer, "%s %f %f %f %d %d",
                   name, &kills, &deaths, &assists, &wins, &losses) != 6)
        {
            continue;
        }

        kda = (kills + assists) / deaths;
        win_ratio = wins * 100.0f / (wins + losses);

        printf("%8s %7.2f %7.2f %7.2f %7.2f %6d %6d %8.2f%%\n",
               name, kills, deaths, assists, kda, wins, losses, win_ratio);

        tokenizer = strtok(NULL, DELIM);
    }
}
