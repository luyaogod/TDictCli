---
title: "CheckBox style attributes"
source: "fgl-topics/r_fgl_presentation_styles_checkbox_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > CheckBox style attributes"
type: "reference"
---

# CheckBox style attributes

> CheckBox presentation style attributes apply to CHECKBOX elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `customWidget`

Defines the type of widget to be used to render the `CHECKBOX`.

Values can be:

- `"toggleButton"`: The checkbox is rendered as a toggle button (also known as
  "toggle switch" in HTML/CSS). Note that `NULL` values cannot be managed with the
  `toggleButton` widget type.

Default is to render the checkbox with a classical box using a check mark when set.

The `customWidget` CheckBox style cannot be changed dynamically, once the widget
has been displayed.

## Related links

**Related concepts**  

[CHECKBOX item type](1686-checkbox-item-type.md "Defines a boolean or three-state checkbox field.")
