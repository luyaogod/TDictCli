---
title: "Form rendering basics"
source: "fgl-topics/c_fgl_form_rendering_basics.html"
breadcrumb: "User interface > Form definitions > Form rendering > Form rendering basics"
type: "concept"
---

# Form rendering basics

> Get the essentials about form rendering.

In graphical mode (GUI mode), application forms can display complex layouts. During program
execution, a window may be resized by the end user, or is resized because the front-end platform
allows different screen orientations like with mobile devices.

There are two approaches to define forms for your application:

- For desktop-only applications, design adaptive-layout forms, where the resizing of an
  application window only affects some resizable form elements like
  `TABLE`/`TREE`/`SCROLLGRID` containers,
  `WEBCOMPONENT`, `TEXTEDIT` and `IMAGE` fields. Such
  type of forms are typically defined with several [layout tags](1678-layout-tags.md "Layout tags define layout areas for containers inside the frame of a grid-based container.") in a single `GRID` container.
- For mobile, web and desktop applications, design responsive-layout forms, where all form
  elements can adapt to the screen size. Such type of forms are defined with a [tree of `HBOX`/`VBOX`
  containers](1543-layout-structure-for-responsive.md "Use the LAYOUT structure to define forms that adapt to screen sizes.") with containes such as `GRID`, and where element attributes are
  defined with a screen-size specifier, to control their visibility, stretchability and
  orientation.

When developing with command line tools, forms are designed with .per form
specification files, which are text files. In order to display text-based forms in graphical
screens, the .per form definitions must be converted to graphical widgets, implying specific layout
rules explained in this chapter.

With the GBC front end, use GBC Themes to adapt the form rendering to the user needs and
preferences. A [GBC Theme specific front call API](../15_library-reference/3428-theme-front-calls.md "This section describes theme handling front calls.") is
available to manage themes. For more details about GBC Themes, see "Theme reference" chapter in the
Genero Browser Client User Guide.

In GUI mode, the size of a window depends in its type. See [Window geometry in GUI mode](1566-position-and-size-of-a-window.md) for more
details.

## Related links

**Related concepts**  

[Form specification files](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
