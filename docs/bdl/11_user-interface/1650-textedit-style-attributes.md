---
title: "TextEdit style attributes"
source: "fgl-topics/r_fgl_presentation_styles_textedit_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > TextEdit style attributes"
type: "reference"
---

# TextEdit style attributes

> Textedit presentation style attributes apply to the TEXTEDIT element.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `aiWritingAssistant`

Enables the AI writing assistant in the `TEXTEDIT` field, if AI support is
enabled in GBC.

> **Important:**
>
> In order to activate AI features, the GBC theme variable `gbc-ai-enabled` needs to
> be set to `true`. This feature requires a configured AI provider, model, and API key.
> Consider the security and confidentiality implications of sharing data with third-party AI providers
> before proceeding. See GBC manual for more details.

Values can be:

1. `"yes"` (default): Activate the AI writing assistant.
2. `"no"`: Do not enable AI writing assistant.

## `showVirtualKeyboard`

Defines how the virtual keyboard must appear on mobile devices.

Values can be:

- `"onFocus"` (default): The virtual keyboard shows up when the edit widget gets
  the focus on a user tap or a program `NEXT FIELD` instruction.
- `"onTap"`: The virtual keyboard only shows up when the user taps on the edit
  widget.

## `sanitize`

By default, to avoid "Stored XSS" attacks, the front-end cleans the HTML sent to form elements to
ensure no malicious script can be executed. This security control prevents for example to use HTML
content such as `"<a href='mailto: …"`.

The default for the `sanitize` attribute is `"yes"`.

To disable the checking of HTML content send to form elements, set the `sanitize`
style attribute to `"no"`.

The `sanitize` style attribute makes only sense for `TextEdit` form
items, when used with the `textFormat` style attribute set to
`"html"`:

```
<Style name="TextEdit.relax">
     <StyleAttribute name="textFormat" value="html"/>
     <StyleAttribute name="sanitize" value="no"/>
</Style>
```

## `showEditToolBox`

Defines if the toolbox for rich text editing is shown or not.

Only available when the `textFormat` style attribute is set to `"html"`.

Possible values are:

- `"no"` (default): The toolbox is hidden.
- `"yes"`: The toolbox is visible.

For more details, see [Rich Text Editing in TEXTEDIT](2252-rich-text-editing-in-textedit.md "The TEXTEDIT form item provides a rich text editing feature based on HTML.").

## `textFormat`

Defines the rendering of the content of the `TEXTEDIT` widget.

Values can be:

- `"plain"` (default): the value assigned to this widget is interpreted as plain
  text.
- `"html"`, the value is interpreted as HTML (with hyperlinks), with rich text
  input feature enabled.
  > **Important:**
  >
  > The HTML content displayed inside a form element
  > using the `textFormat=html` style must not be a complete HTML document (using CSS
  > styles for example). The content must be simple HTML, with basic tags such as text decoration like
  > `<b/>` for bold, `<ul/>+<li/>` for bullet lists, and inline
  > styles.

  Do not mix `textFormat=html` style attribute with [`UPSHIFT`](1838-upshift-attribute.md "The UPSHIFT attribute forces character input to uppercase letters.") or [`DOWNSHIFT`](1776-downshift-attribute.md "The DOWNSHIFT attribute forces character input to lowercase letters.") form field attributes.

  For more details, see [Rich Text Editing in TEXTEDIT](2252-rich-text-editing-in-textedit.md "The TEXTEDIT form item provides a rich text editing feature based on HTML.").

## `wrapPolicy`

Defines where the text can be wrapped in word wrap mode.

Values can be:

- `"atWordBoundary"` (default): the text will wrap at word boundaries.
- `"anywhere"`: the text breaks anywhere, splitting words if needed.

## Related links

**Related concepts**  

[TEXTEDIT item type](1704-textedit-item-type.md "Defines a multi-line edit field.")
