---
title: "Background/foreground modes"
source: "fgl-topics/c_fgl_mobile_enter_background_foreground.html"
breadcrumb: "Mobile applications > App execution > Background/foreground modes"
type: "concept"
---

# Background/foreground modes

> Describes how to handle background or foreground modes in mobile apps.

## Mobile apps foreground and background modes

Mobile apps can change their state from background mode to foreground mode and vice-verse. For
example, when the user switches to another app, or when going back to the home screen, the current
app goes to background mode. An app goes to foreground mode, when it is re-selected from the active
apps list.

## Detecting foreground and background mode

Genero BDL provides two predefined actions, to detect when the state of a mobile app changes, or
when the browser window/tab gets hidden or shown to the end user:

1. The action `enterbackground` is fired, when the mobile app goes to background
   mode, or when the browser window or current tab gets hidden to the user, because a browser
   window/tab switch occurred, or the browser window got minimized. With GDC/UR, the
   `enterbackground` action is fired when the window container is minimized.
2. The action `enterforeground` is fired, when the mobile app goes to foreground
   mode, or when the browser window or current tab is shown to the user, because a browser window/tab
   switch occurred, or the browser window is restored. With GDC/UR, the
   `enterforeground` action is fired when the window container is restored.

The `enterforeground` action is not fired when the app starts. This action is only
fired when returning to foreground mode, after it was in background mode.

To execute code when the app goes to background or foreground mode, use on `ON
ACTION` handler:

```
   ON ACTION enterbackground
      LET skip_timers = TRUE
   ON ACTION enterforeground
      IF NOT ask_password() THEN
         EXIT PROGRAM
      ENF IF
```

For example, when the mobile app enters background mode, it is recommended that the program
suspend any activity, and skip code that would be executed in [`ON TIMER`](../11_user-interface/2227-get-program-control-on-a-regular-timed-basis.md "Execute some code after a given number of seconds, with or without user interaction with the program.") triggers.

On the other hand, when the mobile app enters foreground mode, the program can for example ask
the user's login/password again, for security reasons, to make sure that the mobile device did not
end up in other hands while the app was in background mode.

To control action view rendering defaults and current field validation behavior when the
`enterforeground` / `enterbackground` actions are used, consider
setting action default attributes for these action in your [.4ad file](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") as follows (like done in
FGLDIR/lib/default.4ad):

```
<ActionDefaultList>
  ...
  <ActionDefault name="enterbackground" validate="no" defaultView="no" contextMenu="no"/>
  <ActionDefault name="enterforeground" validate="no" defaultView="no" contextMenu="no"/>
  ...
</ActionDefaultList>
```

Another option is to define these action defaults attributes in the `ON ACTION`
handlers:

```
ON ACTION enterbackground ATTRIBUTES(VALIDATE=NO, DEFAULTVIEW=NO)
   ...
ON ACTION enterforeground ATTRIBUTES(VALIDATE=NO, DEFAULTVIEW=NO)
   ...
```

See also [List of predefined actions](../11_user-interface/2257-list-of-predefined-actions.md) and [Background/foreground modes](5094-background-foreground-modes.md "Describes how to handle background or foreground modes in mobile apps.").

## Checking if the app is currently in foreground mode

To know the current state of a mobile app, use the [`standard.isForeground`](../15_library-reference/3405-standard-isforeground.md "Indicates if the app is in foreground mode.")
front call:

```
DEFINE fg BOOLEAN
CALL ui.Interface.frontCall("standard", "isForeground", [], [fg] )
IF fg THEN
   ...
END IF
```

## What does the app do when in background mode?

On iOS, when the app is in background mode, the program cannot do anything meaningful outside
push notifications or audio play.

On Android™, when the app is in
background mode, the program can continue its execution.

## Deny Android to terminate app when in background

Genero programs running on servers are typically not prepared to be stopped at any time (except
in case of major failure): it's the program that decides when it terminates.

On Android devices, an app can switch between foreground
to background states, and the system can decide to stop an app in background state, to release
resources.

By default, when the Genero mobile app goes to background state, a notification is shown by GMA
to keep the app running, and avoid Android stopping the app. The notification disappears, when the app returns to foreground
state.

To allow Android to stop your app
when it is in background state, set the `androidKeepForeground` style attribute to
`"no"` at the [`UserInterface`
element type](../11_user-interface/1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface."). In this case, GMA will not display a notification, when the app switches to
background mode.

When using `androidKeepForeground=no`, make sure that your code is ready to be
stopped at any time.
