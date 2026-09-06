---
title: "ui.Dialog.setArrayLength"
source: "fgl-topics/c_fgl_ClassDialog_setArrayLength.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setArrayLength"
type: "concept"
---

# ui.Dialog.setArrayLength

> Sets the number of rows in a DISPLAY ARRAY using paged mode.

## Syntax

```
setArrayLength(
   name STRING,
   length INTEGER )
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).
2. length is the new size of the array.

## Usage

The `setArrayLength()` method is used to specify the total number of rows when
using the `DISPLAY ARRAY` [paged
mode](../11_user-interface/2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY."). The name of the screen array is passed to identify the list, followed by an integer
expression defining the number of rows.

When using a regular dynamic array without paged mode (without the `ON FILL
BUFFER` clause), you don't need to specify the total number of rows to the
`DIALOG` instruction. It is defined by the number of elements in the array. However,
when using the paged mode in a `DISPLAY ARRAY`, the total number of rows does not
correspond to the elements in the program array, because the program array holds only a page of the
whole list. In any other case, a call to this method is just ignored.

A call to `setArrayLength()` will not trigger the execution of the `ON FILL
BUFFER` clause immediately: This trigger will be executed when the control goes back to the
dialog instruction, after all user code following `setArrayLength()` has been
executed.

In a paged mode `DISPLAY ARRAY` using `COUNT=-1`, before calling
`setCurrentRow(screen-array, row-index)`, be
sure to provide the actual number of rows with [`ui.Dialog.setArrayLength( screen-array, count
)`](3220-ui-dialog-setarraylength.md "Sets the number of rows in a DISPLAY ARRAY using paged mode.") where count > row-index. Otherwise, the
`setCurrentRow()` call will have no effect, if the dialog has not yet seen
row-index rows through `ON FILL BUFFER`.

## Related links

**Related concepts**  

[ON FILL BUFFER block](../11_user-interface/1968-on-fill-buffer-block.md "ON FILL BUFFER block")
