---
title: "Trailing blanks in CHAR/VARCHAR"
source: "fgl-topics/c_fgl_sql_programming_065.html"
breadcrumb: "SQL support > SQL programming > SQL portability > CHAR and VARCHAR types > Trailing blanks in CHAR/VARCHAR"
type: "concept"
---

# Trailing blanks in CHAR/VARCHAR

> How to cope with trailing blanks in CHAR(N) and VARCHAR(N) SQL columns and program variables?

## Trailing blanks in CHAR/VARCHAR database columns

With all kinds of databases servers, `CHAR` columns are always filled with
blanks up to the size of the column (this is called blank padding).

With IBM®
Informix®, trailing blanks are not significant in
comparisons:

```
CHAR('abc ') = CHAR('abc')
```

With other database servers, trailing blanks are significant when comparing
`VARCHAR` values:

```
VARCHAR('abc ') != VARCHAR('abc')
```

This is a major issue if you mix `CHAR` and `VARCHAR` columns
and variables in your SQL statements, because the result of
an SQL query can be different depending on whether you are
using IBM Informix or another database server.

Furthermore, the semantics of the SQL `LIKE` operator
regarding trailing blanks and CHAR/VARCHAR types can differ
from database to database. For example, try the following
expressions with your database, with a CHAR(5) column containing a
row with the value `'abc'`:

```
CREATE TABLE t1 ( k INT, c CHAR(5), vc VARCHAR(5) )
INSERT INTO t1 VALUES ( 1, 'abc', 'abc' )
SELECT * FROM t1 WHERE c LIKE 'ab_'
SELECT * FROM t1 WHERE vc LIKE 'ab_' 
SELECT * FROM t1 WHERE RTRIM(c) LIKE 'ab_'
SELECT * FROM t1 WHERE c LIKE '%c'
SELECT * FROM t1 WHERE vc LIKE '%c'
SELECT * FROM t1 WHERE RTRIM(c) LIKE '%c'
```

See
discussion about MATCHES and LIKE operators in adaption guides for
more details.

## CHAR blank padding versus VARCHAR

In all database engines, `CHAR(N)` data is blank padded. This means that the
database engine fills the column value with trailing blanks when needed.

Because of the IBM
Informix SQL `VARCHAR()` limit of 255
bytes, there may be legacy code using the `CHAR` data type for larger sizes, such as
`CHAR(400)` or `CHAR(2000)`.

The problem with large `CHAR()` columns is blank padding. For example, with
`CHAR(500)`, if you store and pass the value `"abc"` in program
variables, SQL parameters, and SQL statements, you pass abc + 497 blanks, because
`CHAR` values are blank-padded.

On the other hand, `VARCHAR` types only store the actual value that was provided
(`'abc'` in our case), and trailing blanks only if explicitly provided (`'abc
'`).

In order to optimize your application, consider replacing large `CHAR()` columns
by `VARCHAR()`. The `CHAR()` type can be used to store small,
fixed-size, character string values (such as phone or credit card numbers).

Informix IDS supports the
`LVARCHAR()` type with a larger limit. Genero BDL supports this Informix SQL type.

See also [Passing CHAR parameters to functions](../09_advanced-features/0975-passing-char-parameters-to-functions.md).

## Trailing blanks in CHAR/VARCHAR program variables

In
programs, `CHAR` variables are filled with blanks,
even if the value used does not contain all spaces.

For example:

```
DEFINE c CHAR(5)
LET c = "abc"
DISPLAY c || "."  -- shows the value "abc  ." (5 chars + dot)
```

`VARCHAR` variables
are assigned with the exact value specified, with significant
trailing blanks.

For example:

```
DEFINE v VARCHAR(5)
LET v = "abc "
DISPLAY v || "."  -- shows the value "abc ." (4 chars + dot)
```

Assigning an empty string to
a `CHAR` or `VARCHAR` variable will
set the variable to `NULL`:

```
DEFINE v VARCHAR(5)
LET v = ""
IF v IS NULL THEN
   DISPLAY "is null"   -- will be displayed
END IF
```

When comparing `CHAR` or `VARCHAR` variables
in an expression, the trailing blanks are not significant:

```
DEFINE c CHAR(5)
DEFINE v1, v2 VARCHAR(5)
LET c = "abc"
LET v1 = "abc "
LET v2 = "abc  "
IF c == v1 THEN
   DISPLAY "c==v1" 
END IF
IF c == v2 THEN
   DISPLAY "c==v2" 
END IF
IF v1 == v2 THEN
   DISPLAY "v1==v2" 
END IF
```

All three messages are shown.

Additionally, when you assign a `VARCHAR` variable from a `CHAR`,
the target variable gets the trailing blanks of the `CHAR`
variable:

```
DEFINE pc CHAR(50)
DEFINE pv VARCHAR(50)
LET pc = "abc"
LET pv = pc
DISPLAY pv || "."  -- shows "abc<47 spaces>." ( 50 chars + dot )
```

To avoid this, use
the `CLIPPED` operator:

```
LET pv = pc CLIPPED
```

## Trailing blanks in SQL statement parameters

When you insert a row containing a `CHAR` variable
into a `CHAR` or `VARCHAR` column,
the database interface removes the trailing blanks to avoid
overflow problems, (insert `CHAR(100)` into `CHAR(20)`
when value is `"abc"` must work).

In
this example:

```
DEFINE c CHAR(5)
LET c = "abc"
CREATE TABLE t ( v1 CHAR(10), v2 VARCHAR(10) )
INSERT INTO tab VALUES ( c, c )
```

The value in
column v1 and v2 would be `"abc"` ( 3 chars in both
columns ).

When you insert a row containing a `VARCHAR` variable
into a `VARCHAR` column, the `VARCHAR` value
in the database gets the trailing blanks as set in the variable.
When the column is a `CHAR(N)`, the database
server fills the value with blanks so that the size of the string
is N characters.

For example:

```
DEFINE vc VARCHAR(5)
LET vc = "abc  "  -- note 2 spaces at end of string
CREATE TABLE t ( v1 CHAR(10), v2 VARCHAR(10) )
INSERT INTO tab VALUES ( vc, vc )
```

The value in column v1 would be `"abc<7 blanks>"` ( 10 chars ) and v2 would
be `"abc<2 blanks>"` (5 chars ).

## Related links

**Related concepts**  

[MATCHES and LIKE operators](1043-matches-and-like-operators.md "Use the standard LIKE operator instead of the MATCHES operator.")

[CHAR(size)](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type.")

[VARCHAR(size)](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.")

[CLIPPED](../08_language-basics/0635-clipped.md "The CLIPPED operator removes trailing blank spaces (ASCII 32) of a string expression.")
