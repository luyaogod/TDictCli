---
title: "DYNAMIC ARRAY.sortByComparisonFunction"
source: "fgl-topics/c_fgl_Arrays_ARRAY_sortByComparisonFunction.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.sortByComparisonFunction"
type: "concept"
---

# DYNAMIC ARRAY.sortByComparisonFunction

> Sorts the rows in the array by using a comparison function.

## Syntax

```
sortByComparisonFunction(
    key STRING,
    reverse BOOLEAN,
    compar FUNCTION (
                 s1 STRING,
                 s2 STRING
             ) RETURNS INTEGER
 )
```

1. key is the name of a member of a structured array (`DYNAMIC ARRAY OF
   RECORD`), or `NULL` if the array is not structured.
2. reverse is `FALSE` for ascending order, `TRUE`
   for descending order.
3. compar is a [function
   reference](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") of a user-defined sort function.

## Usage

The `sortByComparisonFunction()` method sorts the rows of an array with a
user-defined function to compare element values.

The first parameter defines the record member, when the array is a `DYNAMIC ARRAY OF
RECORD`. It must be `NULL` for non-structured arrays.

The second parameter defines the sort order as ascending (`FALSE`) or descending
(`TRUE`).

The last parameter is a [function
reference](../08_language-basics/0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.") to a user-defined function that compares two values s1 and
s2. The comparison function must return a negative `INTEGER`, if
value s1 is less then s2, a positive `INTEGER`,
if s1 is greater than s2, and return zero when
s1 and s2 are equal. Take care with `INTEGER`
overflows, when substracting large numbers to produce the negative or positive return value.
> **Tip:**
>
> The comparison function can be a reference to the [`util.Strings.collate()`](3557-util-strings-collate.md "Compares two strings using locale collation rules.") or [`util.Strings.collateNumeric()`](3558-util-strings-collatenumeric.md "Compares two strings using locale collation rules and sequences of numerical digits.") functions: `FUNCTION
> util.Strings.collateNumeric`

Pay attention to `NULL` values: When both s1
and s2 are `NULL`, they can be considered as equal and the
comparison function should return zero. If only one of the values is `NULL`, it's up
to the comparison function to consider `NULL` as lower or as greater than any
non-null value. The [`NVL()`](../08_language-basics/0615-nvl-function.md "The NVL() operator returns the second parameter if the first argument evaluates to NULL.") operator is very useful in such case.

## Example

In this example, clothing sizes are expressed as standard codes like XS, M, L. Sorting by these
codes in the alphabetical order would not reflect the actual meaning of clothing sizes. To solve
this problem, a user-defined comparison function is provided that will map a clothing size code to
an integer value, and order the items from the smallest to the largest size:

```
IMPORT util

PRIVATE DEFINE size2int DICTIONARY OF INTEGER =
     ("XXS": -3, "XS": -2, "S": -1, "M": 0, "L": 1, "XL": 2, "XXL": 3)

PUBLIC TYPE t_item RECORD
        item_id INTEGER,
        item_size STRING
    END RECORD

MAIN
    DEFINE items DYNAMIC ARRAY OF t_item =
        [
          (item_id:1024, item_size:"XL"),
          (item_id:1026, item_size:"S"),
          (item_id:1028, item_size:"L")
        ]

    CALL items.sortByComparisonFunction("item_size", FALSE, FUNCTION compareSize)
    DISPLAY util.JSON.format( util.JSON.stringify(items) )

END MAIN

PUBLIC FUNCTION compareSize(s1 STRING, s2 STRING) RETURNS INTEGER
    RETURN NVL(size2int[s1],-100) - NVL(size2Int[s2],-100)
END FUNCTION
```

Output:

```
[
    {
        "item_id": 1026,
        "item_size": "S"
    },
    {
        "item_id": 1028,
        "item_size": "L"
    },
    {
        "item_id": 1024,
        "item_size": "XL"
    }
]
```

## Related links

**Related concepts**  

[util.Strings.collateNumeric](3558-util-strings-collatenumeric.md "Compares two strings using locale collation rules and sequences of numerical digits.")

[DYNAMIC ARRAY.sort](2953-dynamic-array-sort.md "Sorts the rows in the array.")
