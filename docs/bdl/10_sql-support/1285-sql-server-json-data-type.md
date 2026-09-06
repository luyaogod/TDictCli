---
title: "SQL Server JSON data type"
source: "fgl-topics/c_fgl_odiagmsv_020.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > SQL Server JSON data type"
type: "concept"
description: "Informix® JSON (JavaScript Object Notation) is a widely used standard for data serialization. Informix supports the BSON and JSON data types, to store JSON documents. With Genero BDL, you can use ..."
---

# SQL Server JSON data type

## Informix®

[JSON (JavaScript
Object Notation)](https://www.json.org/json-en.html) is a widely used standard for data serialization.

Informix supports the `BSON`
and `JSON` data types, to store JSON documents.

With Genero BDL, you can use `VARCHAR`, `STRING` or
`TEXT` variables, as well as build-in classes such as [`util.JSON`](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation."), to store and manipulate JSON
data.

## SQL SERVER

SQL Server 2025 introduced support for a native `JSON` data
type:

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

As the size of a `JSON` object is undefined, you must fetch `JSON`
data into `TEXT` variables, when the `SELECT` statement retrieves
`JSON` data without any type conversion.

In order to fetch `JSON` data into a `VARCHAR` or
`STRING` variable, convert the `JSON` data to a character string with
`CAST(jdoc AS VARCHAR(MAX))`:

```
DEFINE k INTEGER
DEFINE s STRING
DECLARE c1 CURSOR FOR
    SELECT pkey, CAST(jdoc AS VARCHAR(MAX)) FROM tab1 ORDER BY pkey
FOREACH c1 INTO k, s
    DISPLAY k, ": ", NVL(s, "<null>")
END FOREACH
```

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with SQL Server `JSON` type are
converted to the FGL `TEXT` data type.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
