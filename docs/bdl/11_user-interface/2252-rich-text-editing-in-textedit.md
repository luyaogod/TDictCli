---
title: "Rich Text Editing in TEXTEDIT"
source: "fgl-topics/c_rich_text_feature.html"
breadcrumb: "User interface > User interface programming > Input fields > Rich Text Editing in TEXTEDIT"
type: "concept"
---

# Rich Text Editing in TEXTEDIT

> The TEXTEDIT form item provides a rich text editing feature based on HTML.

![The figure shows a screenshot of a textedit field implementing the rich text editing feature.](../_images/richtext_gbc.jpg)

*Rich text editing interface*

To enable rich text editing, set the [`textFormat`](1650-textedit-style-attributes.md) style attribute to
`html`:

```
<Style name="TextEdit.richText">
  <StyleAttribute name="textFormat" value="html" />
</Style>
```

## Richtext toolbox

By default, the rich text editor toolbox is hidden. Use the [`showEditToolBox`](1650-textedit-style-attributes.md) style attribute, to show a toolbox at the top the the text
editor:

```
<Style name="TextEdit.richText">
  <StyleAttribute name="textFormat" value="html" />
  <StyleAttribute name="showEditToolBox" value="yes" />
</Style>
```

![The figure is a screenshot of a textedit field implementing the rich text editing feature with the toolbox shown at the top.](../_images/richtext_gbc_toolbox.jpg)

*Rich text editing interface with toolbox always displayed.*

## Related links

**Related reference**  

[TextEdit style attributes](1650-textedit-style-attributes.md "Textedit presentation style attributes apply to the TEXTEDIT element.")
