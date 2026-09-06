---
title: "Dynamic arrays"
source: "fgl-topics/c_fgl_Arrays_010.html"
breadcrumb: "Language basics > Arrays > Dynamic arrays"
type: "concept"
description: "Defining dynamic arrays Dynamic arrays are defined with the DYNAMIC ARRAY syntax and specify an array with a variable size. Dynamic arrays have no theoretical size limit. The elements of dynamic ..."
---

# Dynamic arrays

## Defining dynamic arrays

Dynamic arrays are defined with the `DYNAMIC ARRAY` syntax and specify an
array with a variable size. Dynamic arrays have no theoretical size limit. The elements of
dynamic arrays are allocated automatically by the runtime system, based on the indexes
used.

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF INTEGER
  LET arr[5000] = 12456  -- Automatic allocation for element 5000
END MAIN
```

## Dynamic array starting index

The first element in a dynamic array is at index position
1:

```
DISPLAY arr[1].name
```

## Default values when adding new elements

A dynamic array is by default empty (its length is zero). When creating a new element with the
`appendElement()` or `insertElement()` methods, or when using an index
for a non-existing element, the element values are initialized to `NULL`, no matter
the data type:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF RECORD
               f_integer INTEGER,
               f_varchar VARCHAR(100)
           END RECORD
    DISPLAY "arr length: ", arr.getLength() -- shows 0
    CALL arr.appendElement()
    DISPLAY "arr length: ", arr.getLength() -- shows 1
    DISPLAY arr[1].f_integer IS NULL  -- shows 1 (TRUE)
    DISPLAY arr[1].f_varchar IS NULL  -- shows 1 (TRUE)
END MAIN
```

## Initializing arrays

Arrays can be initialized in their definitions with [variable
initializers](0691-variable-initializers.md "Variables can be initialized in their definition."):

```
DEFINE num DYNAMIC ARRAY OF STRING = [ "One", "Two", "Three" ]
```

## Sorting arrays

The rows of an array can be sorted by comparing a specific field of a structured array, or by
comparing the values in a simple flat array.

The basic [`sort()`](../15_library-reference/2953-dynamic-array-sort.md "Sorts the rows in the array.") sorting
method compares elements with the implicit ordering rules of the primitive data
types:

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF RECORD
               key INTEGER,
               name VARCHAR(30)
       END RECORD
  -- (fill array with data)
  CALL arr.sort("name",FALSE)
END MAIN
```

You can also implement your own comparison function and sort rows with the
[`sortByComparisonFunction()`](../15_library-reference/2954-dynamic-array-sortbycomparisonfunction.md "Sorts the rows in the array by using a comparison function.") method.

## Element types

The elements of a dynamic array variable are typically defined as a structured
record:

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF RECORD
               key INTEGER,
               name VARCHAR(30),
               address VARCHAR(200),
               contacts DYNAMIC ARRAY OF VARCHAR(20)
       END RECORD
  LET arr[1].key = 12456
  LET arr[1].name = "Scott"
  LET arr[1].contacts[1] = "Bryan COX"
  LET arr[1].contacts[2] = "Mike FLOWER"
END MAIN
```

Dynamic arrays can also be combined with another collection type, such as a dictionary,
or another dynamic array:

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF DICTIONARY OF DECIMAL(10,2)
    LET arr[1]["march"] = 456.25
    LET arr[1]["april"] = 188.30
    LET arr[3]["may"]   = 206.99
END MAIN
```

Dynamic arrays can also be defined with built-in classes such as [`base.Channel`](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions."):

```
MAIN
    DEFINE arr DYNAMIC ARRAY OF base.Channel
    LET arr[1] = base.Channel.create()
    CALL arr[1].openFile("myfile.txt","r")
END MAIN
```

Imported Java classes can also be used to define a dynamic
array:

```
IMPORT JAVA java.lang.String
MAIN
    DEFINE arr DYNAMIC ARRAY OF java.lang.String
    LET arr[1] = "hello"
    DISPLAY arr[1]
END MAIN
```

## Automatic element allocation

When a dynamic array element does not exist, it is automatically allocated before it is used.
For example, when you assign an array element with the `LET` instruction by
specifying an array index greater than the current length of the array, the new element is
created automatically before assigning the value. This is also true when using a dynamic
array in a `FOREACH` loop or when dynamic array elements are used as r-values,
for example in a `DISPLAY`.

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF INTEGER
  LET arr[50] = 33 -- Extends array size to 50 and assigns 33 to element #50
  DISPLAY arr[100] -- Extends array size to 100 and displays NULL
END MAIN
```

> **Important:**
>
> Pay attention to automatic element allocation in dynamic arrays. The
> following code example creates an additional element because at each iteration, the runtime system
> must allocate a new element to fetch the row from the database. As result, you need to remove the
> last element of the array after the `FOREACH`
> loop:
>
> ```
> DEFINE arr DYNAMIC ARRAY OF RECORD
>                key INTEGER,
>                name VARCHAR(30)
>        END RECORD,
>        x INTEGER
> DECLARE c1 CURSOR FOR SELECT ckey, cname FROM mytable
> LET x=1
> FOREACH c1 INTO arr[x].*
>    LET x=x+1
> END FOREACH
> CALL arr.deleteElement(x)
>
> -- A more elegant way to fetch rows into an array:
> TYPE my_type RECORD LIKE mytable.*
> DEFINE arr DYNAMIC ARRAY OF my_type,
>        rec my_type,
>        x INTEGER
> DECLARE c1 CURSOR FOR SELECT * FROM mytable
> LET x=1
> FOREACH c1 INTO rec.*
>    LET arr[x:=x+1] = rec
> END FOREACH
> ```

## Passing and returning dynamic arrays to functions

Dynamic arrays are passed (or returned) by reference to/from functions.

The dynamic array can be modified inside the called function, and the caller will see the
modifications:

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF INTEGER
  CALL fill(arr)
  DISPLAY arr.getLength() -- shows 2
END MAIN

FUNCTION fill(p_arr DYNAMIC ARRAY OF INTEGER) RETURNS ()
  CALL p_arr.appendElement()
  CALL p_arr.appendElement()
END FUNCTION
```

See also [Passing dynamic arrays as parameter](../09_advanced-features/0837-passing-dynamic-arrays-as-parameter.md).

The reference to dynamic arrays created inside a function can be returned from the
function:

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF INTEGER
  LET arr = create_array()
  DISPLAY arr.getLength() -- shows 2
END MAIN

FUNCTION create_array() RETURNS DYNAMIC ARRAY OF INTEGER
  DEFINE arr DYNAMIC ARRAY OF INTEGER
  CALL arr.appendElement()
  CALL arr.appendElement()
  RETURN arr
END FUNCTION
```

See also [Returning dynamic arrays from functions](../09_advanced-features/0843-returning-dynamic-arrays-from-functions.md).

## Using multi-dimensional dynamic arrays

Multi-dimensional dynamic arrays can be defined by using the `WITH DIMENSION`
syntax.

Array methods can be used on multi-dimensional arrays with the brackets
notation:

```
MAIN
  DEFINE a2 DYNAMIC ARRAY WITH DIMENSION 2 OF INTEGER
  DEFINE a3 DYNAMIC ARRAY WITH DIMENSION 3 OF INTEGER
  LET a2[50,100]  = 12456
  LET a2[51,1000] = 12456
  DISPLAY a2.getLength()         -- shows 51
  DISPLAY a2[50].getLength()     -- shows 100
  DISPLAY a2[51].getLength()     -- shows 1000
  LET a3[50,100,100]  = 12456
  LET a3[51,101,1000] = 12456
  DISPLAY a3.getLength()         -- shows 51
  DISPLAY a3[50].getLength()     -- shows 100
  DISPLAY a3[51].getLength()     -- shows 101
  DISPLAY a3[50,100].getLength() -- shows 100
  DISPLAY a3[51,101].getLength() -- shows 1000
  CALL a3[50].insertElement(10)  -- inserts at 50,10
  CALL a3[50,10].insertElement(1)-- inserts at 50,10,1
END MAIN
```

## Related links

**Related concepts**  

[Array methods](0737-array-methods.md "Native BDL arrays and Java arrays can be used to invoke built-in methods.")
