---
title: "ui.Dialog.setDialogAttribute"
source: "fgl-topics/c_fgl_ClassDialog_setDialogAttribute.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setDialogAttribute"
type: "concept"
---

# ui.Dialog.setDialogAttribute

> Set an attribute to configure a dynamic dialog.

## Syntax

```
setDialogAttribute(
   attr STRING,
   value STRING )
```

1. attr is the name of the dialog attribute.
2. value is the value of the dialog attribute.

## Usage

The `setDialogAttribute()` method can be used to configure a dynamic dialog
with attributes.

When defining a static dialog instruction, you can specify options with the
`ATTRIBUTES()`
clause:

```
INPUT BY NAME ... ATTRIBUTES(UNBUFFERED)
```

When creating a dynamic dialog, call the `setDialogAttribute()` method after
creating a new dialog object (for example with [`createInputByName()`](3174-ui-dialog-createinputbyname.md "Creates an ui.Dialog object to implement a dynamic INPUT BY NAME.")):

```
LET d = ui.Dialog.createInputByName(fields)
CALL d.setDialogAttribute("UNBUFFERED",TRUE)
```

Or after adding a sub-dialog to a multiple dynamic dialog object (for example with [`addDisplayArrayTo()`](3180-ui-dialog-adddisplayarrayto.md "Adds a sub-dialog of type DISPLAY ARRAY TO to an existing ui.Dialog dynamic dialog.")):

```
LET d = ui.Dialog.createMultipleDialog()
CALL d.addDisplayArrayTo(fields, "sr_custlist")
CALL d.setDialogAttribute("FOCUSONFIELD",TRUE)
```

| Attribute name | Possible values | Description | Dialog block equivalent |
| --- | --- | --- | --- |
| `AUTO APPEND` | `TRUE` / `FALSE` | When `FALSE`, `INPUT ARRAY` dialog denies [automatic temp row creation](../11_user-interface/2310-input-array-row-modifications.md "Controlling row creation and deletion in an editable record list.") when moving after last existing row. | [`AUTO APPEND`](../11_user-interface/2092-input-array-attributes-clause.md) |
| `APPEND ROW` | `TRUE` / `FALSE` | When `FALSE`, `INPUT ARRAY` dialog denies [row addition](../11_user-interface/2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") (at end of list). | [`APPEND ROW`](../11_user-interface/2092-input-array-attributes-clause.md) |
| `COUNT` | `integer` | Defines the [number of rows](../11_user-interface/2302-controlling-the-number-of-rows.md "Methods are provided to set and get the total number of rows in a read-only or editable list of records.") of the `DISPLAY ARRAY` or `INPUT ARRAY`. | [`COUNT`](../11_user-interface/1965-display-array-instruction-configuration.md) |
| `DELETE ROW` | `TRUE` / `FALSE` | When `FALSE`, `INPUT ARRAY` dialog denies [row deletion](../11_user-interface/2310-input-array-row-modifications.md "Controlling row creation and deletion in an editable record list."). | [`DELETE ROW`](../11_user-interface/2092-input-array-attributes-clause.md) |
| `DOUBLECLICK` | `identifier` | Defines the action name for the [row choice](../11_user-interface/2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.") (mouse click or tap) | [`DOUBLECLICK`](../11_user-interface/2091-display-array-attributes-clause.md) |
| `FIELD ORDER` | `"FORM"`, `"CONSTRAINED"`, `"UNCONSTRAINED"` | Defines the [tabbing order mode](../11_user-interface/2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.") of form fields for this dialog. | [`FIELD ORDER`](../11_user-interface/2089-dialog-attributes-clause.md) |
| `FOCUSONFIELD` | `TRUE` / `FALSE` | When set to `TRUE`, enables the [field-focus mode](../11_user-interface/2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.") in a `DISPLAY ARRAY` dialog. | [`FOCUSONFIELD`](../11_user-interface/2091-display-array-attributes-clause.md) |
| `HELP` | `integer` | The help number of the [help message](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs."). | [`HELP`](../11_user-interface/2090-input-attributes-clause.md) |
| `INSERT ROW` | `TRUE` / `FALSE` | When `FALSE`, `INPUT ARRAY` dialog denies [row insertion](../11_user-interface/2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") (before current row). | [`INSERT ROW`](../11_user-interface/2092-input-array-attributes-clause.md) |
| `KEEP CURRENT ROW` | `TRUE` / `FALSE` | When `TRUE`, [current row must remain selected](../11_user-interface/2319-controlling-table-rendering.md) after list dialog execution. | [`KEEP CURRENT ROW`](../11_user-interface/1965-display-array-instruction-configuration.md) |
| `MAXCOUNT` | `integer` | Defines the [maximum number of rows](../11_user-interface/2302-controlling-the-number-of-rows.md "Methods are provided to set and get the total number of rows in a read-only or editable list of records.") to be entered in an `INPUT ARRAY` dialog. | [`MAXCOUNT`](../11_user-interface/2011-input-array-instruction-configuration.md) |
| `NAME` | `identifier` | Defines the [name of the sub-dialog](../11_user-interface/2081-identifying-sub-dialogs-in-dialog.md "Sub-dialogs need to be identified by a name to distinguish the different contexts."). | [`NAME`](../11_user-interface/2090-input-attributes-clause.md) |
| `UNBUFFERED` | `TRUE` / `FALSE` | When `TRUE`, the dialog will use the [unbuffered mode](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields."). | [`UNBUFFERED`](../11_user-interface/2089-dialog-attributes-clause.md) |
| `WITHOUT DEFAULTS` | `TRUE` / `FALSE` | When `TRUE`, `INPUT` or `INPUT ARRAY` dialog must not use [form defaults](../11_user-interface/2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option."). | [`WITHOUT DEFAULTS`](../11_user-interface/1936-input-instruction-configuration.md) |

## Related links

**Related concepts**  

[Configuring a dynamic dialog](../11_user-interface/2429-configuring-a-dynamic-dialog.md "The dynamic dialogs can be configured with attributes.")
