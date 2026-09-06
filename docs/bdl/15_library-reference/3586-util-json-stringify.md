---
title: "util.JSON.stringify"
source: "fgl-topics/c_fgl_ext_util_JSON_stringify.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > util.JSON methods > util.JSON.stringify"
type: "concept"
---

# util.JSON.stringify

> Produces a JSON formatted string from the input, by including empty records and empty arrays.

## Syntax

```
util.JSON.stringify(
     value any-type
   )
  RETURNS STRING
```

1. value is the program variable to be converted to a JSON string.
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `util.JSON.stringify()` class method takes a variable as parameter, and
generates the corresponding data string in JSON format, as defined in the [[RFC4627]](http://www.ietf.org/rfc/rfc4627.txt) specification.

The input parameter can be a `RECORD`, an `ARRAY`
or a `DICTIONARY`, as well as a simple type such as `INTEGER` or
`STRING`.

> **Important:**
>
> Unlike [`util.JSON.stringifyOmitNulls()`](3587-util-json-stringifyomitnulls.md "Produces a JSON formatted string from the input, by excluding empty records and empty arrays."),
> empty records (where all members are `NULL`), and empty arrays will be written in the
> JSON string. For detailed control on null and empty variables when serializing to JSON elements, use
> the `json_null` variable definition attribute.

The method raises error [-8110](4483-genero-bdl-errors.md) if
the JSON string cannot be generated.

For more details about FGL to JSON conversion, see [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.").

## Example

```
IMPORT util
MAIN
    DEFINE rec RECORD
             field1 INTEGER,
             field2 CHAR(1),
             subrec1 RECORD
               field11 INTEGER,
               fiedl12 VARCHAR(30)
             END RECORD,
             subarr1 DYNAMIC ARRAY OF INTEGER
           END RECORD
    INITIALIZE rec TO NULL
    LET rec.field1 = 999
    LET rec.subarr1[3] = 888
    DISPLAY "stringify() : ", util.JSON.stringify(rec)
END MAIN
```

Output:

```
stringify() : {"field1":999,"subrec1":{},"subarr1":[null,null,888]}
```

## Related links

**Related concepts**  

[util.JSON.stringifyOmitNulls](3587-util-json-stringifyomitnulls.md "Produces a JSON formatted string from the input, by excluding empty records and empty arrays.")

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")

[BDL to JSON type conversion rules](../09_advanced-features/0960-bdl-to-json-type-conversion-rules.md "Specific type conversion rules apply when converting a BDL variable to JSON.")
