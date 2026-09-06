---
title: "Which form item has the focus?"
source: "fgl-topics/c_fgl_prog_dialogs_focus_check.html"
breadcrumb: "User interface > User interface programming > Input fields > Which form item has the focus?"
type: "concept"
---

# Which form item has the focus?

> Identify what element of the current form has the focus.

Dialog control blocks such as `BEFORE INPUT`, `BEFORE DISPLAY`,
`BEFORE ROW`, `BEFORE FIELD`, `AFTER FIELD`,
`AFTER ROW`, `AFTER DISPLAY`, `AFTER INPUT` are
executed when the focus enters or leaves a field or a group of fields controlled by a sub-dialog. By
design, these control blocks implicitely identify which form element has the focus: When code
executing in a `BEFORE FIELD cust_id` block, we are in the `cust_id`
field.

Sometimes, you need to identify which form element has currently the focus, in a context
different for a control block such as `BEFORE FIELD`, especially with a
`DIALOG` block, which controls several parts of a form. For example, when several
lists are controlled by multiple `DISPLAY ARRAY` sub-dialogs, you may need to know
which is the current list in the context of a dialog-level `ON ACTION` block.

To get the name of the current form item, use the `DIALOG.getCurrentItem()` method. This
method is the replacement of the former `fgl_dialog_getfieldname()` built-in function.
It has been extended to return identifiers for fields, lists, or actions identifiers.

```
DIALOG ATTRIBUTES(UNBUFFERED)
   DISPLAY ARRAY p_orders TO orders.*
     ...
   END DISPLAY
   DISPLAY ARRAY p_items TO items.*
     ...
   END DISPLAY
   ...
      IF DIALOG.getCurrentItem() == "items" THEN
         ...
      END IF
   ...
END DIALOG
```

> **Note:**
>
> In the context of a `MENU` dialog, the `DIALOG.getCurrentItem()` method returns
> the name of the last triggered action.

## Related links

**Related concepts**  

[The DISPLAY ARRAY sub-dialog](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.")

[Detection of focus changes](2246-detection-of-focus-changes.md "Describes how to detect when the focus goes from field to field or to a read-only list.")

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")

[Field-level focus in DISPLAY ARRAY](2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.")
