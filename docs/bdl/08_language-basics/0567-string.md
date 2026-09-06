---
title: "STRING"
source: "fgl-topics/c_fgl_datatypes_STRING.html"
breadcrumb: "Language basics > Primitive Data types > STRING"
type: "concept"
---

# STRING

> The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.

## Syntax

```
STRING
```

## Usage

The `STRING` data type is typically used to implement utility functions
manipulating character string with unknown size, and in some special cases, in SQL
statements.

`STRING` variables are initialized to [`NULL`](0572-null.md "The NULL constant defines a non-value.") in functions, modules and
globals.

The behavior of a `STRING` variable is similar to the `VARCHAR`
data type, except that there is no theoretical size limit.

`STRING` variables can be initialized from [string literals](0588-text-literals.md "Text literals define a character string in an expression."):

```
MAIN
    DEFINE s STRING
    LET s = "abcdef"
END MAIN
```

Variables declared with the `STRING` data type can be used to call
STRING-type methods such as `getLength()` or `toUpperCase()`.
For more details, read [STRING data type as class](../15_library-reference/2915-string-data-type-as-class.md "The STRING primitive data type provides a set of utility methods to manipulate character strings."):

```
MAIN
    DEFINE s STRING
    LET s = "abc"
    DISPLAY s.toUpperCase()
END MAIN
```

`STRING` variables have significant trailing blanks (i.e. `"abc
"` is different from `"abc"`). However, in comparisons, trailing
blanks do not matter:

```
MAIN
    DEFINE s STRING
    LET s = "abc  " -- a b c + 2 whitespaces
    DISPLAY "1: s.length:", s.getLength()
    DISPLAY "[", s, "]" -- displays "[abc  ]"
    DISPLAY IIF(s=="abc","Equals",NULL)
END MAIN
```

Unlike `CHAR` and `VARCHAR`, a `STRING`
can hold a value of zero length without being `NULL`. For example, if you
trim a string variable with the [`trim()`](../15_library-reference/2932-string-trim.md "Removes leading and trailing blank space (ASCII 32) characters.") method and if
the original value is a set of blank characters, the result is an empty string. But
testing the variable with the `IS NULL` operator will evaluate to
`FALSE`. Using a `VARCHAR` with the `CLIPPED`
operator would give a `NULL` string in this case:

```
MAIN
    DEFINE s STRING
    LET s = "     " -- 5 spaces
    LET s = s.trim()
    DISPLAY "s = [", s, "] len=", s.getLength()
    DISPLAY IIF(s IS NULL, "NULL", "not NULL")
END MAIN
```

Output:

```
s = [] len=          0
not NULL
```

`STRING` typed variables can be used in some special cases to hold SQL
character string data, when the size of the SQL data string is not known (string
expressions, large strings like JSON documents). In order to store character string
data stored in a database, consider using the `CHAR` or
`VARCHAR` types instead of `STRING`.

In `STRING` methods, positions and length parameters (or return values)
can be expressed in bytes or characters, depending on the length semantics used in programs.
For more details, read [Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md).

Character string manipulation with `CHAR`,
`VARCHAR` and `STRING` types can have a cost when accessing
parts of large strings, when using UTF-8 with char length semantics. Consider using
`base.StringBuffer` objects when doing heavy string manipulations. For more
details, read [Manipulating character strings](../09_advanced-features/0976-manipulating-character-strings.md).

## Related links

**Related concepts**  

[STRING data type methods](../15_library-reference/2916-string-data-type-methods.md "STRING data type methods")

[CHAR(size)](0557-char-size.md "The CHAR data type is a fixed-length character string data type.")

[VARCHAR(size)](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.")
