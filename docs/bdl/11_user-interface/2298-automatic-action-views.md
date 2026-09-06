---
title: "Automatic action views"
source: "fgl-topics/c_fgl_prog_dialogs_auto_action_views.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Automatic action views"
type: "concept"
---

# Automatic action views

> Action views can be rendered automatically in some form elements.

Form toolbars and topmenu elements can define specific items, to get automatic rendering of
action views, for actions that exist for the current dialog.

With this feature, actions views can appear at the same time in the topmenu, toolbar and in the
action panel. In a typical use case of automatic action views, you want to hide the action panel
with the [`actionPanelPosition=none`](1657-window-style-attributes-action-panel.md) / [`ringMenuPosition=none`](1658-window-style-attributes-ring-menu.md) style attributes for windows.

In a `TOOLBAR` definition, add an `AUTOITEMS (CONTENT=ACTIONS)`
element in the definition:

```
TOOLBAR
  ...
  AUTOITEMS (CONTENT=ACTIONS)
  ...
END
```

In a `TOPMENU` definition, add an `AUTOCOMMANDS (CONTENT=ACTIONS)`
element in the definition:

```
TOPMENU
  ...
  GROUP actions (TEXT="Actions")
      AUTOCOMMANDS (CONTENT=ACTIONS)
  END
  ...
END
```

Additionally, `CONTENT=WINDOWS` or `CONTENT=PROGRAMS` options can
be used, to respectively get the list of current available windows, or the current running
applications. The end user can then switch between windows and applications from these
options:

```
TOPMENU
  ...
  GROUP windows (TEXT="Windows")
    AUTOCOMMANDS (CONTENT=WINDOWS)
  END
  GROUP programs (TEXT="Programs")
    AUTOCOMMANDS (CONTENT=PROGRAMS)
  END
  ...
END
```

For more details about managing applications and windows, see also [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").

The following screenshot shows the rendering of a `TOOLBAR` with the three kind of
`AUTOITEMS` elements:

```
TOPMENU
  ITEM open ( IMAGE="open", TEXT="Open" )
  SEPARATOR
  AUTOITEMS ( CONTENT=ACTIONS )
  SEPARATOR
  AUTOITEMS ( CONTENT=PROGRAMS )
  SEPARATOR
  AUTOITEMS ( CONTENT=WINDOWS )
  SEPARATOR
  ITEM help ( IMAGE="help", TEXT="Help" )
END
```

![TOOLBAR with AUTOITEMS rendering on desktop](../_images/toolbar_autoitems_1.jpg)

*TOOLBAR with AUTOITEMS*

The following screenshot shows the rendering of a `TOPMENU` with the three kind of
`AUTOITEMS` elements:

```
TOPMENU
GROUP file (TEXT="File")
  COMMAND startapp ( TEXT="App Starter" )
  COMMAND open ( IMAGE="open", TEXT="Open" )
  SEPARATOR
  AUTOCOMMANDS ( CONTENT=ACTIONS )
  SEPARATOR
  AUTOCOMMANDS ( CONTENT=PROGRAMS )
  SEPARATOR
  AUTOCOMMANDS ( CONTENT=WINDOWS )
END
GROUP gz (TEXT="Help")
  COMMAND about ( TEXT="About" )
END
END
```

![TOPMENU with AUTOITEMS rendering on desktop](../_images/topmenu_autocommands_1.jpg)

*TOPMENU with AUTOCOMMANDS*

## Related links

**Related concepts**  

[Topmenus](1866-topmenus.md "Topmenus define typical pull-down menus that appear at the top of application forms.")

[Toolbars](1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")
