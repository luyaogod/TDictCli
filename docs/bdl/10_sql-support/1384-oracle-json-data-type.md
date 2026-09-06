---
title: "Oracle JSON data type"
source: "fgl-topics/c_fgl_odiagora_020.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > Oracle JSON data type"
type: "concept"
description: "Informix® JSON (JavaScript Object Notation) is a widely used standard for data serialization. Informix supports the BSON and JSON data types, to store JSON documents. With Genero BDL, you can use ..."
---

# Oracle JSON data type

## Informix®

[JSON (JavaScript
Object Notation)](https://www.json.org/json-en.html) is a widely used standard for data serialization.

Informix supports the `BSON`
and `JSON` data types, to store JSON documents.

With Genero BDL, you can use `VARCHAR`, `STRING` or
`TEXT` variables, as well as build-in classes such as [`util.JSON`](../09_advanced-features/0954-json-support.md "Genero BDL supports JSON data manipulation."), to store and manipulate JSON
data.

## ORACLE

Oracle 21c introduced support for a native `JSON` data
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

The `JSON` type has been introduced in Oracle 21c, but until Oracle 23c and the
`dbmora_23` ODI driver, Genero provided only a `dbmora_18` driver
based on Oracle 18 client without `SQLT_JSON` type support. Therefore, Genero can
only support the new Oracle `JSON` type with the `dbmora_23` ODI
driver based on the Oracle 23c instant client, connecting to Oracle server 21c or 23c.

Starting with the `dbmora_23` ODI driver, it is possible to use [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") and [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") FGL variables as SQL input parameters
for `JSON` columns.

As the size of a `JSON` object is undefined, you must fetch `JSON`
data into `TEXT` variables, when the `SELECT` statement retrieves
`JSON` data without any type conversion.

In order to fetch `JSON` data into a `VARCHAR` or
`STRING` variable, convert the `JSON` data to a character string with
the `JSON_SERIALIZE()` SQL function:

```
DEFINE k INTEGER
DEFINE s STRING
DECLARE c1 CURSOR FOR
    SELECT pkey, JSON_SERIALIZE(jdoc) FROM tab1 ORDER BY pkey
FOREACH c1 INTO k, s
    DISPLAY k, ": ", NVL(s, "<null>")
END FOREACH
```

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with Oracle `JSON` type are
converted to the FGL `TEXT` data type.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
