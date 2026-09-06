---
title: "Boolean expressions in forms"
source: "fgl-topics/c_fgl_FormSpecFiles_Boolean_expressions_in_form_files.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Boolean expressions in forms"
type: "concept"
---

# Boolean expressions in forms

> Some form item definitions can include boolean expressions with a form file specific syntax.

## Syntax

```
[(] bool-expr {AND|OR} bool-expr [)] [...]
```

where bool-expr
is:

```
[NOT]
{ field-tag
    {  =  expression
    |  <> expression
    |  != expression
    |  <= expression
    |  >= expression
    |  <  expression
    |  >  expression
    |  IS [NOT] NULL
    |  [NOT] BETWEEN expression AND expression
    |  [NOT] MATCHES "string"
    |  [NOT] LIKE "string"
    }
}
```

1. field-tag is the name of the current field tag in form line with the
   attribute definition.
2. expression can be the a character string, numeric or date-time literal.

## Usage

Some form specification file attributes such as [`COLOR WHERE`](1767-color-where-attribute.md "The COLOR WHERE attribute defines a condition to set the foreground color dynamically.")
require a boolean expression. These boolean expressions are different from the
language boolean expressions, and have a limited syntax which is specific to
the form files.

When a field-tag is used in the boolean expression, the runtime
system replaces field-tag at runtime with the current value in
the screen field and evaluates the expression.

## Example

```
EDIT f001 = item.price,
     COLOR=RED
          WHERE f001 >= 100 AND f001 < 1000;
```

## Related links

**Related concepts**  

[Literals](../08_language-basics/0585-literals.md "Describes the syntax of literals (constant values) to be used in sources.")
