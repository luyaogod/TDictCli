---
title: "Header files for ESQL/C typedefs"
source: "fgl-topics/c_fgl_CExtensions_003.html"
breadcrumb: "Extending the language > C-Extensions > Header files for ESQL/C typedefs"
type: "concept"
---

# Header files for ESQL/C typedefs

> C header files (.h) are required to define C structures for complex data types used in a C-Extension.

To compile C-Extensions using data types such as `DECIMAL`,
`DATETIME`/`INTERVAL` or `BYTE`/`TEXT`,
you need IBM® Informix® ESQL/C data type structure definitions such as `dec_t`,
`dtime_t`, `intrvl_t`, as well as macros such as
`TU_ENCODE()`.

These definitions are not required, if you use standard C types such as `short`,
`int` or `char[]`.

The definition of the ESQL/C structures like `dec_t` are provided in individual
header files, under the $FGLDIR/include/f2c directory:
fglDecimal.h, fglDatetime,
fglLocator.h.

In order to include these data type header files, simply include the
fglExt.h header file:

```
#include "f2c/fglExt.h"
```

The other header files are then included automatically.

## Related links

**Related concepts**  

[Data types and structures](2712-data-types-and-structures.md "C types are used to write C-Extensions.")
