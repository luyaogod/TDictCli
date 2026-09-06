---
title: "Oracle VECTOR data type"
source: "fgl-topics/c_fgl_odiagora_021.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > Oracle VECTOR data type"
type: "concept"
description: "Informix® A vector is used in AI models to represent the characteristics of an object, that can be a text document, and image or a sound. Recent database engines provide support to store vector data. ..."
---

# Oracle VECTOR data type

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

## ORACLE

Oracle 23ai introduced support for a native VECTOR data type, to store the data produced by
embedding models along with business data:

```
CREATE TABLE tab1 (
   pkey INTEGER NOT NULL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   vect1 VECTOR(256, FLOAT64),
   ...
);
INSERT INTO tab1 values ( 101, 'Mike STORN',
    '[-3.45, -4.45, 9.234, ... ]'
);
```

Read Oracle documentation for more details about AI vector embeddings, and the VECTOR type.

## Solution

The Oracle `VECTOR` type is supported with `dbmora_23` ODI driver
and Oracle Instant Client version 23.7 or +.

The ODI driver relies on implicit `VECTOR` serialization/deserialization provided
by Oracle.

You can use [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") and [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") FGL variables as SQL input parameters
for `VECTOR` columns, or as vector function parameters.

As the size of a `VECTOR` can be quite large, you must fetch
`VECTOR` data into `TEXT` variables, when the `SELECT`
statement retrieves `VECTOR` data without any type conversion.

In order to fetch `VECTOR` data into a `VARCHAR` or
`STRING` variable, convert the `VECTOR` data to a character string
with the `VECTOR_SERIALIZE()` SQL function:

```
DEFINE k INTEGER
DEFINE s STRING
DECLARE c1 CURSOR FOR
    SELECT pkey, VECTOR_SERIALIZE(vect1) FROM tab1 ORDER BY pkey
FOREACH c1 INTO k, s
    DISPLAY k, ": ", NVL(s, "<null>")
END FOREACH
```

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with Oracle `VECTOR` type are
converted to the FGL `TEXT` data type.

Oracle `VECTOR` functions can be used to query the database, as shown in the next
example with the `VECTOR_DISTANCE()` SQL function:

```
VAR d DECIMAL(10,5)
VAR s STRING
VAR t TEXT
DECLARE c11 CURSOR FOR
   SELECT doc_id,
          VECTOR_DISTANCE(doc_v1,TO_VECTOR(?,5,INT8)) AS distance
     FROM mydoc ORDER BY distance
LET s = "[-4,5,3,-2,6]"
LOCATE t IN MEMORY
LET t = s
FOREACH c11
       USING s  -- or can also use TEXT variable t
       INTO rec.doc_id, d
    DISPLAY rec.doc_id, "    distance = ", d
END FOREACH
```

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with Oracle `VECTOR` type are
converted to the FGL `TEXT` data type.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
