---
title: "mobile.getLastNotificationInteractions"
source: "fgl-topics/c_fgl_frontcall_mobile_getlastnotificationinteractions.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.getLastNotificationInteractions"
type: "concept"
---

# mobile.getLastNotificationInteractions

> Get the last user interactions on mobile app notifications.

## Syntax

```
ui.Interface.frontCall("mobile", "getLastNotificationInteractions",
   [], [notifications] )
```

1. notifications - Is a `DYNAMIC ARRAY OF RECORD` defined as follows:
   - id `STRING` - Identifies a notification created by a [`createNotification`](3449-mobile-createnotification.md "Creates or updates a local notification to be displayed on the mobile device.")
     frontcall.
   - type `STRING` - Indicates the type of action taken by the end
     user. Values can be: `"selected"` or `"deleted"`.

## Usage

The "`getLastNotificationInteractions`" front call returns a list of notification
ids, with the corresponding user interaction that has been done (selected or deleted)

If no notification has been selected or deleted by the user since the last call,
`NULL` is returned.

This front call can be used after receiving a [`notificationselected`](../11_user-interface/2257-list-of-predefined-actions.md) predefined action, or when starting the applications,
to detect if notifications have been deleted or selected while the app was not running. The selected
notification can be a local notification created by a [`mobile.createNotification`](3449-mobile-createnotification.md "Creates or updates a local notification to be displayed on the mobile device.") front call, or a [push notification](../17_mobile-applications/5112-push-notifications.md "This section describes how to implement push notification with Genero.") received from the push
service, when using [`mobile.registerForRemoteNotifications`](3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.").

The front call returns a JSON Array of JSON object with the `id` and the
`type` properties. The returned JSON string can be automatically converted to a
dynamic array defined as follows:

```
TYPE t_nl DYNAMIC ARRAY OF RECORD
  id STRING,
  type STRING
END RECORD
DEFINE nl t_nl
```

The `type` defines how the user processed a notification. Values can be
"selected" or "deleted" of the last notifications which has been selected or deleted by the user
since the last time this frontcall has been called.

## Example

```
TYPE t_nl DYNAMIC ARRAY OF RECORD
  id STRING,
  type STRING
END RECORD
DEFINE nl t_nl
...
  ON ACTION notificationselected
     CALL ui.Interface.frontCall("mobile",
             "getLastNotificationInteractions",[],[nl])
     ...
```

## Related links

**Related concepts**  

[mobile.createNotification](3449-mobile-createnotification.md "Creates or updates a local notification to be displayed on the mobile device.")

[mobile.registerForRemoteNotifications](3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.")
