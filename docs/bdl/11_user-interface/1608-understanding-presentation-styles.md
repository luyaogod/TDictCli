---
title: "Understanding presentation styles"
source: "fgl-topics/c_fgl_presentation_styles_002.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Understanding presentation styles"
type: "concept"
---

# Understanding presentation styles

> Presentation styles centralize the attributes related to the decoration of the graphical user interface elements.

The decoration attributes are defined in a separate file, which can be easily modified to
customize the application.

Presentation styles are only supported for the GUI front-ends. If you design an application for
the TUI mode, you can use TTY attributes.

Styles are applied implicitly by using global styles, or explicitly by naming a specific style in
the `style` attribute of the element.

Common presentation attributes define font properties, foreground colors and background colors.
Some presentation attributes are specific to a given class of widgets (like the first day of the
week in a `DATEEDIT item type`).

> **Tip:**
>
> For the best performance and a consistent look & feel, consider customizing
> the GBC using theme variables to define application-wide decoration settings (instead of
> presentation styles.) For example, the background color of form fields can be defined
> using theme variables.

Presentation styles are defined in a resource file having an extension of
.4st. The .4st file must be distributed with the other
runtime files.

Presentation styles are inspired by the cascading style sheets (CSS) used in HTML,
with the following deviations:

1. The elements using style definitions are AUI tree elements; CSS styles apply to HTML
   elements.
2. To specify a style for an AUI tree element, you must use the "style" attribute; HTML/CSS use the
   "class" attribute.
3. Inline-style definition is not supported in the AUI tree.
4. Some pseudo selectors, such as "query, are specific to Genero.

The next screenshot shows a desktop application with default rendering:

![Screenshot of a form without presentation styles](../_images/styles_gbc_001.jpg)

*Form without presentation styles*

The next screenshot shows a desktop application using presentation styles:

![Screenshot of a form using presentation styles](../_images/styles_gbc_002.jpg)

*Form using presentation styles*
