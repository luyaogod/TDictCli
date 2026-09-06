---
title: "Understanding topmenus"
source: "fgl-topics/c_fgl_topmenus_005.html"
breadcrumb: "User interface > Form definitions > Topmenus > Understanding topmenus"
type: "concept"
---

# Understanding topmenus

> This is an introduction to topmenu definitions.

A topmenu defines a graphical menu that holds views for actions controlled in programs with
`ON ACTION` handlers.

The appearance of a topmenu can take several forms:

- The classic rendering (this is the default in desktop browsers and in desktop front-ends):

  ![TopMenu rendering on desktop](../_images/TopMenu1_gbc.jpg)

  *Classic TopMenu rendering (default on desktop)*
- The sidebar rendering (this is the default in mobile browsers and in mobile
  front-ends):

  With the sidebar rendering, a topmenu can be accessed from a hamburger
  button on the left of the chromebar. A tap on the hamburger button will display the topmenu in the sidebar:

  ![TopMenu rendering on mobile](../_images/TopMenu2_gbc.jpg)

  *Sidebar TopMenu rendering*

The rendering of a topmenu can be controlled specifically for desktop and/or mobile front-ends,
respectively with the [topmenuDesktopRendering](1655-window-style-attributes-basics.md) and [topmenuMobileRendering](1655-window-style-attributes-basics.md) style attributes.

A topmenu can be defined with XML in .4tm files, or in forms with the [`TOPMENU` section](1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions."), as
form-specific topmenus. The XML/.4tm fopmenu files can be [loaded](1870-loading-4tm-topmenu-definition-files.md "Topmenu XML definition files can be loaded at runtime.") by program with the methods
`ui.Interface.loadTopMenu()` (for default topmenus) or
`ui.Form.loadTopMenu()` (for form-initializers).

The topmenu options are controlled with the `ON ACTION` handlers defined by the
current interactive instruction. A topmenu option is bound to an action handler by name. Selecting
the topmenu option will execute the user code in the action handler.

Topmenu commands can be automatically hidden when inactive, if the [`AUTOHIDE`](1761-autohide-attribute.md "The AUTOHIDE attribute hides automatically the form element when the related action gets inactive.") attribute is specified for
the command.

Topmenu elements can get a `style` attribute in order to use a specific
rendering/decoration following [presentation style definitions](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.").

When [binding to an action](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?"),
make sure that you are using the right value in the `name` attribute. As `ON
ACTION` and `COMMAND` generate lowercase identifiers, it is recommended to
use lowercase names.

The decoration of topmenu options for common actions can be centralized with [action defaults](1600-action-defaults-files.md "Action defaults files allow to centralize action configuration parameters such as text, icon, accelerators and behavior options in XML format."). For example, to define the icon and
text for a "help" topmenu option that repeats in many topmenus.

Topmenus can get automatic options with the [`AUTOCOMMANDS`](1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions.") element, to show action views for all default action views,
running applications list and open windows list. For more details, see [Automatic action views](2298-automatic-action-views.md "Action views can be rendered automatically in some form elements.").

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
