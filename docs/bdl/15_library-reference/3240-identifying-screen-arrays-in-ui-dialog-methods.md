---
title: "Identifying screen-arrays in ui.Dialog methods"
source: "fgl-topics/c_fgl_ClassDialog_identify_scrarr.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Usage > Identifying screen-arrays in ui.Dialog methods"
type: "concept"
description: "In ui.Dialog methods such as setCurrentRow() , the first parameter is the name of the screen array to identify the list container in the form. Screen arrays are defined in form specification files ..."
---

# Identifying screen-arrays in ui.Dialog methods

In `ui.Dialog` methods such as [`setCurrentRow()`](3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list."), the first parameter is the name of the screen array to
identify the list container in the form.

Screen arrays are defined in form specification files with the `SCREEN RECORD`
clause, and used in `DISPLAY ARRAY` and `INPUT ARRAY` instructions to
bind program array variables to the list container.

In the form file:

```
LAYOUT
...
INSTRUCTIONS
SCREEN RECORD custlist ( ... );
END
```

In the program
code:

```
DISPLAY ARRAY custarr TO custlist.*
   ...
   ON ACTION set_row
      CALL DIALOG.setCurrentRow("custlist", row_index)
   ...
```

The name of the screen array passed as parameter can use the same letter case as in the
`SCREEN RECORD` definition: The lookup is case-insensitive:

```
-- In the form file:
SCREEN RECORD CustList ( ... );
-- In the program code:
CALL DIALOG.setCurrentRow("CustList", row_index)
```

## Related links

**Related concepts**  

[Screen records / arrays](../11_user-interface/1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")

[Variable binding in DISPLAY ARRAY](../11_user-interface/1964-variable-binding-in-display-array.md "Variable binding in DISPLAY ARRAY")

[Variable binding in INPUT ARRAY](../11_user-interface/2010-variable-binding-in-input-array.md "Variable binding in INPUT ARRAY")
