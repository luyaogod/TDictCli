---
title: "arr_count()"
source: "fgl-topics/c_fgl_BuiltInFunctions_ARR_COUNT.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > arr_count()"
type: "concept"
---

# arr_count()

> Returns the number of rows entered during an INPUT ARRAY statement.

## Syntax

```
FUNCTION arr_count()
  RETURNS INTEGER
```

## Usage

Use `arr_count()` to determine the number of program records
that are currently stored in a static program array used by the
`INPUT ARRAY` instruction.

This function is typically called inside or after `INPUT ARRAY`
or `DISPLAY ARRAY` statement.

`arr_count()` returns
a positive integer, corresponding to the index of the furthest record
within the static program array that the user accessed. Not all the
rows counted by `arr_count()` necessarily contain data
(for example, if the user presses the Down key more times than there
are rows of data.

This function is not required when using
dynamic arrays. In such case, the total number of rows in defined
by the [`array.getLength()`](2949-dynamic-array-getlength.md "Returns the length of the array.") method
after the dialog, or by the [`ui.Dialog.getArrayLength()`](3193-ui-dialog-getarraylength.md "Returns the total number of rows in the specified list.")
method during the dialog execution.

## Related links

**Related concepts**  

[Editable record list (INPUT ARRAY)](../11_user-interface/2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")

[Record list (DISPLAY ARRAY)](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.")

[arr\_curr()](2728-arr-curr.md "Returns the current row in a DISPLAY ARRAY or INPUT ARRAY.")
