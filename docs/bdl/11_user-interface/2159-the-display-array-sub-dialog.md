---
title: "The DISPLAY ARRAY sub-dialog"
source: "fgl-topics/c_fgl_DIALOG_subdialog_DISPLAY_ARRAY_2.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > Structure of a declarative DIALOG block > The DISPLAY ARRAY sub-dialog"
type: "concept"
---

# The DISPLAY ARRAY sub-dialog

> The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.

## Program array to screen array binding

The `DISPLAY ARRAY` sub-dialog binds the members of the flat record (or the
primitive member) of an array to the [screen-array or screen-record](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") fields specified with the `TO` keyword.

The number of variables in each record of the program array must be the same as the number of
fields in each screen record (that is, in a single row of the screen array).

In most cases you want to bind a [program array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") to a
screen-array, in order to display a page of records. However, the `DISPLAY ARRAY`
instruction can also bind the program array to a simple flat screen-record, to show one record at a
time.

The next code example defines an
array with a flat record and binds it to a screen array:

```
PUBLIC TYPE t_items DYNAMIC ARRAY OF RECORD
       item_num INTEGER,
       item_name VARCHAR(50),
       item_price DECIMAL(6,2)
   END RECORD
   ...
DIALOG items_list(p_items t_items)
   DISPLAY ARRAY p_items TO sa.*
      BEFORE ROW
      ...
   END DISPLAY
   ...
END DIALOG
```

If the screen array is defined with one
field only, you can bind an array defined with a primitive type:

```
PUBLIC TYPE t_names DYNAMIC ARRAY OF VARCHAR(50)
   ...
DIALOG names_list(p_names t_names)
   DISPLAY ARRAY p_names TO sa.*
      BEFORE DELETE
      ...
   END DISPLAY
   ...
END DIALOG
```

## Identifying a DISPLAY ARRAY sub-dialog

The name of the screen array specified with the `TO` clause identifies the list.
The dialog class method takes the name of the screen array as the parameter, identifying the list.
For example, you would use [`DIALOG.getCurrentRow("screen-array")`](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.") to query for the current row in the
list identified by 'screen-array'. The name of the screen-array is also used to qualify [sub-dialog actions](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.") with a prefix.

## Control blocks in DISPLAY ARRAY

Read-only record lists declared with the `DISPLAY
ARRAY` sub-dialog can raise the following triggers:

- [BEFORE
  DISPLAY](1973-before-display-block.md)
- [BEFORE ROW](1975-before-row-block.md)
- [AFTER ROW](1976-after-row-block.md)
- [AFTER DISPLAY](1974-after-display-block.md)

In the singular `DISPLAY ARRAY`instruction, `BEFORE DISPLAY` and
`AFTER DISPLAY` blocks are typically used as initialization and
finalization blocks. In a `DISPLAY ARRAY` sub-dialog of a
`DIALOG` block, `BEFORE DISPLAY` and `AFTER
DISPLAY` blocks will be executed each time the focus goes to
(`BEFORE`) or leaves (`AFTER`) the group of fields
defined by this sub-dialog.

## Related links

**Related concepts**  

[DISPLAY ARRAY ATTRIBUTES clause](2091-display-array-attributes-clause.md "DISPLAY ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")
