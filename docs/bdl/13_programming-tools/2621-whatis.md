---
title: "whatis"
source: "fgl-topics/c_fgl_Debugger_whatis.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > whatis"
type: "concept"
---

# whatis

> The whatis command prints the data type of a variable.

## Syntax

```
whatis variable-name
```

1. variable-name is the name of the variable.

## Usage

The `whatis` command can be used to show the data type of a program variable.

The program variable must exist in the current scope.

## Example

```
(fgldb) run 
Breakpoint 1, main() at t.4gl:4
4         LET i = 1
(fgldb) whatis i
type = INTEGER
(fgldb)
```
