---
title: "DYNAMIC ARRAY.sort"
source: "fgl-topics/c_fgl_Arrays_ARRAY_sort.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.sort"
type: "concept"
---

# DYNAMIC ARRAY.sort

> Sorts the rows in the array.

## Syntax

```
sort( key STRING, reverse BOOLEAN )
```

1. key is the name of a member of a structured array (`DYNAMIC ARRAY OF
   RECORD`), or `NULL` if the array is not structured.
2. reverse is `FALSE` for ascending order, `TRUE`
   for descending order.

## Usage

A dynamic array can be sorted with the `sort()` method.

- For non-structured dynamic arrays (`DYNAMIC ARRAY OF
  simple-type`), the first argument of `sort()` must be
  `NULL`. The array will be sorted by the single-typed elements.
- With structured arrays (`DYNAMIC ARRAY OF RECORD`), this method sorts the array
  by the member passed as first parameter.

Use the second parameter to define the sort order as ascending (`FALSE`) or
descending (`TRUE`).

Character string data is sorted by following the collation rules defined by [current application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."): The `sort()`
method uses the operating system collation functions. Note that collation rules can slightly differ
from platform to platform and therefore produce different ordering.

When making subsequent calls to the `sort()` method using different record members
of the array, the rows will be ordered by all of the record members specified for the cumulative
sorts, with the most recent call defining the main sort field.

Another way to think of this is in terms of the `ORDER BY` clause of a SQL
statement: If your dynamic array contained the variables A, B and C, and you included the
following calls to the `sort()` method:

```
CALL a.sort("C",false)
CALL a.sort("B",false)
CALL a.sort("A",false)
```

This would be equivalent to writing an `ORDER BY` clause that states:

```
ORDER BY A, B, C
```

## Example

In this example, the first call to the `sort()` method sorts the rows by name in
ascending order, and the second call will sort the rows by key in descendant order, then by name
within each key. The last sort becomes the main sort field.

```
MAIN
  DEFINE a DYNAMIC ARRAY OF RECORD
               key INTEGER,
               name VARCHAR(30)
         END RECORD
  LET a[1].key = 776236    LET a[1].name = "aaaaa"
  LET a[2].key = 273434    LET a[2].name = "cccccccc"
  LET a[3].key = 934092    LET a[3].name = "bbbbb"
  CALL a.sort("name",FALSE) -- Sorted by name (asc order)
  CALL a.sort("key",TRUE) -- Sorted by key (desc), then by name
END MAIN
```

## Related links

**Related concepts**  

[DYNAMIC ARRAY.sortByComparisonFunction](2954-dynamic-array-sortbycomparisonfunction.md "Sorts the rows in the array by using a comparison function.")
