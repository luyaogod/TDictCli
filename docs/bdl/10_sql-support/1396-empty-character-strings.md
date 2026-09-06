---
title: "Empty character strings"
source: "fgl-topics/c_fgl_odiagora_018.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data manipulation > Empty character strings"
type: "concept"
description: "Informix® Informix SQL consisders empty strings ( '' or \"\" ) as a non-NULL string with a length of zero. Note: In Genero BDL, when setting a variable with an empty string constant, it is automatically ..."
---

# Empty character strings

## Informix®

Informix SQL consisders empty strings (
`''` or `""` ) as a non-NULL string with a length of zero.

> **Note:**
>
> In Genero BDL, when setting a variable with an empty string constant, it is automatically set to
> a `NULL` value:
>
> ```
> DEFINE x char(10)
> LET x = ""
> IF x IS NULL THEN -- evaluates to TRUE
>    ...
> END IF
> ```

## ORACLE

Oracle® SQL considers empty string
literals (`''`) as `NULL`.

Using literal string values that are empty (`''`) for `INSERT` or
`UPDATE` statements will result in the storage of `NULL` values with
Oracle, while Informix stores the value as a string with a length of zero:

```
INSERT INTO tab1 ( col1, col2 ) VALUES ( NULL, '' )
```

Using comparison operators (`col=''`) with Oracle makes no sense, because an empty string is equivalent to
`NULL`: The correct SQL expression is (`col IS NULL`).

```
SELECT * FROM tab1 WHERE col2 IS NULL
```

With Oracle, we must also distinguish `CHAR` and `VARCHAR2`.
`CHAR` values are blank padded. In SQL, `''` empty string literals are
`NULL`, but in PL/SQL, assigning a `CHAR` variable with
`''` results as `NOT NULL`, because `CHAR` is
blank-padded. However, a `VARCHAR2` variable set with `''` becomes
`NULL`:

```
SET SERVEROUTPUT ON;

DROP TABLE tab1;
CREATE TABLE tab1 ( pk NUMBER, c1 CHAR(10), c2 VARCHAR2(10) );

DECLARE
    v1 CHAR(10);
    v2 VARCHAR2(10);
BEGIN
    v1 := '';
    v2 := '';
    IF v1 IS NULL THEN
       dbms_output.put_line('v1 is null');
    ELSE
       dbms_output.put_line('v1 = ' || v1);
    END IF;
    IF v2 IS NULL THEN
       dbms_output.put_line('v2 is null');
    ELSE
       dbms_output.put_line('v2 = ' || v2);
    END IF;
    INSERT INTO tab1 VALUES ( 101, '', '' );
    INSERT INTO tab1 VALUES ( 102, v1, v2 );
END;
/

SELECT pk, NVL(c1,'NULL'), NVL(c2,'NULL') FROM tab1 ORDER BY pk;

QUIT
```

Output:

```
v1 =
v2 is null
PL/SQL procedure successfully completed.
 PK NVL(C1,'NULL') NVL(C2,'NULL')
---------- ------------------------------ ------------------------------
 101 NULL NULL
 102 NULL
```

## Solution

To increase portability, it is recommended that you avoid the usage of literal string values with
a length of zero in SQL statements. Instead, use the `NULL` constant, or program
variables.
