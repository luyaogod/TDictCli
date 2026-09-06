---
title: "Front calls changes"
source: "fgl-topics/c_fgl_Migrate_to_600_frontcalls.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 6.00 upgrade guide > Front calls changes"
type: "concept"
---

# Front calls changes

> Modifications to consider when using front calls.

## New front calls

URL `#` anchor handling (GBC 5.01.08):

- [`browser.setApplicationState`](../15_library-reference/3433-browser-setapplicationstate.md "Sets the # anchor of the URL in the browser address bar.")
- [`browser.getApplicationState`](../15_library-reference/3434-browser-getapplicationstate.md "Gets the # anchor of the current URL in the browser address bar.")

Local notifications control (GBC 6.00.01):

- [`standard.clearNotifications`](../15_library-reference/3394-standard-clearnotifications.md "Drops notifications displayed on the host system of the front end.")
- [`standard.createNotification`](../15_library-reference/3397-standard-createnotification.md "Creates a new local notification to be displayed on the host system of the front end.")
- [`standard.getLastNotificationInteractions`](../15_library-reference/3402-standard-getlastnotificationinteractions.md "Report all user interactions with notifications since last call to this function.")

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Front
call changes in BDL 5.01](0106-front-calls-changes.md "Modifications to consider when using front calls.").

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
