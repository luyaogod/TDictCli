---
title: "fgldbutl.db_get_sequence()"
source: "fgl-topics/c_fgl_utility_functions_DB_GET_SEQUENCE.html"
breadcrumb: "Library reference > Utility modules > fgldbutl: Database utility module > fgldbutl.db_get_sequence()"
type: "concept"
---

# fgldbutl.db_get_sequence()

> Generates a new sequence for a given identifier.

## Syntax

```
FUNCTION db_get_sequence(
   id STRING )
  RETURNS BIGINT
```

1. id is the identifier of the sequence.

## Usage

This function generates a new sequence from a register table created in
the current database.

> **Important:**
>
> 1. Needs a database table called SEQREG.
> 2. The function must be called inside a transaction block.

The table must be created as follows:

```
CREATE TABLE seqreg (
  sr_name VARCHAR(30) NOT NULL,
  sr_last BIGINT NOT NULL,
  PRIMARY KEY (sr_name)
)
```

Each time you call this function, the sequence
is incremented in the database table and returned by the function.

It is mandatory to use this function inside a transaction block, in
order to generate unique sequences.

## Example

```
IMPORT FGL fgldbutl
MAIN
  DEFINE ns BIGINT, s INTEGER
  DATABASE mydb
  BEGIN WORK
  LET ns = db_get_sequence("mytable")
  INSERT INTO mytable VALUES ( ns, 'a new sequence' )
  COMMIT WORK
END MAIN
```

## Related links

**Related concepts**  

[Auto-incremented columns (serials)](../10_sql-support/1023-auto-incremented-columns-serials.md "How to implement automatic record keys.")
