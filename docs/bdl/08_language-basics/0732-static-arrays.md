---
title: "Static arrays"
source: "fgl-topics/c_fgl_Arrays_006.html"
breadcrumb: "Language basics > Arrays > Static arrays"
type: "concept"
---

# Static arrays

> Static arrays have a predefined and limited size.

## Defining static arrays

Static arrays can store a one-, two- or three-dimensional array of variables, all of the same
type. An array member can be any type except another array (`ARRAY ... OF
ARRAY`).

```
MAIN
  DEFINE custlist ARRAY[100] OF RECORD
             id INTEGER,
             name VARCHAR(50)
         END RECORD
  LET custlist[50].id = 12456
  LET custlist[50].name = "Beerlington"
END MAIN
```

## Static array starting index

The first element in a static array is at index position
1:

```
DISPLAY arr[1].name
```

## Static array size is fixed

Static arrays have a fixed size. Using an index greater than the static array size will produce a
runtime error:

```
MAIN
  DEFINE a1 ARRAY[100] OF INTEGER
  LET a1[50] = 12456
  LET a1[101] = 12456  -- Runtime error -1326
END MAIN
```

## Default values of static array elements

A static array is by default filled with elements up to its maximum size, where each element
field is initialized according to its data type, like when defining a `RECORD`
variable. See [Variable default values](0697-variable-default-values.md "Variables get a default value when defined.") for more details.

```
MAIN
    DEFINE arr ARRAY[10] OF RECORD
               f_integer INTEGER,
               f_varchar VARCHAR(100)
           END RECORD
    DISPLAY "arr length: ", arr.getLength()  -- shows 10
    DISPLAY arr[1].f_integer IS NULL  -- shows 0 (FALSE)
    DISPLAY arr[1].f_varchar IS NULL  -- shows 1 (TRUE)
END MAIN
```

## Multi-dimensional static arrays

The multi-dimensional array syntax (`ARRAY[i[,j[,k]]]`) specifies static arrays defined with an explicit
size for all dimensions. Static arrays have a size limit. The biggest static array size you can
define is 65535.

A single array element can be referenced by specifying its coordinates in each dimension of the
array:

```
MAIN
  DEFINE a1 ARRAY[100,100,100] OF INTEGER
  LET a1[50,25,45] = 12456
END MAIN
```

## Element types

The elements of a static array variable can be defined as a structured
record:

```
MAIN
  DEFINE arr ARRAY[50] OF RECORD
               key INTEGER,
               name CHAR(10),
               address VARCHAR(200),
               contacts ARRAY[50] OF VARCHAR(20)
       END RECORD
  LET arr[1].key = 12456
  LET arr[1].name = "Scott"
  LET arr[1].contacts[1] = "Bryan COX"
  LET arr[1].contacts[2] = "Mike FLOWER"
END MAIN
```

Static arrays can also be combined with another collection type, such as a dictionary, or another
static or dynamic array:

```
MAIN
    DEFINE a ARRAY[12] OF DICTIONARY OF DECIMAL(10,2)
    LET a[1]["march"] = 456.25
    LET a[1]["april"] = 188.30
    LET a[3]["may"]   = 206.99
END MAIN
```

## Passing static arrays to functions

Static arrays are passed by value to functions. This is not recommended, as all array
members will be copied on the stack.

A static array cannot be returned from a function.

Consider using dynamic arrays if you need to pass/return a list of elements to/from
functions.

## Using array methods

Array methods can be used on static arrays; however, these methods are designed for dynamic
arrays and are not appropriate for static arrays.

## Related links

**Related concepts**  

[Array methods](0737-array-methods.md "Native BDL arrays and Java arrays can be used to invoke built-in methods.")

[Controlling out of bound in static arrays](0733-controlling-out-of-bound-in-static-arrays.md "Controlling out of bound in static arrays")

## Child topics

- [Controlling out of bound in static arrays](0733-controlling-out-of-bound-in-static-arrays.md)
