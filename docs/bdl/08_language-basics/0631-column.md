---
title: "COLUMN"
source: "fgl-topics/c_fgl_operators_COLUMN.html"
breadcrumb: "Language basics > Operators > List of expression elements > Character string operators > COLUMN"
type: "concept"
---

# COLUMN

> The COLUMN operator generates blanks.

## Syntax

```
COLUMN position
```

1. position is the column position (starts at 1).

## Usage

The `COLUMN` operators generates space characters up to the given position.

The `COLUMN` operator is typically used in report routines to align data in [`PRINT`](../12_reports/2491-print.md "Formats and prints a row of data in a report routine.") statements and move the character
position forward within the current line.

This operator makes sense when used in an expression with the comma append operator: Spaces will
be generated depending on the number of characters that have been used in the expression, before the
`COLUMN` operator.

The `COLUMN` operator is aware of `\n` characters in the string
preceding the `COLUMN` keyword. If the preceding string contains a
`\n`, the position count starts at the first character after the last
`\n`:

```
MAIN
    DISPLAY "line 1", COLUMN 20, ".\n",
            "line 2", COLUMN 20, ".\n",
            "line 3", COLUMN 20, "."
END MAIN
```

Shows:

```
line 1             .
line 2             .
line 3             .
```

The `COLUMN` operator can be used outside report routines, in order to align data
to be displayed with a proportional font, typically in a TUI context. For example, the next lines
will always [`DISPLAY`](../11_user-interface/1878-display-to-stdout.md "The DISPLAY instruction displays text in line mode to the standard output channel.") the
content of the `lastname` variable starting from column 30 of the terminal, no
matters the number of characters contained in the `firstname` variable. The example
defines [`VARCHAR`](0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") variables, since
[`CHAR`](0557-char-size.md "The CHAR data type is a fixed-length character string data type.") variables are blank-padded, we
would need to use the [`CLIPPED`](0635-clipped.md "The CLIPPED operator removes trailing blank spaces (ASCII 32) of a string expression.")
operator:

```
DEFINE firstname, lastname VARCHAR(50)
DISPLAY firstname, COLUMN(30), lasttname
```

The position operand must be a non-negative integer that specifies a character
position offset (from the left margin) no greater than the line width (that is, no greater than the
difference (right margin - left margin). This designation moves the character position to a
left-offset, where 1 is the first position after the left margin. If current position is greater
than the operand, the `COLUMN` specification is ignored.

## Example

```
PAGE HEADER
  PRINT "Number", COLUMN 12,"Name", COLUMN 35,"Location"
ON EVERY ROW
  PRINT customer_num, COLUMN 12,cust_fname, COLUMN 35,cust_city
```

## Related links

**Related concepts**  

[CLIPPED](0635-clipped.md "The CLIPPED operator removes trailing blank spaces (ASCII 32) of a string expression.")
