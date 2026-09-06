---
title: "PostgreSQL TSVECTOR/TSQUERY data types"
source: "fgl-topics/c_fgl_odiagpgs_019.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > PostgreSQL TSVECTOR/TSQUERY data types"
type: "concept"
description: "Informix® Informix supports basic text search with the bts index and the bts_contains() SQL function. A bts index requires an sbspace . For example: CREATE INDEX desc_idx ON products (brands ..."
---

# PostgreSQL TSVECTOR/TSQUERY data types

## Informix®

Informix supports basic text search with
the `bts` index and the `bts_contains()` SQL function. A
`bts` index requires an sbspace.

For example:

```
CREATE INDEX desc_idx ON products (brands bts_char_ops)
   USING bts IN sbsp1;
```

For more details about basic text search, refere to the Informix documentation.

## PostgreSQL

PostgreSQL 9+ support native `TSVECTOR` and `TSQUERY` data types to
implement full text search capabilities in your application.

`TSVECTOR` is the type to store text search vector data, that can be compared to
`TSQUERY` values with the `@@` matching operator.

Create a dedicated `TSVECTOR` column that is automatically populated using the
`to_tsvector()` function. The standard approach is to use a `GENERATED ALWAYS
AS` clause, which ensures the vector is automatically updated whenever the source text
columns change. To make searches fast, a `GIN` index is then applied to this
generated column:

```
CREATE TABLE mydocs (
       pkey INTEGER,
       text_chunk VARCHAR(300),
       ts_index_col TSVECTOR GENERATED ALWAYS AS
          (to_tsvector('english',text_chunk)) STORED
    );

CREATE INDEX mydocs_ix1 ON mydocs USING GIN (ts_index_col);
```

## Solution

The purpose of `TSVECTOR` and `TSQUERY` data types is to perform
text search, where the program code provides the search condition as a `TSQUERY`,
that will be compared to a `TSVECTOR` column with the `@@`
operator:

```
DEFINE search_filter STRING
DECLARE c1 CURSOR
   FROM "SELECT * FROM mydocs WHERE ts_index_col @@ ?::tsquery ORDER BY pkey"
LET search_filter = "cats & dogs"
FOREACH c1 USING search_filter INTO rec.*
    DISPLAY rec.*
END FOREACH
```

For specific cases, you can use [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size."), [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.") and [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") FGL variables as SQL input parameters for `TSVECTOR`
columns, as long as the `::tsvector` cast operator follows the `?`
question mark placeholder:

```
DEFINE tsv STRING
PREPARE stmt
   FROM "UPDATE mydocs SET ts_index_col = ?::tsvector WHERE pkey = ?"
LET pkey = 101
LET tsv = "'cat':2 'kitchen':6"
EXECUTE stmt USING tsv, pkey
```

However, the standard approach is to let PostgreSQL generate the `TSVECTOR` data
automatically with the `GENERATED ALWAYS AS` clause.

If needed, since `TSVECTOR` or `TSQUERY` types have no size,
consider fetching `TSVECTOR` or `TSQUERY` data into
`TEXT` variables:

```
DEFINE k INTEGER
DEFINE s STRING
DECLARE c1 CURSOR FOR
    SELECT pkey, ts_index_col FROM tab1 ORDER BY pkey
FOREACH c1 INTO k, s
    DISPLAY k, ": ", NVL(s, "<null>")
END FOREACH
```

Fetching `TSVECTOR` or `TSQUERY` data into `VARCHAR`
or `STRING` variable is also possible.

When extracting database schemas with the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool, columns with PostgreSQL `TSVECTOR` or
`TSQUERY` type are converted to the FGL `TEXT` data type.

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
