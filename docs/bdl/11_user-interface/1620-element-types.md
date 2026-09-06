---
title: "Element types"
source: "fgl-topics/c_fgl_presentation_styles_014.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Element types"
type: "concept"
---

# Element types

> Which AUI tree elements can get a style?

Styles may apply to any graphical elements of the user interface, such as
`Button`, `Edit`, `ComboBox`,
`ButtonEdit`, `Table`, `Window`.

The name of the element when used in a style file is case sensitive (use
`CheckBox`, not `checkbox`).

For example, the following style definition uses the "Window" element type in the style
name:

```
<Style name="Window.dialog">
   <StyleAttribute name="position" value="center" />
</Style>
```

The supported element types is defined by the style attributes, for more details, see
[Style attributes reference](1628-style-attributes-reference.md "A presentation style attribute may be a common attribute that can be applied to any graphical element. Most presentation style attributes apply only to a specific graphical element.").
