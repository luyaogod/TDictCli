---
title: "Concatenation operator"
source: "fgl-topics/c_fgl_preprocessor_008.html"
breadcrumb: "Programming tools > Source preprocessor > Concatenation operator"
type: "concept"
---

# Concatenation operator

> Concatenates two parameters of a preprocessor macro.

## Syntax

```
token1 ## token2
```

1. token1 is a parameter of the macro or a simple token.
2. token2 is a parameter of the macro or a simple token.

## Usage

The double-hash operator `##` can be used to merge two tokens while expanding a
macro and create a single token.

All tokens can not be merged. Usually these tokens are identifiers, or numbers.

The concatenation result produces an identifier.

Source: File
A

```
&define COMMAND(NAME) #NAME, NAME ## _command 
COMMAND(quit)
```

Result (fglcomp -E
output):

```
& 1 "A"

"quit", quit_command
```
