---
title: "Reacting to field value changes"
source: "fgl-topics/c_fgl_prog_dialogs_on_change.html"
breadcrumb: "User interface > User interface programming > Input fields > Reacting to field value changes"
type: "concept"
---

# Reacting to field value changes

> This section describes the purpose of the ON CHANGE interaction block.

The `ON CHANGE` interaction block
can be used in different ways:

- With form fields allowing only entire value input such as `CHECKBOX`,
  or using an additional widget such as a calendar in a `DATEEDIT`:
  `ON CHANGE` can be used to detect an immediate value change, or the
  selection of a value in the additional widget, without leaving the field.
- With text fields like `EDIT` (allowing incomplete values), defined
  with the `COMPLETER` attribute to implement autocompletion: In this case the
  `ON CHANGE` trigger is used to fill the list of matching values, when the user types
  characters in (`ON CHANGE` is fired after a short delay).
- With text fields like `EDIT` (allowing incomplete values): `ON
  CHANGE` can be used to detect a value change, when the field is left.

A typically usage of `ON CHANGE` is for example with a `CHECKBOX`,
to enable/disable other form elements depending on the value of the checkbox field:

```
INPUT BY NAME rec.* ...
    ...
    ON CHANGE input_details -- can be TRUE or FALSE
       CALL DIALOG.setFieldActive("address1", rec.input_details)
       CALL DIALOG.setFieldActive("address2", rec.input_details)
    ...
END INPUT
```

> **Important:**
>
> The `dialogtouched` predefined action can also be used to
> detect field changes immediately, but with this action you cannot get the data in the target
> variables. Use of this special action is only recommended to detect if the user has started to
> modify data in the current dialog.

## Related links

**Related concepts**  

[Immediate detection of user changes](2239-immediate-detection-of-user-changes.md "This section describes the dialogtouched predefined action.")

[Enabling autocompletion](2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field.")
