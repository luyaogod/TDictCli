---
title: "UserInterface style attributes"
source: "fgl-topics/r_fgl_presentation_styles_userinterface_style_attributes.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Style attributes reference > UserInterface style attributes"
type: "reference"
---

# UserInterface style attributes

> UserInterface presentation style attributes define general options related to the application user interface.

> **Note:**
>
> This topic lists presentation style attributes for a specific class of form
> element, [common
> presentation style attributes](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") can also be used for this type of element.

## `androidKeepForeground` (GMA / Android™ specific)

Use the `androidKeepForeground` style attribute to control the way the GMA forces
Android to keep your app alive:

An Android app can switch between
foreground to background states.

The runtime system part (fglrun) of a Genero Android app can be in the foreground state and still active,
while the front-end part (GMA) is visually in the background.

By default, when the app goes to background state, a notification is shown by GMA to give a
higher priority to the app, and to prevent Android stopping the app when resources are required for other apps. The notification
disappears, when the app returns to foreground state.

Values can be:

- `"yes"` (default): The app remains in foreground state: to keep it in foreground
  state, a notification is displayed when the visual part of the app goes background.
- `"no"`: The app can switch between foreground and background state. No
  notification is displayed to force the app to remain in foreground state, and the Android system may stop the app at any
  time.

> **Important:**
>
> When using `androidKeepForeground=no`, Android may stop the app at any time. Make
> sure that your application code is ready for this.

## `applicationListPosition`

The application list is can be used to navigate between
programs running concurrently started from a parent program with `RUN
command WITHOUT WAITING`. For more details about the application list,
read [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").

The `applicationListPosition` style attribute controls the position of the
application list.

Values can be:

- `"left"` (default): Application list displayed on the left of the window
  container, in the SideBarRail
- `"top"`: Application list is displated on the top of the window
  container, in the ChromeBar, when the ChromeBar is visible.

> **Important:**
>
> This style attribute is taken into account for the first program that displays to the front-end.
> All sub-sequent settings for this attribute in new application runs do not impact the
> application list visibility and position.

Notes:

1. The `applicationListPosition` attribute is ignored, when [applicationListVisible](1652-userinterface-style-attributes.md) is set to `"no"`.
2. With the application list is displayed on top, the topmenu rendering ([topmenuDesktopRendering](1655-window-style-attributes-basics.md), [topmenuMobileRendering](1655-window-style-attributes-basics.md)) is ignored and defaults to a classic topmenu rendering.
3. Application list on top isn’t available on mobile.
4. If the ChromeBar is hidden, the application list displays on the left in the
   SideBarRail.
5. The windows list isn’t displayed when the application list is on top.
6. The application list is always hidden in the
   following conditions:
   1. On GDC/Desktop, when each program displays in a dedicated window container (the [`desktopMultiWindow`](1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface.") style attribute is set to `"yes"`)
   2. On GAS/Browser, when each program displays in a dedicated browser tab (the [`browserMultiPage`](1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface.") style attribute is set to "yes")

## `applicationListVisible`

The application list is can be used to navigate between
programs running concurrently started from a parent program with `RUN
command WITHOUT WAITING`. For more details about the application list,
read [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").

The `applicationListVisible` style attribute defines if the application
list is visible to the end user.

> **Important:**
>
> This style attribute is taken into account for the first program that displays to the front-end.
> All sub-sequent settings for this attribute in new application runs do not impact the
> application list visibility and position.

Values can be:

- `"auto"` (default): The application list becomes visible, when there is more than
  one running programs displayed by the front-end, except when the conditions described below
  apply.
- `"yes"`: The application list is always visible, even if only one program is
  displayed by the front-end, except when the conditions described below apply.
- `"no"`: The application list is always hidden.

Notes:

1. The applicationListPosition style
   attribute can be used to specify the position of the application list, when it is
   visible.
2. The application list is always hidden in the
   following conditions:
   1. On GDC/Desktop, when each program displays in a dedicated window container (the [`desktopMultiWindow`](1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface.") style attribute is set to `"yes"`)
   2. On GAS/Browser, when each program displays in a dedicated browser tab (the [`browserMultiPage`](1652-userinterface-style-attributes.md "UserInterface presentation style attributes define general options related to the application user interface.") style attribute is set to "yes")

## `browserMultiPage`

When using a GAS/web browser front-end, defines whether child programs started with `RUN
command WITHOUT WAITING` are displayed in the current browser tab, or in
a new browser tab.

The `browserMultiPage` style attribute only works, if the application having this
style is the first of the session.

When starting a child program with a `RUN` instruction not using the
`WITHOUT WAITING` clause, the child program windows are always displayed in the same
browser tab as the parent program.

Values can be:

- `"no"` (default): All started child programs display in the current browser
  tab.
- `"yes"`: All started child programs get a dedicated browser tab.

When using `browserMultiPage=yes`, by default, when the main/parent program is
terminated while child programs (started with a `RUN command WITHOUT
WAITING`) are still running and displaying in their dedicated browser tabs, the browser tab
of the parent program will remain open and show a message, until all child programs are terminated.
The GAS XCF configuration parameter `CloseAllOnMainExit` can be set to
`TRUE`, to force all child programs to terminate automatically, when the main/parent
program terminates.

See also [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.") and the [`RUN`](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.") instruction.

## `desktopMultiWindow`

When using the GDC/desktop front-end, defines if child programs started with `RUN ...
WITHOUT WAITING` are displayed in the same window container or get a dedicated window
container of the operating system.

When starting a child program with a `RUN` instruction not using the
`WITHOUT WAITING` clause, the child program windows are always displayed in the same
desktop window container as the parent program.

Values can be:

- `"no"` (default): A single window container is created for all applications. All
  started child programs display in the same container. The application list can be used to switching
  between applications.
- `"yes"`: Multiple window containers are created for each running program. All
  started child programs get a dedicated window container.

See also [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.") and the [`RUN`](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.") instruction.

## `reverse`

Global configuration option to display forms in reverse mode and enable right to left input,
for right-to-left language support.

Values can be:

- `"no"` (default): Display forms for left-to-right languages.
- `"yes"`: Display mirrored forms for right-to-left languages.

## `windowListVisible`

The window list is defined by the set of regular windows opened by the current
program and opened by all child programs started with a waiting `RUN`. The purpose of
the window list is to show the content of inactive windows, opened before the current
active window.

By default, the window list option is visible on the top left of the chromebar.

The `windowListVisible` attribute defines if the window list is
visible to the end user.

Values can be:

- `"yes"` (default): The window list is visible, when there more than one window is
  opened by the program or child programs started with a waiting `RUN`.
- `"no"`: The window list is always hidden. The stack of inactive windows cannot be
  seen. Only the current active window is available.

See also [Containers for program windows](1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end.").

## Related links

**Related concepts**  

[The abstract user interface tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")

[Right-to-left languages support](../09_advanced-features/0893-right-to-left-languages-support.md "Genero supports right-to-left languages, such as Arabic and Hebrew.")

[Background/foreground modes](../17_mobile-applications/5094-background-foreground-modes.md "Describes how to handle background or foreground modes in mobile apps.")
