---
title: "Giving the focus to a form element"
source: "fgl-topics/c_fgl_prog_dialogs_set_focus.html"
breadcrumb: "User interface > User interface programming > Input fields > Giving the focus to a form element"
type: "concept"
---

# Giving the focus to a form element

> How to force the focus to move or stay in a specific form element using program code.

Use the [`NEXT FIELD`](1955-next-field-instruction.md) instruction
to force the focus to a specific field or screen record (list). The `NEXT FIELD`
instruction expects a form field name.

In a `DIALOG` block, when the specified field is the first column identifier
of a sub-dialog driven by a `DISPLAY ARRAY` block, the read-only list gets
the focus. If the field name is not known at compile time, you can alternatively use the
`ui.Dialog.nextfield()`
method.

```
DIALOG ATTRIBUTES(UNBUFFERED)
   INPUT BY NAME p_cust ATTRIBUTES(NAME="cust")
     ...
   END DISPLAY
   DISPLAY ARRAY p_orders TO orders.*
     ...
   END DISPLAY
   ON ACTION go_to_header 
      NEXT FIELD cust_num 
   ON ACTION go_to_detail 
      NEXT FIELD order_lineno 
   ...
END DIALOG
```

When a `BUTTON` exists in the form layout, it can get the focus if the
`DIALOG` block defines a `COMMAND` clause as action
handler. Currently there is no way to give the focus to a `BUTTON` by
program.

```
DIALOG ATTRIBUTES(UNBUFFERED)
   ...
   COMMAND "print" 
      CALL print_order()
   ...
END DIALOG
```

In rare cases (especially when using folder tabs), it may be required to show a part of the form
that is not controlled by the dialog, when there is no active field or button that can get the focus
in that part of the form, and when the above techniques cannot work. In this case, in order to show
temporarily a given part of the form that cannot get the focus, you use the `ui.Form.ensureFieldVisible` or [`ui.Form.ensureElementVisible`](../15_library-reference/3148-ui-form-ensureelementvisible.md "Ensure the visibility of a form element.") methods.

```
DEFINE form ui.Form
...
DIALOG ATTRIBUTES(UNBUFFERED)
   ...
   BEFORE DIALOG
      LET form = DIALOG.getForm()
   ...
   ON ACTION show_image1 
      CALL form.ensureElementVisible("image1")
   ...
END DIALOG
```

When using the `FOCUSONFIELD` attribute of `DISPLAY ARRAY`, you can
set the focus to a specific cell by using the `NEXT
FIELD` instruction or the [`ui.Dialog.nextField()`](../15_library-reference/3210-ui-dialog-nextfield.md "Registers the next field to go to.") method, in conjunction with the [`ui.Dialog.setCurrentRow()`](../15_library-reference/3224-ui-dialog-setcurrentrow.md "Sets the current row in the specified list.")
method. For more details, see [Field-level focus in DISPLAY ARRAY](2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.").

## Related links

**Related concepts**  

[The DISPLAY ARRAY sub-dialog](2085-the-display-array-sub-dialog.md "The DISPLAY ARRAY sub-dialog is the controller to implement the navigation in a list of records, with option data modification actions.")

[Which form item has the focus?](2244-which-form-item-has-the-focus.md "Identify what element of the current form has the focus.")

[Implementing dialog action handlers](2279-implementing-dialog-action-handlers.md "How to execute user code in ON ACTION blocks when an action is fired.")
