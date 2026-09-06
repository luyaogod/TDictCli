---
title: "The INPUT ARRAY sub-dialog"
source: "fgl-topics/c_fgl_DIALOG_subdialog_INPUT_ARRAY_2.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > Structure of a declarative DIALOG block > The INPUT ARRAY sub-dialog"
type: "concept"
---

# The INPUT ARRAY sub-dialog

> The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.

## Program array to screen array binding

The `INPUT ARRAY` sub-dialog binds the members of the flat record (or the
primitive member) of an array to the [screen-array or screen-record](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.") fields specified with the `FROM` keyword. The
number of variables in each record of the program array must be the same as the number of fields in
each screen record (that is, in a single row of the screen array).

In most cases you want to bind a [program array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") to a
screen-array, in order to display a page of records. However, the `INPUT ARRAY`
instruction can also bind the program array to a simple flat screen-record, to show one record at a
time.

The next code example defines an array with a flat record and binds it to a screen array:

```
PUBLIC TYPE t_items DYNAMIC ARRAY OF RECORD
       item_num INTEGER,
       item_name VARCHAR(50),
       item_price DECIMAL(6,2)
   END RECORD
   ...
DIALOG items_input(p_items t_items)
   INPUT ARRAY p_items FROM sa.*
      BEFORE INSERT
      ...
   END INPUT
   ...
END DIALOG
```

If the screen array is defined with one field only, you can bind an array defined with a
primitive type:

```
PUBLIC TYPE t_names DYNAMIC ARRAY OF VARCHAR(50)
   ...
DIALOG names_input(p_names t_names)
   INPUT ARRAY p_names FROM sa.*
      BEFORE DELETE
      ...
   END INPUT
   ...
END DIALOG
```

## Identifying an INPUT ARRAY sub-dialog

The name of the screen array specified with the `FROM` clause will be used to
identify the list. For example, the dialog class method such as [`DIALOG.getCurrentRow("screen-array")`](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction.") takes the name of the screen array as
the parameter, to identify the list you want to query for the current row. The name of the
screen-array is also used to qualify [sub-dialog actions](2284-sub-dialog-actions-in-procedural-dialog-blocks.md "This topic describes how action are differentiated with handlers defined in a procedural DIALOG block.") with a prefix.

## Control blocks in INPUT ARRAY

Editable record lists declared with the `INPUT ARRAY` sub-dialog can raise the
following triggers:

- [BEFORE INPUT](1940-before-input-block.md)
- [BEFORE ROW](1975-before-row-block.md)
- [BEFORE FIELD](1942-before-field-block.md)
- [ON CHANGE](1943-on-change-block.md)
- [AFTER FIELD](1944-after-field-block.md)
- [ON ROW CHANGE](2018-on-row-change-block.md)
- [AFTER ROW](1976-after-row-block.md)
- [BEFORE DELETE](2022-before-delete-block.md)
- [AFTER DELETE](2023-after-delete-block.md)
- [BEFORE INSERT](2020-before-insert-block.md)
- [AFTER INSERT](2021-after-insert-block.md)
- [AFTER INPUT](1941-after-input-block.md)

In the singular `INPUT ARRAY` instruction, `BEFORE INPUT` and
`AFTER INPUT` blocks are typically used as initialization and finalization blocks. In
the `INPUT ARRAY` sub-dialog of a `DIALOG` block, `BEFORE
INPUT` and `AFTER INPUT` blocks are executed each time the focus goes to
(`BEFORE`) or leaves (`AFTER`) the group of fields defined by this
sub-dialog.

## Related links

**Related concepts**  

[INPUT ARRAY ATTRIBUTES clause](2092-input-array-attributes-clause.md "INPUT ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.")
