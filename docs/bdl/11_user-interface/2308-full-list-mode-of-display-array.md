---
title: "Full list mode of DISPLAY ARRAY"
source: "fgl-topics/c_fgl_prog_dialogs_dafill_full.html"
breadcrumb: "User interface > User interface programming > List dialogs > Populating a DISPLAY ARRAY > Full list mode of DISPLAY ARRAY"
type: "concept"
---

# Full list mode of DISPLAY ARRAY

> In order to handle short/medium result sets, use the full list mode of DISPLAY ARRAY.

## Understanding the full list mode

In full list mode, [`DISPLAY ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.") uses a complete copy of the result set to be displayed in the
form array. The full list mode is typically used for a short or medium row set (10 - 100 rows).

In full list mode, the `DISPLAY ARRAY` instruction uses a static or dynamic program
array defined with a record structure corresponding to (or to a part of ) a screen-array in the
current form.

The program array is filled with data rows before `DISPLAY ARRAY` is executed,
typically with a `FOREACH` loop when rows come from the database.

![Full list mode diagram](../_images/DARFig01.jpg)

*Full list mode in DISPLAY ARRAY diagram*

Consider using a dynamic array instead of a static array: By using a dynamic array the program
will only use the required memory resources, and the dialog will automatically detect the number of
rows from the dynamic array (`array.getLength()`)

## Full list mode example

The following example implements a `DISPLAY ARRAY` in its simpler form: a dynamic
array is filled with database rows and contains the whole result set to be displayed in the
table:

```
MAIN
  DEFINE arr DYNAMIC ARRAY OF RECORD
            id INTEGER,
            fname CHAR(30),
            lname CHAR(30)
        END RECORD 
  DEFINE i INTEGER

  DATABASE stores

  OPEN FORM f1 FROM "custlist"
  DISPLAY FORM f1

  DECLARE c1 CURSOR FOR
         SELECT customer_num, fname, lname FROM customer 
  LET i=1
  FOREACH c1 INTO arr[i].*
      LET i = i+1
  END FOREACH
  CALL arr.deleteElement(i)

  DISPLAY ARRAY arr TO sr.* ATTRIBUTES(UNBUFFERED)
      BEFORE ROW
        MESSAGE "Moved to row ", arr_curr() 
  END DISPLAY

END MAIN
```

## Related links

**Related concepts**  

[ARRAY](../08_language-basics/0731-array.md "An array defines a vector variable with a list of elements.")

[Dynamic arrays](../08_language-basics/0734-dynamic-arrays.md "Dynamic arrays")

[Multiple row selection](2314-multiple-row-selection.md "Multiple row selection allows the end user to select several rows within a list of records.")

[Tree views](2352-tree-views.md "Describes how to implement tree views.")

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")

[Paged mode of DISPLAY ARRAY](2309-paged-mode-of-display-array.md "In order to handle very large result sets, use the paged mode of DISPLAY ARRAY.")
