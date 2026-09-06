---
title: "CHAR(size)"
source: "fgl-topics/c_fgl_datatypes_CHAR.html"
breadcrumb: "Language basics > Primitive Data types > CHAR(size)"
type: "concept"
---

# CHAR(size)

> The CHAR data type is a fixed-length character string data type.

## Syntax

```
CHAR[ACTER] [ (size) ]
```

1. size defines the maximum length of the character string, in byte or char
   units (depending on the character length semantics)
2. The maximum size of a `CHAR` type is 65534.
3. If no size is specified, it defaults to 1.

## Usage

The `CHAR` type is typically used to store fixed-length character
strings such as short codes (XB124), phone numbers (650-23-2345), vehicle
identification numbers.

`CHAR` and `CHARACTER` are synonyms.

The size can be expressed in bytes or characters, depending on the length
semantics used in programs. For more details about character length semantics, see
[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md).

When size is not specified, the default length is 1.

`CHAR` variables are initialized to [NULL](0572-null.md "The NULL constant defines a non-value.")
in functions, modules and globals.

[Text literals](0588-text-literals.md "Text literals define a character string in an expression.") can be assigned to character
string variables:

```
MAIN
  DEFINE c CHAR(10)
  LET c = "abcdef"
END MAIN
```

When assigning a non-NULL value, `CHAR` variables are always
blank-padded:

```
MAIN
  DEFINE c CHAR(10)
  LET c = "abcdef"
  DISPLAY "[", c ,"]"   -- displays [abcdef    ]
END MAIN
```

Trailing blanks of a `CHAR` value are not significant in
comparisons:

```
MAIN
  DEFINE c CHAR(5)
  LET c = "abc"
  IF c == "abc" THEN    -- evaluates to TRUE
     DISPLAY "equals"
  END IF
END MAIN
```

Numeric and date-time values can be directly assigned the character
strings:

```
MAIN
  DEFINE c CHAR(50), da DATE, dec DECIMAL(10,2)
  LET da = TODAY
  LET dec = 345.12
  LET c = da, " : ", dec
END MAIN
```

When you insert character data from `CHAR` variables into
`CHAR` columns in a database table, the column-value is blank-padded
to the size of the column. Likewise, when you fetch `CHAR` column
values into `CHAR` variables, the program variable is blank-padded to
the size of the variable.

```
MAIN
  DEFINE c CHAR(10)
  DATABASE test1
  CREATE TABLE table1 ( k INT, x CHAR(10) )
  LET c = "abc"
  INSERT INTO table1 VALUES ( 1, c )
  SELECT x INTO c FROM table1 WHERE k = 1
  DISPLAY "[", c ,"]"     -- displays [abc  ]
END MAIN
```

In SQL statements, the behavior of the comparison operators when using
`CHAR` values may vary from one database to the other. However, most
database engines ignore trailing blanks when compating `CHAR` values.
For more details, see [SQL portability](../10_sql-support/1000-sql-portability.md "Writing portable SQL is mandatory, to support different kind of database servers.").

Character string manipulation with `CHAR`,
`VARCHAR` and `STRING` types can have a cost when accessing
parts of large strings, when using UTF-8 with char length semantics. Consider using
`base.StringBuffer` objects when doing heavy string manipulations. For more
details, read [Manipulating character strings](../09_advanced-features/0976-manipulating-character-strings.md).

## Related links

**Related concepts**  

[VARCHAR(size)](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.")

[STRING](0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.")
