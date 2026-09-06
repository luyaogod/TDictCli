---
title: "Understanding toolbars"
source: "fgl-topics/c_fgl_toolbars_005.html"
breadcrumb: "User interface > Form definitions > Toolbars > Understanding toolbars"
type: "concept"
---

# Understanding toolbars

> This is an introduction to toolbars definition.

A toolbar defines action views presented as a set of buttons that can trigger events in an
interactive instruction.

![Toolbar rendering on desktop](../_images/ToolBar1_gbc.jpg)

*Toolbar rendering on desktop*

This section describes how to define toolbars with XML in files or in programs as
global/default toolbars; it is also possible to define toolbars in forms with the [`TOOLBAR` section](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions."), as
form-specific toolbars.

Toolbar files can be [loaded](1859-loading-4tb-toolbar-definition-files.md "Toolbar XML definition files can be loaded at runtime.") by a program with
the methods `ui.Interface.loadToolBar()` (for global/default toolbars) or
`ui.Form.loadToolBar()` (to load form-specific toolbars).

The global/default toolbar is displayed in all windows, while form-specific toolbar is displayed
in the form where it is defined.

The position and overall rendering of toolbar can be defined with presentation style attributes
of Window elements, such as [`toolBarPosition`](1655-window-style-attributes-basics.md). Typical "modal windows"
(`windowType="modal"`) do not display toolbars. See
FGLDIR/lib/default.4st, where `toolBarPosition` is set to
`"none"` for that kind of windows. Toolbar position can also be specified
individually with the [`position`](1651-toolbar-style-attributes.md) style attribute for ToolBar elements.

The toolbar items (or buttons) are enabled/disabled based on the [`ON ACTION`](2279-implementing-dialog-action-handlers.md "How to execute user code in ON ACTION blocks when an action is fired.") handlers defined by
the current interactive instruction. A toolbar item is bound to an action handler by name. A click
on the toolbar button will execute the user code in the action handler.

Toolbar items can be automatically hidden when inactive, if the [`AUTOHIDE`](1761-autohide-attribute.md "The AUTOHIDE attribute hides automatically the form element when the related action gets inactive.") attribute is specified for
the item.

Toolbar elements can get a `style` attribute in order to use a specific
rendering/decoration following [presentation style definitions](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items."). Specific decoration can be achieved with the
`aspect` and `size` toolbar-specific style attributes. For more
details, see [toolbar style
attributes](1651-toolbar-style-attributes.md "ToolBar presentation style attributes apply to the TOOLBAR element.").

The DOM tag names are case sensitive; `Toolbar` is different from `ToolBar`.

When [binding to an action](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?"), make sure
that you are using the right value in the `name` attribute. As `ON
ACTION` and `COMMAND` generate lowercase identifiers, it is recommended to
use lowercase names.

It is recommended that you define the decoration of toolbar items for common actions with action
defaults.

Toolbars can get automatic options with the [`AUTOITEMS`](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.") element, to show action views for all default action views,
running applications list and open window list. For more details, see [Automatic action views](2298-automatic-action-views.md "Action views can be rendered automatically in some form elements.").

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
