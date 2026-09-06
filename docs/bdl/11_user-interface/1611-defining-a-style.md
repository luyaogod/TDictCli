---
title: "Defining a style"
source: "fgl-topics/c_fgl_presentation_styles_005.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Defining a style"
type: "concept"
---

# Defining a style

> Styles can be defined to be global (for all elements), for an element in general, or for specific types of an element.

The style is identified by the `name` attribute, that can be a combination of
element type, style name and pseudo selector, or the star character. See [Syntax of presentation styles file (.4st)](1609-syntax-of-presentation-styles-file-4st.md "A .4st presentation styles file is an XML file defining style attributes to be applied by front-ends.") for a complete description of the presentation style
definition syntax.

In the definition of a style, the `name` attribute is used as a selector to
apply style attributes to graphical elements.

The identifiers used in the definition of a style are case-sensitive.

You can define a style as global or specific to a class of graphical object:

- A style identified by a star (`*`) is a global style that is automatically
  applied to all elements:

  ```
  <Style name="*" >
  ```
- A style identified by an element-type is a global style that is
  automatically applied to all objects of this type:

  ```
  <Style name="ComboBox" >
  ```
- A style identified by a style-name is a specific style that can be applied
  to any element types using that style name in a `STYLE` attribute:

  ```
  <Style name=".important" >
  ```
- A style identified by an element-type followed by a dot and a
  style-name is a specific style that will only be applied to elements of the
  given type and using the style name in a `STYLE` attribute:

  ```
  <Style name="Window.main" >
  ```
- A style identified by an element-type followed by a colon and a
  pseudo-selector is a style that will only be applied to elements of the
  given type, if the condition defined by the pseudo-selector is satisfied:

  ```
  <Style name="Edit:focus" >
  ```
- A style identified by an element-type followed by a dot and a
  style-name, and a colon with a pseudo-selector, is a specific style that
  will only be applied to elements of the given type, using the style name in a
  `STYLE` attribute, if the condition defined by the pseudo-selector is
  satisfied:

  ```
  <Style name="Edit.important:focus" >
  ```
- It is possible to combine pseudo-selectors:

  ```
  <Style name="Edit:query:focus" >
  ```
