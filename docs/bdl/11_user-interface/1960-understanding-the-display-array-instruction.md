---
title: "Understanding the DISPLAY ARRAY instruction"
source: "fgl-topics/c_fgl_DisplayArray_002.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Understanding the DISPLAY ARRAY instruction"
type: "concept"
---

# Understanding the DISPLAY ARRAY instruction

> The DISPLAY ARRAY dialog controls a list of records.

`DISPLAY ARRAY` is designed to browse a list of records, binding a static or
dynamic array model to a screen array of the current displayed form.

A `DISPLAY ARRAY` instruction supports additional features such as drag
& drop, tree-view management, built-in sort and search, multi-row selection and list
modification triggers. For a detailed description of these features, see [Table views](2315-table-views.md "Describes how to implement table/list views.").

Use the `DISPLAY ARRAY` instruction to let the end user browse in a list of rows
after fetching a result set from the database. The result set is produced with a database cursor
executing a `SELECT` statement. The `SELECT` SQL statement is usually
completed at runtime with a `WHERE` clause produced from a `CONSTRUCT`
dialog. When the `DISPLAY ARRAY` statement completes execution, the program must test
the `int_flag` variable to check if the dialog was
validated or canceled. If `int_flag` is `FALSE`, the program can get
the current row from `arr_curr()`.

Depending on the type of `DISPLAY ARRAY`, you must implement [dialog data blocks](1967-display-array-data-blocks.md).

When the user browses the list, [dialog control
blocks](1971-display-array-control-blocks.md) such as `BEFORE ROW` are executed.

When the user clicks on an action view (button), or when an asynchronous event occurs, [dialog interaction blocks](1979-display-array-interaction-blocks.md) like `ON ACTION`
are executed.

The code inside a `DISPLAY ARRAY` dialog can use [control instructions](1995-display-array-control-instructions.md), [dialog control functions](2223-dialog-control-functions.md "The language provides several built-in functions and operators to be used in a dialog instruction."), and the [`ui.Dialog`](2222-the-dialog-control-class.md "This topic explains the purpose of the ui.DIALOG class.") class, to implement
the dialog behavior.

> **Tip:**
>
> A `DISPLAY ARRAY` dialog is in read-only mode for navigating the list of records
> and needs a user action to perform a modification to the record list. An `INPUT
> ARRAY` dialog offers record list navigation with direct modification; no user action is
> required to modify the data.
>
> `DISPLAY ARRAY` is better adapted to modern graphical user interfaces, while
> `INPUT ARRAY` is better adapted to TUI mode ergonomics. It is also easier to program
> a `DISPLAY ARRAY` with modification triggers compared to the implementation of an
> `INPUT ARRAY`, especially when the underlying SQL table must be updated during the
> dialog execution.
>
> For modern GUI applications, we recommend using `DISPLAY ARRAY` with modification
> triggers ([`ON APPEND`](1984-on-append-block.md), [`ON INSERT`](1985-on-insert-block.md), [`ON UPDATE`](1986-on-update-block.md), [`ON DELETE`](1987-on-delete-block.md)).

## Related links

**Related concepts**  

[Dialog programming basics](2218-dialog-programming-basics.md "This section describes basic dialog programming concepts.")
