---
title: "Action views in chromebar"
source: "fgl-topics/c_fgl_ui_mobile_action_with_ur.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Action views in chromebar"
type: "concept"
---

# Action views in chromebar

> Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.

Distinguish the different concepts using the "chrome" term:

- The web-browser chrome includes the control widgets of the web browser window
  (menus, toolbars, scroll bars and URL address bar), surrounding the HTML content.
- The Genero front-end chrome includes all the decoration arround the application
  forms, including the chromebar, the sidebars, etc.
- The chromebar is the control bar of the GBC front-end, which is by default
  displayed with a blue background on the top of the application forms.

On iOS and Android™ mobile devices,
the [default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") of the
action panel and ring menu panel, as well as the [toolbar action
views](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.") can be displayed in the chromebar.

On mobile devices, this will follow the Material Design specification.

![Screenshot of an Android device, with toolbar action views displayed in the chromebar](../_images/gbc_chromebar_01.jpg)

*Android device display with default action views in the chromebar*

By default, on mobile devices, these action views are implicitly displayed in the
chromebar. This rendering corresponds to the following style attribute
settings:

```
<Style name="Window">
  <StyleAttribute name="actionPanelPosition"  value="chrome" />
  <StyleAttribute name="ringMenuPosition"     value="chrome" />
  <StyleAttribute name="toolBarPosition"      value="chrome" />
</Style>
```

In the chromebar, action views will be rendered in the following order:

1. Toolbar action views
2. Default action views of the action panel or ring menu panel
3. Common front-end options (Application information, Settings, Bookmarks, Close window)

When there is not enough room in the chromebar, the action views will be rendered in
a vertical drop down menu that can be opened from a three-dots button on the right. This drop down
menu will replace the default drop down menu that shows up on small webviews.

On mobile, to get the same default rendering as on a desktop browser, use following settings:

```
<Style name="Window">
  <StyleAttribute name="actionPanelPosition"  value="right" />
  <StyleAttribute name="ringMenuPosition"     value="right" />
  <StyleAttribute name="toolBarPosition"      value="top" />
</Style>
```

When defining an action with the name "`back`", and the chromebar is visible, a
back button will be automatically displayed on the left of the chromebar. For more
details, see [Implementing the back action](2290-implementing-the-back-action.md "The back action is a predefined action dedicated to move back in the stack of windows/forms.").

## Related links

**Related concepts**  

[Toolbars on mobile devices](1861-toolbars-on-mobile-devices.md "Toolbars can be used to control action view rendering on mobile devices.")

**Related reference**  

[Window style attributes: Basics](1655-window-style-attributes-basics.md "Basic presentation style attributes for window elements.")

[Window style attributes: Action Panel](1657-window-style-attributes-action-panel.md "Presentation style attributes that apply to the window action panel.")

[Window style attributes: Ring Menu](1658-window-style-attributes-ring-menu.md "Presentation style attributes that apply to a window ring menu.")
