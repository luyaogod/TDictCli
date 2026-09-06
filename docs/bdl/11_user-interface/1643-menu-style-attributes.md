---
title: "Menu style attributes"
source: "fgl-topics/r_fgl_presentation_styles_menu_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Menu style attributes"
type: "reference"
---

# Menu style attributes

> Menu presentation style attributes apply to MENU elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

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

[Ring menus (MENU)](1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.")
