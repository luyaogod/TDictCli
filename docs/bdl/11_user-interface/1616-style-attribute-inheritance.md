---
title: "Style attribute inheritance"
source: "fgl-topics/c_fgl_presentation_styles_009.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Style attribute inheritance"
type: "concept"
---

# Style attribute inheritance

> A style attribute may be inherited by the descendants of a given node in the abstract user interface tree.

The style inheritance is implicitly defined by the attribute. Common font and color related style
attributes are typically inherited.

For example, when defining a style using the `fontFamily` in a window or group
container, all the children of this container will get the same font family. However, most style
attributes (such as "border" for the Window element type) are specific to a type of element, and are
not inherited.

Style attribute inheritance can be overwritten by defining the same attribute for the type of
elements that are used inside the parent container. In the example, the .4st
style file defines a text color for groups and all its descendant nodes. The same attribute is then
redefined explicitly for labels and edit
nodes:

```
<StyleList>
  <Style name="Group">
     <StyleAttribute name="textColor" value="red" />
  </Style>
  <Style name="Label">
     <StyleAttribute name="textColor" value="windowText" />
  </Style>
  <Style name="Edit">
     <StyleAttribute name="textColor" value="windowText" />
  </Style>
</StyleList>
```

For more details, see [Style attributes reference](1628-style-attributes-reference.md "A presentation style attribute may be a common attribute that can be applied to any graphical element. Most presentation style attributes apply only to a specific graphical element.").
