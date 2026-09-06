---
title: "VARCHAR(size)"
source: "fgl-topics/c_fgl_datatypes_VARCHAR.html"
breadcrumb: "Language basics > Primitive Data types > VARCHAR(size)"
type: "concept"
---

# VARCHAR(size)

> The VARCHAR data type is a variable-length character string data type, with a maximum size.

## Syntax

```
VARCHAR [ ( size [,reserve] ) ]
```

1. size defines the maximum length of the character string, in byte or char
   units (depending on the character length semantics)
2. The maximum size of a `VARCHAR` type is 65534.
3. When no size is specified, it defaults to 1.
4. reserve is ignored; Its inclusion in the syntax
   is permitted for compatibility with the SQL data type.

## Usage

The `VARCHAR` type is typically used to store variable-length
character strings such as names, addresses and comments.

The size can be expressed in bytes or characters, depending on the length
semantics used in programs. For more details about character length semantics, see
[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md).

When size is not specified, the default length is 1.

`VARCHAR` variables are initialized to [NULL](0572-null.md "The NULL constant defines a non-value.")
in functions, modules and globals.

[Text literals](0588-text-literals.md "Text literals define a character string in an expression.") can be assigned to character
string variables:

```
MAIN
  DEFINE c VARCHAR(10)
  LET c = "abcdef"
END MAIN
```

`VARCHAR` variables store trailing blanks (trailing blanks are displayed
or printed in reports, and stored in database columns):

```
MAIN
  DEFINE vc VARCHAR(10)
  LET vc = "abc  "       -- a b c + 2 whitespaces
  DISPLAY "[", vc ,"]"   -- displays [abc  ]
END MAIN
```

Trailing blanks of a `VARCHAR` value are not significant in
comparisons:

```
MAIN
  DEFINE vc VARCHAR(10)
  LET vc = "abc  "        -- a b c + 2 whitespaces
  IF vc == "abc " THEN    -- evaluates to TRUE
     DISPLAY "equals"
  END IF
END MAIN
```

Numeric and date-time values can be directly assigned the character strings:

```
MAIN
  DEFINE vc VARCHAR(50), da DATE, dec DECIMAL(10,2)
  LET da = TODAY
  LET dec = 345.12
  LET vc = da, " : ", dec
END MAIN
```

When you insert character data from `VARCHAR` variables
into `VARCHAR` columns in a database table, the trailing
blanks are kept. Likewise, when you fetch `VARCHAR` column
values into `VARCHAR` variables, trailing blanks are
kept.

```
MAIN
  DEFINE vc VARCHAR(10)
  DATABASE test1
  CREATE TABLE table1 ( k INT, x VARCHAR(10) )
  LET vc = "abc  "         -- two trailing blanks
  INSERT INTO table1 VALUES ( 1, vc )
  SELECT x INTO vc FROM table1 WHERE k = 1
  DISPLAY "[", vc ,"]"     -- displays [abc  ]
END MAIN
```

In SQL statements, the behavior of the comparison operators when using `VARCHAR`
values differs from one database to the other. IBM®
Informix® is ignoring trailing blanks, but most other
databases take trailing blanks of `VARCHAR` values into account. For more details,
see [SQL portability](../10_sql-support/1000-sql-portability.md "Writing portable SQL is mandatory, to support different kind of database servers.").

Character string manipulation with `CHAR`,
`VARCHAR` and `STRING` types can have a cost when accessing
parts of large strings, when using UTF-8 with char length semantics. Consider using
`base.StringBuffer` objects when doing heavy string manipulations. For more
details, read [Manipulating character strings](../09_advanced-features/0976-manipulating-character-strings.md).

## Related links

**Related concepts**  

[CHAR(size)](0557-char-size.md "The CHAR data type is a fixed-length character string data type.")

[STRING](0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.")
