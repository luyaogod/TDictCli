---
title: "ASCII"
source: "fgl-topics/c_fgl_operators_ASCII.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > ASCII"
type: "concept"
---

# ASCII

> The ASCII operator produces a character from its ordinal value.

## Syntax

```
ASCII int-expr
```

1. int-expr is an integer expression. The range of possible values depends on
   the current application locale and length semantics.

## Usage

The `ASCII` operator returns the character corresponding to the ASCII code or
16bit UNICODE code point passed as a parameter, in the current character encoding defined by the
[application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."), and depends on the current [length semantics](../09_advanced-features/0881-length-semantics-settings.md).
> **Note:**
>
> `ASCII` is a real operator in the Genero BDL language: The right operand is an
> integer expression that does not need to be enclosed in parentheses. However, you can find legacy
> source code using the `ASCII` operator as a function, with parentheses around the
> character code: `ASCII(expr)`. This is legal, because parentheses
> will be interpreted as part of the integer expression, just like in the `(34)+(67)`
> arithmetic expression. Note that this is not the case of the [`ORD()`](0636-ord-function.md "The ORD() operator returns the code point of a character in the current locale.") syntax element where parentheses are mandatory.

The range of possible values of the int-expr parameter depends on the current
application locale / character set:

- For single byte encodings (like ISO8859-15), the argument must be in the range of 0 to 255.
- For UTF-8, using char length semantics, the argument must be any valid 16bit code point (in the
  range 0-65535).
- For any other locale setting (any multibyte character set, or UTF-8 using byte length
  semantics), the argument must be in the range 0 to 127.

The `ASCII` function can be also used to produce special characters such as escape
(`ASCII 27`), newline (`ASCII 10`), horizontal tab (`ASCII
9`).

When the argument is zero, `ASCII` has a different behavior, depending on the
context:

- `ASCII 0` only displays the `NUL/0` null character within the
  [`PRINT`](../12_reports/2491-print.md "Formats and prints a row of data in a report routine.") statement.
- If you specify `ASCII 0` in other contexts, it returns a blank space.

If the character encoding is single byte, or when using UTF-8 and char length semantics,
`ASCII` is the inverse of [`ORD()`](0636-ord-function.md "The ORD() operator returns the code point of a character in the current locale.").

## Example

```
MAIN
  DISPLAY ASCII 65, ASCII 66, ASCII 7
END MAIN
```

## Related links

**Related concepts**  

[Defining the application locale](../09_advanced-features/0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")

[ORD() [function]](0636-ord-function.md "The ORD() operator returns the code point of a character in the current locale.")
