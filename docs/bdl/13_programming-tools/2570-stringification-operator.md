---
title: "Stringification operator"
source: "fgl-topics/c_fgl_preprocessor_007.html"
breadcrumb: "Programming tools > Source preprocessor > Stringification operator"
type: "concept"
---

# Stringification operator

> Transforms a preprocessor macro element to a string.

## Syntax

```
#param
```

1. param is a parameter of the macro

## Usage

The stringification operator `#` converts a preprocessor macro
parameter to a string.

When a macro parameter is used with a preceding `#`, it is replaced by
a string containing the literal text of the argument.

The argument is not macro expanded before the substitution.

Source: File A

```
&define disp(x) DISPLAY #x
disp(abcdef)
```

Result (fglcomp -E
output):

```
& 1 "A"
DISPLAY "abcdef"
```
