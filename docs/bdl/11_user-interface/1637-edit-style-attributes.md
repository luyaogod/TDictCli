---
title: "Edit style attributes"
source: "fgl-topics/r_fgl_presentation_styles_edit_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Edit style attributes"
type: "reference"
---

# Edit style attributes

> Edit presentation style attributes apply to an EDIT element.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `customWidget`

Defines the type of widget to be used to render the `EDIT`.

Values can be:

- `"phoneEdit"`: The edit field provides international
  phone number decoration and input features, by displaying flags representing the country code. The
  phone edit widget displayed value will follow the country’s phone format. When typing or when the
  field contains the value `+41446681800`, the widget will display the flag of
  Switzerland, and the formatted phone number `44 668 18 00`. In the
  .per file, consider to define the form field that is large enough for those
  additional spaces. This type of widget can be used in conjunction with the [field autocompletion](1770-completer-attribute.md "The COMPLETER attribute enables autocompletion for the edit field.") feature, if needed. If the
  dialog controlling the field is a `CONSTRUCT`, the widget will be a regular edit box,
  to allow any search criteria input. Read also [Phone number fields](2249-phone-number-fields.md "Phone number fields allow to display and input formatted telephone numbers, with international call prefix selection.").
- `"tagEdit"`: The edit field allows multiple-value
  input in the same editor widget. To propose values to the end user, this type of widget can be used
  in conjunction with the [field autocompletion](1770-completer-attribute.md "The COMPLETER attribute enables autocompletion for the edit field.")
  feature. The tag edit widget renders multiple labels in individual graphical elements, that can be
  removed with the `[x]` cross buttons. New values can be added by the end user, if the
  autocompletion code allows it. If the dialog controlling the field is a `CONSTRUCT`,
  the widget will be a regular edit box, to allow any search criteria input. See also [Multi-valued fields](2248-multi-valued-fields.md "Multi-valued form fields allow to display and input a set of values in the same character string field.").

## `allowTagCreation`

The `allowTagCreation` style attribute can be used to control the drop down list
rendering for additional tag creation in a tag field defined with the `customWidget`
`"tagEdit"`.

Values can be `"yes"` (default), `"no"`.

> **Important:**
>
> The `allowTagCreation` style attribute does not deny tag creation: It is mainly
> provided to define the behavior of the drop down list. New tag creation must be controlled with the
> autocompletion code (`COMPLETER`/`ON CHANGE`) implement for the
> field.

A tag entered by the end user is considered as new, if the returned `COMPLETER`
proposals are empty.

By default, or when set to `"yes"`, when a new tag is entered, the tag edit shows
the drop down box with the new element followed by the `(+)` indicator.

When set to `"no"`, the tag edit does not show the drop down box if a new tag is
entered. The tag field must implement autocompletion with the `COMPLETER` attribute
and limit the proposals to a predefined list of tags, to force the end user to choose only values
shown in the drop down list. When autocompletion request occurs with an `ON CHANGE`
trigger, the program code should also reset the field value to remove the invalid tag.

## `showVirtualKeyboard`

Defines how the virtual keyboard must appear on mobile devices.

Values can be:

- `"onFocus"` (default): The virtual keyboard shows up when the edit widget gets
  the focus on a user tap or a program `NEXT FIELD` instruction.
- `"onTap"`: The virtual keyboard only shows up when the user taps on the edit
  widget.

## Related links

**Related concepts**  

[EDIT item type](1690-edit-item-type.md "Defines a simple line-edit field.")
