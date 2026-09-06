---
title: "Data types and structures"
source: "fgl-topics/c_fgl_CExtensions_010.html"
breadcrumb: "Extending the language > C-Extensions > Data types and structures"
type: "concept"
---

# Data types and structures

> C types are used to write C-Extensions.

The following C types are used to write C-Extensions.

| Type name | Description |
| --- | --- |
| `bigint` | signed integer with a size of 8 bytes |
| `int4` | signed integer with a size of 4 bytes |
| `uint4` | unsigned integer with a size of 4 bytes |
| `int2` | signed integer with a size of 2 bytes |
| `uint2` | unsigned integer with a size of 2 bytes |
| `int1` | signed integer with a size of 1 byte |
| `uint1` | unsigned integer with a size of 1 byte |
| `mint` | signed machine-dependent C int |
| `muint` | unsigned machine-dependent C int |
| `mlong` | signed machine-dependent C long |
| `mulong` | unsigned machine-dependent C long |
| `dec_t` | `DECIMAL` data type structure |
| `dtime_t` | `DATETIME` data type structure |
| `intrvl_t` | `INTERVAL` data type structure |
| `ifx_loc_t` | `TEXT` / `BYTE` locator structure |

## Basic data types

Basic data types such as `bigint`, `int4` and `int2`
are provided to define variables that must hold `BIGINT` (`bigint`),
`SMALLINT` (`int2`), `INTEGER` (`int4`)
and `DATE` (`int4`) values. Standard `char` array can
be used to hold `CHAR` and `VARCHAR` data.

## DATE

No specific typedef exists for the `DATE` type; you can use the `int4`
type to store a `DATE` value.

## DECIMAL/MONEY

The `dec_t` structure is provided to hold `DECIMAL` and
`MONEY` values.

The internals of `dec_t` structure can be ignored
during C-Extension programming, because decimal API functions are provided to manipulate any aspects
of a decimal.

## DATETIME

The `dtime_t` structure holds a `DATETIME` value.

Before manipulating a `dtime_t`, you must initialize its qualifier
`qt_qual`, by using the `TU_DTENCODE` macro:

```
dtime_t dt;
dt.dt_qual = TU_DTENCODE(TU_YEAR, TU_SECOND);
dtcvasc( "2004-02-12 12:34:56", &dt );
```

## INTERVAL

The `intrvl_t` structure holds an `INTERVAL` value.

Before manipulating an `intrvl_t`, you must initialize its qualifier
`in_qual`, by using the `TU_IENCODE` macro:

```
intrvl_t in;
in.in_qual = TU_IENCODE(5, TU_YEAR, TU_MONTH);
incvasc( "65234-02", &in );
```

## TEXT/BYTE Locator

The `ifx_loc_t` structure is used to declare host variables for a
`TEXT`/`BYTE` values (simple large objects). Because the potential
size of the data can be quite large, this is a locator structure that contains information about the
size and location of the `TEXT`/`BYTE` data, rather than containing
the actual data.

| Field name | Data type | Description |
| --- | --- | --- |
| `loc_indicator` | `int4` | Null indicator; a value of -1 indicates a null `TEXT`/`BYTE` value. Your program can set the field to indicate the insertion of a null value. Database client libraries set the value for selects and fetches. |
| `loc_type` | `int4` | data type - `SQLTEXT` (for `TEXT` values) or `SQLBYTES` (for `BYTE` values). |
| `loc_size` | `int4` | Size of the `TEXT` / `BYTE` value in bytes; your program sets the size of the large object for insertions. Database client libraries set the size for selects and fetches. |
| `loc_loctype` | `int2` | Location - `LOCMEMORY` (in memory) or `LOCFNAME` (in a named file). Set `loc_loctype` after you declare the locator variable and before this declared variable receives the large object value. |
| `loc_buffer` | `char *` | If `loc_loctype` is `LOCMEMORY`, this is the location of the `TEXT`/`BYTE` value; your program must allocate space for the buffer and store its address here. |
| `loc_bufsize` | `int4` | If `loc_loctype` is `LOCMEMORY`, this is the size of the buffer loc\_buffer; If you set `loc_bufsize` to -1, database client libraries will allocate the memory buffer for selects and fetches. Otherwise, it is assumed that your program will handle memory allocation and de-allocation. |
| `loc_fname` | `char *` | If `loc_loc_type` is `LOCFNAME`, this is the address of the path name string that contains the file. |

## Example

C Extension source (ext1.c):

```
#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>

#include "f2c/fglExt.h"

int lob_size(int);

UsrFunction usrFunctions[]={
  { "lob_size",     lob_size,       1, 1 },
  { NULL,           NULL,           0, 0 }
};

int lob_size(int pc)
{
    ifx_loc_t *pb1;
    double ratio;
    char *source = NULL;
    char *psource = NULL;
    int size;

    if (pc != 1) exit(1);
    poplocator(&pb1);

    if (pb1->loc_loctype == LOCMEMORY) {
        psource = pb1->loc_buffer;
        size = pb1->loc_size;
    } else if (pb1->loc_loctype == LOCFNAME) {
        int fd;
        struct stat st;
        fd = open(pb1->loc_fname, O_RDONLY);
        fstat(fd, &st);
        size = st.st_size;
        psource = source = (char *) malloc(size);
        read(fd, source, size);
        close(fd);
    }
    pushint(size);
    return 1;
}
```

Genero program (main.4gl):

```
IMPORT libext1
MAIN
    DEFINE t TEXT
    LOCATE t IN MEMORY
    LET t = "aaaaaaaaaaaaaa"
    DISPLAY lob_size(t)
END MAIN
```

Commands to compile and execute (on
Linux):

```
$ gcc -fPIC -c ext1.c -I $FGLDIR/include
$ gcc --shared -o libext1.so ext1.o -L$FGLDIR/lib -lfgl
$ fglcomp main.4gl
$ fglrun main.42m
         14
```

## Related links

**Related concepts**  

[Header files for ESQL/C typedefs](2705-header-files-for-esql-c-typedefs.md "C header files (.h) are required to define C structures for complex data types used in a C-Extension.")
