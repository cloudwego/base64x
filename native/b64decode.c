#include "native.h"

ssize_t b64decode(struct slice_t *out, const char *src, size_t nb, int mode) {
    return do_b64decode(out, src, nb, mode);
}