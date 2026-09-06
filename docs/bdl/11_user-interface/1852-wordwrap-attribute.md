---
title: "WORDWRAP Attribute"
source: "fgl-topics/c_fgl_FSFAttributes_WORDWRAP.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > WORDWRAP Attribute"
type: "concept"
---

# WORDWRAP Attribute

> The WORDWRAP attribute enables a multiple-line editor in TUI mode.

## Syntax

```
WORDWRAP [ { COMPRESS | NONCOMPRESS } ]
```

## Usage

This attribute is provided for backward compatibility with character-based forms, to support word
wrapping in multi-line text input.

In GUI mode, it is recommended that you use a `TEXTEDIT` form item instead. When used, the
`WORDWRAP` attribute is ignored, because text input and display is managed by the
text editor widget. The text data is not automatically modified by the editor by adding blank spaces
to put words on the next line.

In TUI mode, the `WORDWRAP` attribute has following effects:

- During input and display, the runtime system treats all segments that have that field tag as
  segments of a single field.
- The multi-line editor can wrap long character strings to the next line of a multiple-segment
  field for data entry, data editing, and data display.
- The `COMPRESS` option prevents blanks produced by the editor from being included
  in the program variable. `COMPRESS` is applied by default and can cause truncation to
  occur if the sum of intentional characters exceeds the field or column size. Because of editing
  blanks in the `WORDWRAP` field, the stored value might not correspond exactly to its
  multiple-line display.
- Specifying `NONCOMPRESS` after the `WORDWRAP` keyword causes any
  editor blanks to be saved when the string value is saved in a database column, in a variable, or in
  a file.

Using `WORDWRAP` fields with character-based terminals results in quite different
behavior than with graphical front-ends. With character-based terminals, the text input and display
is modified by the multi-line editor. This editor can automatically modify the text data by adding
blank spaces to put words to the next line, in order to make the text fit into the form field. In
GUI mode, the text input and display is managed by a multi-line edit control.

The maximum number of bytes a user can enter is the width of the form-field multiplied by the
height of the form-field. Blank characters may be intentional blanks or fill blanks. Intentional
blanks are initially stored in the target variable where entered by the user. Fill blanks are
inserted at the end of a line by the editor when a newline or a word-alignment forces a line-break.
It is not possible to set the cursor at a fill blank. Intentional blanks are always displayed (even
on the beginning of a line; the word-wrapping method used in reports with `PRINT
WORDWRAP` works differently).

When entering characters with Japanese locales, special characters are prohibited from being the
first or the last character on a line. If the last character is prohibited from ending a line, this
character is wrapped down to the next line. If the first character is prohibited from beginning a
line, the preceding character will also wrap down to the next line. This method is called kinsoku.
The test for prohibited characters will be done only once for the first and the last character on
each line.

Word-wrapping is disabled on the last row of a `WORDWRAP` field. The last word on
the last row may by truncated. The `WORDWRAP COMPRESS` attribute instructs the editor
to remove fill blanks before assigning the field-buffer to the target variable. The `WORDWRAP
NONCOMPRESS` attribute instructs the editor to store fill blanks to the target variable.

The `WORDWRAP` attribute is not used by the `CONSTRUCT`
instruction.

## Example

A field with `WORDWRAP` attribute must be defined with repeating [item tags](1679-item-tags.md "Item tags define the position and size in a grid-based container."), defining the number of lines of the
input field:

```
SCREEN
{
[f1                       ]
[f1                       ]
[f1                       ]
}
END
ATTRIBUTES
f1 = FORMONLY.longtext, WORDWRAP;
END
```

## Related links

**Related concepts**  

[Query by example (CONSTRUCT)](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")

[Text mode rendering (TUI mode)](1517-text-mode-rendering-tui-mode.md "Text mode rendering (TUI mode)")
