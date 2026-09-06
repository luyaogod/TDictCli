---
title: "Conditional compilation"
source: "fgl-topics/c_fgl_preprocessor_011.html"
breadcrumb: "Programming tools > Source preprocessor > Conditional compilation"
type: "concept"
---

# Conditional compilation

> Integrate code lines conditionally.

## Syntax 1

```
&ifdef identifier
...
[&else
...]
&endif
```

1. identifier is a preprocessor constant.

## Syntax 2

```
&ifndef identifier
...
[&else
...]
&endif
```

1. identifier is a preprocessor constant.

## Usage

The `&ifdef` and `&ifndef` preprocessor macros can be used
to integrate code lines conditionally depending on the existence of a preprocessor constant.

The constant is defined with a `&define` or with the `-D`
option in the command line.

Even if the condition is evaluated to false, the content of the `&ifdef` block
is still scanned and tokenized. Therefore, it must be lexically correct.

Sometimes it is useful to use some code if a macro is not defined. You can use
`&ifndef`, that evaluates to true if the macro is not defined.

Source: File
A

```
&define IS_DEFINED
&ifdef IS_DEFINED
DISPLAY "The macro is defined"
&endif  /* IS_DEFINED */
```

Result (fglcomp -E
output):

```
& 1 "A"

DISPLAY "The macro is defined"
```
