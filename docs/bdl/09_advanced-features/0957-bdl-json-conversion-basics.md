---
title: "BDL/JSON conversion basics"
source: "fgl-topics/c_fgl_json_utils_conv.html"
breadcrumb: "Advanced features > JSON support > BDL/JSON conversion basics"
type: "concept"
---

# BDL/JSON conversion basics

> When the data structures and member names match, Genero BDL variables can be converted from/to JSON data with the util.JSON* utility classes.

## Matching BDL and JSON data structures

In order to convert a BDL variable to/from a JSON string, the program `RECORD` or
`DYNAMIC ARRAY` and the JSON data string must have the same structure.

JSON object elements and BDL `RECORD` member are associated by name, not by
position. Elements in the JSON string and in the BDL variable can be at a different ordinal
position.

JSON array elements and BDL `DYNAMIC ARRAY` elements are associated by
position.

A JSON object can also be converted to a BDL `DICTIONARY`, when the JSON object is
a list of named elements, using the same structure as the dictionary.

Example of BDL data structure:

```
DEFINE rec RECORD
             pkey INT,
             name VARCHAR(50),
             arr DYNAMIC ARRAY OF STRING,
             dic DICTIONARY OF DECIMAL
         END RECORD

LET rec.pkey = 999
LET rec.name = "Tim Birton"
LET rec.arr[1] = "item1"
LET rec.arr[2] = "item2"
LET rec.dic["abc"] = 14.45
LET rec.dic["def"] = 18.11
```

JSON
equivalent:

```
{
    "pkey": 999,
    "name": "Tim Birton",
    "arr": ["item1","item2"
    ],
    "dic": {
        "def": 18.11,
        "abc": 14.45
    }
}
```

## BDL to JSON conversion

BDL variables can be converted to JSON strings for example with the [`util.JSON.stringify()`](../15_library-reference/3586-util-json-stringify.md "Produces a JSON formatted string from the input, by including empty records and empty arrays.") method.

The JSON elements get the same names of the record members, as defined in the program source. For
more details about BDL to JSON names handling, see [BDL names and JSON element names](0958-bdl-names-and-json-element-names.md "To identify elements, JSON standards use different format as Genero BDL variable names.")

Program array members in the record are converted to JSON arrays delimited by square brackets
(`[]`).

Special consideration needs to be taken regarding empty dynamic arrays records where all elements
are null. The Genero JSON API provides options to control the production of JSON elements for empty
records and array. For more details, see [NULLs and empty structures](0959-nulls-and-empty-structures.md "Unlike Genero BDL, JSON distinguishes NULL, empty and undefined elements.")

JSON string values are double-quoted, and the escape character is backslash. When
creating a JSON string from a single-quoted FGL string, keep in mind to escape the
backslash characters. Another option is to use back-quotes for the FGL string, to
define the character string with the same backslash sequences as in the native JSON
string:

```
-- The expected value is:   C:\Users\Mike\
-- JSON representation  :   C:\\Users\\Mike\\  (where backslashes are escaped)
-- Use single quotes in FGL:
LET rec = '{ path: "C:\\\\Users\\\\Mike\\\\}'
-- or use back-quotes and escape backslashes like in JSON representation:
LET rec = `{ path: "C:\\Users\\Mike\\}`
```

When using a variable defined as `BOOLEAN`, the JSON API knows the type and can
write `true` or `false` values in the JSON result. However, the
predefined constants `TRUE` and `FALSE` has a historical type of
`INTEGER`, and the resulting JSON value will be `1` or
`0`. In order to get true or false JSON values, use a variable instead of using the
TRUE or FALSE constants directly in the API call:

```
IMPORT util
MAIN
   DEFINE jo util.JSONObject
   DEFINE v BOOLEAN
   LET jo = util.JSONObject.create()
   CALL jo.put( "active", v:=TRUE )
   DISPLAY jo.toString()
END MAIN
```

By default, [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") and [`INTERVAL`](../08_language-basics/0563-interval-qual1-to-qual2.md "The INTERVAL data type stores spans of time as Year/Month or Day/Hour/Minute/Second/Fraction units.") values are converted to JSON
strings using the FGL formatting rules. In order to produce RFC 3339 / ISO 8601 formatted date/time
and duration values, use the following configuration methods:

- To produce RFC 3339 formatted date/time values, use the [`util.JSON.setDatetimeSerializationMode()`](../15_library-reference/3584-util-json-setdatetimeserializationmode.md "Defines the JSON formatting mode for DATETIME values.") method, with format mode
  `"RFC3339/UTC"` or `"RFC3339/TZO"`.
- To produce ISO 8601 formatted duration values, use the [`util.JSON.setIntervalSerializationMode()`](../15_library-reference/3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values.") method, with format mode
  `"ISO8601"`.

For details about BDL to JSON data type conversion rules, see [BDL to JSON type conversion rules](0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.").

## JSON to BDL conversion

When conversion from JSON to BDL, elements in the JSON string that do not match an Genero BDL
record member are ignored; no error is thrown if there is no corresponding Genero BDL member.

Genero BDL record members that have no matching JSON element are initialized to
`NULL`.

The JSON value must match the data format of the destination member. If the value does not
correspond to the type (for example, if the JSON value is a character string while the target record
member is defined with a numeric type), the target member will be set to `NULL`.

JSON arrays delimited by square brackets are used to fill a program array of the destination
record. The destination array should be a dynamic array. If the array is defined as static, the
additional elements of the source JSON array will be discarded, while missing elements will be
initialized to `NULL`.

The JSON source string must follow the JSON format specification. It can contain multilevel
structured data. If the source string is not well formatted, the runtime system will throw error
[-8109](../15_library-reference/4483-genero-bdl-errors.md).

For details about JSON to BDL data type conversion rules, see [JSON to BDL type conversion rules](0961-json-to-bdl-type-conversion-rules.md "Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.")

## JSON number limitations

The JSON specification defines numbers as a sequence of digits, with optional sign, dot, and
exponent notation such as `874523` or `-8.346E-5`. There is
theoretically no limitation regarding the precision and range for a number, in pure JSON grammar.
However, the specification suggests that implementations use the IEEE 754-2008 binary64 (double
precision) type to handle JSON numbers. See [JSON
specification](https://www.rfc-editor.org/rfc/rfc7159#section-6) for more details.

When parsing JSON numbers, or when writing JSON numbers with the [`util.JSONObject`](../15_library-reference/3590-the-util-jsonobject-class.md "The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax.") and [`util.JSONArray`](../15_library-reference/3604-the-util-jsonarray-class.md "The util.JSONArray class provides methods to handle an array of values, following the JSON string syntax.") classes, the numbers
must be in the precision and range that a binary64 can provide. For example, the integer
9007199254740997 can be stored in a `BIGINT` variable, but with
`JSONObject` and `JSONArray` classes, this number will approximate to
9007199254740996 as a binary64. However, this binary64 limitation does not apply to the [`util.JSON`](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") class, which can handle any value
of any BDL type, including `BIGINT` and `DECIMAL(P,S)`.
