#ifndef HASH_FUNCTIONS_H
#define HASH_FUNCTIONS_H

#include <stddef.h>

size_t hash_int(int value);

size_t hash_float(float value);

size_t hash_data(const void *data, size_t length);

#endif /* HASH_FUNCTIONS_H */
