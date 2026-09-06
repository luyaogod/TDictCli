---
title: "DYNAMIC ARRAY.search"
source: "fgl-topics/c_fgl_Arrays_ARRAY_search.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.search"
type: "concept"
---

# DYNAMIC ARRAY.search

> Scans the array to find an element that matches the search parameter.

## Syntax

```
search( key STRING, value value-type )
   RETURNS INTEGER
```

1. key is the name of the `RECORD` member, when the array is
   structured.
2. value is the value to look for.
3. value-type is the type of the searched value. It can be a [primitive-type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") like `STRING`, or a reference
   to an object such as a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.").

## Usage

The `search()` method will scan the whole array, to find a specific element that
matches the value passed as second parameter.

The method returns the index of the first occurrence found.

The method returns zero, if no matching element is found.

Elements set to `NULL` cannot be found: If the search value passed as parameter is
`NULL`, the method returns zero.

The next code example shows a user-written function which is equivalent to the built-in array
`search()` method:

```
TYPE t_cust DYNAMIC ARRAY OF RECORD
         cust_id INTEGER,
         cust_name VARCHAR(50)
     END RECORD

FUNCTION search_customer_by_name(
    custlist t_cust,
    a_name VARCHAR(50)
) RETURNS INTEGER
    DEFINE x INT
    FOR x=1 TO custlist.getLength()
        IF custlist[x].cust_name = a_name THEN
            RETURN x
        END IF
    END FOR
    RETURN 0
END FUNCTION

FUNCTION main()
    DEFINE arr t_cust
    LET arr[1].cust_id = 9999
    LET arr[1].cust_name = "Scott"
    DISPLAY search_customer_by_name(arr, "Scott")
    DISPLAY arr.search("cust_name","Scott")
END FUNCTION
```

If the array is structured (`DYNAMIC ARRAY OF RECORD`), the method returns the
index of the first occurrence of an element where the record member specified by
key is equal to the value parameter.

If the array is a flat array, such as `DYNAMIC ARRAY OF INTEGER`, the method
returns the index of the first occurrence of the element that matches the value
parameter. The key parameter is ignored.

## Example

```
MAIN
    DEFINE a DYNAMIC ARRAY OF RECORD
                 name STRING
             END RECORD

    LET a[1].name = "Mike"
    LET a[2].name = "Phil"
    LET a[3].name = "John"

    DISPLAY a.search("name", "Marc")  -- Shows 0
    DISPLAY a.search("name", "John")  -- Shows 3

END MAIN
```

## Related links

**Related concepts**  

[DYNAMIC ARRAY.searchRange](2952-dynamic-array-searchrange.md "Scans the array to find an element that matches the search parameter.")
