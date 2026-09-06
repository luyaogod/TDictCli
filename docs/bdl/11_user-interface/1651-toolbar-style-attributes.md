---
title: "ToolBar style attributes"
source: "fgl-topics/r_fgl_presentation_styles_toolbar_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > ToolBar style attributes"
type: "reference"
---

# ToolBar style attributes

> ToolBar presentation style attributes apply to the TOOLBAR element.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `aspect`

Defines the aspect of the toolbar, by specifying the visibility or
the relative position of the icon and text of a toolbar item.

Values can be:

- `"iconAboveText"` (default): The icons are displayed on top of the texts.
- `"icon"`: Only toolbar item icons are displayed. When no icon is available,
  displays the text.
- `"text"`: Only toolbar item texts is displayed.
- `"iconAndText"`: The icons are displayed to the left of the texts.

If the condition requires that a text must be displayed, but not text value is available,
the action name will be used as text.

> **Note:**
>
> If a continuous dropdown is displayed (the 3-dots icon) because the window is too small to
> show all toolbar buttons, both icon and text are displayed in the items of the drop-down list (the
> `aspect` attribute is ignored for this list).

## `itemsAlignment`

Defines how toolbar items must be aligned for the available space.

Values can be:

- `"left"` (default): Items are compacted to the left.
- `"center"`: Items are compacted and centered.
- `"right"`: Items are compacted to the right.
- `"justify"`: Items are equally stretched to occupy all available space.

## `position`

Defines the position of the toolbar element.

Values can be:

- `"default"` (default): The toolbar position is defined by the [`toolbarPosition`](1655-window-style-attributes-basics.md) style attribute at Window level.
- `"top"`: The toolbar is displayed at the top of the window.
- `"bottom"`: The toolbar is displayed at the bottom of the window.
- `"none"`: The toolbar is hidden.
- `"chrome"`: The toolbar items are merged in the chromebar. See [Action views in chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.").

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

## `size`

Specifies an abstract size for the toolbar.

Values can be:

- `"small"`: Displays toolbar items with a small size, allowing a maximum of items
  to be displayed.
- `"medium"` (default): Displays toolbar items with a regular size.
- `"large"`: Displays large toolbar items, showing only a handful of buttons.

The number of toolbar items displayed with these different size specifications depends on the
size of the screen. For example, on a mobile device screen, `"small"` will show about
seven toolbar items (buttons). This will also depend on the screen orientation.

## Related links

**Related concepts**  

[TOOLBAR section](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.")
