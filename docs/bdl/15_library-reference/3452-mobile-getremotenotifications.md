---
title: "mobile.getRemoteNotifications"
source: "fgl-topics/c_fgl_frontcall_mobile_getremotenotifications.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.getRemoteNotifications"
type: "concept"
---

# mobile.getRemoteNotifications

> This front call retrieves push notification messages.

## Syntax

```
ui.Interface.frontCall("mobile","getRemoteNotifications",
   [], [data] )
```

1. data - `STRING` containing a JSON array of notifications.

## Usage

After registering for push notifications with the [mobile.registerForRemoteNotifications](3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.") front call, the
`getRemoteNotifications` front call can be called in the context of an
`ON ACTION notificationpushed` action handler.

The GMI or GMA front-end will send the [`notificationpushed` special action](../11_user-interface/2257-list-of-predefined-actions.md), when it receives notifications from the
push notification server. When this action is fired, use the `getRemoteNotifications`
front call to get notification data.
> **Important:**
>
> When an app restarts, if notifications are pending and the app has already registered for push
> notification in a previous execution, the `notificationpushed` action will be raised
> as soon as a dialog with the corresponding `ON ACTION` handler activates. The app
> then performs a [mobile.getRemoteNotifications](3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") front call as in
> the usual way, to get the pending notifications pushed to the device while the app was off.
>
> However, special consideration needs to be given to iOS devices. When push notification arrives
> for an iOS app that has not started, there is no mechanism to wake up the app and get the push data.
> Therefore, when the user starts the app from the springboard, there will never be any push data
> available. Depending on the context, implement the following programming patterns to solve this
> problem:
>
> 1. If the push notification contains a badge number, the app can verify if the badge is greater
>    than 0 (with the [`getBadgeNumber`](3470-ios-getbadgenumber.md "Returns the current badge number associated to the app.") front call) in order to perform a
>    `getRemoteNotifications` front call. Even if there is no data available with the
>    front call, it is recommended that the app sends a request directly to the server push provider to
>    get last push data.
> 2. If the push notification does not contain badge numbers, it is still recommended that the app
>    performs a `getRemoteNotifications` front call when it starts. If there is no push
>    data available from the front call, the recommendation is that the app sends a request to the server
>    push provider to see if there is push data available. This is by the way also recommended when
>    receiving a `notificationpushed` action during application life time.
> 3. If the user starts the app from the Notification Center, the app is launched with push data
>    transmitted from the system, and the `notificationpushed` action is sent. It is
>    recommended that the app perform a `getRemoteNotifications` front call and get the
>    push data.

The "`getRemoteNotifications`" front call returns a list of notification records
as a JSON array string. Use the [util.JSONArray](3604-the-util-jsonarray-class.md "The util.JSONArray class provides methods to handle an array of values, following the JSON string syntax.")
or [util.JSON](3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") class to extract notification data from
the returned string. The structure of a push notification is platform specific. See below for
details.
> **Important:**
>
> When an iOS app is
> in background, silent push notifications can occur, but notification message data (i.e. the payload)
> may not be available. In such case, GMI is able to detect that a notification arrived (i.e. when
> the app badge number is greater than zero) and raise the `notificationpushed` action,
> but the `getRemoteNotifications` front call will return no message data
> (data return param is `NULL`). If such case, implement a fallback
> mechanism (based on RESTFul web services for example), to contact the push notification provider
> and retrieve the message information.

## Push notification records with GMA / Android™

The returned JSON string from a FCM notification server contains an array of notification
records.

A notification record contains the following JSON keys:

- `"type"` - can be `"message"` or `"token"`.
- `"data"` - Contains notification data.
  - When `"type":"message"`, the notification record is a FCM application message,
    and the `data` attribute contains custom notification information, as well as the
    `"id"` field, with the notification id to be used for example in the [mobile.clearNotification](3448-mobile-clearnotifications.md "Drops notifications displayed on the mobile device.") front
    call.
  - When `"type":"token"`, the notification record is a registration token update,
    and the `"data"` attribute contains the new registration token, that is required to
    be re-sent to the push notification server.
- `"from"` - Contains the FCM project number.

JSON push notification data example for GMA:

```
[
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
 },
 {
  "type": "token",
  "data": "a2ef...........",
  "from": "764001239714"
 },
 ...
]
```

For more details, see Firebase Could Messaging (FCM) documentation at <https://firebase.google.com>.

## Push notification records with GMI / iOS

The returned JSON string from an Apple®
Push Notification contains an array of notification records.

A push notification record contains the following JSON attributes:

- `"aps"` - key to be recognized by devices as an Apple Push Notification
  - `"alert"` (required) - key of the push notification content. If not specified as
    a single value, the `alert` key can hold:
    - `"title"` - title of the alert.
    - `"body"` - the message to be displayed.
  - `"badge"` (optional) - the number to display as the badge of the app icon. If
    this property is absent, the badge is not changed. You need to manage it through your push
    notification provider.
  - `"sound"` (optional) - the sound played by the alert (aiff,
    wav, or caf format) default value: "default". To use a
    custom file you will need to use the GMI extension project and be familiar with Objective-C. The
    file must bundled with the app.
  - `"content-available"` (required) - The content-available property with a value of
    1 allows the remote notification to act as a "silent" notification. The recommendation is that
    notifications received in background mode are stored for delivery when the app enters foreground
    mode.
- `"id"` - the notification id managed by GMI, for example, to be used in the [mobile.clearNotification](3448-mobile-clearnotifications.md "Drops notifications displayed on the mobile device.") front
  call.
- `"sysId"` - the notification iOS system id (provided for information and debug
  purpose)

JSON push notification data example for GMI:

```
[
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
 },
 {
  "aps" :
  {
    "alert" :
    {
      "title" : "Push",
      "body" : "My second push"
    }
    "badge" : 2,
    "sound" : "default",
    "content-available" : 1
  },
  "id":2,
  "sysId":"6A31016D-051E-9834-8734-F873489812CE",
  "new_ids" : [ "XV234", "ZF452", "RT563" ],
  "updated_ids" : [ "AC634", "HJ153" ]
 }
]
```

In the second record, custom information is provided in the "`new_ids`" and
"`updated_ids`" attributes, as a JSON array of identifiers.

For more details, see the Apple Push Notification service (APNs) documentation on
<https://developer.apple.com>.

## Example

```
IMPORT util -- JSON API

DEFINE notif_list STRING

DIALOG ...
  ...
  ON ACTION notificationpushed

     CALL ui.Interface.frontCall(
             "mobile", "getRemoteNotifications", 
             [ ], [ notif_list ] )

     DISPLAY util.JSON.format(notif_list)
  ...
```

## Related links

**Related concepts**  

[mobile.registerForRemoteNotifications](3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.")
