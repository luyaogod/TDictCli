---
title: "util.Regexp.compile"
source: "fgl-topics/c_fgl_ext_util_Regexp_compile.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Regexp class > util.Regexp methods > util.Regexp.compile"
type: "concept"
---

# util.Regexp.compile

> Compiles a regular expression and returns a util.Regexp object.

## Syntax

```
util.Regexp.compile(
  expr STRING
 )
  RETURNS util.Regexp
```

1. expr is the regular expression to be compiled.

## Usage

The `util.Regexp.compile()` class method compiles the pattern passed as parameter,
and creates an object of type `util.Regexp`, that can then be used to process
strings.

If the regular expression is invalid, error [-8136](4483-genero-bdl-errors.md) is raised and the method
returns `NULL`.

Since this method returns an object that can be used in an expression, you must use a
`TRY/CATCH` block to handle exceptions. See [Understanding exceptions](../09_advanced-features/0849-understanding-exceptions.md "Exceptions are abnormal runtime events that can be trapped for control.").

Define a variable of type `util.Regexp` to hold a reference to the object returned
by the `compile()` method.

> **Note:**
>
> When using single or double quoted string literals, backslash is
> interpreted as escape character. Consider using [back
> quotes](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.") as delimiters for regular expressions strings, and use backslash characters directly
> as required by the regexp syntax ( `` `\b` `` instead of `"\\b"` or
> `'\\b'` )

See [Regular expression patterns](3545-regular-expression-patterns.md "Summary of util.Regexp.compile() pattern syntax.") for details about the regular
expression patterns accepted by the compile() method.

## Example

```
IMPORT util
MAIN
    DEFINE re util.Regexp
    TRY
        LET re = util.Regexp.compile(`[abcdef]`)
    CATCH
        DISPLAY "Invalid regular expression!"
        EXIT PROGRAM 1
    END TRY
END MAIN
```

## Related links

**Related concepts**  

[Text literals](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.")
