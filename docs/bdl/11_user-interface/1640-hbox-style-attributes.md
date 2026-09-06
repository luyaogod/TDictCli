---
title: "HBox style attributes"
source: "fgl-topics/r_fgl_presentation_styles_hbox_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > HBox style attributes"
type: "reference"
---

# HBox style attributes

> HBox presentation style attributes apply to an HBox element.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `navigationArrows`

When the HBox is defined with the [`SPLIT`](1818-split-attribute.md "The SPLIT attribute forces a horizontal box to show only one child container.") attribute,the `navigationArrows` style attribute
controls left and right navigation arrows for the view, to let the end user navigate between child
containers of the box.

Values can be `"yes"`, `"no"` (default).

## `navigationDots`

When the HBox is defined with the [`SPLIT`](1818-split-attribute.md "The SPLIT attribute forces a horizontal box to show only one child container.")
attribute,the `navigationDots` style attribute controls the navigation dots created
for the view, to let the end user navigate between child containers of the box.

Values can be `"yes"` (default), `"no"`.

## `packed`

Allows to pack the elements of a `VBOX` or `HBOX`, when there is no
stretchable child element such as a `TABLE`.

Values can be `"yes"`, `"no"` (default).

When set to `"yes"`, for a vertical box, elements are packed to the top. For an
horizontal box, elements are packed to the left.

By default, or when this attribute is set to `"no"`, space is inserted between any
two elements, if none of the elements can stretch in the alignment direction of the box.

If the `HBOX` or `VBOX` is defined with a `SPLITTER`
attribute, the `packed` style attribute is ignored once the user set manually the
splitter positions.

## Related links

**Related concepts**  

[HBOX item type](1694-hbox-item-type.md "Defines a layout area to render child elements in horizontal direction.")
