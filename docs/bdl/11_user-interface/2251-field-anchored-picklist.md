---
title: "Field-anchored picklist"
source: "fgl-topics/c_fgl_prog_dialogs_win_pos_field.html"
breadcrumb: "User interface > User interface programming > Input fields > Field-anchored picklist"
type: "concept"
---

# Field-anchored picklist

> Drop-down picklist windows can be displayed under the current field.

In order to show a drop-down picklist below the current field, set the
"`position`" style attribute for Window elements to
`"field"`:

```
<StyleList>
  <Style name="Window.dropdown">
     <StyleAttribute name="position" value="field" />
     ...
  </Style>
<StyleList>
```

Additional Window style attributes are required to get the expected rendering. For example, we
remove the border, system menu, statusbar, action panel, etc, to get a minimal window frame.

For a complete example, see [Field-anchored windows](1570-field-anchored-windows.md "The Window style attribute \"position\" can be set to \"field\" in order to display the window under the current field.").
