---
title: "Image style attributes"
source: "fgl-topics/r_fgl_presentation_styles_image_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Image style attributes"
type: "reference"
---

# Image style attributes

> Image style presentation attributes apply to an IMAGE element.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `alignment`

Defines the image alignment when the container is bigger than the image itself.

Possible values are a pair of horizontal (`"left"`,
`"horizontalCenter"`, `"right"`) and vertical alignments
(`"top"`, `"verticalCenter"`, `"bottom"`). To combine
alignment options, use a space as separator.

Value can also be `"center"`, which is equivalent to `"horizontalCenter
verticalCenter"`.

The default value is `"top left"`.

## `imageBorderRadius`

The `imageBorderRadius` style attribute defines the
rounder border radius for images. This attribute accepts the same values as the CSS
`border-radius` property. For example, you can make circular images with a value of
`50%`, or get a little rounded corner with `12px`.

## Related links

**Related concepts**  

[IMAGE item type](1695-image-item-type.md "Defines an area that can display an image resource.")
