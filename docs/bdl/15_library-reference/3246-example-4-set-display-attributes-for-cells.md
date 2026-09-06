---
title: "Example 4: Set display attributes for cells"
source: "fgl-topics/c_fgl_ClassDialog_example_4.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > Examples > Example 4: Set display attributes for cells"
type: "concept"
description: "Note: This code example uses a DYNAMIC ARRAY with the same RECORD structure as the data array. You could also use a two-dimensional dynamic array or a simple flat array of strings. For more details, ..."
---

# Example 4: Set display attributes for cells

> **Note:**
>
> This code example uses a DYNAMIC ARRAY with the same RECORD structure
> as the data array. You could also use a two-dimensional dynamic array or a
> simple flat array of strings. For more details, see [ui.Dialog.setArrayAttributes](3219-ui-dialog-setarrayattributes.md "Define cell decoration attributes array for the specified list (singular or multiple dialogs).")

```
FUNCTION display_items()
   DEFINE i INTEGER
   DEFINE items DYNAMIC ARRAY OF RECORD
                key INTEGER,
                name CHAR(10)
              END RECORD
   DEFINE attributes DYNAMIC ARRAY OF RECORD
                key STRING,
                name STRING
              END RECORD

   FOR i=1 TO 10
     CALL items.appendElement()
     LET items[i].key = i 
     LET items[i].name = "name " || i 
     CALL attributes.appendElement()
     IF i MOD 2 = 0 THEN
       LET attributes[i].key = "red"
       LET attributes[i].name = "blue reverse"
     ELSE
       LET attributes[i].key = "green"
       LET attributes[i].name = "magenta reverse"
     END IF
   END FOR

   DISPLAY ARRAY items TO sr.* ATTRIBUTES(UNBUFFERED)
       BEFORE DISPLAY 
         CALL DIALOG.setCellAttributes(attributes)
       ON ACTION att_modify_cell 
         LET attributes[2].key = "red reverse"
       ON ACTION att_clear_cell 
         LET attributes[2].key = NULL
   END DISPLAY

END FUNCTION
```
