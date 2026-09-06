---
title: "Input length of form fields"
source: "fgl-topics/c_fgl_prog_dialogs_input_length.html"
breadcrumb: "User interface > User interface programming > Input fields > Input length of form fields"
type: "concept"
---

# Input length of form fields

> Field input length defines the amount of characters the user can type in a form field.

## Input length basics

The field input length is used by interactive instructions to limit the
size of the data that can be entered by the user.

The field input length also matters when displaying a program variable to a form field with the
`DISPLAY TO` or `DISPLAY BY NAME` instruction, that may truncate the
text resulting from the data conversion.

The field input length is also used to define the input limit when using the [`AUTONEXT`](1763-autonext-attribute.md "The AUTONEXT attribute forces the focus to automatically leave the current field when completed.") attribute.

For non-character types, if the resulting text does not fit into the input length, the field will
show `*` stars to indicate an overflow.

## Input length definition

The field input length is defined from:

1. The data type of the program variable bound to the field by the interactive instruction.
2. The size (number of cells) of the form item tag.
3. The usage of the `SCROLL` attribute for fields using character string types.

## Trailing blanks in input fields

When the user enters data in a form field, the runtime system automatically truncates the
trailing blanks.

The resulting value in the corresponding program variable or in the field input buffer ([`DIALOG.getFieldBuffer()`](../15_library-reference/3197-ui-dialog-getfieldbuffer.md "Returns the input buffer of the specified field.")) is
right-trimmed.

## Length semantics and input length

Character string input limit depends also from the [Byte
Length Semantics (BLS) or Char Length Semantics (CLS)](../09_advanced-features/0881-length-semantics-settings.md) usage.

In this section, the term "character width" relates to the horizontal sizes of
characters, especially in a fixed font size. This concept is also known as "fullwidth"
versus "halfwidth" characters. Typically, the width of a Chinese logogram is twice the
width of characters of occidental languages such as A,B,C. For more details, see [Character size unit and length semantics](../09_advanced-features/0874-character-size-unit-and-length-semantics.md).

**Using Byte Length Semantics**

When using BLS (the default), the input length represents the maximum number of bytes that the
field can hold, in the current character set of the runtime system.

For example:

- When using a Single Byte Character Set (SBCS) like ISO-8859-15 (and BLS): Each characters of
  this codeset uses one byte (and has a width of 1). A field with an input length of 5 cells can hold
  5 characters of this codeset.
- When using a Chinese BIG5 encoding (and BLS): Latin characters (a,b,c) use one byte each, while
  Chinese characters use 2 bytes. A field with an input length of 5 cells can hold:
  - 5 Latin characters (`5x1B=5B`),
  - 3 Latin characters and 1 Chinese character (`3x1B+1x2B=5B`),
  - 1 Latin charater and 2 Chinese characters (`1x1B+2x2B=5B`),
  - it cannot hold 3 Chinese characters (`3x2B=6B > 5B`).
- When using UTF-8 and BLS ([not recommended, use CLS with
  UTF-8](../09_advanced-features/0881-length-semantics-settings.md)): ASCII characters (like "e") use one byte, Latin acute characters (like "é") use 2
  bytes and Chinese characters use 3 bytes. A field with an input length of 5 cells can hold:
  - 5 ASCII characters (`5x1B=5B`),
  - 1 ASCII character and 2 Latin acute characters (`1x1B+2x2B=5B`),
  - 2 ASCII characters and 1 Chinese character (`2x1B+1x3B=5B`),
  - it cannot hold 3 Chinese characters (`3x3B=9B > 5B`).

**Using Char Length Semantics**

When using character length semantics ([FGL\_LENGTH\_SEMANTICS=CHAR](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs.")), the input length is expressed in character width. The runtime
system will truncate the entered text by computing the total width of the string. Trailing
characters that do not fit into the field input length (interpreted as a maximum width) will be
rejected by the runtime and an input error will be displayed to the user.

For example:

- When using UTF-8 and CLS: ASCII characters (like "e") use one byte, Latin acute characters (like
  "é") use 2 bytes and Chinese characters use 3 bytes. A field with an input length of 5 cells can hold:
  - 5 Latin characters (`5x1W=5W`),
  - 3 Latin characters and one Chinese character (`3x1W+1x2W=5W`),
  - 2 Chinese characters (`2x2W=4W < 5W`),
  - it cannot hold 3 Chinese characters (`3x2W=6W > 5W`).

## Field width definition

The input length is by default defined by the width of the field item tag in the
`LAYOUT` section. The width of a field item tag is defined by the number of cell
positions used between the square brackets:

```
LAYOUT
GRID
{
   123
  [f1 ]    -- width = 3 cells
   123456
  [f2    ] -- width = 6 cells
```

As a general rule, forms must define fields that can hold all possible values that the
corresponding program variable can contain. For example, for a `DATE` data type, an
`EDIT` field must be defined with 10 cells, to hold date values in the format
`DD/MM/YYYY`.

If the program variable is defined with a numeric data type like `INTEGER` or
`DECIMAL`, the input length is defined by the form field width. For example, if the
form item tag defines a width of 5 cells and is bound to an `INTEGER` variable by the
input dialog, even if the integer variable can hold larger values, the user can only enter 5 digits
or the negative sign and 4 digits. As result, the value range will be -9999 to 99999.

If the program variable is defined with a `DATE`, `DATETIME` or
`INTERVAL` data type, the input length is defined by the form field width. The user
can potentially enter any kind of characters. However, when the date/time field is checked by the
dialog instruction, it must represent a valid date/time value.

If the program variable is defined with character data type such as `CHAR`,
`VARCHAR`, `TEXT` or `STRING`, by default, the input
length is defined by the form field width, without exceeding the length of the character string
program variable. The `SCROLL`
attribute can be used to bypass this limit and force the input length to be as large as the program
variable can hold. For example, when using a `VARCHAR(20)` variable with a form field
defined with 6 cells, the maximum input length is by default a character width of 6 (= 6 Latin chars
or 3 Chinese ideograms). When using `SCROLL`, the maximum input length will be 20
bytes or 20 chars, depending on [FGL\_LENGTH\_SEMANTICS](../07_configuration/0518-fgl-length-semantics.md "Defines the length semantics to be used in programs."). Using the `SCROLL` attribute must be an exception:
define form fields with a number of cells that can hold all possible characters, that can fit in the
corresponding program variable. However, for some item types like `TEXTEDIT`, the
`SCROLL` attribute behavior is implicit when the element is stretchable or allows
scrollbars.

To compute the width of the form item in the AUI tree, specific rules apply for some kind of form
item types like `DATEEDIT`, `COMBOBOX` and
`BUTTONEDIT`, see [Widget width inside hbox tags](1559-widget-width-inside-hbox-tags.md) for more
details.

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")

[FORMAT attribute](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.")

[SAMPLE attribute](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget.")
