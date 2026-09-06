---
title: "Message style attributes"
source: "fgl-topics/r_fgl_presentation_styles_message_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Message style attributes"
type: "reference"
---

# Message style attributes

> Message presentation style attributes apply to an ERROR or MESSAGE instruction.

The `Message` presentation style element type can be used to set style attributes
for both [`ERROR`](1886-error.md "The ERROR instruction displays an error message to the user.") and [`MESSAGE`](1885-message.md "The MESSAGE instruction displays a message to the user.") instructions.

To distinguish `ERROR` from `MESSAGE`, use the
"`:error`" or "`:message`" pseudo-selectors:

- "`Message:error`" corresponds to the `ERROR` instruction
- "`Message:message`" corresponds to the `MESSAGE` instruction.

The `ERROR` and `MESSAGE` instructions can use a
`STYLE` attribute in the `ATTRIBUTES` clause:

```
MESSAGE "No rows have been found." ATTRIBUTES(STYLE="info")
```

A limited set of common style attributes are supported for error/message display. In
addition to the attributes described in the section, you can only define [font style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.")
for messages.

Like simple form fields, [TTY
attributes](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.") have a higher priority than style attributes. By default,
`ERROR` has the TTY attribute `REVERSE`, which explains why
`ERROR` messages have a reverse background, even when you use a [backgroundColor](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") style
attribute. Use the `NORMAL` attribute in `ERROR`, to avoid the
default `REVERSE` TTY attribute and define your own background color with a
style.

> **Tip:**
>
> Consider centralizing your `ERROR` and `MESSAGE`
> instruction calls in a function, to simplify global
> modifications:
>
> ```
> FUNCTION my_error(m, s)
>    DEFINE m, s STRING
>    IF s IS NULL THEN
>        ERROR m ATTRIBUTES(NORMAL)
>    ELSE
>        ERROR m ATTRIBUTES(NORMAL, STYLE=s)
>    END IF
> END FUNCTION
> ```

## `sanitize`

By default, to avoid "Stored XSS" attacks, the front-end cleans the HTML sent to form elements to
ensure no malicious script can be executed. This security control prevents for example to use HTML
content such as `"<a href='mailto: …"`.

The default for the `sanitize` attribute is `"yes"`.

To disable the checking of HTML content send to form elements, set the `sanitize`
style attribute to `"no"`.

The `sanitize` style attribute makes only sense for `Message` form
elements, when used with the `textFormat` style attribute set to
`"html"`:

```
<Style name="Message.relax">
     <StyleAttribute name="textFormat" value="html"/>
     <StyleAttribute name="sanitize" value="no"/>
</Style>
```

## `textFormat`

Defines the rendering of the content of the widget.

Possible values are:

- `"plain"` (default): the value assigned to this widget is interpreted as plain
  text.
- `"html"`, it is interpreted as HTML (with hyperlinks).
  > **Important:**
  >
  > The HTML content displayed inside a form element
  > using the `textFormat=html` style must not be a complete HTML document (using CSS
  > styles for example). The content must be simple HTML, with basic tags such as text decoration like
  > `<b/>` for bold, `<ul/>+<li/>` for bullet lists, and inline
  > styles.

## Related links

**Related concepts**  

[Pseudo selectors](1612-pseudo-selectors.md "Pseudo selectors can be used to apply style only when some conditions are fulfilled.")
