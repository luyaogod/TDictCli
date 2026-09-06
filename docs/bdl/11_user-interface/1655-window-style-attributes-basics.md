---
title: "Window style attributes: Basics"
source: "fgl-topics/r_fgl_presentation_styles_window_style_attributes_basics.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > Window style attributes > Window style attributes: Basics"
type: "reference"
---

# Window style attributes: Basics

> Basic presentation style attributes for window elements.

Style attributes listed in this page are general and supported on several front-end types, when
the platform allows it. See also other pages about window style attributes.

## `forceDefaultSettings`

Indicates if the window content must be initialized with the saved positions and sizes. By default,
windows are reopened at the position and with the size they had when they were closed. You can force
the use of the initial settings with this attribute. This applies also to column position and width
in tables.

Values can be `"yes"` or `"no"` (default).

## `position`

Indicates the position of the window.

Values can be:

- `"default"` (default): the window is displayed in the window container, as
  defined by the container rules. For more details, see [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").
- `"field"`: the window is displayed below the current field. If there is no
  current field, `"default"` position applies. The `windowType` style
  attribute must be "`modal`". For more details, see [Field-anchored picklist](2251-field-anchored-picklist.md "Drop-down picklist windows can be displayed under the current field.").
- `"center"`: the window is displayed in the center of the window container.
  Typically used when the `windowType` style attribute set to "`modal`".

## `sizable`

Defines if the window can be resized by the user.

Values can be `"yes"` (default), `"no"` or `"auto"`.

> **Note:**
>
> - On a normal window, the behavior is applied to the form instead of the window. When set to
>   `"no"`, the form content is not stretched, even if the form contains stretchable
>   items.
> - On a modal window, when set to `"no"`, the end-user cannot resize the modal
>   window.

## `tabbedContainer` (deprecated)

> **Important:**
>
> `tabbedContainer` is deprecated and replaced by the new user interface elements to
> control application windows. See [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.") for more
> details.

Values can be `"yes"` or `"no"` (default).

When the main window of the first started application uses `tabbedContainer=yes`,
the parent application and every child application started with `RUN ... WITHOUT
WAITING` gets an app button in the ChromeBar, like when using the [`applicationListPosition=top`](1652-userinterface-style-attributes.md) style attribute for the
`UserInterface` class.

## `thinScrollbarDisplayTime`

Defines the display time (in seconds) of the automatic scrollbar displayed when scrolling on
fixed screen array (a.k.a. "Matrix") and SCROLLGRIDs. After the delay, the scrollbar disappears.

A value of zero specifies an infinite time: The thin scrollbar remains visible while the record
list can be scrolled (during dialog execution).

A negative value specifies that the scrollbar must always be hidden.

Default is 1 second.

## `toolBarPosition`

Indicates the default position of the [toolbar](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms."), when
such element is defined by the form, and no specific [`position`](1651-toolbar-style-attributes.md) style attribute is defined for toolbar elements.

Values can be one of:

- `"top"` (default desktop): Toolbar is displayed at the top of the window.
- `"bottom"`: Toolbar is displayed at the bottom of the window.
- `"none"`: The toolbar is not displayed.
- `"chrome"` (default on mobile): Action views are displayed
  in the chromebar. See [Action views in chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.").

## `topmenuDesktopRendering`

Indicates the rendering of the [topmenu](1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.") on desktop
front-ends, when such element is defined by the form.

Values can be one of:

- `"classic"` (default on desktop): Topmenu is displayed as a tree of options on
  top of the window, all items of the first group appear horizontally in the topmenu bar.
- `"sidebar"` (default on mobile): Topmenu is displayed as a popup treeview, when
  pressing the dedicated hamburger button on the top left of the chromebar, assuming the chromebar is
  displayed.

This style attribute applies on window topmenus, but has an impact on global
topmenus as well.

## `topmenuMobileRendering`

Indicates the rendering of the [topmenu](1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.") on mobile
front-ends, when such element is defined by the form.

Values can be one of:

- `"classic"` (default on desktop): Topmenu is displayed as a tree of options on
  top of the window, all items of the first group appear horizontally in the topmenu bar.
- `"sidebar"` (default on mobile): Topmenu is displayed as a popup treeview, when
  pressing the dedicated hamburger button on the top left of the chromebar, assuming the chromebar is
  displayed.

This style attribute applies on window topmenus, but has an impact on global
topmenus as well.

## `windowOptionClose`

Defines if the window can be closed with a system menu option or window header button.

Values can be `"yes"`, `"no"` or `"auto"` (default).

When value is `"auto"`, the option is enabled depending on the window type.

This attribute may have different behavior depending on the front-end operating system. For example,
when no system menu is used, it may not be possible to have this option enabled.

## `windowState`

Defines the display state of the window container when using the GDC front-end.

The `windowState` style attribute is defined for BDL `WINDOW`
objects. However, it applies to the [window container](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), following the `windowState` style attribute
defined for the first `WINDOW` displayed by the program.

Values can be:

- The following options apply, unless stored settings have a different window state registered,
  or when stored settings are disabled:
  - `"normal"` (default): The window container displays in normal state (first
    window/form layout defines the window container size)
  - `"fullscreen"`: The window container uses the whole screen: The window decoration
    is removed, the size is set to the screen size, and its position is set to the top-left corner of
    the screen.
  - `"maximized"`: The window container is maximized: The window size (including its
    decoration) is set to the screen size and its position is set to the top-left corner of the
    screen.
  - `"minimized"`: The window container is minimized.
- The following options always apply, even when stored settings are active and define a different
  window state:
  - `"forceFullscreen"`: The window container uses the whole screen.
  - `"forceMaximized"`: The window container is maximized.
  - `"forceMinimized"`: The window container is minimized.

## `windowType`

Defines the basic type of the window.

Values can be:

- `"normal"` (default): Normal windows are displayed as typical application
  windows.
- `"modal"`: Modal windows are "popup" windows displayed over all
  other windows, typically used for temporary dialogs.
- `"modalOnLargeScreen"`: The window type adapts to the size of the screen, and
  displays as `"modal"` on screens with a size that matches the [`@LARGE`](1542-understanding-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities.") size modifier of form
  definition attributes. For smaller sizes, the window is displayed as `"normal"`, to
  fill all the available container space. The window will keep its type (`"normal"` or
  `"modal"`) until it is closed, even if the container is resized or the device
  orientation changes.

## Related links

**Related concepts**  

[Windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.")
