---
title: "Cursors WITH HOLD"
source: "fgl-topics/c_fgl_odiagpgs_033.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > Cursors WITH HOLD"
type: "concept"
description: "Informix® Informix closes opened cursors automatically when a transaction ends, unless the WITH HOLD option is used in the DECLARE instruction: DECLARE c1 CURSOR WITH HOLD FOR SELECT ... OPEN c1 BEGIN ..."
---

# Cursors WITH HOLD

## Informix®

Informix closes opened cursors automatically when a
transaction ends, unless the `WITH HOLD` option is used in the
`DECLARE`
instruction:

```
DECLARE c1 CURSOR WITH HOLD FOR SELECT ...
OPEN c1
BEGIN WORK
FETCH c1 ...
COMMIT WORK
FETCH c1 ...
CLOSE c1
```

## PostgreSQL

With PostgreSQL, opened cursors using `SELECT` statements without a `FOR
UPDATE` clause are not closed when a transaction ends. All PostgreSQL cursors are
`WITH HOLD` cursors, unless the `FOR UPDATE` clause issued in the
`SELECT` statement.
> **Note:**
>
> Native PostgreSQL `WITH HOLD` cursors are
> automatically closed, if the cursor is opened inside a transaction block, and the transaction is
> canceled with a rollback.

BDL `WITH HOLD` cursors declared with a `SELECT ... FOR UPDATE`
cannot be supported with PostgreSQL: Native holdable cursors declared for update produce the
following SQL
error:

```
DECLARE CURSOR WITH HOLD ... FOR UPDATE is not supported
DETAIL:  Holdable cursors must be READ ONLY.
```

## Solution

For consistency with other database brands, database cursors that are
not declared `WITH HOLD` are automatically closed, when a `COMMIT
WORK` or `ROLLBACK WORK` is performed.

> **Important:**
>
> Opening a `WITH HOLD` cursor
> declared with a `SELECT FOR UPDATE` results in an SQL error; in the same conditions,
> this does not normally appear with Informix. Review the
> program logic in order to find another way to set locks.

Cursors declared `WITH HOLD` and using a `SELECT` without the
`FOR UPDATE` clause can be used, as long as the cursor is opened outside
a transaction block.
> **Important:**
>
> Since PostgreSQL automatically closes all `WITH
> HOLD` cursors opened in a transaction which is canceled by a rollback, do not open such
> cursor after a `BEGIN WORK`, that can potentially be terminated by a `ROLLBACK
> WORK`.

The following code can be used with PostgreSQL:

```
MAIN
    DEFINE rec RECORD
               pkey INT,
               name VARCHAR(50)
           END RECORD

    CONNECT TO "test1+driver='dbmpgs'" USER "pgsuser" USING "fourjs"

    WHENEVER ERROR CONTINUE
    DROP TABLE tab1
    WHENEVER ERROR STOP
    CREATE TABLE tab1 ( pkey INT, name VARCHAR(50) )
    FOR rec.pkey=1 TO 10
        LET rec.name = SFMT("Item %1", rec.pkey)
        INSERT INTO tab1 VALUES ( rec.pkey, rec.name )
    END FOR

    DECLARE c1 CURSOR WITH HOLD FOR SELECT * FROM tab1 ORDER BY pkey

    OPEN c1 -- Outside TX block!
    FETCH c1 INTO rec.*     DISPLAY rec.*
    BEGIN WORK
    FETCH c1 INTO rec.*     DISPLAY rec.*
    ROLLBACK WORK
    BEGIN WORK
    FETCH c1 INTO rec.*     DISPLAY rec.*
    COMMIT WORK
    FETCH c1 INTO rec.*     DISPLAY rec.*
    BEGIN WORK
    FETCH c1 INTO rec.*     DISPLAY rec.*
    ROLLBACK WORK
    FETCH c1 INTO rec.*     DISPLAY rec.*
    CLOSE c1

    FOREACH c1 INTO rec.*
        BEGIN WORK
        DISPLAY rec.* -- Do some real SQL here...
        IF rec.pkey MOD 2 == 0 THEN
            COMMIT WORK
        ELSE
            ROLLBACK WORK
        END IF
    END FOREACH

END MAIN
```

Since PostgreSQL automatically closes `FOR UPDATE` cursors when the transaction
ends, opening cursors declared `FOR UPDATE` and the `WITH HOLD` option
results in an SQL error that does not normally appear with Informix under the same conditions.

Review the program logic in order to find another way to set locks.

## Related links

**Related concepts**  

[Cursors WITH HOLD](1029-cursors-with-hold.md "Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.")
