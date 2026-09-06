---
title: "LAYOUT section"
source: "fgl-topics/c_fgl_FormSpecFiles_LAYOUT_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section"
type: "concept"
---

# LAYOUT section

> The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.

## Syntax

```
LAYOUT [ ( layout-attribute [,...] ) ]
  root-container
     child-container
     [...] 
  END
[END]
```

1. layout-attribute is an attribute for the whole form.
2. root-container is the first container that holds child-containers.

## Form attributes

[`IMAGE`](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item."), [`MINHEIGHT`](1799-minheight-attribute.md "The MINHEIGHT attribute defines the minimum height of a form."), [`MINWIDTH`](1800-minwidth-attribute.md "The MINWIDTH attribute defines the minimum width of a form."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."), [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element."), [`TEXT`](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item."), [`TAG`](1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string."), [`VERSION`](1846-version-attribute.md "The VERSION attribute is used to specify a user version string for an element."), [`WINDOWSTYLE`](1851-windowstyle-attribute.md "The WINDOWSTYLE attribute defines the style to be used by the parent window of a form.").

## Style attributes (form)

Common: [`backgroundColor`](1629-style-attributes-common-to-all-elements.md), [`border`](1629-style-attributes-common-to-all-elements.md), [`fontFamily`](1629-style-attributes-common-to-all-elements.md), [`fontSize`](1629-style-attributes-common-to-all-elements.md), [`fontStyle`](1629-style-attributes-common-to-all-elements.md), [`fontWeight`](1629-style-attributes-common-to-all-elements.md), [`textColor`](1629-style-attributes-common-to-all-elements.md), [`textDecoration`](1629-style-attributes-common-to-all-elements.md).

See also: [Window style attributes](1654-window-style-attributes.md "Window presentation style attributes apply to a window element."), [UserInterface style attributes](1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface.").

## Can hold

`FORM`, `VBOX`, `HBOX`, `GROUP`, `FOLDER`, `GRID`, `SCROLLGRID`, `TABLE`, `TREE`.

## Usage

The `LAYOUT` section defines a tree of layout containers such as
`VBOX`, `HBOX`, `GRID`, `TABLE`.

The `LAYOUT` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

This section is mandatory, unless you use a [`SCREEN` section](1714-screen-section.md "The SCREEN section defines the form layout for TUI mode forms."). However, the purpose of the `SCREEN` section
is to design forms for the TUI mode.

Indentation is supported in the `LAYOUT` section.

The `END` keyword is optional.

The layout tree of the form is defined by associating layout containers. Different
kinds of layout containers are provided, each of them having a specific role. Some
containers such as `VBOX`, `HBOX` and
`FOLDER` can hold children containers, while others such as
`GRID` and `TABLE` define a screen area. Containers
using a screen area define a formatted region containing static text labels, item
tags and layout tags. External form files can be included in the current layout with
the `FORM` clause.

```
LAYOUT (VERSION="12", STYLE="regular")
  VBOX
    GRID grid1
      grid-area
    END
    GROUP group1
      HBOX
        GRID grid2
          grid-area
        END
        TABLE table1
          table-area
        END
      END
    END
  END 
END
```

The definition would result in a layout tree that looks like
this:

```
-- VBOX
   |
   +-- GRID grid1
   |
   +-- GROUP group1
       |
       +-- HBOX
           |
           +-- GRID grid2
           |
           +-- TABLE table1
```

The layout section can also contain a simple `GRID` container
(equivalent to a V3 SCREEN
definition):

```
LAYOUT
  GRID
    grid-area
  END 
END
```

## Description of LAYOUT attributes

The `VERSION` attribute can be used to specify a version for the form.
This allows you to indicate that the form content has changed. Typically used to
avoid having the front-end reload the saved window settings.

The `MINHEIGHT`, `MINWIDTH` attributes can be used to
specify a minimum width and height for the form. You typically use these attributes to force the
form to have a bigger size than the default when it is first rendered. If the front-end stores
window sizes, these attributes will only be significant the first time the form is opened, or each
time the `VERSION` attribute is changed.

The `IMAGE` attribute can be used to define the icon of the window that will
display the form. This attribute is automatically applied to the parent window node when a form is
loaded. The window icon can be changed dynamically with the [`ui.Window.setImage)`](../15_library-reference/3138-ui-window-setimage.md "Set the window icon.") method.

The `TEXT` attribute can be used to define the title of the window that will
display the form. This attribute is automatically applied to the parent window node when a form is
loaded. The window title can be changed dynamically with the [`ui.Window.setText()`](../15_library-reference/3139-ui-window-settext.md "Set the window title.") method.

The `SPACING` attribute can be used to give a hint to the front-end to define the
gap between form elements.

The `STYLE` attribute defines the presentation style for form elements, you can
for example define a font property for all form elements.

With the `WINDOWSTYLE` attribute, you can define the window style as if it was
specified in the [`OPEN
WINDOW`](1572-open-window.md "Creates and displays a new window.") instruction: The `WINDOWSTYLE` attribute is automatically
applied to the parent window when a form is loaded. For backward compatibility, in the
`LAYOUT` element definition, the `STYLE` attribute is used as the
default `WINDOWSTYLE`, if this attribute is not used.

## Related links

**Related concepts**  

[SCREEN section](1714-screen-section.md "The SCREEN section defines the form layout for TUI mode forms.")

[ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")

[Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.")

## Child topics

- [FORM clause](1716-form-clause.md): Reuse the definition of a form in the current form.
- [HBOX container](1717-hbox-container.md): Packs child layout elements horizontally.
- [VBOX container](1718-vbox-container.md): Packs child layout elements vertically.
- [GROUP container](1719-group-container.md): Defines a layout area to group other layout elements together.
- [FOLDER container](1720-folder-container.md): Defines the parent container for folder pages.
- [PAGE container](1721-page-container.md): Defines the content of a folder page.
- [GRID container](1722-grid-container.md): Defines a layout area based on a grid of cells.
- [SCROLLGRID container](1723-scrollgrid-container.md): Defines a scrollable grid view widget.
- [TABLE container](1724-table-container.md): Defines a re-sizable table designed to display a list of records.
- [TREE container](1725-tree-container.md): The TREE container defines the presentation of a list of ordered records in a tree-view widget.
