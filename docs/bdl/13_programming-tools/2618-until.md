---
title: "until"
source: "fgl-topics/c_fgl_Debugger_until.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > until"
type: "concept"
---

# until

> The until command continues running the program until the specified location is reached.

## Syntax

```
until [ location ]
```

where location
is:

```
{ [module.]function
| [module:]line
}
```

1. module is the name of a source file, without extension.
2. function is a function name.
3. line is a source code line.

## Usage

The `until` command continues running your program until either the specified
location is reached, or the current stack frame returns.

This command can be used to avoid stepping through a loop more than once.

## Example

```
(fgldb) until add_customer()
```
