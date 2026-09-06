---
title: "Group style attributes"
source: "fgl-topics/r_fgl_presentation_styles_group_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Group style attributes"
type: "reference"
---

# Group style attributes

> Group presentation style attributes apply to an GROUP element.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `collapserPosition`

Indicates the position of the collapser icon, when using collapsible groups.

Values can be `"left"` (default), `"right"`.

## `collapsible`

Defines if the group element can be collapsed and expanded.

The group needs to be defined with a TEXT, in order to see the effect of the
`collapsible` style attribute.

By default groups are not collapsible.

Values can be `"yes"`, `"no"` (default).

## `initiallyCollapsed`

Defines if a collapsible group is collapsed or expanded when the form is displayed.

> **Note:**
>
> This style attribute is ignored, if the `collapsible` attribute is not defined
> to `"yes"` for this group.

Values can be:

- `"no"` (default): When displayed for the first time (no stored settings exist),
  the group appears expanded.
- `"yes"`: When displayed for the first time (no stored settings exist), the group
  appears collapsed.
- `"never"`: Each time it is displayed, the group appears expanded (stored settings
  and any previous display state are ignored).
- `"always"`: Each time it is displayed, the group appears collapsed (stored
  settings and any previous display state are ignored).

## Related links

**Related concepts**  

[GROUP item type](1693-group-item-type.md "Defines a layout area to group other layout elements together.")
