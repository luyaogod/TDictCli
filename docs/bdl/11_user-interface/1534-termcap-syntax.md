---
title: "Termcap syntax"
source: "fgl-topics/c_fgl_DynamicUI_033.html"
breadcrumb: "User interface > User interface basics > Using a text terminal > TERMCAP terminal capabilities > Termcap syntax"
type: "concept"
description: "All termcap entries contain a list of terminal names, followed by a list of terminal capabilities, in the following format: Each capability, including the last one in the entry, is followed by a colon ..."
---

# Termcap syntax

All termcap entries contain a list of terminal names, followed
by a list of terminal capabilities, in the following format:

- Each capability, including the last one in the entry, is followed
  by a colon ( : ).
- ESCAPE is specified as a backslash ( \ ) followed by the letter
  E. CTRL is specified as a caret (^). Do not use the ESCAPE or CTRL
  keys to indicate escape sequences or control characters in a termcap
  entry.
- Entries must be defined on a single logical line; a backslash
  ( \ ) appears at the end of each line that wraps to the next line.
- Comment lines begin with a hash sign ( # ).

Example: xterm terminal definition:

```
xterm|xterm terminal emulator:\
:km:mi:ms:xn:pt:\
:co#80:li#24:\
:is=\E[r\E[m\E[2J\E[H\E[?7h\E[?1;3;4;6l:\
...
```

## Terminal Names

Termcap
entries begin with one or more names for the terminal, each separated
by a vertical ( | ) bar. Any one of these names can be used for access
to the termcap entry.

## Boolean capabilities

Boolean
capabilities are two-character codes indicating whether a terminal
has a specific feature. If the boolean capability exists in the termcap
entry, the terminal has that particular feature.

For example:

```
:bs:am:
# bs backspace with CTRL-H
# am automatic margins
```

## Numeric Capabilities

Numeric
capabilities are two-character codes followed by a hash symbol ( # ) and a value.

For example:

```
:co#80:li#24:
# co number of columns in a line
# li number of lines on the screen
```

The runtime system assumes that the value is zero for any numeric
capabilities that are not listed.

## String Capabilities

String capabilities specify a sequence that can be used to perform a terminal
operation.

A string capability is a two-character code, followed
by an equal sign ( = ) and a string ending at the next delimiter (
: ).

Most termcap entries include string capabilities for clearing
the screen, arrow keys, cursor movement, underscore, function keys,
etc.

For example, this shows some string capabilities for a
Wyse 50 terminal:

```
:ce=\Et:cl=\E*:\
:nd=^L:up=^K:\
:so=\EG4:se=\EG0:
# ce=\Et clear to end of line
# cl=\E* clear the screen
# nd=^L non-destructive cursor right
# up=^K up one line
# so=\EG4 start stand-out
# se=\EG0 end stand-out
```
