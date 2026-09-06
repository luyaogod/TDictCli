---
title: "Controlling the number of rows"
source: "fgl-topics/c_fgl_prog_dialogs_array_count.html"
breadcrumb: "User interface > User interface programming > List dialogs > Controlling the number of rows"
type: "concept"
---

# Controlling the number of rows

> Methods are provided to set and get the total number of rows in a read-only or editable list of records.

## Set the number of rows when using a static array

When using a static array in `DISPLAY ARRAY` or `INPUT
ARRAY`, you must specify the actual number of rows with the [`SET_COUNT()`](../15_library-reference/2788-set-count.md "Defines the number of rows containing explicit data in a static array used by the next dialog.")
built-in function or with the [`COUNT`](2091-display-array-attributes-clause.md "DISPLAY ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.") dialog attribute. Both of them are only taken
into account when the interactive instruction starts.

```
DEFINE arr ARRAY[100] OF ...
... (fill the array with x rows)
CALL set_count(x)
DISPLAY ARRAY arr TO sa.*
   ...
END DISPLAY
```

When using multiple list subdialogs in a `DIALOG` block, the
`SET_COUNT()` built-in function is unusable, as it defines the
total number of rows for all lists. The only way to define the number of rows when
using a static array in multiple dialogs is to use the `COUNT`
attribute.

Consider using dynamic arrays instead of static arrays.

## Set the number of rows when using a dynamic array

When using a dynamic array in `DISPLAY ARRAY` or `INPUT
ARRAY`, the total number of rows is automatically defined by the array
variable
(`array.getLength()`).

```
DEFINE arr DYNAMIC ARRAY OF ...
... (fill the array with x rows)
DISPLAY ARRAY arr TO sa.*
   ...
END DISPLAY
```

However, special consideration has to be taken when using the paged mode of
`DISPLAY ARRAY`. In this mode, the dynamic array only holds a
page of the complete row set shown to the user: In paged mode, you must specify the
total number of rows with the [`ui.Dialog.setArrayLength()`](../15_library-reference/3220-ui-dialog-setarraylength.md "Sets the number of rows in a DISPLAY ARRAY using paged mode.") method.

## Defining the maximum number of rows in `INPUT ARRAY`

In an `INPUT ARRAY` dialog allowing row creation with the insert or append
actions, there is by default no limit for the total number of rows. To specify the maximum number of
rows an INPUT ARRAY can accept, define the [`MAXCOUNT`](2011-input-array-instruction-configuration.md)
attribute:

## Get the number of rows in a list

To get the current number of rows in a `DISPLAY ARRAY` or
`INPUT ARRAY`, use either the [`ui.Dialog.getArrayLength()`](../15_library-reference/3193-ui-dialog-getarraylength.md "Returns the total number of rows in the specified list.") or the [`ARR_COUNT()`](../15_library-reference/2727-arr-count.md "Returns the number of rows entered during an INPUT ARRAY statement.")
function.

The `getArrayLength()` method can be used inside or outside the
context of the list dialog, as it takes the screen array as parameter to identify the list dialog.
For example, when implementing a `DIALOG` block with two `DISPLAY
ARRAY` subdialogs, you can query the number of rows of a list in the code block of another
list
controller:

```
DIALOG ...
   DISPLAY ARRAY arr1 TO sa1.*
      ON ACTION check
         IF DIALOG.getArrayLength("sa2")] > 1 THEN
            ...
         END IF
   END DISPLAY
   DISPLAY ARRAY arr2 TO sa2.*
   END DISPLAY
END DIALOG
```

The `ARR_COUNT()` function must be used in the context of the
`DISPLAY ARRAY` or `INPUT ARRAY` dialog, or just
after executing such dialog. For example, it can be used just after an `INPUT
ARRAY` dialog, to get the number of rows left in the
list:

```
INPUT ARRAY arr FROM sa.*
   ...
END INPUT
IF NOT int_flag THEN
   FOR i=1 TO arr_count()
     ...
   END FOR
END IF
```

The `ARR_COUNT()` function returns the number of rows for the last
executed dialog, until a new list dialog is started.

## Related links

**Related concepts**  

[ARRAY](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.")

[Populating a DISPLAY ARRAY](2307-populating-a-display-array.md "The program array must be filled with rows to populate the DISPLAY ARRAY dialog.")
