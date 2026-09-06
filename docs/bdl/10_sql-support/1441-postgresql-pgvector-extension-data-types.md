---
title: "PostgreSQL pgvector extension data types"
source: "fgl-topics/c_fgl_odiagpgs_035.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > PostgreSQL pgvector extension data types"
type: "concept"
description: "Informix® A vector is used in AI models to represent the characteristics of an object, that can be a text document, and image or a sound. Recent database engines provide support to store vector data. ..."
---

# PostgreSQL pgvector extension data types

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

## PostgreSQL

The `pgvector` extension can be installer in PostgreSQL 13+ to get vector storage
and similarity search functions. Additional extensions such as `pgvectorscale` and
`pgai` should be installed to get full benefit of this feature. For more detail,
check out the open source <https://github.com/pgvector/pgvector> project.

Once `pgvector` extension installed in your PostgreSQL database, you can use the
following new SQL types to store vector data:

- `vector(dim)`: 64-bit floating point vector.
- `halfvec(dim)`: 32-bit floating point vector.
- `sparsevec(dim)`: 64-bit floating point sparse vector.
- `bit(size)`: native bit-array (very specific use cases.

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

## Solution

The ODI driver relies on implicit vector/text conversion and cast operators provided by
PostgreSQL.

> **Important:**
>
> The `pgvector` extension support is disabled by default. When enabled, at
> connection time, the ODI driver will query the PostgreSQL system catalog to get the OID identifier
> of the `vector`, `halfvec` and `sparsevec` types. In
> order to enable `pgvector` support, set the following [FGLPROFILE
> entry](1084-postgresql-specific-fglprofile-parameters.md) to
> `true`:
>
> ```
> dbi.database.test1.pgs.pgvector = true
> ```

Program variables of type [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") and [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") can be used as SQL input parameters for vector columns, or as vector
function parameters, as long as the SQL statement contains the cast operators such as
`::vector` for the `?` SQL parameter:

```
DEFINE tsv STRING
PREPARE stmt
   FROM "UPDATE mydocs SET emb = ?::vector WHERE pkey = ?"
LET pkey = 101
LET tsv = "[0.34,1.04,0.89,...]"
EXECUTE stmt USING tsv, pkey
```

As the size of a `VECTOR` can be quite large, consider fetching vector data into
`TEXT` variables. When fetching vector data into a `TEXT`,
`VARCHAR` or `STRING` variable, PostgreSQL can implicitely convert the
vector data to a JSON array of numbers:

```
DEFINE k INTEGER
DEFINE s STRING
DECLARE c1 CURSOR FOR
    SELECT pkey, emb FROM mydocs ORDER BY pkey
FOREACH c1 INTO k, s
    DISPLAY k, ": ", NVL(s, "<null>")
END FOREACH
```

PostgreSQL vector similarity search functions and operators can be used to query the database, as
shown in the next example:

```
VAR d DECIMAL(10,5)
VAR s STRING
VAR t TEXT
DECLARE c11 CURSOR FROM
   "SELECT pkey, ( emb <-> ?::vector ) AS distance
     FROM mydoc ORDER BY distance"
LET s = "[-4,5,3,-2,6]"
FOREACH c11 USING s INTO rec.doc_id, d
    DISPLAY rec.pkey, "    distance = ", d
END FOREACH
```

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with PostgreSQL `VECTOR`,
`HALFVEC` and `SPARSEVEC` types are converted to the FGL
`TEXT` data type.

> **Important:**
>
> The PostgreSQL `bit` data type is not supported by Genero BDL. You must use
> `vector`, `halfvec` or `sparsevec` PostgreSQL pgvector
> types.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
