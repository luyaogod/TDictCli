---
title: "DICTIONARY.copyTo"
source: "fgl-topics/c_fgl_DICTIONARY_method_copyTo.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DICTIONARY as class > DICTIONARY methods > DICTIONARY.copyTo"
type: "concept"
---

# DICTIONARY.copyTo

> Copies all elements of the dictionary into another dictionary.

## Syntax

```
copyTo( dst dictionary-type )
```

1. dst is the destination dictionary, with the same type definition as the
   source dictionary.
2. dictionary-type is a `DICTIONARY OF ...` type that corresponds
   to the source dictionary.

## Usage

The `copyTo()` method clones the complete dictionary into the destination
dictionary passed as parameter.

The destination dictionary will be cleared before the copy operation starts.

The source and destination dictionary must be defined with equivalent data structures
(`RECORD` field names and types must match). If the type has methods, these are
ignored. If the type structures do not match, error [-8112](4483-genero-bdl-errors.md) is thrown.

Avoid making a copy of huge dictionaries, consider using a database temporary table instead.

The elements and sub-elements of the dictionary are copied as follows:

- Basic data types such as `INTEGER` and `VARCHAR` are copied by
  value.
- For `TEXT/BYTE` fields, only the locator is copied by value: The LOB data is not
  cloned. See [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") reference page.
- Static arrays are cloned (copied by value)
- Reference types such as `base.Channel`, `DYNAMIC ARRAY` and
  `DICTIONARY` are copied by reference.

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
    DEFINE d1 DICTIONARY OF t1
    DEFINE d2 DICTIONARY OF t1

    CALL initFiles()

    CALL init(d1["key1"], 1, "old", "file1.txt")
    CALL d1.copyTo(d2)
    CALL init(d1["key1"], 2, "new", "file2.txt")
    DISPLAY "d1:", util.JSON.stringifyOmitNulls(d1)
    DISPLAY "d2:", util.JSON.stringifyOmitNulls(d2)
    DISPLAY "d1['key1'].ch.readline(): ", d1['key1'].ch.readLine()
    DISPLAY "d2['key1'].ch.readline(): ", d2['key1'].ch.readLine()

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
d1:{"key1":{"i":2,"vc":"new","t":"new","ch":"","sa":[2,null,null],"da":[2],"di":{"key1":"new"}}}
d2:{"key1":{"i":1,"vc":"old","t":"new","ch":"","sa":[1,null,null],"da":[2],"di":{"key1":"new"}}}
d1['key1'].ch.readline(): file2 line1
d2['key1'].ch.readline(): file2 line2
```

The first line `"d1:"` shows the values in the first dictionary, after
`d1.copyTo(d2)`, and after setting new values into `d1`.

In the second line `"d2:"`, showing the content of the dictionary
`d2`:

- The fields `"i"` (`INTEGER`) and `"vc"`
  (`VARCHAR`) still have the initial values.
- The field `"t"` (`TEXT`) locator is copied by value, but points to
  the same LOB data, so changing the value in `a1[1].t` or `a2[1].t`
  affects the same data.
- The field `"sa"` (static `ARRAY[]`) is copied by value: The
  elements `[1,null,null]` are the initial values.
- The field `"da"` (`DYNAMIC ARRAY`) is copied reference: The
  elements `[2]` are the new values.
- The field `"di"` (`DICTIONARY`) is copied by reference: The
  elements `{"key1":"new"}` are the new values.
- The field `"ch"` (`base.Channel`) is copied by reference and
  points to the same channel object: `d2[1].ch.realLine()` returns a line from
  file2.txt.
