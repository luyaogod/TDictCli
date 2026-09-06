---
title: "Empty strings and NULLs"
source: "fgl-topics/c_fgl_sql_programming_068.html"
breadcrumb: "SQL support > SQL programming > SQL portability > CHAR and VARCHAR types > Empty strings and NULLs"
type: "concept"
---

# Empty strings and NULLs

> Depending on the context, an empty string ( '' ) can be considered as NULL or NOT NULL.

Most SQL databases distinguish `''` empty strings from `NULL` (with
some exceptions like Oracle® DB).

With Genero BDL character string variables, an empty string is the equivalent to
`NULL`:

```
MAIN
    DEFINE vc10 VARCHAR(10)
    LET vc10 = ''
    DISPLAY (vc10 IS NULL)   -- shows 1 / TRUE
END MAIN
```

When inserting an empty string SQL literal into a DB row, and fetching the value back into a
variable, the result might be null or non-null, depending on the database type and target variable
type, since `CHAR` values are blank-padded.

See for example the following
code:

```
DEFINE rec RECORD
           pk INTEGER,
           c10 CHAR(10),
           vc10 VARCHAR(10)
       END RECORD

MAIN

    -- CONNECT TO ... (your connection instruction goes here)
    DISPLAY "DB driver: ", fgl_db_driver_type()

    CREATE TEMP TABLE tt1 ( pk INT, c10 CHAR(10), vc10 VARCHAR(10) )

    INSERT INTO tt1 VALUES ( 101, ' ', ' ' )
    INSERT INTO tt1 VALUES ( 102, NULL, NULL )
    INSERT INTO tt1 VALUES ( 103, '', '' )

    SELECT tt1.* INTO rec.* FROM tt1 WHERE pk=101
    CALL show_record()

    SELECT tt1.* INTO rec.* FROM tt1 WHERE pk=102
    CALL show_record()

    SELECT tt1.* INTO rec.* FROM tt1 WHERE pk=103
    CALL show_record()

END MAIN

FUNCTION show_record()
    DISPLAY rec.pk, "  [", NVL(rec.c10,"<NULL>"),"]",
            COLUMN 30, "[",NVL(rec.vc10,"<NULL>"),"]"
END FUNCTION
```

Result with Informix IDS 14:

```
DB driver: ifx
        101  [          ]    [ ]
        102  [<NULL>]        [<NULL>]
        103  [          ]    [<NULL>]
```

Result with Oracle DB 19c:

```
DB driver: ora
        101  [          ]    [ ]
        102  [<NULL>]        [<NULL>]
        103  [<NULL>]        [<NULL>]
```

> **Tip:**
>
> Mixing empty strings and `NULL` values in database rows makes
> queries more difficult: Is an empty string considered as null or not null by your application? To
> avoid empty strings in database columns, always use program variables to insert data into SQL
> tables: If a `CHAR` or `VARCHAR` program variable is assigned with an
> empty string, it will be set to `NULL`, and inserted as `NULL` into
> the database.
