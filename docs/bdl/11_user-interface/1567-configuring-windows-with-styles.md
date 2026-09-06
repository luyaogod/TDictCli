---
title: "Configuring windows with styles"
source: "fgl-topics/c_fgl_windows_and_forms_010.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Configuring windows with styles"
type: "concept"
---

# Configuring windows with styles

> Use the STYLE attribute to set a style for a window.

The window style defines the type of the window (regular or popup/modal) and its decoration, via
the presentation style attribute (`STYLE`). The presentation style specifies a set of
style attributes in an external file (.4st).

By default, windows are displayed as regular (`"normal"`) application windows. The
type of window is defined by the [`windowType`](1655-window-style-attributes-basics.md) style attribute. This type can be set to
`"normal"` (for regular windows), `"modal"` (for popup windows) and
`"modalOnLargeScreen"` (to adapt to the screen size)

There are different ways to define the style of a window: The [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") can be used in the [`OPEN WINDOW`](1572-open-window.md "Creates and displays a new window.") instruction to
define the default style for a window, or it can be specified in the form file, with the [`WINDOWSTYLE`](1851-windowstyle-attribute.md "The WINDOWSTYLE attribute defines the style to be used by the parent window of a form.") attribute of the
`LAYOUT` section.

> **Important:**
>
> In GUI mode, form elements can also be decorated with
> presentation styles. Pay attention to the specific rules that apply when [combining TTY attributes and presentation
> styles](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.").

Genero BDL provides a set of predefined window styles in the
FGLDIR/lib/default.4st file. It is recommended that you not change the
default settings of windows styles in the $FGLDIR/lib/default.4st file. If you
create your own style file, copy the default styles into your own file in a different directory.

It is not possible to change the presentation style attributes of windows dynamically in the AUI
tree. The style is applied when the window and form are loaded.

If you open and display a second form in an existing window, the window style of the second form
is not applied.

| Style name in 4st file | Description |
| --- | --- |
| `Window` | Defines presentation attributes for regular application windows. |
| `Window.main`, `Window.main2` | Defines presentation attributes for starter applications, where the main window shows a start menu if one is defined by the application. |
| `Window.dialog`, `Window.dialog2`, `Window.dialog3`, `Window.dialog4` | Defines presentation attributes for modal windows. |
| `Window.naked` | Defines presentation attributes for windows that do not show the default view for ring menus and action buttons (OK/Cancel). |
| `Window.viewer` | Defines presentation attributes for viewers as the report pager (fglreport.per). |

## Related links

**Related concepts**  

[Start menus](2446-start-menus.md "Start menus define a tree of application programs that can be started.")

[The abstract user interface tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")

[Understanding presentation styles](1608-understanding-presentation-styles.md "Presentation styles centralize the attributes related to the decoration of the graphical user interface elements.")

**Related reference**  

[Window style attributes](1654-window-style-attributes.md "Window presentation style attributes apply to a window element.")
