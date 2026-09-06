---
title: "Data validation at action invocation"
source: "fgl-topics/c_fgl_prog_dialogs_action_validate.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Data validation at action invocation"
type: "concept"
---

# Data validation at action invocation

> The validate action attribute controls field validation when an action is fired.

When using the `UNBUFFERED` mode of interactive instructions such
as `INPUT` or `DIALOG`, if the user triggers an action, the current
field data is checked and loaded in the target variable bound to the form field. For example, if the
user types a wrong date (or only a part of a date) in a field using a `DATE` variable
and then clicks on a button to invoke an action, the runtime system will display an invalid input
error and will not execute the `ON ACTION` block corresponding to the button.

To prevent data validation for some actions, use the [`VALIDATE` action
attribute](2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.").

The `validate` action attribute can be set in the global action default file, or
at the form level, with the `VALIDATE` attribute in a line of the `ACTION
DEFAULTS` section, and in the `ATTRIBUTES()` clause of the `ON
ACTION` action handler.

This `validate` attribute instructs the runtime not to copy the input buffer text
into the program variable (requiring input buffer text to match the target data type).

```
ACTION DEFAULTS
   ...
   ACTION undo ( TEXT="Undo", VALIDATE=NO )
   ...
END
```

This is especially needed in `DIALOG` instructions; in singular dialogs like
`INPUT`, predefined actions like cancel do not validate the current field value when
`UNBUFFERED` mode is used.

## Related links

**Related concepts**  

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")

[ACTION DEFAULTS section](1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.")
