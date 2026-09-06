---
title: "INSTRUCTIONS section"
source: "fgl-topics/c_fgl_FormSpecFiles_INSTRUCTIONS_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > INSTRUCTIONS section"
type: "concept"
---

# INSTRUCTIONS section

> The INSTRUCTIONS section is used to define screen arrays, non-default screen records, and global form properties.

## Syntax

```
INSTRUCTIONS
{ screen-record-definition [;] }
[ DELIMITERS "AB" [;] ]
[ DEFAULT SAMPLE = "string" ]
[END]
```

1. screen-record-definition is the definition of a screen record or screen
   array.
2. A and B define the opening and closing field delimiters
   for character based terminals.

## Usage

The `INSTRUCTIONS` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `INSTRUCTIONS` section is optional in a form definition.

The `END` keyword is optional.

This section is mainly used to define screen records, to group fields using tables, tree views,
scrollgrids, or traditional static field arrays.

## Screen records (or screen arrays)

A screen record is a named group of form
fields:

```
SCREEN RECORD sr(customer.*);
```

See [Screen records / arrays](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") for more details.

## Field delimiters

Use the DELIMITER keyword to specify the characters to be displayed as field delimiters on the
screen. This option is especially used for TUI mode applications:

```
DELIMITERS "[]";
```

## Default sample

The `DEFAULT SAMPLE` directive defines the default sample text for all
fields:

```
DEFAULT SAMPLE = "MMM"
```

See [SAMPLE attribute](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget.") for more details.

## Example

```
SCHEMA stores 
LAYOUT
GRID
{
  ...
}
END
TABLES
 stock, items 
END
ATTRIBUTES
...
END
INSTRUCTIONS
  SCREEN RECORD s_items[10]
    ( stock.*,
      items.quantity,
      FORMONLY.total_price )
  DELIMITERS "[]"
END
```

## Related links

**Related concepts**  

[Binding tables to arrays in dialogs](2318-binding-tables-to-arrays-in-dialogs.md "Program arrays act as data model that are bound to form tables, when implementing list dialogs.")

[Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.")
