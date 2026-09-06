---
title: "ptype"
source: "fgl-topics/c_fgl_Debugger_ptype.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > ptype"
type: "concept"
---

# ptype

> The ptype command prints the data type or structure of a variable.

## Syntax

```
ptype variable-name
```

1. variable-name is the name of the variable.

## Usage

The ptype command displays the type of the variable passed as
parameter.

## Example

```
MAIN
    DEFINE rec RECORD
                 pkey INTEGER,
                 name VARCHAR(50)
           END RECORD

    LET rec.pkey = 101
    LET rec.name = "Mike TORN"
    DISPLAY rec.*

END MAIN
```

```
$ fglcomp -M prog.4gl && fglrun -d prog.42m

(fgldb) break 9
Breakpoint 1 at 0x00000000: file prog.4gl, line 9.

(fgldb) run
Breakpoint 1, main() at prog.4gl:9
   6
   7         LET rec.pkey = 101
   8         LET rec.name = "Mike TORN"
-> 9         DISPLAY rec.*
   10
   11    END MAIN
   12

(fgldb) ptype rec
type = RECORD
    pkey INTEGER,
    name VARCHAR(50)
END RECORD
```
