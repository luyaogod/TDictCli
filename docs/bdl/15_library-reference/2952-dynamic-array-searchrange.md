---
title: "DYNAMIC ARRAY.searchRange"
source: "fgl-topics/c_fgl_Arrays_ARRAY_searchRange.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.searchRange"
type: "concept"
---

# DYNAMIC ARRAY.searchRange

> Scans the array to find an element that matches the search parameter.

## Syntax

```
searchRange( key STRING, value value-type,
                 from INTEGER, to INTEGER )
   RETURNS INTEGER
```

1. key is the name of the `RECORD` member, when the array is
   structured.
2. value is the value to look for.
3. value-type is the type of the searched value. It can be a [primitive-type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") like `STRING`, or a reference
   to an object such as a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.").
4. from is the start index (inclusive).
5. to is the end index (inclusive).

## Usage

The `searchRange()` method is similar to the [search()](2951-dynamic-array-search.md "Scans the array to find an element that matches the search parameter.") method, except that array rows are scanned
from a given index to an ending index.

The method returns the index of the first occurrence found in the specified range.

The method returns zero, if no matching element is found in the specified range.

If the start index is lower than 1, the search starts at index 1.

If the end index is lower than the start index, the method returns zero.

If the end index is greater than the array length, it is ignored and the search is done until the
last element of the array.

## Example

```
MAIN
    DEFINE a DYNAMIC ARRAY OF RECORD
                 name STRING
             END RECORD

    LET a[1].name = "Mike"
    LET a[2].name = "Phil"
    LET a[3].name = "John"
    LET a[4].name = "Phil"

    DISPLAY a.searchRange("name", "John", 1,  2) -- Shows 0
    DISPLAY a.searchRange("name", "Phil", 1,  2) -- Shows 2
    DISPLAY a.searchRange("name", "Phil", 3,  4) -- Shows 4
    DISPLAY a.searchRange("name", "Phil", 0, 10) -- Shows 2

END MAIN
```

## Related links

**Related concepts**  

[DYNAMIC ARRAY.search](2951-dynamic-array-search.md "Scans the array to find an element that matches the search parameter.")
