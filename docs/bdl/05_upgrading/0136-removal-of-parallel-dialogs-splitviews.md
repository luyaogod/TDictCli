---
title: "Removal of parallel dialogs / splitviews"
source: "fgl-topics/c_fgl_Migrate_to_400_paradlg_desupport.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Removal of parallel dialogs / splitviews"
type: "concept"
---

# Removal of parallel dialogs / splitviews

> Parallel dialog and splitview features for mobile are no longer supported.

Starting with version 4.00, the concept of parallel dialogs and windows as
splitviews is desupported.

Parallel dialog programming involve the following language elements:

- `fgl_eventloop()`
- `START DIALOG`
- `TERMINATE DIALOG`
- `OPEN WINDOW ... ATTRIBUTES(TYPE=...)`

> **Note:**
>
> Parallel dialogs also involve the concept of [declarative dialogs](../11_user-interface/2150-declarative-dialogs-dialog-at-module-level.md "DIALOG/END DIALOG defined at module level implement declarative dialogs that can be used in procedural dialogs."). However, this
> concept is still valid and supported, to implement reusable dialog components with the [`SUBDIALOG`](../11_user-interface/2087-the-subdialog-clause.md) clause in
> `DIALOG` blocks.

Parallel dialogs were introduced to control windows as splitviews for mobile applications.

Splitviews are specific to native mobile GUI rendering. With GBC Universal Rendering, splitviews
are no longer supported, making also parallel dialogs obsolete (parallel dialogs could only be used
to implement splitviews)

As a replacement for parallel dialogs and splitviews, use regular multiple-dialog instructions
with [`DIALOG` blocks](../11_user-interface/2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form."), and
define forms with [`HBOX`
containers using the `SPLIT` attribute](../11_user-interface/1547-horizontal-box-splitting.md "HBOX/VBOX containers can be defined to display a single child container."):

```
HBOX (SPLIT)
GRID
...
END
TABLE
...
END
END
```

> **Note:**
>
> With the former splitview solution, you could only define a left (`TYPE=LEFT`)
> and a right (`TYPE=RIGHT`) side pane. With `HBOX` +
> `SPLIT`, one can now define several sub-containers for the horizontal box. For [responsive layout suppport](../11_user-interface/1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities."), the
> `SPLIT` attribute can also get a `@screen-size`
> modifier, to show all sub-containers simultaneously on large screens.

Windows defined with the `TYPE=POPUP` attribute can be replaced by a window style
attribute `windowType="modal"` for example by defining the [`STYLE="dialog"`
attribute](../11_user-interface/1567-configuring-windows-with-styles.md "Use the STYLE attribute to set a style for a window."):

```
OPEN WINDOW w_zoom WITH FORM "zoom" ATTRIBUTES(STYLE="dialog")
```

Windows defined with the `TYPE=NAVIGATOR` attribute must be replaced by [`TOOLBAR`](../11_user-interface/1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms."), with buttons bound to [`ON ACTION`](../11_user-interface/2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") handlers, that give the
focus to a field ([`NEXT
FIELD`](../11_user-interface/2245-giving-the-focus-to-a-form-element.md "How to force the focus to move or stay in a specific form element using program code.")), in one of the child containers of the `HBOX` with
`SPLIT` attribute. To get a similar look as the former
`TYPE=NAVIGATOR` window, use style attributes to put the [toolbar at the bottom](../11_user-interface/1655-window-style-attributes-basics.md) of the window, and [equally justify](../11_user-interface/1651-toolbar-style-attributes.md) toolbar buttons:

```
  <Style name="Window.splitview">
     <StyleAttribute name="windowType" value="normal" />
     <StyleAttribute name="actionPanelPosition" value="none" />
     <StyleAttribute name="ringMenuPosition" value="none" />
     <StyleAttribute name="toolBarPosition" value="bottom" />
  </Style>
  <Style name="ToolBar">
     <StyleAttribute name="itemsAlignment" value="justify" />
  </Style>
```

> **Note:**
>
> The syntax and instructions for parallel dialogs implementation is still available. This
> includes the `START DIALOG`, `TERMINATE DIALOG`,
> `fgl_eventloop()`, `WINDOW TYPE` attribute. These must now be
> considered as an experimental feature and are no longer documented.
