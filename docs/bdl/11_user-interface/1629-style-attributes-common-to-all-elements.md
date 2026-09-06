---
title: "Style attributes common to all elements"
source: "fgl-topics/r_fgl_presentation_styles_common_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Style attributes common to all elements"
type: "reference"
---

# Style attributes common to all elements

> Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.

> **Important:**
>
> Common style attribute apply to basic layout elements such as containers (Group) and form widgets
> (Label, Button, Edit, CheckBox).

## `backgroundColor`

Defines the color to be used to fill the background of the object.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

Default is no value (default background color of the object, or inherited background color).

## `border`

Defines if the widget gets a border.

Values can be `"yes"` (default) or `"no"`/`"none"`.
The value `"none"` is equivalent to `"no"`.

Default is no value (the widget gets its default appearance).

This style attribute can for example be used for the [`IMAGE`](1695-image-item-type.md "Defines an area that can display an image resource.") form item type.

## `defaultTTFColor`

Defines the default color used for TTF icons.

All TTF icons displayed in the form element using this style get the color specified in the
attribute.

The value for this attribute must be an RGB specification or a named color as listed in [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

When specified for a container such as Window or Group, child elements inherit the color for TTF
icons defined by the `defaultTTFColor` style attribute.

For more details about TTF icon usage see [Using a simple image name (centralized TTF icons)](1586-providing-the-image-resource.md).

## `fontFamily`

Defines the name of the font.

For possible values, see [Font families](1624-font-families.md "Use the fontFamily style attribute to define a font family.").

Default is no value (default object font or inherited font).

## `fontSize`

Defines the size of the characters.

For possible values, see [Font sizes](1625-font-sizes.md "Use the fontSize style attribute to influence the size of a font.").

Default is no value (default object font or inherited font).

## `fontStyle`

Defines the style of characters.

For possible values, see [Font styles](1626-font-styles.md "Use the fontStyle style attribute to define the style of a font.").

Default is no value (default object font or inherited font).

## `fontWeight`

Defines the weight of the characters.

Possible values for font weights depend on the front-end native font names, see [Font weights](1627-font-weights.md "Use the fontWeight style attribute to define the aspect of a font.") for details.

Default is no value (default object font or inherited font).

## `textColor`

Defines the color to be used to paint the text of the object.

For possible values, see [Colors](1622-colors.md "When providing a value for style attributes that define color, you can specify a generic color name or its RGB value.").

Default is no value (default object color or inherited color).

## `textDecoration`

Defines the decoration for the text.

Values can be `"overline"`, `"underline"` or
`"line-through"`.

Default is no value (default object font or inherited font).

## Related links

**Related concepts**  

[STYLE attribute](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.")
