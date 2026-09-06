---
title: "ui.Dialog.setArrayAttributes"
source: "fgl-topics/c_fgl_ClassDialog_setArrayAttributes.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setArrayAttributes"
type: "concept"
---

# ui.Dialog.setArrayAttributes

> Define cell decoration attributes array for the specified list (singular or multiple dialogs).

## Syntax

```
setArrayAttributes(
   name STRING,
   attributes dynamic-array-type )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. attributes is a program array defining the cell attributes.
3. dynamic-array-type is a `DYNAMIC ARRAY OF ...` type, which can
   be:
   - A `DYNAMIC ARRAY OF RECORD ... END RECORD` (with the same structure as the data
     array)
   - A `DYNAMIC ARRAY WITH DIMENSION 2 OF STRING` (to define attributes in dynamic
     dialog when the row structure is defined at runtime)
   - A `DYNAMIC ARRAY OF STRING` (to define attributes for complete lines instead of
     individual cells)

## Usage

In an [`INPUT ARRAY`](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.") or [`DISPLAY ARRAY`](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") dialog, the
`setArrayAttributes()` method can be used to specify display attributes for each
cell, or for the complete row.

The `setArrayAttributes()` is typically used in a `DIALOG` block
where several screen arrays are defined. The method takes the name of the screen array as first
parameter, to identify the list to be decorated with colors. An equivalent method called [`setCellAttributes()`](3221-ui-dialog-setcellattributes.md "Define cell decoration attributes array for the specified list (singular dialog only).") can be
used, for dialogs where only one screen array is defined.

Cell attributes are equivalent to TTY attributes defined at the form item level, such as the
[`COLOR`](../11_user-interface/1766-color-attribute.md "The COLOR attribute defines the foreground color of the text displayed by a form element.") attribute, except that
cell attribute can change dynamically.

Possible values for cell attributes are a combination of the following values:

- One of the [supported color names](../11_user-interface/1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value."), or an
  #RRGGBB value.
- The `blink` attribute: The text in the cell blinks.
- The `bold` attribute: The text in the cell renders with a bold font.
- The `underline` attribute: The text in the cell is underlined.
- The `reverse` attribute: The cell color it is used as background color instead of
  foreground text color. If no color is specified, it defaults to “black”, which produces a gray
  background.

The cell attributes must be specified in lowercase, the color names must be in camelCase, and the
separator is a blank space, for example:

> **Note:**
>
> For the predefined (TTY)
> color names (like `blue`, `red`, `yellow`), the actual
> rendering color used by the front-ends is adapted to the reverse or normal mode: In normal mode,
> colors like `"yellow"` must be darker, to make the foreground text easily readable on
> an white background. In reverse mode (where the color specification applies to the background), the
> actual background color must be lighter, so that the black foreground text can be read more
> easily.

- `"green"`: Foreground text color is green.
- `"lightRed reverse"`: Background cell color is light red.
- `"blue underline"`: Foreground text color is blue and the text font is
  underlined.
- `"red bold reverse"`: Background cell color is red and the text font is
  bold.

The structure of the dynamic array containing the cell attributes can be:

- A `DYNAMIC ARRAY OF RECORD`, with the same structure as the data array
- A `DYNAMIC ARRAY WITH DIMENSION 2 OF STRING`, to define a flexible set of cell
  attributes (for dynamic dialogs)
- A `DYNAMIC ARRAY OF STRING`, to define attributes for complete lines

The following example defines a dynamic array with the same structure as the data array. Note
however that the members of the attributes array use the `STRING` data
type:

```
DEFINE data DYNAMIC ARRAY OF RECORD
         pkey INTEGER,
         name VARCHAR(50)
       END RECORD
DEFINE attributes DYNAMIC ARRAY OF RECORD
         pkey STRING,
         name STRING
       END RECORD
```

Similarly, the cell attributes array can be defined with a two-dimensional dynamic array of
strings:

```
DEFINE attributes DYNAMIC ARRAY WITH DIMENSION 2 OF STRING
```

The advantage of a two-dimensional array is the flexibility, as it can define an unlimited number
of cells for each row. This solution is typically used when implementing a [dynamic dialog](../11_user-interface/2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.").

Finally, if you want to decorate complete lines instead of individual cells, use a simple dynamic
array of strings:

```
DEFINE attributes DYNAMIC ARRAY OF STRING
```

Fill the display attributes array with color and video
attributes:

```
FOR i=1 TO data.getLength()  -- length from data array!
    LET attributes[i].name = "blue reverse"
END FOR
```

Then, attach the array to the dialog with the `setArrayAttributes()` method, in a
`BEFORE DIALOG`, `BEFORE INPUT` or `BEFORE DISPLAY`
block:

```
BEFORE DIALOG
   CALL DIALOG.setArrayAttributes( "sr", attributes )
```

Like data values, if you change the cell attributes during the dialog, these are not displayed
automatically unless the [`UNBUFFERED`](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.") mode is
used.

```
ON ACTION modify_cell_attribute 
   LET attributes[arr_curr()].name = "red reverse"
```

If you set `NULL` to an element, the default TTY attributes will be
reset:

```
ON ACTION clean_cell_attribute 
   LET attributes[arr_curr()].name = NULL
```

## Related links

**Related concepts**  

[Cell color attributes](../11_user-interface/2313-cell-color-attributes.md "List controllers can display every cell in a specific color.")

[Combining TTY and style attributes](../11_user-interface/1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.")

[ui.Dialog.setCellAttributes](3221-ui-dialog-setcellattributes.md "Define cell decoration attributes array for the specified list (singular dialog only).")

[Example 4: Set display attributes for cells](3246-example-4-set-display-attributes-for-cells.md "Example 4: Set display attributes for cells")
