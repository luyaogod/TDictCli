---
title: "List of predefined actions"
source: "fgl-topics/c_fgl_prog_dialogs_predact_list.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Predefined actions > List of predefined actions"
type: "concept"
description: "Important: The predefined actions described in these tables are created depending on the usage context described in the Context descriptions notes. Automatic actions (implicitly created) Automatic ..."
---

# List of predefined actions

> **Important:**
>
> The predefined actions described in these tables are created depending
> on the usage context described in the Context descriptions
> notes.

## Automatic actions (implicitly created)

Automatic actions are implicitly created by dialogs, and may be overwritten by `ON
ACTION` handlers.

| Name | Description | Context |
| --- | --- | --- |
| `accept` | Validates the current interactive instruction | (3) |
| `cancel` | Cancels the current interactive instruction | (3) |
| `close` | Triggers a cancel key in the current interactive instruction (by default)See also [Implementing the close action](2289-implementing-the-close-action.md "The close action is a predefined action dedicated to close graphical windows (for example, with the X cross button)."). | (8) |
| `insert` | Inserts a new row before current row | (9) |
| `append` | Appends a new row at the end of the list | (9) |
| `delete` | Deletes the current row | (9) |
| `update` | Updates the current row | (9) |
| `find` | Opens the fglfind dialog window to let the user enter a search value, and seeks to the row matching the value | (4) |
| `findnext` | Seeks the next row matching the value entered during the fglfind dialog | (4) |
| `nextrow` | Moves to the next row | (4) |
| `prevrow` | Moves to the previous row | (4) |
| `firstrow` | Moves to the first row | (4) |
| `lastrow` | Moves to the last row | (4) |
| `help` | Shows the help topic defined by the `HELP` clause | (1) |
| `editcopy` | Copy selected rows (or current row if MRS is off) to the clipboard | (5) |
| `expandall` | A tree node is completly expanded | (6) |
| `collapseall` | A tree node is completly collapsed | (6) |
| `left` | Moves to the previous cell | (7) |
| `right` | Moves to the next cell | (7) |

## Particular actions (ON ACTION is not required)

Particular actions do not need an `ON ACTION` handler. They are automatically
enabled depending on the context, and produce a specific event in the AUI protocol.

| Name | Description | Context |
| --- | --- | --- |
| `interrupt` | Sends an interruption request to the program when processing. For more details, see [User interruption handling](2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure."). | (2) |

## Special actions (ON ACTION is required)

Special actions have a specific behavior and need an `ON ACTION` handler to be
activated.

| Name | Description | Context |
| --- | --- | --- |
| `applicationstatechanged` | With a web browser as front-end, this action is triggered when the end user changes the `#` anchor in the URL field.Read [browser.getApplicationState](../15_library-reference/3434-browser-getapplicationstate.md "Gets the # anchor of the current URL in the browser address bar."). | (1) |
| `back` | Displays implicitely a back button in the chromebar when an action handler with this name exists and the action is active.See [Implementing the back action](2290-implementing-the-back-action.md "The back action is a predefined action dedicated to move back in the stack of windows/forms."). | (1) |
| `browser_back` | Sent when the user hits the back button in a web browser (web front-end only). | (1) |
| `browser_forward` | Sent when the user hits the forward button in a web browser (web front-end only). | (1) |
| `dialogtouched` | Sent by the front-end each time the user modifies the value of a field. For more details, see [Immediate detection of user changes](2239-immediate-detection-of-user-changes.md "This section describes the dialogtouched predefined action."). | (8) |
| `enterbackground` | On Mobile devices, this action is fired when the app goes to background mode.With GBC in a web browser, this action is fired when the browser window or current tab gets hidden to the user, because a browser window/tab switch occurred, or the browser window got minimized.With GDC/UR, the `enterbackground` action is fired when the window container is minimized.For more details, see [Foreground and background modes](../09_advanced-features/0829-executing-programs.md). | (1) |
| `enterforeground` | On Mobile devices, this action is fired when the app goes to foreground mode.With GBC in a web browser, this action is fired when the browser window or current tab is shown to the user, because a browser window/tab switch occurred, or the browser window is restored.With GDC/UR, the `enterforeground` action is fired when the window container is restored.For more details, see [Foreground and background modes](../09_advanced-features/0829-executing-programs.md). | (1) |
| `windowresized` | **Important:**This feature is deprecated, its use is discouraged although not prohibited.On Mobile devices, this action is sent when changing the orientation of the device. On other front-ends, it is sent when the current active window is resized. For more details, see [Adapting to viewport changes](1548-adapting-to-viewport-changes.md "Application forms and functions can be adapted to the front-end viewport size or mobile device orientation."). | (1) |
| `notificationpushed` | On Mobile devices, this action is fired when receiving a push notification message. See [mobile.getRemoteNotifications](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") | (1) |
| `notificationselected` | This action is sent when the user selects a local or push notification. See [standard.getLastNotificationInteractions](../15_library-reference/3402-standard-getlastnotificationinteractions.md "Report all user interactions with notifications since last call to this function."), [mobile.getLastNotificationInteractions](../15_library-reference/3450-mobile-getlastnotificationinteractions.md "Get the last user interactions on mobile app notifications."). | (1) |
| `cordovacallback` | On Mobile devices, this action is fired when receiving Cordova plugin results. See [cordova.callWithoutWaiting](../15_library-reference/3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result."). | (1) |

## Context descriptions

1. Possible in any kind of interactive instruction ([`MENU`](1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.") included).
2. Only possible when no interactive instruction is active.
3. Singular [`CONSTRUCT`](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form."), [`INPUT`](1930-record-input-input.md "The INPUT instruction provides single record input control in an application form."), [`PROMPT`](1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window."), [`INPUT
   ARRAY`](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.") and [`DISPLAY
   ARRAY`](1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.").
4. `INPUT ARRAY` and `DISPLAY ARRAY`.
5. `DISPLAY ARRAY` only.
6. `DISPLAY ARRAY` controlling a [`TREE`](2352-tree-views.md "Describes how to implement tree views.").
7. `DISPLAY ARRAY` with [`FOCUSONFIELD`](2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.") option.
8. [`DIALOG`](2076-multiple-dialogs-dialog-inside-functions.md "The procedural DIALOG instruction allows for the combination of record list, record input, and query criteria input in the same application form."),
   `CONSTRUCT`, `INPUT`, `PROMPT`, `INPUT
   ARRAY` and `DISPLAY ARRAY`.
9. `INPUT ARRAY` (for `insert`, `append`,
   `delete`) or `DISPLAY ARRAY` when using [modification triggers](2312-display-array-modification-triggers.md "Using dedicated interaction blocks to allow the user to modify a read-only record list.") (for
   `insert`, `append`, `delete` and
   `update`).

## Related links

**Related concepts**  

[Dialog actions](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.")
