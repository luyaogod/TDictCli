---
title: "ButtonEdit style attributes"
source: "fgl-topics/r_fgl_presentation_styles_buttonedit_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > ButtonEdit style attributes"
type: "reference"
---

# ButtonEdit style attributes

> ButtonEdit presentation style attributes apply to BUTTONEDIT elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `customWidget`

Defines the type of widget to be used to render the `BUTTONEDIT`.

Values can be:

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

[BUTTONEDIT item type](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.")
