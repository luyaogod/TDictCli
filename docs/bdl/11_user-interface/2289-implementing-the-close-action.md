---
title: "Implementing the close action"
source: "fgl-topics/c_fgl_prog_dialogs_close_action.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Implementing the close action"
type: "concept"
---

# Implementing the close action

> The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button).

## Purpose of the `close` action

With GDC/UR on desktop, the "`close`" action can be triggered by pressing Alt+F4
or by clicking the `X` cross button in the upper-right corner of the desktop
window.

![Window close button on desktop](../_images/WindowCloseButton1_gbc.jpg)

*Window close button on desktop*

With GBC in a web browser, an `X` button is displayed on the right of the [chromebar](2287-action-views-in-chromebar.md "Default action views and toolbar action views can be displayed in the chromebar, to save space on small screens."), to trigger the "`close`"
action:

![Chromebar close button in web browser](../_images/ChromeBarCloseButton1_gbc.jpg)

*Chromebar close button in web browser*

Genero BDL has a predefined "`close`" action name dedicated to this specific
window closing event.

When the end user triggers an `X` close button, the event can be managed in the
program code with an `ON ACTION close` block. If this action handler is not defined
in the current dialog, the VM implements a default behavior for the close action, depending on the
type of dialog.

The "`close`" action can be disabled with
`DIALOG.setActionActive("close",FALSE)`. With GDC/UR desktop, a click on the x-cross
button of the desktop window will have no effect. If the chromebar is visible, the action views for
the "`close`" action is made invisible, while the default action view button in the
action panel is disabled, if present.

Regarding the [default action
view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") handling of the "`close`" action: Independently from the chromebar
visibility, the [`DEFAULTVIEW`](2273-defaultview-action-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.") attribute is taken into account to hide or
show the default action view button in the action panel. Note that the action panel can be shown in
the chromebar when [`actionPanelPosition="chrome"`](1657-window-style-attributes-action-panel.md) style attribute.

## The `close` action in a `DIALOG` multiple dialog block

The `close` action executes an `ON ACTION close` block, if defined.
Otherwise, the `close` action is mapped to the `cancel` action, when
an `ON ACTION cancel` handler is defined.

If neither `ON ACTION close`, nor `ON ACTION cancel` are defined,
nothing will happen if the user tries to close the window with the `X` cross button
or an `ALT+F4` keystroke.

The `int_flag` register will not
be set in the context of `DIALOG`.

## The `close` action in singular data handling dialogs

By default, for `INPUT`, `INPUT ARRAY`, `CONSTRUCT`,
`DISPLAY ARRAY` or `PROMPT` dialogs, the VM will create a
`close` action, except if the singular dialog is defined with
`ATTRIBUTES(CANCEL=FALSE)`.

When an `ON ACTION close` handler is explicitly defined in an
`INPUT`, `INPUT ARRAY`, `CONSTRUCT`, `DISPLAY
ARRAY` or `PROMPT` dialog, the handler code will be executed if the
`close` action is fired.

If no explicit `ON ACTION close` handler is defined, the `close`
action acts the same as the `cancel` predefined action: When the user clicks the
`X` cross button, the dialog stops and the `int_flag` is set to 1.

If there is an explicit `ON ACTION cancel` block defined,
`int_flag` is set to 1 and the user code under `ON ACTION cancel`
will be executed.

If the `CANCEL=FALSE` option is set, no `cancel` and no
`close` action will be created, and you must write an `ON ACTION
close` handler to proceed with the `close` action. In this case, the
`int_flag` register will not be set when the `close` action is
invoked.

## The `close` action in MENU dialogs

By default, for the `MENU` dialog, the VM will create a `close`
action, if an `COMMAND KEY(INTERRUPT)` or `ON ACTION cancel` block is
defined by the `MENU`.

When an `ON ACTION close` handler is explicitly defined in a `MENU`
dialog, the handler code will be executed if the `close` action is fired.

If no explicit `ON ACTION close` action handler is defined, the code of the
`COMMAND KEY(INTERRUPT)` or `ON ACTION cancel` will be executed,
if defined.

If neither `COMMAND KEY(INTERRUPT)` nor `ON ACTION cancel` is
defined:

- If the `MENU` uses the [default action views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") rendering, or with the attribute `STYLE="dialog"`,
  nothing happens and the program stays in the `MENU` instruction.
- If the `MENU` is defined with the attribute [`STYLE="popup"`](1909-rendering-modes-of-a-menu.md), there is no `X` cross button to click.
  However, the Escape key, or a click outside the popup menu will terminate the `MENU`
  dialog.

The `int_flag` register is not
touched by a `MENU` instruction.

## The `close` action on mobile devices

When displaying on Android™ and iOS mobile device, the
`close` action can be bound to the "Back button". For more details, see [The Back button](2297-action-views-on-mobile-devices.md)

## Example

Implementing a `close` action handler, to open a confirmation dialog box, before
leaving the dialog:

```
DIALOG ...
  INPUT BY NAME cust_rec.*
  END INPUT
  ...
  ON ACTION close
    IF msg_box_yn("Close this window?") == "y" THEN
      EXIT DIALOG
    END IF
  ...
END DIALOG
```

## Related links

**Related concepts**  

[Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")
