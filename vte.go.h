#pragma once

#include <stdlib.h>

static inline char ** make_strings(int count)
{
    return (char **)malloc(sizeof(char *) * count);
}

static inline void destroy_strings(char **strings, int count)
{
    for (int i = 0; i < count; i++) {
        free(strings[i]);
    }
    free(strings);
}

static inline void set_string(char **strings, int n, char *str)
{
    strings[n] = str;
}

static VteTerminal *toVteTerminal(void *p)
{
    return (VTE_TERMINAL(p));
}