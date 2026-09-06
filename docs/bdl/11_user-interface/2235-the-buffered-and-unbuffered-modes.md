---
title: "The buffered and unbuffered modes"
source: "fgl-topics/c_fgl_prog_dialogs_unbuffered.html"
breadcrumb: "User interface > User interface programming > Input fields > The buffered and unbuffered modes"
type: "concept"
---

# The buffered and unbuffered modes

> The buffered and unbuffered mode control the synchronization of program variables and form fields.

## Data model / view / controller paradigm

When bound to an interactive instruction (such as dialog), program variables act as a [data model](2219-the-model-view-controller-paradigm.md "The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.") to display data or to get user input.
To change the values of form fields by program code, the corresponding variables must be set and
displayed.

Synchronization of program variables with the form fields depends on the buffer mode used by the
dialog. Use the unbuffered mode to get automatic data model / form field synchronization.

## Configuring the buffer mode

By default, singular dialogs (`INPUT`, `DISPLAY ARRAY`) and
procedural `DIALOG` blocks use the buffered mode.

The unbuffered mode can be set per (modal) dialog instruction, with the `UNBUFFERED`
dialog attribute:

```
INPUT BY NAME p_site.* ATTRIBUTES(UNBUFFERED) 
   ...
END INPUT
```

When using a [procedural DIALOG](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form.")
block, all sub dialogs defined locally or included with the [`SUBDIALOG`](2087-the-subdialog-clause.md) clause inherit the
buffer mode of the parent procedural dialog block:

```
DIALOG ATTRIBUTES(UNBUFFERED)
  INPUT BY NAME p_site.* -- unbuffered
     ...
  END INPUT
  DISPLAY ARRAY a_events TO sr_events.* -- unbuffered
     ...
  END DISPLAY
  SUBDIALOG d_comments -- unbuffered
END DIALOG
```

The unbuffered mode can also be set globally with the [ui.Dialog.setDefaultUnbuffered()](../15_library-reference/3177-ui-dialog-setdefaultunbuffered.md "Set the default unbuffered mode for all dialogs.")
method, for singular and procedural dialogs:

```
CALL ui.Dialog.setDefaultUnbuffered(TRUE)
...
INPUT BY NAME rec_cust.* WITHOUT DEFAULTS -- uses unbuffered mode 
   ...
END INPUT
```

## The buffered mode

When you use the default "buffered" mode, program variable changes are not automatically
displayed to form fields; you need to execute [`DISPLAY TO`](1880-display-to.md "The DISPLAY ... TO instruction displays data to specific form fields.") or [`DISPLAY BY NAME`](1881-display-by-name.md "The DISPLAY BY NAME instruction displays data to form fields corresponding to the variable names."). Additionally, if an action is triggered, the value of the
current field is not validated and is not copied into the corresponding program variable. The only
way to get the text of a field is to use [GET\_FLDBUF()](../08_language-basics/0671-get-fldbuf-function.md "The GET_FLDBUF() operator returns as character strings the current values of the specified fields.") or [DIALOG.getFieldBuffer()](../15_library-reference/3197-ui-dialog-getfieldbuffer.md "Returns the input buffer of the specified field."). These functions return the current text, which might not be a valid
representation of a value of the field data type:

```
INPUT BY NAME p_item.*
   ON ACTION zoom 
      CALL select_item()
           RETURNING p_item.code, p_item.desc
         DISPLAY BY NAME p_item.code, p_item.desc
      END IF
   ...
END INPUT
```

## The unbuffered mode

With the unbuffered mode, program variables and form fields are automatically synchronized, and
the dialog instruction is sensitive to program variable changes: You don't need to display
values explicitly with `DISPLAY TO` or `DISPLAY BY NAME`. When
an action is triggered, the value of the current field is validated and is copied into the
corresponding program variable. If you need to display new data during the dialog execution,
just assign the values to the program variables; the runtime system will automatically
display the values to the screen after user code of the current control or interaction block
has been executed:

```
INPUT BY NAME p_site.* ATTRIBUTES(UNBUFFERED)
   ON ACTION zoom 
      CALL select_item()
           RETURNING p_item.code, p_item.desc
         -- no need to display desc.
      END IF
   ...
END INPUT
```

## Actions configuration for field validation

During data input, values entered by the user in form fields are automatically validated and
copied into the program variables. Actually the value entered in form fields is first available in
the form field buffer. This buffer can be queried with built-in functions or [dialog class methods](../15_library-reference/3169-the-dialog-class.md "The ui.Dialog class provides a set of methods to configure, query and control the current interactive instruction."). With the unbuffered mode, the field
buffer is used to synchronize program variables each time control returns to the runtime system -
for example, when the user clicks on a button to execute an action.

With the unbuffered mode, data validation must be prevented for some actions such as cancel or
close. To avoid field validation for a given action, set the [`validate`](2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.") action default
attribute to "no", in the .4ad file or in the `ACTION DEFAULTS`
section of the form file:

```
ACTION DEFAULTS
   ACTION undo (TEXT = "Undo", VALIDATE = NO)
   ...
END
```

Actions such as `dialogtouched`, `cancel`, `delete`,
`close`, `help` are by default defined with the
`validate=no` attribute in the $FGLDIR/lib/default.4ad file.

If field validation is disabled for an action, the code executed in the [`ON ACTION` block](1918-on-action-block.md) acts as if the dialog was
in buffered mode. The program variable is not set. However, the input buffer of the current field is
updated. When returning from the user code, the dialog will not synchronize the form fields with
program variables, and the current field will display the input buffer content. Therefore, if you
change the value of the program variable during an `ON ACTION` block where validation
is disabled, you must explicitly display the values to the fields with `DISPLAY TO / BY
NAME`.

To illustrate this case, imagine that you want to implement an undo action to allow the
modifications done by the user to be reverted (before these have been saved to the database of
course). You typically copy the current record into a clone variable when the dialog starts, and
copy these old values back to the input record when the undo action is invoked. An undo action is a
good candidate to avoid field validation, since you want to ignore current values. If you don't
re-display the values, the input buffer of the current field will remain when returning from the
`ON ACTION` block:

```
DIALOG ATTRIBUTES(UNBUFFERED)
 INPUT BY NAME p_cust.*
   BEFORE INPUT
     LET p_cust_copy = p_cust
   ON ACTION undo -- Defined with VALIDATE=NO
     LET p_cust = p_cust_copy
     DISPLAY BY NAME p_cust.*
 END INPUT
END DIALOG
```

For more details, see [Data validation at action invocation](2281-data-validation-at-action-invocation.md "The validate action attribute controls field validation when an action is fired.").

## Related links

**Related concepts**  

[Binding variables to form fields](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.")
