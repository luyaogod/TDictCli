---
title: "DYNAMIC ARRAY.copyTo"
source: "fgl-topics/c_fgl_Arrays_ARRAY_copyTo.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DYNAMIC ARRAY as class > DYNAMIC ARRAY methods > DYNAMIC ARRAY.copyTo"
type: "concept"
---

# DYNAMIC ARRAY.copyTo

> Copies a complete array to the destination array passed as parameter.

## Syntax

```
copyTo( dst dynamic-array-type )
```

1. dst is the destination dynamic array, with the same type definition as the
   source array.
2. dynamic-array-type is a `DYNAMIC ARRAY OF ...` type that
   corresponds to the source array.

## Usage

The `copyTo()` method copies all elements of the source array into the destination
array.

The method truncates the length of the destination array to the length of the source array.

The source and destination array must be defined with equivalent data structures
(`RECORD` field names and types must match). If the type has methods, these are
ignored. If the type structures do not match, error [-8112](4483-genero-bdl-errors.md) is thrown.

Avoid making a copy of huge arrays, consider using a database temporary table instead.

The elements and sub-elements of the array are copied as follows:

- Basic data types such as `INTEGER` and `VARCHAR` are copied by
  value.
- For `TEXT/BYTE` fields, only the locator is copied by value: The LOB data is not
  cloned. See [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") reference page.
- Static arrays are cloned (copied by value)
- Reference types such as `base.Channel` and `DICTIONARY` are copied
  by reference.
- `DYNAMIC ARRAY` typed fields is an exception: These are cloned.

## Example

```
IMPORT util

TYPE t1 RECORD
    i INTEGER,
    vc VARCHAR(200),
    t TEXT,
    ch base.Channel,
    sa ARRAY[3] OF INTEGER,
    da DYNAMIC ARRAY OF INTEGER,
    di DICTIONARY OF STRING
END RECORD

MAIN
    DEFINE a1 DYNAMIC ARRAY OF t1
    DEFINE a2 DYNAMIC ARRAY OF t1

    CALL initFiles()

    CALL init(a1[1], 1, "old", "file1.txt")
    CALL a1.copyTo(a2)
    CALL init(a1[1], 2, "new", "file2.txt")
    DISPLAY "a1:", util.JSON.stringifyOmitNulls(a1)
    DISPLAY "a2:", util.JSON.stringifyOmitNulls(a2)
    DISPLAY "a1[1].ch.readline(): ", a1[1].ch.readLine()
    DISPLAY "a2[1].ch.readline(): ", a2[1].ch.readLine()

END MAIN

FUNCTION init(r t1 INOUT, i INT, s STRING, fn STRING)
    LET r.i = i
    LET r.vc = s
    LOCATE r.t IN MEMORY
    LET r.t = s
    IF (r.ch IS NULL) THEN
        LET r.ch = base.Channel.create()
    END IF
    CALL r.ch.openFile(fn, "r")
    LET r.sa[1] = i
    LET r.da[1] = i
    LET r.di["key1"] = s
END FUNCTION

FUNCTION initFiles()
    DEFINE f base.Channel

    LET f = base.Channel.create()

    CALL f.openFile("file1.txt", "w")
    CALL f.writeLine("file1 line1")
    CALL f.writeLine("file1 line2")
    CALL f.close()

    CALL f.openFile("file2.txt", "w")
    CALL f.writeLine("file2 line1")
    CALL f.writeLine("file2 line2")
    CALL f.close()

END FUNCTION
```

Output:

```
a1:[{"i":2,"vc":"new","t":"new","ch":"","sa":[2,null,null],"da":[2],"di":{"key1":"new"}}]
a2:[{"i":1,"vc":"old","t":"new","ch":"","sa":[1,null,null],"da":[1],"di":{"key1":"new"}}]
a1[1].ch.readline(): file2 line1
a2[1].ch.readline(): file2 line2
```

The first line `"a1:"` shows the values in the first array, after
`a1.copyTo(a2)`, and after setting new values into `a1`.

In the second line `"a2:"`, showing the content of the array `a2`:

- The fields `"i"` (`INTEGER`) and `"vc"`
  (`VARCHAR`) still have the initial values.
- The field `"t"` (`TEXT`) locator is copied by value, but points to
  the same LOB data, so changing the value in `a1[1].t` or `a2[1].t`
  affects the same data.
- The field `"sa"` (static `ARRAY[]`) is copied by value: The
  elements `[1,null,null]` are the initial values.
- The field `"da"` (`DYNAMIC ARRAY`) is copied by value (not by
  reference!): The elements `[1]` are the initial values.
- The field `"di"` (`DICTIONARY`) is copied by reference: The
  elements `{"key1":"new"}` are the new values.
- The field `"ch"` (`base.Channel`) is copied by reference and
  points to the same channel object: `a2[1].ch.realLine()` returns a line from
  file2.txt.
