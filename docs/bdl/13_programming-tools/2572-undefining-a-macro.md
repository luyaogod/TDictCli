---
title: "Undefining a macro"
source: "fgl-topics/c_fgl_preprocessor_010.html"
breadcrumb: "Programming tools > Source preprocessor > Undefining a macro"
type: "concept"
---

# Undefining a macro

> Undefines a preprocessor macro.

## Syntax

```
&undef identifier
```

1. identifier is a preprocessor constant.

## Usage

If a macro is redefined without having been undefined previously, the preprocessor issues a
warning and replaces the existing definition with the new one. First undefine a macro with the
`&undef` directive.

Source: File A

```
&define HELLO "hello"
DISPLAY HELLO
&undef HELLO
DISPLAY HELLO
```

Result (fglcomp -E
output):

```
& 1 "A"

DISPLAY "hello"
DISPLAY HELLO
```

It is also possible to undefine a macro with the `-U` command line option of
compilers. However, predefined macros cannot be undefined with this option.
