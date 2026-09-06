---
title: "WORDWRAP"
source: "fgl-topics/c_fgl_reports_WORDWRAP.html"
breadcrumb: "Reports > Report operators > WORDWRAP"
type: "concept"
---

# WORDWRAP

> Splits a character string to match a given margin limit.

## Syntax

```
WORDWRAP [ RIGHT MARGIN position ]
```

1. position defines the temporary right margin, as a number of characters,
   counting from the left.

## Usage

The `WORDWRAP` operator automatically wraps successive segments of long
character strings onto successive lines of report output. Any string value that is too long
to fit between the current position and the right margin is divided into segments and
displayed between temporary margins:

- The current character position becomes the temporary left margin.
- Unless you specify `RIGHT MARGIN`, the right margin defaults to 132, or
  to the size value from the `RIGHT MARGIN` clause of the `OUTPUT` section or
  `START REPORT`
  instruction.

Specify `WORDWRAP RIGHT MARGIN integer` to set a temporary right
margin as a number of characters, counting from the left edge of the page. This value cannot
be smaller than the current character position or greater than right margin defined for the
report. The current character position becomes the temporary left margin. These temporary
values override the specified or default left and right margins of the report.

After the `PRINT` statement
has executed, any explicit or default margins defined in the `RIGHT MARGIN`
clause of the `OUTPUT` section or `START REPORT` instruction
are restored.

The following `PRINT` statement specifies a temporary left margin in
column 10 and a temporary right margin in column 70 to display the character string that
is stored in the variable called `mynovel`:

```
  PRINT COLUMN 10, mynovel WORDWRAP RIGHT MARGIN 70
```

The data string can include printable ASCII characters. It can also include the TAB
(ASCII 9), LINEFEED (ASCII 10), and ENTER (ASCII 13) characters to partition the string
into words that consist of substrings of other printable characters. Other non-printable
characters might cause runtime errors. If the data string cannot fit between the margins
of the current line, the report engine breaks the line at a word division, and pads the
line with blanks at the right.

From left to right, the report engine expands any TAB character to enough blank spaces
to reach the next tab stop. By default, tab stops are in every eighth column, beginning
at the left-hand edge of the page. If the next tab stop or a string of blank characters
extends beyond the right margin, the report engine takes these actions:

1. Prints blank characters only to the right margin.
2. Discards any remaining blanks from the blank string or tab.
3. Starts a new line at the temporary left margin.
4. Processes the next word.

The report engine starts a new line when a word plus the next blank space cannot fit
on the current line. If all words are separated by a single space, this action creates
an even left margin. The following rules are applied (in descending order of precedence)
to the portion of the data string within the right margin:

- Break at any LINEFEED, or ENTER, or LINEFEED, ENTER pair.
- Break at the last blank (ASCII 32) or TAB character before the
  right margin.
- Break at the right margin, if no character farther to the left is a space, ENTER,
  TAB, or LINEFEED character.

The report engine maintains page discipline under the `WORDWRAP` option.
If the string is too long for the current page, the report engine executes the statements
in any page trailer and header control blocks before continuing output onto a new page.

For Japanese locales, a suitable break can also be made between the Japanese characters.
However, certain characters must not begin a new line, and some characters must not end a
line. This convention creates the need for KINSOKU processing, whose purpose is to format
the line properly, without any prohibited word at the beginning or ending of a line.

Reports use the wrap-down method for WORDWRAP and KINSOKU processing. The wrap-down
method forces down to the next line characters that are prohibited from ending a line.
A character that precedes another that is prohibited from beginning a line can also wrap
down to the next line. Characters that are prohibited from beginning or ending a line must
be listed in the locale. The runtime system tests for prohibited characters at the
beginning and ending of a line, testing the first and last visible characters. The KINSOKU
processing only happens once for each line. That is, no further KINSOKU processing occurs,
even if prohibited characters are still on the same line after the first KINSOKU processing.
