---
title: "SQL Server VECTOR data type"
source: "fgl-topics/c_fgl_odiagmsv_035.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > SQL Server VECTOR data type"
type: "concept"
description: "Informix® A vector is used in AI models to represent the characteristics of an object, that can be a text document, and image or a sound. Recent database engines provide support to store vector data. ..."
---

# SQL Server VECTOR data type

## Informix®

A vector is used in AI models to
represent the characteristics of an object, that can be a text document, and image or a sound.
Recent database engines provide support to store vector data.

Informix version 15 provides the "Vector Blade" extension to
store and search for vector data. See Informix documentation for more details.

With Genero BDL,
you can use `VARCHAR`, `STRING` or `TEXT` variables, to
store vector data as a JSON array of
numbers:

```
[-3.45, -4.45, 9.234, ... ]
```

Use Genero [JSON APIs](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation.") to manipulate vectors as JSON arrays.

## SQL SERVER

SQL Server 2025 introduced support for a native VECTOR data type, to store the data produced by
embedding models along with business data:

```
CREATE TABLE tab1 (
   pkey INTEGER NOT NULL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   vect1 VECTOR(256),
   ...
);
INSERT INTO tab1 values ( 101, 'Mike STORN',
    '[-3.45, -4.45, 9.234, ... ]'
);
```

Read SQL Server documentation for more details about AI vector embeddings, and the VECTOR
type.

## Solution

The SQL Server `VECTOR` type is supported with `dbmsnc_18` ODI
driver and Microsoft ODBC for SQL Server version 18.6.1.1 or +.

The ODI driver relies on implicit `VECTOR` serialization/deserialization provided
by SQL Server.

You can use [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") and [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") FGL variables as SQL input parameters
for `VECTOR` columns, or as vector function parameters.

`VARCHAR` and `STRING` variable types are suitable for
`VECTORS` with small dimensions such as 128. For large vectors with dimensions such
as 1024, use `TEXT` variables.

As the size of a `VECTOR` can be quite large, you must fetch
`VECTOR` data into `TEXT` variables, when the `SELECT`
statement retrieves `VECTOR` data without any type conversion.

In order to fetch `VECTOR` data into a `VARCHAR` or
`STRING` variable, convert the `VECTOR` data to a character string
with a `CAST()` SQL operator:

```
DEFINE k INTEGER
DEFINE s STRING
DECLARE c1 CURSOR FOR
    SELECT pkey, CAST(vect1 AS VARCHAR(200)) FROM tab1 ORDER BY pkey
FOREACH c1 INTO k, s
    DISPLAY k, ": ", NVL(s, "<null>")
END FOREACH
```

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with SQL Server `VECTOR` type are
converted to the FGL `TEXT` data type.

SQL Server `VECTOR` functions can be used to query the database, as shown in the
next example with the `VECTOR_DISTANCE()` SQL function.

> **Important:**
>
> For vector functions such as `VECTOR_DISTANCE()`, SQL Server expects expressions
> evaluating to a `VECTOR` data type. Consequently, it is mandatory to use a
> `CAST()` operator to convert the `STRING`, `VARCHAR` or
> `TEXT` parameter to a
> `VECTOR(nnn)`:
>
> ```
> VECTOR_DISTANCE( 'cosine', vector_column , CAST(? AS VECTOR(128) )
> ```

```
FUNCTION find_matching_rows() RETURNS ()
  DEFINE dist DECIMAL(10,5)
  DEFINE vect TEXT
  DEFINE doc_id INTEGER
  DECLARE curs CURSOR FROM
     "SELECT doc_id,
        VECTOR_DISTANCE('cosine',vect1,CAST(? AS VECTOR(128)) AS distance
     FROM mydoc ORDER BY distance"
  LOCATE vect IN MEMORY
  LET vect = "[-4,5,3,-2,6,...]"
  FOREACH curs USING vect INTO doc_id, dist
    DISPLAY doc_id, ": distance = ", dist
  END FOREACH
END FUNCTION
```

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
