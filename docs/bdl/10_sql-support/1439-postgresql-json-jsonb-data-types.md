---
title: "PostgreSQL JSON/JSONB data types"
source: "fgl-topics/c_fgl_odiagpgs_020.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > PostgreSQL JSON/JSONB data types"
type: "concept"
description: "Informix® JSON (JavaScript Object Notation) is a widely used standard for data serialization. Informix supports the BSON and JSON data types, to store JSON documents. With Genero BDL, you can use ..."
---

# PostgreSQL JSON/JSONB data types

## Informix®

[JSON (JavaScript
Object Notation)](https://www.json.org/json-en.html) is a widely used standard for data serialization.

Informix supports the `BSON`
and `JSON` data types, to store JSON documents.

With Genero BDL, you can use `VARCHAR`, `STRING` or
`TEXT` variables, as well as build-in classes such as [`util.JSON`](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation."), to store and manipulate JSON
data.

## PostgreSQL

PostgreSQL 9+ support native `JSON` and `JSONB` data types.
`JSONB` is nothing but a different way to store JSON data in a PostgreSQL
database:

```
CREATE TABLE tab1 (
   pkey INTEGER NOT NULL PRIMARY KEY,
   doc1 JSON,
   ...
);
INSERT INTO tab1 values ( 101,
    '{"user":{"creadate":"2023-03-14","scorelist":[700, 650, 720]}}'
);
```

## Solution

With Genero BDL, it is possible to use [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") and [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") FGL variables as SQL input parameters for `JSON` or
`JSONB` columns, as long as the `::json` or `::jsonb`
cast operator follows the `?` question mark
placeholder:

```
DEFINE rec RECORD
        pkey INTEGER,
        doc1 STRING,
        ...
    END RECORD
PREPARE stmt FROM "INSERT INTO tab1 VALUES ( ?, ?::json, ... )
EXECUTE stmt USING rec.*
```

As the size of a `JSON` object is undefined, you must fetch `JSON`
and `JSONB` data into `TEXT` variables, when the
`SELECT` statement retrieves `JSON` and `JSONB` data
without any type conversion.

To fetch `JSON` and `JSONB` data into a `VARCHAR` or
`STRING` variable, convert the `JSON/JSONB` data to a character string
with the `::text` cast operator:

```
DEFINE k INTEGER
DEFINE s STRING
DECLARE c1 CURSOR FOR
    SELECT pkey, jdoc::text FROM tab1 ORDER BY pkey
FOREACH c1 INTO k, s
    DISPLAY k, ": ", NVL(s, "<null>")
END FOREACH
```

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with PostgreSQL `JSON` or
`JSONB` type are converted to the FGL `TEXT` data type.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
