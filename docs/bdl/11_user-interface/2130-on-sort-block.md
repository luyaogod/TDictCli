---
title: "ON SORT block"
source: "fgl-topics/c_fgl_dialog_ON_SORT_3.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG interaction blocks > ON SORT block"
type: "concept"
description: "Syntax ON SORT instruction [...] Basics The ON SORT interaction block is used to detect row sort specification change in a DISPLAY ARRAY or INPUT ARRAY dialog. ON SORT is used in two different ..."
---

# ON SORT block

## Syntax

```
ON SORT
   instruction [...]
```

## Basics

The `ON SORT` interaction block is used to detect row sort specification change in
a [`DISPLAY ARRAY`](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.") or
[`INPUT ARRAY`](2086-the-input-array-sub-dialog.md "The INPUT ARRAY sub-dialog is the controller to implement the navigation and edition in a list of records.") dialog.

`ON SORT` is used in two different contexts:

1. In a [regular full-list](2308-full-list-mode-of-display-array.md "In order to handle short/medium result sets, use the full list mode of DISPLAY ARRAY.") `DISPLAY
   ARRAY` / `INPUT ARRAY` dialog, the `ON SORT` trigger can be
   used to detect that a list sort was performed.
2. In a [paged mode](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.") `DISPLAY
   ARRAY` using `ON FILL BUFFER`, the `ON SORT` trigger is
   mandatory to detect a sort request from the user and re-fetch the rows from the database in the
   required order.

## ON SORT in regular full-list DISPLAY ARRAY or INPUT ARRAY

In a regular `DISPLAY ARRAY` / `INPUT ARRAY` dialog (not using
paged mode), the `ON SORT` trigger can be used to detect that a list sort was
performed.

When the `ON SORT` block executes in this context, the (visual) sort is already
done by the runtime system and the `ON SORT` block is only used to execute post-sort
tasks, such as displaying current row information.

To display the row position information, use the [arrayToVisualIndex()](../15_library-reference/3187-ui-dialog-arraytovisualindex.md "Converts the program array index to the visual index for a given screen array.") dialog method to
convert the current program row number to the visual row
number:

```
DISPLAY ARRAY arr TO sr.* ...
   ...
   ON SORT
      MESSAGE SFMT( "Row: %1/%2",
         DIALOG.arrayToVisualIndex( "sr", DIALOG.getCurrentRow("sr") ),
         DIALOG.getArrayLength( "sr" )
      )
   ...
```

The current sort specification is awailable with the [`getSortKeyAt()`](../15_library-reference/3203-ui-dialog-getsortkeyat.md "Returns the name of field used as grouping or sorting column, for a given sort column position.")/[`isSortReverseAt()`](../15_library-reference/3208-ui-dialog-issortreverseat.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending), for a given sort column position.") dialog
methods:

```
DISPLAY ARRAY arr TO sr.* ...
   ...
   ON SORT
      MESSAGE SFMT( "First sort column: %1, %2 order",
           DIALOG.getSortKeyAt("sr",1),
           IIF( DIALOG.isSortReverseAt("sr",1), "descending", "ascending" )
      )
   ...
```

## ON SORT in DISPLAY ARRAY using the paged mode

In a `DISPLAY ARRAY` implementing paged mode with `ON FILL BUFFER`
trigger, built-in row sorting is not available because data is provided by pages.

The `ON SORT` trigger must be defined, to detect a sort request and perform a new
SQL query to re-order the rows. In this context, the sort column and sort order are available with
the [`getSortKeyAt()`](../15_library-reference/3203-ui-dialog-getsortkeyat.md "Returns the name of field used as grouping or sorting column, for a given sort column position.")/[`isSortReverseAt()`](../15_library-reference/3208-ui-dialog-issortreverseat.md "Indicates the sort order direction (FALSE=ascending, TRUE=descending), for a given sort column position.") dialog
methods:

```
DEFINE key STRING, rev BOOLEAN

DISPLAY ARRAY arr TO sr.* ...
   ...
    ON SORT
      -- Assuming that form field names match table column names
      LET i = 1
      LET order_by = NULL
      WHILE TRUE
          LET key = DIALOG.getSortKeyAt("sr",i)
          LET rev = DIALOG.isSortReverseAt("sr",i)
          IF key IS NULL THEN EXIT WHILE END IF
          IF order_by IS NULL THEN
             LET order_by = " ORDER BY ", key
          ELSE
             LET order_by = order_by, ", ", key
          END IF
          LET order_by = order_by, " ", IIF(rev,"DESC","ASC")
          LET i = i + 1
      END WHILE
      CALL execute_sql(order_by)
```

See [Paged mode of DISPLAY ARRAY](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.") for more details about the paged mode in
`DISPLAY ARRAY` and how to implement sort in this type of record list dialog.

## Related links

**Related concepts**  

[Sorting rows in a list](2324-sorting-rows-in-a-list.md "List controllers implement a built-in sort. This feature can be disabled if not required.")
