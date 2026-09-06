---
title: "Form field deactivation"
source: "fgl-topics/c_fgl_prog_dialogs_field_activation.html"
breadcrumb: "User interface > User interface programming > Input fields > Form field deactivation"
type: "concept"
description: "The form fields bound to a dialog are by default active (they can get the focus). When needed, disable the fields that do not require user input, and reactivate them later during the dialog execution. ..."
---

# Form field deactivation

The form fields bound to a dialog are by default active (they can get the focus). When needed,
disable the fields that do not require user input, and reactivate them later during the dialog
execution.

For example, imagine a form containing an "Industry" `COMBOBOX` field , with the
options *Healthcare*, *Education*, *Government*, *Manufacturing*, and
*Other*. If the user selects "Other", a secondary `EDIT` field is expected to be
activated automatically, to let the user input the specific description of the industry. But if one
of the predefined values is selected, there is no need for the additional field, so the secondary
field can be left disabled.

This can be achieved by enabling/disabling fields with the
`ui.Dialog.setFieldActive()` method depending on the context. The "Industry" field
case described can be implemented as
follows:

```
DIALOG ATTRIBUTES(UNBUFFERED)
   INPUT BY NAME rec.*
      ON CHANGE industry 
         -- A value of 99 corresponds to the "Other" item 
         CALL DIALOG.setFieldActive( "cust.industry", (rec.industry!=99) )
     ...
   END INPUT
   BEFORE DIALOG
     CALL DIALOG.setFieldActive( "cust.industry", FALSE )
     ...
END DIALOG
```

Consider centralizing field activation / deactivation in a setup function specific to the dialog,
passing the `DIALOG` object as parameter.

Form fields can be directly hidden with the `ui.Form.setFieldHidden()` method, or
indirectly hidden, when the field is a child element of an object hidden with the `ui.Form.setElementHidden()` method. A
dialog considers hidden fields as disabled. There is no need to disable fields that are already
hidden.

Do not disable or hide all fields of a dialog, otherwise the dialog
execution stops: At least one field must be able to get the focus during a dialog execution.

If you disable or hide the current field having the focus, the dialog will execute the
`AFTER FIELD` block of the current field and the `BEFORE FIELD` block
on the next field in the tabbing order. This can unnecessarily fire validation code implemented in
`AFTER FIELD`. As a general pattern, do not disable the current field having the
focus.

## Related links

**Related concepts**  

[COMBOBOX item type](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.")

[EDIT item type](1690-edit-item-type.md "Defines a simple line-edit field.")
