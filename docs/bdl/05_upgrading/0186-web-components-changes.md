---
title: "Web components changes"
source: "fgl-topics/c_fgl_Migrate_to_310_web_components.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Web components changes"
type: "concept"
---

# Web components changes

> Modifications to consider when using web components.

## Only SIZEPOLICY=FIXED is supported for WEBCOMPONENT fields

When defining a `WEBCOMPONENT` field in the form, the only valid option for the
`SIZEPOLICY` attribute is `FIXED`. This is by the way the default for
`WEBCOMPONENT` fields.

For more details, see [Controlling the web component layout](../11_user-interface/2381-controlling-the-web-component-layout.md).

## fglrichtext: New implementation

Up to FGLGWS version 3.10.09 (WCG 1.00.11), the fglrichtext built-in web component was based on
`draft.js` Rich Text Editor Framework.

Starting with FGLGWS 3.10.11 (WCG 1.00.12), the fglrichtext web component is implemented with the
`quilljs` Rich Text Editor.

Available features and configuration options have changed.

> **Note:**
>
> Starting with FGL 5.00, the fglrichtext webcomponent is now deprecated,
> in favor to `TEXTEDIT` with `textFormat="html"` style attribute.

## fglrichtext: Support of emojis

Starting with FGLGWS version 3.10.13 (WCG 1.00.13), the `emoji`
toolbar option allows users to include an emoji in the text of an fglrichtext web component.

> **Important:**
>
> Since emojis are Unicode characters and not plain images, your application
> needs to use UTF-8. Emojis rendering will differ depending on the device you're running it on. You
> might need to install fonts that handle emoji characters as well if your device doesn't handle it.

Starting with FGL 5.00, the fglrichtext webcomponent is now deprecated,
in favor to `TEXTEDIT` with `textFormat="html"` style attribute.

## fglrichtext: Defining a default font

Starting with FGLGWS version 3.10.16 (WCG 1.00.15), a default font family and font size can be
specified for a fglrichtext web component with new properties.

Starting with FGL 5.00, the fglrichtext webcomponent is now deprecated,
in favor to `TEXTEDIT` with `textFormat="html"` style attribute.

## fglrichtext: Localizing texts

Starting with FGLGWS version 3.10.16 (WCG 1.00.15), it is possible to localize the texts used by
the fglrichtext web component.

Several properties have been added to the fglrichtext web component, to define the strings for
each toolbar button tooltips, combobox items, and popup dialog labels.

Starting with FGL 5.00, the fglrichtext webcomponent is now deprecated,
in favor to `TEXTEDIT` with `textFormat="html"` style attribute.

## fglsvgcanvas: Mouse hovering events

Starting with FGLGWS version 3.10.11 (WCG 1.00.12), the fglsvgcanvas built-in web component
supports new properties to detect `onmouseover` / `onmouseout` mouse
hovering SVG events, and trigger `ON ACTION` blocks.

For more details, see [The fglsvgcanvas web component](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.").

## fglsvgcanvas: title() and createChars() functions

Starting with FGLGWS version 3.10.18 (WCG 1.00.16), the fglsvgcanvas library provides the [`title()`](../15_library-reference/2889-fglsvgcanvas-title.md "Produces an SVG \"title\" element.") function to
create a title SVG element, and the [`createChars()`](../15_library-reference/2851-fglsvgcanvas-createchars.md "Produces an SVG DOM text node.")
function to create a text node.

For more details, see [fglsvgcanvas: SVG drawing module](../15_library-reference/2841-fglsvgcanvas-svg-drawing-module.md).

## fglgallery: Aspect ratio for image elements

Using FGLGWS version 3.10.18 (WCG 1.00.16), image elements are now aligned properly by using the
same size, in case if picture resources have different sizes.

By default, the images are displayed with a square (`1:1`) aspect ratio. To define
a different aspect ratio, the fglgallery library provides the `setImageAspectRatio()`
function.

For more details see [fglgallery image aspect ratio handling](../11_user-interface/2418-the-fglgallery-web-component.md).

## fglgallery: Custom CSS file handling

Starting with FGLGWS version 3.10.18 (WCG 1.00.16), to avoid missing resource errors in web
viewers, the custom CSS file
$FGLDIR/webcomponents/fglgallery/css/fglgallery-custom.css is no longer included
by default in the fglgallery.html file.

If you want to provide you own CSS file for fglgallery, uncomment the `<link
/>` line in fglgallery.html.

Note also that CSS names have changed, since the HTML content of fglgallery has been reviewed for
images aspect ratio handling.

## Related links

**Related concepts**  

[Web components](../11_user-interface/2378-web-components.md "This section describes how to use web components in your application.")
