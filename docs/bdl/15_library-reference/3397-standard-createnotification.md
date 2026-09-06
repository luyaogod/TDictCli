---
title: "standard.createNotification"
source: "fgl-topics/c_fgl_frontcall_standard_createnotification.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.createNotification"
type: "concept"
---

# standard.createNotification

> Creates a new local notification to be displayed on the host system of the front end.

## Syntax

```
ui.Interface.frontCall("standard", "createNotification",
   [options], [id] )
```

1. options - A `RECORD` containing options to create the
   notification, where members are:
   - id `INTEGER` - When `NULL` or when providing an
     unexisting identifier, the front call creates a new notification. Otherwise, if the notification
     identifier exists already, the call will produce an error.
   - title `STRING` - Defines the title of the notification.
   - content `STRING` - Defines the body / text of the
     notification.
   - icon `STRING` - Defines the icon of the notification.
2. id - is an `INTEGER` variable receiving the unique
   notification id returned by the frontcall. Returns `NULL` if notification cannot be
   processed on the front-end device.

## Usage

The "`createNotification`" front call creates a new local notification.

The options parameter must be a record defined as follows:

```
DEFINE options RECORD
    id INTEGER,
    title STRING,
    content STRING,
    icon STRING
END RECORD
```

The id member must be `NULL` or must be an unexisting
identifier for this front end session. Otherwise, if the notification identifier exists already, the
call will produce an error.

If the user selects a notification on the front-end device, the predefined [`notificationselected`](../11_user-interface/2257-list-of-predefined-actions.md) action will
be sent to the current dialog instruction. For more details, see [`standard.getLastNotiticationInteractions`](3402-standard-getlastnotificationinteractions.md "Report all user interactions with notifications since last call to this function.").

In order to clear local notifications, use the [`standard.clearNotifications`](3394-standard-clearnotifications.md "Drops notifications displayed on the host system of the front end.") front call.

## Example

```
DEFINE options RECORD
    id INTEGER,
    title STRING,
    content STRING,
    icon STRING
END RECORD
LET options.id = NULL
LET options.title = "Arrival"
LET options.content = "Shipment arrived"
CALL ui.Interface.frontCall("standard","createNotification",[options],[options.id])
DISPLAY "Returned notification id: ", options.id
```

## Related links

**Related concepts**  

[standard.getLastNotificationInteractions](3402-standard-getlastnotificationinteractions.md "Report all user interactions with notifications since last call to this function.")

[standard.clearNotifications](3394-standard-clearnotifications.md "Drops notifications displayed on the host system of the front end.")

[mobile.createNotification](3449-mobile-createnotification.md "Creates or updates a local notification to be displayed on the mobile device.")
