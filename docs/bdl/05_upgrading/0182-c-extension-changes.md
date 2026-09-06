---
title: "C Extension changes"
source: "fgl-topics/c_fgl_Migrate_to_310_cext.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > C Extension changes"
type: "concept"
---

# C Extension changes

> Modifications to consider when using C Extensions

## Informix ESQL/C header files no longer distributed

Prior to version 3.10.11, Informix ESQL/C header files such as decimal.h
were provided in $FGLDIR/include/esql directory.

Starting with BDL 3.10.11, the Informix ESQL/C header files are no longer distributed in the BDL
packages: Genero provides its own header files defining the C structures for
`DECIMAL`, `DATETIME`, `INTERVAL` and
`TEXT/BYTE` types.

Simply include the fglExt.h header file in your C extension source.

If you need additional type definitions that are not provided in Genero C Extension header files,
install the latest Informix CSDK, and include the Informix header files before
fglExt.h.

For more details, see [Header files for ESQL/C typedefs](../14_extending-the-language/2705-header-files-for-esql-c-typedefs.md "C header files (.h) are required to define C structures for complex data types used in a C-Extension.").

## C Extension API functions for bigint

Genero BDL 2.51 has desupported C Extension stack functions, that were introduced again in 3.10
(also backported in 3.00.10).

The following C API functions are available again:

- `popbigint(bigint *dst)`
- `pushbigint(bigint val)`

The following C type definitions are available again:

- `bigint`: Defines a 8-byte signed integer

For more details, see [Runtime stack functions](../14_extending-the-language/2711-runtime-stack-functions.md "To pass values between a C function and a program, the C function and the runtime system use the runtime stack.")

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
