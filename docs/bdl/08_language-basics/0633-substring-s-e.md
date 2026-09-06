---
title: "Substring ([s,e])"
source: "fgl-topics/c_fgl_operators_SUBSTRING.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > Substring ([s,e])"
type: "concept"
---

# Substring ([s,e])

> The [] (square brackets) operator extracts a substring from a variable.

## Syntax

> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

Single character position form:

```
variable [ pos ]
```

Start/end position form:

```
variable [ start-pos , end-pos ]
```

1. variable must be a variable defined with the `CHAR` or
   `VARCHAR` type.
2. pos defines the position of the character to be extracted.
3. start-pos defines the position of the first character of the substring to be
   extracted.
4. end-pos defines the position of the last character of the substring to be
   extracted.

## Usage

The `[]` (square brackets) notation following a [`CHAR`](0557-char-size.md "The CHAR data type is a fixed-length character string data type.") or [`VARCHAR`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") variable extracts a substring
from that character variable.

The pos, start-pos and end-pos arguments
can be expressed in bytes or characters, depending on the [length semantics](../09_advanced-features/0881-length-semantics-settings.md) used in your programs.

> **Important:**
>
> Substring expressions in SQL statements are evaluated by the database
> server. This may have a different behavior than the substring operator of the language.

Substring expression will return `NULL`, if there are no characters at the
specified positions in the source string. When the source variable is a `CHAR`, the
value is always blank-padded and substring can return spaces. When using a `VARCHAR`
type, the value may have trailing blanks or no characters at the specified position(s). When there
are no characters at the specified position(s), substring returns `NULL`. See the
code sample below using `CHAR` and `VARCHAR` variables.

## Example

```
MAIN
  DEFINE c5 CHAR(5)
  DEFINE v5 VARCHAR(5)
  LET c5 = "abc"
  LET v5 = "abc"
  DISPLAY "c5[3,3] = [",c5[3,3],"]"
  DISPLAY "v5[2,3] = [",v5[2,3],"]"
  DISPLAY "At position 4, c5 has a space and v5 has nothing:"
  DISPLAY "c5[4] = [",c5[4],"]"
  DISPLAY "c5[4] IS NULL : ", (c5[4] IS NULL)
  DISPLAY "v5[4] = [",v5[4],"]"
  DISPLAY "v5[4] IS NULL : ", (v5[4] IS NULL)
END MAIN
```

Output:

```
c5[3,3] = [c]
v5[2,3] = [bc]
At position 4, c5 has a space and v5 has nothing:
c5[4] = [ ]
c5[4] IS NULL :      0
v5[4] = []
v5[4] IS NULL :      1
```

## Related links

**Related concepts**  

[Defining the application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")
