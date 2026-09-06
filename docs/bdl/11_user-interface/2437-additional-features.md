---
title: "Additional features"
source: "fgl-topics/c_fgl_dynamic_dialogs_misc.html"
breadcrumb: "User interface > User interface programming > Dynamic Dialogs > Additional features"
type: "concept"
---

# Additional features

> Miscellaneous features of dynamic dialogs.

## Defining cell attributes

In order to [display array cells with
colors](../15_library-reference/3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs).") in a list controlled by a dynamic dialog, you can define a two-dimensional dynamic
array to hold cell attributes:

```
DEFINE attrs DYNAMIC ARRAY WITH DIMENSION 2 OF STRING
...
LET attrs[row,col] = "red reverse"
...
CALL d.setArrayAttributes("custlist", attrs)
```

If all cells of the array must get the same display attributes, use a simple `DYNAMIC
ARRAY OF STRING`.

## Handling rowset sorting

With a dynamic dialog controlling a list container that allows rowset sorting (like a
`TABLE`), the sort event can be detected with the `"ON SORT"` event
name.

When the `"ON SORT"` event occurs, the dynamic dialog can perform the same
operations as when in a static `DISPLAY ARRAY` or `INPUT ARRAY` using
the [`ON SORT`](1989-on-sort-block.md) interaction
block:

```
    CALL d.addTrigger("ON SORT")
    ...
    WHILE (t := d.nextEvent()) IS NOT NULL
        CASE t
            ...
            WHEN "ON SORT"
                MESSAGE SFMT( "Row: %1/%2",
                    d.arrayToVisualIndex( "sr", d.getCurrentRow("sr") ),
                    d.getArrayLength( "sr" )
                )
    ...
```
