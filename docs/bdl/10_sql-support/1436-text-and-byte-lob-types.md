---
title: "TEXT and BYTE (LOB) types"
source: "fgl-topics/c_fgl_odiagpgs_032.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > TEXT and BYTE (LOB) types"
type: "concept"
description: "Informix® Informix provides the TEXT , BYTE , CLOB and BLOB data types to store very large texts or binary data. Legacy Informix 4GL applications typically use the TEXT and BYTE types. Genero BDL does ..."
---

# TEXT and BYTE (LOB) types

## Informix®

Informix provides the `TEXT`,
`BYTE`, `CLOB` and `BLOB` data types to store very
large texts or binary data.

Legacy Informix 4GL applications typically use the
`TEXT` and `BYTE` types.

Genero BDL does not support the Informix
`CLOB` and `BLOB` types.

## PostgreSQL

PostgreSQL provides the `TEXT` and `BYTEA` data types for large
objects storage. With these data types, large objects are handled as a whole.

PostgreSQL also supports LOB storage through the large objects facility based on stream-style
access. The large object facility is provided as a set of C and SQL API functions to create / delete
/ modify large objects identified by a unique object id (OID). For example, the
`lo_create(-1)` SQL function will create a new large object and return a new object
id that will be used to handle the LOB. See PostgreSQL documentation for more details.

## Solution

`TEXT` and `BYTE` data can be stored in PostgreSQL
`TEXT` and `BYTEA` columns.

The `TEXT` and
`BYTE` types translation can be controlled with the following FGLPROFILE
entries:

```
dbi.database.dsname.ifxemul.datatype.text = { true | false }
dbi.database.dsname.ifxemul.datatype.byte = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

Genero BDL does not interface automatically with the PostgreSQL Large Object facility. However,
the identifier returned by the lo\_import() function can be stored in
`BIGINT` variables, and you can use server-side LOB functions to convert large
objects to `BYTEA` data, that can be fetched into `BYTE` variables.
The next code example creates a table with an `"image"` column defined with the
`OID` type, imports a LOB from an image file, and then fetches the LOB back into a
`BYTE`:

```
MAIN
    DEFINE img BYTE, obj_id BIGINT

    CONNECT TO "test1+driver='dbmpgs'" USER "postgres" USING "fourjs"

    # Need superuser privileges to create the LOB....
    WHENEVER ERROR CONTINUE
    DROP TABLE t1
    WHENEVER ERROR STOP
    EXECUTE IMMEDIATE "create table t1 ( k int, image oid )"
    GRANT SELECT ON t1 TO PUBLIC
    INSERT INTO t1 VALUES ( 1, lo_import("/var/images/landscape.png") )
    SELECT image INTO obj_id FROM t1 WHERE k=1
    DISPLAY "obj_id = ", obj_id
    EXECUTE IMMEDIATE "grant select on large object "||obj_id||" to public"

    # Next block can be executed by any user:
    LOCATE img IN FILE -- a temp file will be used
    SELECT loread(lo_open(image, 262144), 1000000)
      INTO img FROM t1 WHERE k=1
    DISPLAY length(img)

    # Delete the object
    SELECT lo_unlink(obj_id) FROM t1 WHERE k=1

    DROP TABLE t1

END MAIN
```

## Related links

**Related concepts**  

[Using portable data types](1008-using-portable-data-types.md "Only a limited set of data types are really portable across several database engines.")
