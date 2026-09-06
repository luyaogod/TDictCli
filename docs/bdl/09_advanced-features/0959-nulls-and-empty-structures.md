---
title: "NULLs and empty structures"
source: "fgl-topics/c_fgl_json_utils_nulls.html"
breadcrumb: "Advanced features > JSON support > NULLs and empty structures"
type: "concept"
---

# NULLs and empty structures

> Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.

## JSON notation for NULL, empty and undefined

In JSON notation, a `NULL` element is defined with the `null`
keyword. In the next example, the `"name"` element is
null:

```
{ "key":345, "name":null }
```

Empty JSON objects are represented with an opening followed by a closing curly
brace:

```
{ "key":345, "address":{} }
```

Empty JSON arrays are represented with an opening followed by a closing square
brace:

```
{ "key":345, "list":[] }
```

Undefined elements do not appear in the JSON string representation. In the next example, the
corresponding BDL record could have a "key" and "name" member. In the JSON notation, the "name"
element is just omitted:

```
{ "key":345 }
```

## States of Genero BDL variables

Genero BDL variables defined with a primitive type such as `INTEGER` can have a
value or can be `NULL`, but they cannot be undefined as in JSON.

`RECORD` variables cannot be `NULL`. However, all elements of a
`RECORD` can be `NULL` and thus the record can be considered as empty
or null.

A `DYNAMIC ARRAY` or `DICTIONARY` containing zero elements can be
considered as empty or null.

## Controlling JSON serialization from BDL

By default, the [`util.JSON.stringify()`](../15_library-reference/3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays.") method writes all elements of the BDL variable.

Sub-records and dynamic arrays are written to the JSON output:

```
IMPORT util
MAIN
    DEFINE rec RECORD
           field1 INTEGER,
           subarr DYNAMIC ARRAY OF INTEGER
       END RECORD
    INITIALIZE rec TO NULL
    LET rec.field1 = 999
    DISPLAY util.JSON.stringify(rec)
END MAIN
```

Produces following output:

```
{"field1":999,"subarr":[]}
```

If you want to omit all empty `RECORD` or `DYNAMIC ARRAY` elements
in the JSON string, use the [`util.JSON.stringifyOmitNulls()`](../15_library-reference/3587-util-json-stringifyomitnulls.md "Produces a JSON formatted string from the input, by excluding empty records and empty arrays.") method:

```
IMPORT util
MAIN
    DEFINE rec RECORD
           field1 INTEGER,
           subarr DYNAMIC ARRAY OF INTEGER
       END RECORD
    INITIALIZE rec TO NULL
    LET rec.field1 = 999
    DISPLAY util.JSON.stringifyOmitNulls(rec)
END MAIN
```

Produces following output (the `subarr` element is omitted because it is
empty):

```
{"field1":999}
```

If you want to have a fine-grained control on the JSON serialization of null values and empty
records and dynamic arrays, use the `util.JSON.stringify()` method, in conjunction
with the `json_null` attribute when defining the BDL variable. The
`json_null` variable definition attribute can be set to `"null"` or
`"undefined"`.

## Using the json\_null="null" attribute

When defining a BDL variable with the `json_null="null"` attribute, the following
BDL elements will be represented with the `null` keyword in the resulting JSON
string:

1. Simple primitive typed variables that are `NULL`,
2. `RECORD` variables where all members are `NULL`,
3. Empty `DYNAMIC ARRAY` that contain no elements.
4. Empty `DICTIONARY` that contain no elements.

In the next example, the sub-record and sub-array are written to the JSON output as
nulls:

```
IMPORT util
MAIN
    DEFINE rec RECORD
           field1 INTEGER ATTRIBUTES(json_null="null"),
           field2 CHAR(1) ATTRIBUTES(json_null="null"),
           subrec1 RECORD ATTRIBUTES(json_null="null")
               field11 INTEGER,
               fiedl12 VARCHAR(30)
           END RECORD,
           subarr1 DYNAMIC ARRAY ATTRIBUTES(json_null="null") OF INTEGER,
           subdic1 DICTIONARY ATTRIBUTES(json_null="null") OF INTEGER
       END RECORD
    INITIALIZE rec TO NULL
    DISPLAY util.JSON.format(util.JSON.stringify(rec))
END MAIN
```

Produces following output:

```
{
    "field1": null,
    "field2": null,
    "subrec1": null,
    "subarr1": null,
    "subdic1": null
}
```

## Using the json\_null="undefined" attribute

When defining a BDL variable with the `json_null="undefined"` attribute, the
following BDL elements will be omitted in the resulting JSON string (with the exception of root
elements):

1. Simple primitive typed variables that are `NULL` (in fact,
   `json_null="undefined"` is the default for primitive types),
2. `RECORD` variables where all members are `NULL`,
3. Empty `DYNAMIC ARRAY` that contains no elements.
   > **Note:**
   >
   > If some elements of a dynamic array are `NULL`, the JSON serialization class must
   > produce a `null` keyword for these element, even when using
   > `json_null="undefined"`. This is mandatory because JSON requires
   > `null` keyword for undefined array elements. Otherwise, when omitted, the positions
   > of array elements would be shifted.
4. Empty `DICTIONARY` that contain no elements.

In the next example, some elements of the record are omitted in the resulting JSON
output:

```
IMPORT util
MAIN
    DEFINE rec RECORD
       field1 INTEGER ATTRIBUTES(json_null="undefined"),
       field2 INTEGER, -- default is json_null="undefined" for primitives
       subrec1 RECORD ATTRIBUTES(json_null="undefined")
           field11 INTEGER,
           fiedl12 VARCHAR(30)
       END RECORD,
       subrec2 RECORD -- will be serialized as {}
           field21 INTEGER,
           fiedl22 VARCHAR(30)
       END RECORD,
       subarr1 DYNAMIC ARRAY ATTRIBUTES(json_null="undefined") OF INTEGER,
       subarr2 DYNAMIC ARRAY OF INTEGER, -- will be serialized as []
       subarr3 DYNAMIC ARRAY ATTRIBUTES(json_null="undefined") OF INTEGER,
       subdic1 DICTIONARY ATTRIBUTES(json_null="undefined") OF INTEGER
    END RECORD
    INITIALIZE rec TO NULL
    LET rec.field1 = 999
    LET rec.subarr3[1] = NULL
    LET rec.subarr3[2] = 888
    DISPLAY util.JSON.format(util.JSON.stringify(rec))
END MAIN
```

Produces following output:

```
{
    "field1": 999,
    "subrec2": {},
    "subarr2": [],
    "subarr3": [
        null,
        888
    ]
}
```

Variables defined with a type using `json_null="undefined"` attribute at the root
element, will be serialized as JSON `null`, even if all members of the variable are
`NULL`, because the root element cannot be omitted:

```
IMPORT util
MAIN

    VAR i INTEGER -- json_null="undefined" is the default for primitive types
    LET i = NULL
    DISPLAY "i:   ", util.JSON.stringify(i)

    VAR rec RECORD ATTRIBUTES(json_null="undefined")
           field1 INTEGER
       END RECORD
    INITIALIZE rec.* TO NULL
    DISPLAY "rec: ", util.JSON.stringify(rec)

    VAR arr DYNAMIC ARRAY ATTRIBUTES(json_null="undefined") OF INTEGER
    DISPLAY "arr: ", util.JSON.stringify(arr)

    VAR sar ARRAY[10] ATTRIBUTES(json_null="undefined") OF INTEGER
    INITIALIZE sar TO NULL
    DISPLAY "sar: ", util.JSON.stringify(sar)

    VAR dic DICTIONARY ATTRIBUTES(json_null="undefined") OF INTEGER
    DISPLAY "dic: ", util.JSON.stringify(dic)

END MAIN
```

Output:

```
i:   null
rec: null
arr: null
sar: null
dic: null
```
