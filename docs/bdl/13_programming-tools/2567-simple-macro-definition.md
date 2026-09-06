---
title: "Simple macro definition"
source: "fgl-topics/c_fgl_preprocessor_005.html"
breadcrumb: "Programming tools > Source preprocessor > Simple macro definition"
type: "concept"
---

# Simple macro definition

> A simple macro is identified by its name and body.

## Syntax

```
&define identifier body
```

1. identifier is the name of the macro. Any valid identifier can be used.
2. body is any sequence of tokens until the end of the line.

After substitution, the macro definition is replaced with blank lines.

A preprocessor directive must start at the beginning of the line. The `&`
ampersand character can be preceded by blanks.

## Usage

As the preprocessor scans the text, it substitutes the macro body for the name identifier.

The following example show macro substitution with 2 simple macros:

Source: File A

```
&define MAX_TEST 12
&define HW "Hello world"

MAIN
  DEFINE i INTEGER
  FOR i=1 TO MAX_TEST 
    DISPLAY HW
  END FOR
END MAIN
```

Result (fglcomp -E
output):

```
& 1 "A"

MAIN
  DEFINE i INTEGER
  FOR i=1 TO 12
    DISPLAY "Hello world"
  END FOR
END MAIN
```

The macro definition can be continued on multiple lines, but when the macro is expanded,
it is joined to a single line as follows:

Source: File A

```
&define TABLE_VALUES 1, \
                     2, \
                     3
DISPLAY TABLE_VALUES
```

Result (fglcomp -E
output):

```
& 1 "A"
 
 
 
DISPLAY 1, 2, 3
```

The source file is processed sequentially, so a macro takes effect at the place it has
been written:

Source: File A

```
DISPLAY X
&define X "Hello"
DISPLAY X
```

Result (fglcomp -E
output):

```
& 1 "A"
DISPLAY X

DISPLAY "Hello"
```

The macro body is expanded only when the macro is applied:

Source: File A

```
&define AA BB
&define BB 12
DISPLAY AA
```

Result (fglcomp -E
output):

```
& 1 "A"
 
 
DISPLAY 12
```

- AA is first expanded to BB.
- The text is re-scanned and BB is expanded to 12.
- When the macro AA is defined, BB is not known yet; but it is known when the macro AA is used.

In order to prevent infinite recursion, a macro cannot be expanded recursively.

Source: File A

```
&define A B
&define B A
&define C C
A C
```

Result (fglcomp -E
output):

```
& 1 "A"

A C
```

- A is first expanded to B.
- B is expanded to A.
- A is not expanded again as it appears in its own expansion.
- C expands to C and can not be expanded further.

It is also possible to define a macro with the `-D` command line option of
compilers.
