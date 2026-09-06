---
title: "Toolbars on mobile devices"
source: "fgl-topics/c_fgl_toolbars_mobile.html"
breadcrumb: "User interface > Form definitions > Toolbars > Toolbars on mobile devices"
type: "concept"
---

# Toolbars on mobile devices

> Toolbars can be used to control action view rendering on mobile devices.

## Toolbar action views in chromebar

By default, on a mobile device, the toolbar action views are displayed in the front-end
chromebar, which is equivalent to set the `toolBarPosition` window style
attribute to `"chrome"`:

```
<Style name="Window">
    ...
    <StyleAttribute name="toolBarPosition" value="chrome" />
</Style>
```

![Screenshot of the of an Android device, with toolbar action views displayed in the chromebar](../_images/gbc_chromebar_02.jpg)

*Android display with toolbar action views in the chromebar*

For more details about chromebar related style attributes, see [Action views in chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens.").

The toolbar can also be displayed at other positions, like the top or bottom of the
window:

```
<Style name="Window">
    ...
    <StyleAttribute name="toolBarPosition" value="top" />
</Style>
```

![Screenshot of an Android device, with toolbar at top of the windpw](../_images/gbc_toolbar_01.jpg)

*Android device display with toolbar at top of the window*

Use style attributes to control the rendering of the toolbar. See [ToolBar style attributes](1651-toolbar-style-attributes.md "ToolBar presentation style attributes apply to the TOOLBAR element.") for more details.

## Related links

**Related concepts**  

[Action views on mobile devices](2297-action-views-on-mobile-devices.md "Action views are rendered following mobile specific standards.")
