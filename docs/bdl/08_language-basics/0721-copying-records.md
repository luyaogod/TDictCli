---
title: "Copying records"
source: "fgl-topics/c_fgl_records_009.html"
breadcrumb: "Language basics > Records > Copying records"
type: "concept"
---

# Copying records

> Records can be assigned to each other with the = operator.

It is possible to copy all members of a record to the members of another record, as long as the
types match exactly, by using the assignment operator:

```
LET rec2 = rec1
```

> **Important:**
>
> Record members of type `DYNAMIC ARRAY`,
> `DICTIONARY` and the LOB data of `TEXT/BYTE` members are not cloned
> when assigning records. A full copy of such members must be done by hand, otherwise both records
> will reference the same dynamic array, dictionary or LOB data.

When assigning records, both variables must have been defined with the same
`RECORD` specification, or with the same [user-type](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."). If the types do not match exactly, the compiler will produce the error [-4325](../15_library-reference/4483-genero-bdl-errors.md):

```
DEFINE rec1 RECORD
            id INTEGER,
            name VARCHAR(50)
       END RECORD

DEFINE rec2 RECORD
            id INTEGER,
            name VARCHAR(100)  -- Note the difference with rec1
       END RECORD

MAIN
    LET rec2 = rec1
| The source and destination records in this record assignment statement
| are not compatible in types and/or length.
| See error number -4325.
END MAIN
```

The `.*` (dot star) notation `LET rec2.* = rec1.*` is supported for
backward compatibility and is equivalent to `LET rec2 = rec1`. When copying records,
avoid the `.*` after the record names: The result at runtime is the same, but your
code will benefit from more compiler verifications.

If the record structure contains sub-records on several levels, the sub-members are copied
recursively:

```
TYPE t_person RECORD
            ...
            address RECORD
                ...
            END RECORD
       END RECORD
...
    DEFINE r1, r2 t_person
...
    LET r1 = r2   -- address sub-record is copied as well
...
```

If the record structure contains [dynamic array](0734-dynamic-arrays.md) or [dictionary](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") members, the runtime system will copy the reference
of such members. Similarly, if the record uses [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.")/[`TEXT`](0569-text.md "The TEXT data type stores large text data.")
members, the (locator) handler of the large object will be copied, but the actual LOB data is not
cloned. As result, after the copy, both records will point to the same array, dictionary or
`TEXT/BYTE` data.

The next code example implements a `copyTo()` method for the
`t_reader` record type, which makes a proper clone of the `book_ids`
dynamic array:

```
IMPORT util

TYPE t_reader RECORD
            id INTEGER,
            name VARCHAR(100),
            book_ids DYNAMIC ARRAY OF INTEGER
       END RECORD

MAIN
    DEFINE r1, r2 t_reader

    LET r1.id = 999
    LET r1.name = "Mike Rutberg"
    LET r1.book_ids[1] = 98458
    LET r1.book_ids[2] = 98111
    DISPLAY "1: r1 books: ", util.JSON.stringify(r1.book_ids)

    CALL r1.copyTo( r2 )
    DISPLAY "2: r2 books: ", util.JSON.stringify(r2.book_ids)

    LET r1.book_ids[3] = 18234
    DISPLAY "3: r1 books: ", util.JSON.stringify(r1.book_ids)

    DISPLAY "4: r2 books: ", util.JSON.stringify(r2.book_ids)

END MAIN

FUNCTION (r t_reader) copyTo(dst t_reader INOUT)
    DEFINE bl DYNAMIC ARRAY OF INTEGER
    LET dst = r -- makes a copy of the reference of book_ids
    CALL r.book_ids.copyTo(bl) -- clone the array
    LET dst.book_ids = bl -- copy the reference to the new array
    -- Don't do this:
    -- CALL r.book_ids.copyTo(dst.book_ids)
    -- dst.book_ids and r.book_ids reference the same array!
END FUNCTION
```

Output:

```
1: r1 books: [98458,98111]
2: r2 books: [98458,98111]
3: r1 books: [98458,98111,18234]
4: r2 books: [98458,98111]
```

Unlike dynamic array members, [static array](0732-static-arrays.md "Static arrays have a predefined and limited size.") members are
fully copied, as when assigning static arrays directly.

For more details about copying `TEXT/BYTE` variables, see [the `TEXT` data type](0569-text.md "The TEXT data type stores large text data.").

## Related links

**Related concepts**  

[Equal to (== or =)](0609-equal-to-or.md "The == operator checks for equality of two expressions or for two record variables. A single = can be used as alias for ==.")

[Different from (!= or <>)](0610-different-from-or.md "The != operator checks for non-equality of two expressions or for two record variables. The <> can be used as alias for !=.")
