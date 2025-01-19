#include "native.h"

void b64encode(struct slice_t *out, const struct slice_t *src, int mode) {
     do_b64encode(out, src, mode);
}


