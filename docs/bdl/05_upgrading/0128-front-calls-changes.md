---
title: "Front calls changes"
source: "fgl-topics/c_fgl_Migrate_to_401_frontcalls.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.01 upgrade guide > Front calls changes"
type: "concept"
---

# Front calls changes

> Modifications to consider when using front calls.

## New front calls

Starting with GMA 4.01.04 and GMI 4.01.04:

- [`mobile.isEmulator`](../15_library-reference/3454-mobile-isemulator.md "Indicates if the mobile app runs or displays forms on an emulator/simulator.")

Starting with GMA 4.01.06 and GMI 4.01.04:

- [`mobile.createNotification`](../15_library-reference/3449-mobile-createnotification.md "Creates or updates a local notification to be displayed on the mobile device.")
- [`mobile.getLastNotificationInterfactions`](../15_library-reference/3450-mobile-getlastnotificationinteractions.md "Get the last user interactions on mobile app notifications.")
- [`mobile.clearNotifications`](../15_library-reference/3448-mobile-clearnotifications.md "Drops notifications displayed on the mobile device.")

Starting with GBC 4.01.20:

- [`standard.feInfo`](../15_library-reference/3399-standard-feinfo.md) supports now the `colorScheme`
  parameter.

Starting with GBC 4.01.22, these mobile front calls are now also available as standard front
calls, and can be used on a desktop/browser front end platform:

- [`standard.composeMail`](../15_library-reference/3395-standard-composemail.md "Invokes the user's default mail application for a new mail to send.") (as alias for
  `mobile.composeMail`)
- [`standard.connectivity`](../15_library-reference/3396-standard-connectivity.md "Returns the type of network available for the device.") (as alias for
  `mobile.connectivity`)
- [`standard.getGeolocation`](../15_library-reference/3401-standard-getgeolocation.md "Returns the Global Positioning System (GPS) location of a device.") (as alias for
  `mobile.getGeolocation`)
- [`standard.isForeground`](../15_library-reference/3405-standard-isforeground.md "Indicates if the app is in foreground mode.") (as alias for
  `mobile.isForeground`)

## JSON data from `mobile.getRemoteNotifications`

**JSON structure returned with GMA/Android**

Starting with GMA 4.01.07, with the `mobile.getRemoteNotifications` front call,
the structure of the returned JSON data for an FCM notification of type `"message"`
has changed, maching the Firebase Cloud Messaging API V1 JSON used by the push notification send
from the push server.

Before GMA 4.01.07, the structure of a JSON element a serialized JSON object in the
`"data"` field, using a `"genero_notification"`
member:

```
 {
  "type": "message",
  "data": "\"genero_notification\" : { \"title\" : ... }",
  "from": "764001239714"
 }
```

Since GMA 4.01.07, the `"genero_notification"` member is no longer used, and the
`"title"`, `"content"`, `"icon"` (as well as the new
`"id"`) fields appear now directly under `"data"`:

```
 {
  "type": "message",
  "data": {
           "id"      : "82837497234",
           "title"   : "Game Request!",
           "content" : "Bob wants to play poker...",
           "icon"    : "smiley",
           ...
          },
  "from": "764001239714"
 }
```

**JSON structure returned with GMI/iOS**

Starting with GMI 4.01.05, the JSON structure returned by the
`mobile.getRemoteNotifications` front call contents two new members to identify
notifications `"id"` and `"sysId"`:

```
 {
  "aps" :
  {
    "alert" : "My first push",
    "badge" : 1,
    "sound" : "default",
    "content-available" : 1
  },
  "id":1,
  "sysId":"6A31016D-051E-4943-9111-E959241D3DCD"
 }
```

For more details, see [mobile.getRemoteNotifications](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.").

## GBC supports mobile front calls

Starting with GBC 4.01.20, some front calls of the "mobile" module are implemented in the GBC
JavaScript code, and can be used when displaying applications forms in a browser on the mobile
device, using GBC through GAS or in a WASM app. Before this support, it was not possible to use the
mobile front calls in this configuration.

Notes:

- Some front calls such as [`mobile.chooseContact`](../15_library-reference/3442-mobile-choosecontact.md "Lets the user choose a contact from the mobile device contact list and returns the vCard.") need a secure context (connection through HTTPS or on
  localhost)
- The behavior and return values of some front calls have changed, such as [`mobile.composeMail`](../15_library-reference/3445-mobile-composemail.md "Invokes the user's default mail application for a new mail to send.") or [`mobile.composeSMS`](../15_library-reference/3446-mobile-composesms.md "Sends an SMS text to one or more phone numbers."), as the
  control of the messaging/SMS app is limited from GBC JavaScript.

For more details, see [Genero Mobile common front calls](../15_library-reference/3441-genero-mobile-common-front-calls.md "This section describes common front calls provided by all mobile front-ends.").

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Front
call changes in BDL 4.00](0144-front-calls-changes.md "Modifications to consider when using front calls.").

Notable changes introduced in maintenance releases:

- No particular change to consider.

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
