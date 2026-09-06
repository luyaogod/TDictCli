---
title: "undisplay"
source: "fgl-topics/c_fgl_Debugger_undisplay.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > undisplay"
type: "concept"
---

# undisplay

> The undisplay command cancels expressions to be displayed when the program execution stops.

## Syntax

```
undisplay itemnum
```

1. itemnum is the number of the expressions for which the display is
   canceled.

## Usage

When the [`display`](2595-display.md "The display command displays the specified expression's value each time program execution stops.") command is
used, each expression displayed is assigned an item number.

The `undisplay` command allows you to remove expressions from the list to be
displayed, using the item number to specify the expression to be removed.

## Example

```
(fgldb) step 
2: i = 2 
1: a = 20 
9     FOR  i = 1 TO 10
(fgldb) undisplay 2
(fgldb) step 
1: a = 20 
10   LET cont = TRUE
(fgldb)
```
