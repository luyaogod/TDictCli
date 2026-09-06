---
title: "ComboBox style attributes"
source: "fgl-topics/r_fgl_presentation_styles_combobox_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > ComboBox style attributes"
type: "reference"
---

# ComboBox style attributes

> ComboBox presentation style attributes apply to COMBOBOX elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `customWidget`

Defines the type of widget to be used to render the `COMBOBOX`.

Values can be:

- `"comboBoxSimpleConstruct"`: When the dialog is a `CONSTRUCT`, the
  combobox does not allows multiple item selection to produce a pipe-separatred QBE condition, and the
  `QUERYEDITABLE` attribute is ignored.

The `customWidget` ComboBox style cannot be changed dynamically, once the widget
has been displayed.

## Related links

**Related concepts**  

[COMBOBOX item type](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.")
