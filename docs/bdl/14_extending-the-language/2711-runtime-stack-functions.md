---
title: "Runtime stack functions"
source: "fgl-topics/c_fgl_CExtensions_009.html"
breadcrumb: "Extending the language > C-Extensions > Runtime stack functions"
type: "concept"
---

# Runtime stack functions

> To pass values between a C function and a program, the C function and the runtime system use the runtime stack.

## Stack function basics

The parameters passed to the C function must be popped from the stack at the
beginning of the C function, and the return values expected by the Genero BDL
call must be pushed on the stack before leaving the C function.

The `int` parameter of the C function defines the number of input
parameters passed on the stack, and the function must return an `int`
value defining the number of values returned on the stack.

If you don't pop / push the specified number of parameters / return values, you corrupt the stack
and get a fatal error.

## Pop parameters from the stack

The runtime system library includes a set of functions to retrieve
the values passed as parameters on the stack. This table shows the
library functions provided to pop values from the stack into C buffers:

| Function | Data type | Details |
| --- | --- | --- |
| `void popdate(int4 *dst);` | `DATE` | 4-byte integer value corresponding to days since 12/31/1899. |
| `void popbigint(bigint *dst);` | `BIGINT` | 8-byte integer value. |
| `void popint(mint *dst);` | `INTEGER` | System dependent integer value (int) |
| `void popshort(int2 *dst);` | `SMALLINT` | 2-byte integer value |
| `void poplong(int4 *dst);` | `INTEGER` | 4-byte integer value |
| `void popflo(float *dst);` | `SMALLFLOAT` | 4-byte floating point value |
| `void popdub(double *dst);` | `FLOAT` | 8-byte floating point value |
| `void popdec(dec_t *dst);` | `DECIMAL` | See structure definition in $FGLDIR/include/f2c headers |
| `void popquote(char *dst, int size);` | `CHAR(n)` | The `size` parameter defines the size of the char buffer (with the '\0').The trailing blanks are kept. |
| `void popvchar(char *dst, int size);` | `VARCHAR(n)` | The `size` parameter defines the size of the char buffer (with the '\0').The trailing blanks are kept. |
| `void popstring(char *dst, int size);` | `VARCHAR(n)` | The `size` parameter defines the size of the char buffer (with the '\0').This function trims all the trailing spaces, even the last one. There is no way to distinguish from NULL if the string has only spaces. |
| `void popdtime(dtime_t *dst, int size);` | `DATETIME` | See structure definition in $FGLDIR/include/f2c headerssize = TU_DTENCODE(start, end) |
| `void popinv(intrvl_t *dst, int size);` | `INTERVAL` | See structure definition in $FGLDIR/include/f2c headerssize = TU_IENCODE(len, start, end) |
| `void poplocator(ifx_loc_t **dst);` | `BYTE`, `TEXT` | See structure definition in $FGLDIR/include/f2c headers.Nota that this function pops the pointer of a `ifx_loc_t` object. |

When using a pop function, the value is copied from the stack to
the local C variable and the value is removed from the stack.

In a Genero program, strings (`CHAR`, `VARCHAR`) are not terminated
by `'\0'`. Therefore, the C variable must have one additional character to store the
`'\0'`. For example, the equivalent of a `VARCHAR(100)` in Genero BDL
programs is a `char x[101]` in C.

## Stack introspection

A set of C API functions are provided to query information on the parameters passed
on the stack to a C function. Query for the parameter type and the actual size of a character string
value, to adapt the buffer receiving the parameter.

| Function | Description |
| --- | --- |
| `const char *fglcapi_peekStackType(void)` | Returns the type name of the topmost value on the stack as a string.For example, if the value on the stack is a `CHAR(100)`, the function returns the string `"CHAR(100)"`.If the current value on the stack is a string literal, the returned type name is `"STRING"`. |
| `int fglcapi_peekStackBufferSize(void)` | Returns the size (in bytes) for the topmost value on the stack, to allocate a C `char` buffer, when using pop\* function to get character strings.String pop functions such as `popquote()` and `popvchar()` require a C `char` buffer to be allocated. To allocate the buffer dynamically, use the `fglcapi_peekStackBufferSize()` function to get the required buffer size.Allocating `char` buffers with the proposed size avoids truncating values returned from the stack.When the argument type is different from `CHAR`, `VARCHAR` and `STRING`, the size is the total number of bytes needed for the string representation of the maximal possible value. For example, with a `DECIMAL(32,5)`, the size would be 35 bytes (including sign, decimal separator and string terminator). |

Stack introspection
example:

```
int my_function(int n)
{
    int sz;
    char *buf;
    sz = fglcapi_peekStackBufferSize();
    buf = malloc(sz);
    popstring(buf, sz);
    // ...
    free(buf);
    return 0;
}
```

## Push returns on the stack

To return a value from the C function, you must use one of the functions provided in the runtime
system library.

| Function | Data type | Details |
| --- | --- | --- |
| `void pushdate(int4 val);` | `DATE` | 4-byte integer value corresponding to days since 12/31/1899. |
| `void pushbigint(bigint val);` | `BIGINT` | 8-byte integer value. |
| `void pushdec(const dec_t *val, const unsigned decp);` | `DECIMAL` | See structure definition in $FGLDIR/include/f2c headers |
| `void pushint(mint val);` | `INTEGER` | System dependent integer value (int) |
| `void pushlong(int4 val);` | `INTEGER` | 4-byte integer value |
| `void pushshort(int2 val);` | `SMALLINT` | 2-byte integer value |
| `void pushflo(float *val);` | `SMALLFLOAT` | 4-byte floating point value.Note that this function takes a `(float*)` pointer as argument. |
| `void pushdub(double *val);` | `FLOAT` | 8-byte floating point value.Note that this function takes a `(double*)` pointer as argument. |
| `void pushquote(const char *val, int len);` | `CHAR(n)` | len = strlen(val) (without '\0') |
| `void pushvchar(const char *val, int len);` | `VARCHAR(n)` | len = strlen(val) (without '\0') |
| `void pushdtime(const dtime_t *val);` | `DATETIME` | See structure definition in $FGLDIR/include/f2c headers |
| `void pushinv(const intrvl_t *val);` | `INTERVAL` | See structure definition in $FGLDIR/include/f2c headers |

When using a push function, the value of the C variable is copied at the top of the stack;
therefore the scope and lifespan of the C variable does not matter.

To simplify migration of IBM® I4GL legacy
C extensions using `ret*()` style functions, Genero supports the following
synonyms:

| Function | Equivalent |
| --- | --- |
| `void retdate(int4 val)` | `pushdate` |
| `void retdec(const dec_t *val)` | `pushdec` |
| `void retmoney(const dec_t *val)` | `pushdec` |
| `void retint(int val)` | `pushint` |
| `void retlong(int4 val)` | `pushlong` |
| `void retshort(int2 val)` | `pushshort` |
| `void retflo(float *val)` | `pushflo` |
| `void retdub(double *val)` | `pushdub` |
| `void retquote(const char *val)` | `pushquote` |
| `void retstring(const char *val)` | `pushquote` |
| `void retvchar(const char *val)` | `pushvchar` |
| `void retdtime(const dtime_t *val)` | `pushdtime` |
| `void retinv(const intrvl_t *val)` | `pushinv` |

Pay attention to the `retdec()`, `retmoney()`,
`retquote()` and `retvchar()` functions. These do not have the same
signature as the equivalent `push*()` functions.

## Related links

**Related concepts**  

[Runtime stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions.")

[Header files for ESQL/C typedefs](2705-header-files-for-esql-c-typedefs.md "C header files (.h) are required to define C structures for complex data types used in a C-Extension.")
