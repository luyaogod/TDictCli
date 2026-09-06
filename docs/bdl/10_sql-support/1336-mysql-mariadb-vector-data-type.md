---
title: "MySQL/MariaDB VECTOR data type"
source: "fgl-topics/c_fgl_odiagmys_037.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > MySQL/MariaDB VECTOR data type"
type: "concept"
description: "Informix® A vector is used in AI models to represent the characteristics of an object, that can be a text document, and image or a sound. Recent database engines provide support to store vector data. ..."
---

# MySQL/MariaDB VECTOR data type

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

## Oracle® MySQL and MariaDB

MySQL 9 and MariaDB 11.7 introduced support for a native VECTOR data type, to store the data
produced by embedding models along with business data:

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

ODI drivers for `VECTOR` support:

- The MySQL `VECTOR` type is supported with the `dbmmys_9_7` ODI
  driver.
- The MariaDB `VECTOR` type is supported with the `dbmmdb_10_2` ODI
  drivers.

Use [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") or [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") FGL variables as SQL input parameters
for `VECTOR` columns, or as vector function parameters.

When using `TEXT` variables as SQL parameters, use the MySQL
`STRING_TO_VECTOR()` function (or MariaDB `VEC_FromText()` function),
to cast the character string data to a `VECTOR` value:

```
DEFINE v1 TEXT
LOCATE v1 IN FILE "myvector.txt"
INSERT INTO tab VALUES ( 101, STRING_TO_VECTOR(v1) )
```

When using `BYTE` variables as SQL parameters, the variable can be used directly
without casting. However, the `BYTE` data must be a valid binary representation of a
MySQL or MariaDB vector:

```
DEFINE v1 BYTE
LOCATE v1 IN FILE "myvector.bin"
INSERT INTO tab VALUES ( 101, v1 )
```

`VECTOR` data can be fetched into `TEXT` variables, as long as you
cast the `VECTOR` values to character strings with the MySQL
`VECTOR_TO_STRING()` function or MariaDB `VEC_FromText()`
function:

```
DEFINE v1 TEXT
LOCATE v1 IN MEMORY
SELECT VECTOR_TO_STRING(vector_column) INTO v1 FROM tab WHERE pkey = 101
```

Vector distance calculation functions can be used to query the database. With MySQL, use
`DISTANCE(vector1,vector2,distance-metric)`
SQL function. With MariaDB, use
`VEC_DISTANCE(vector1,vector2)`. See respective
database engines documentation for more details.
> **Important:**
>
> According to MySQL doc:
>
> <https://dev.mysql.com/doc/refman/9.7/en/vector-functions.html#function_distance>
>
> `DISTANCE()` is available only for users of MySQL HeatWave on OCI and MySQL AI; it
> is not included in MySQL Commercial or Community distributions.

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with MySQL or MariaDB `VECTOR` type
are converted to the FGL `BYTE` data type.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
