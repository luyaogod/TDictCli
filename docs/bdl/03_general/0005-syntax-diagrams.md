---
title: "Syntax diagrams"
source: "common-topics/c_common_DocConv_002.html"
breadcrumb: "General > Documentation conventions > Syntax diagrams"
type: "concept"
---

# Syntax diagrams

> A syntax diagram describes the context-free, abstract grammar of a product function.

A syntax diagram for example describes a language instruction, the structure of a configuration
file, or the options of command-line tool.

The following rules are used in a syntax diagram:

- Invariable syntax elements (keywords) are written in `fixed font`.
- Language keywords are in uppercase, like `INPUT BY NAME`.
- Variable syntax elements are written in italics.
- Wildcard characters are used to indicate syntax elements that can either repeat, be mandatory or
  optional.

| Wildcard characters notation | Description |
| --- | --- |
| `[ element ]` | Square brackets indicate an optional element in the syntax. For example:IS [NOT] NULL |
| `[ element-1 \| element-2 ... ]` | Square brackets with pipe separator indicate an optional element to be selected from the list. For example:ORDER BY colname [ASC\|DESC] |
| `{ element-1 \| element-2 ... }` | Curly brackets with pipe separator indicate a mandatory element to be selected from the list. For example:STRETCH={NONE\|X\|Y\|BOTH} |
| element `[...]` | A sign made of square brackets with three dots indicate that the previous element can appear more than once. For example:child-element [...] |
| element `[,...]`element `[:...]`element `[;...]`element `[\|...]` | A sign made of square brackets including a separator character (comma, colon, semi-colon, pipe) followed by three dots indicate that the previous element can appear more than once, and must be separated by the specified character. For example:SELECT colname [,...] |

The `[ ]` square braces and `{ }` curly braces notations can be
combined and nested, for example (dummy language
element):

```
[ integer-expression { LINES | COLUMNS } ]
```

Example (dummy language
instruction):

```
CREATE [ PRIVATE | PUBLIC ] SHAPE object-name
    WITH integer-expression { ROWS | COLUMNS }
    [ USING option [,...] ]
```

In
this syntax diagram:

- The `CREATE` keyword can be followed by `PRIVATE` or
  `PUBLIC` keywords, but these are not mandatory (square brackets are used)
- object-name is a variable syntax element, to define the name of the object to
  be created.
- integer-expression indicates an expression that must evaluate to an whole
  number.
- The integer-expression must be followed by the mandatory
  `ROWS` or `COLUMNS` keyword (curly braces are used)
- The whole `USING` clause is not required (square braces are used)
- If the `USING` clause is specified, it must be followed by a comma-separated list
  of option elements.

For example, instructions using the above syntax could have the following
forms:

```
CREATE SHAPE shape1
   WITH 5 ROWS

CREATE PRIVATE SHAPE shape2
   WITH 25 COLUMNS
   USING fragmentation, autoload
```
