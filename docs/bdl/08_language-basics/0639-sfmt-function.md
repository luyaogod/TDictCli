---
title: "SFMT() [function]"
source: "fgl-topics/c_fgl_operators_SFMT.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > SFMT() [function]"
type: "concept"
---

# SFMT() [function]

> The SFMT() operator replaces place holders in a string with values.

## Syntax

```
SFMT ( str-expr , param [ , param [...] ] )
```

1. str-expr is a string expression.
2. param is any valid expression used to replace parameter place holders (%*n*).
3. At least one parameter is required.

## Usage

The `SFMT()` operator can be used with parameters that will be automatically
set in the string at the position defined by parameter placeholders. The parameters used with the
`SFMT()` operator can be any valid expression. Numeric and date/time expressions are
evaluated to strings depending on the current format settings ([DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values."), [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.")).

A placeholder is a special marker in the string, that is defined by the percent character
followed by the parameter number. For example, `%4` represents the parameter
`#4`. You are allowed to use the same parameter placeholder several times in the
string. If you want to use the percent sign in the string, you must escape it with
`%%`.

Predefined placeholders can be used to insert information about last runtime system error that
occurred. Note that these are only available in the context of a runtime error trapped with a
[`WHENEVER ERROR GOTO / CALL`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") handler:

| Predefined parameter | Description |
| --- | --- |
| `%(ERRORFILE)` | Name of the module where last runtime error occurred. |
| `%(ERRORLINE)` | Line number in the module where last runtime error occurred. |
| `%(ERRNO)` | Last operating system error number. |
| `%(STRERROR)` | Last operating system error text. |

## Example

```
MAIN
  DEFINE n INTEGER
  LET n = 234
  DISPLAY SFMT("Order #%1 has been %2.",n,"deleted")
END MAIN
```

In this example, `%1` is replaced by the value of the variable
n, while `%2` is replaced by the string
"`deleted`", resulting in: `Order #234 has been
deleted.`

## Related links

**Related concepts**  

[String expressions](0597-string-expressions.md "This section covers string expression evaluation rules.")
