---
title: "Label style attributes"
source: "fgl-topics/r_fgl_presentation_styles_label_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Label style attributes"
type: "reference"
---

# Label style attributes

> Label presentation style attributes apply to LABEL elements.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `sanitize`

By default, to avoid "Stored XSS" attacks, the front-end cleans the HTML sent to form elements to
ensure no malicious script can be executed. This security control prevents for example to use HTML
content such as `"<a href='mailto: …"`.

The default for the `sanitize` attribute is `"yes"`.

To disable the checking of HTML content send to form elements, set the `sanitize`
style attribute to `"no"`.

The `sanitize` style attribute makes only sense for `Label` form
items, when used with the `textFormat` style attribute set to
`"html"`:

```
<Style name="Label.relax">
     <StyleAttribute name="textFormat" value="html"/>
     <StyleAttribute name="sanitize" value="no"/>
</Style>
```

## `textFormat`

Defines the rendering of the content of the `LABEL` widget.

Possible values are:

- `"plain"` (default): the value assigned to this widget is interpreted as plain
  text.
- `"html"`: it is interpreted as HTML (with hyperlinks).
  > **Important:**
  >
  > The HTML content displayed inside a form element
  > using the `textFormat=html` style must not be a complete HTML document (using CSS
  > styles for example). The content must be simple HTML, with basic tags such as text decoration like
  > `<b/>` for bold, `<ul/>+<li/>` for bullet lists, and inline
  > styles.

## Related links

**Related concepts**  

[LABEL item type](1696-label-item-type.md "Defines a simple text area to display a read-only value.")
