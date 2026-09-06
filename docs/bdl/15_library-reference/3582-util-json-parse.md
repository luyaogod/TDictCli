---
title: "util.JSON.parse"
source: "fgl-topics/c_fgl_ext_util_JSON_parse.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > util.JSON methods > util.JSON.parse"
type: "concept"
---

# util.JSON.parse

> Parses a JSON string and fills program variables with the values.

## Syntax

```
util.JSON.parse(
     s STRING,
     variableRef any-type
   )
```

1. s is a string value that contains JSON formatted data.
2. variableRef is the variable to be initialized with values of the JSON string.
   > **Important:**
   >
   > The variableRef parameter is passed by reference to the
   > method.
3. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `util.JSON.parse()` class method scans the JSON source string passed as
parameter and fills the destination variable members by name.

If the provided string is not valid JSON, the `parse()` method will raise error
[-8109](4483-genero-bdl-errors.md). Consider enclosing the
`parse()` method call in a `TRY/CATCH` block, if the source string can
be malformed JSON.

The destination variable is expected to have the same structure as the JSON source data, it can
be a `RECORD`, `DYNAMIC ARRAY` or a `DICTIONARY`.

The `parse()` method initializes the target variable to `NULL`
before the parsing starts.

See [JSON support](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.") for details on how the destination
variable is populated when the structures are not identical.

> **Note:**
>
> When parsing a JSON string to fill a `TEXT` or
> `BYTE` variable, if the data storage for the LOB variable has not been defined with
> the [`LOCATE`](../08_language-basics/0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.") instruction, the JSON
> methods will automatically locate the `TEXT` or `BYTE` in memory. This
> applies also to `TEXT` and `BYTE` elements of records, arrays and
> dictionaries.

## Example

```
IMPORT util
MAIN
    DEFINE cust_rec RECORD
               cust_num INTEGER,
               cust_name VARCHAR(30),
               order_ids DYNAMIC ARRAY OF INTEGER
           END RECORD
    DEFINE js STRING
    LET js='{ "cust_num":2735, "cust_name":"McCarlson",
              "order_ids":[234,3456,24656,34561] }'
    TRY
        CALL util.JSON.parse( js, cust_rec )
        DISPLAY cust_rec.cust_name
        DISPLAY cust_rec.order_ids[4]
    CATCH
        DISPLAY "ERROR:", status
    END TRY
END MAIN
```

## Related links

**Related concepts**  

[Records](../08_language-basics/0715-records.md "Records allow structured program variables definitions.")
