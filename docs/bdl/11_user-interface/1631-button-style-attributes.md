---
title: "Button style attributes"
source: "fgl-topics/r_fgl_presentation_styles_button_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Button style attributes"
type: "reference"
---

# Button style attributes

> Button presentation style attributes apply to BUTTON elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `alignment`

Defines the position of the image and/or text inside a button, when the button is bigger than
the content.

When both image and text are specified for a button, these will be combined as a single graphical
element before applying the `alignment` style attribute. If the image and text render
with different heights, the smallest element will be centered vertically to the element with the
greatest height.

The value can be a combination of a vertical and horizontal alignment hints, separated by a
space.

The value can also be `"center"`, which is equivalent to `"verticalCenter
horizontalCenter"`.

Vertical alignment hints:

- `"top"`: anchor to the top edge.
- `"verticalCenter"`: center in middle.
- `"bottom"`: anchor to the bottom edge.

Horizontal alignment hints:

- `"left"`: anchor to the left edge.
- `"horizontalCenter"`: center in middle.
- `"right"`: anchor to the right edge.

When only a text or only an image is used for the button, the default alignment is
`"center"`. When combining an image and a text, the default is `"left
verticalCenter"`, in order to align icons when several buttons are defined vertically.

## `buttonType`

Defines the rendering of a button.

Values can be:

- `"normal"` (default): The button is rendered as a regular push button.
- `"link"`: the button is rendered as an HTML hyper-link. In contrast to the label
  hyper-link support, clicking on a "link" button does not start the default browser, but triggers the
  corresponding action, like a normal button.
- `"commandLink"`: the button is rendered as a "Command Link" button.

The `buttonType` Button style attribute cannot be changed dynamically, once the
widget has been displayed.

## `scaleIcon`

Defines the scaling behavior of the associated icon. Depending on the image source format and the
`scaleIcon` attribute value, images can be upscaled or downscaled.

The purpose of the `scaleIcon` attribute is to adapt the image source to the size
of the containing widget, which can depend on the layout and font size when text is displayed.

Raster images are never upscaled to avoid blurring. However, SVG images and TTF icons can be
upscaled without any penalty.

The image is centered in the containing button, which adapts to the widget size. This allows a
mix of larger and smaller icons while keeping widget alignment.

If scaling takes place, the aspect ratio of the original image is kept. A non-square source image
displays as a non-square scaled icon.

If the `scaleIcon` attribute is undefined, the behavior depends on the type of
action view: toolbar button icons and action panel button icons are scaled down to match the size of
the widget. For other widgets, by default no scaling occurs, as for
`scaleIcon="no"`.

Values can be:

- `"no"`: No scaling occurs and the image is taken as-is. It is up to the developer
  to resize the source image to avoid misalignment.
- `"yes"`: Images are scaled following the height of the widget. Raster images can only be downscaled, while SVG/TTF icons can be downscaled and upscaled.
- `"nnnpx"`: Images are scaled to the specified size. For example,
  scaleIcon="128px" will make every icon a maximum of 128\*128 pixels. At least one side equal to 128
  pixels, depending if the source image is square or not. Raster images can only be downscaled, while SVG/TTF icons can be downscaled and upscaled.

## Related links

**Related concepts**  

[BUTTON item type](1684-button-item-type.md "Defines a push-button that can trigger an action.")
