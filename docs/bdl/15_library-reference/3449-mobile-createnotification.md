---
title: "mobile.createNotification"
source: "fgl-topics/c_fgl_frontcall_mobile_createnotification.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.createNotification"
type: "concept"
---

# mobile.createNotification

> Creates or updates a local notification to be displayed on the mobile device.

## Syntax

```
ui.Interface.frontCall("mobile", "createNotification",
   [options], [id] )
```

1. options - A `RECORD` containing options to create the
   notification, where members are:
   - id `INTEGER` - Defines a unique notification id.
   - number `INTEGER` - Badge increment (one by default)
   - channel\_name STRING - Android specific: All notifications are created under a
     notification channel. Using a channel name allows to create groups of notifications.
   - title `STRING` - Defines the title of the notification.
   - content `STRING` - Defines the body / text of the
     notification.
   - icon `STRING` - Defines the icon of the notification.
   - show\_badge `BOOLEAN` - TRUE / FALSE: shows or hides the badge.
2. id - is an `INTEGER` variable receiving the unique
   notification id returned by the frontcall. Returns `NULL` if notification cannot be
   processed on the mobile device.

## Usage

The "`createNotification`" front call creates or updates a new (local)
notification.

> **Note:**
>
> - Badge numbers are not handled automatically by iOS. To get a badge number, use the [`ios.setBadgeNumber`](3471-ios-setbadgenumber.md "Sets the current badge number associated to the app.")
>   frontcall.
> - Icons are not planned to be supported by GMI.
> - Channels are specific to Android. Consequently, the behavior with
>   `clearNotifications` is a GMA specific feature.

The options parameter must be a record defined as follows:

```
DEFINE options RECORD
    id INTEGER,
    number INTEGER,
    channel_name STRING,
    title STRING,
    content STRING,
    icon STRING,
    show_badge BOOLEAN
END RECORD
```

If the provided id member identifies an existing notification, it will update
the notification. Otherwise, if `NULL` or if the id is not yet
used, the frontcall will create a new notification. If `NULL` is provided in the
`id` record member, a unique id will be automatically created by the frontcall.

The badge number is automatically increased by one when a notification (remote or locale) is
created (on a same channel) and decreased by one when a notification is removed (by clearing, or
interacting with it). However, the value (one by default) that will be added or subtracted for a
specific notification can be changed by the number member.

The `channel_name` option can be used to group notifications by channels. If
`NULL` is specified, it defaults to `"GMA NOTIFICATION channel"`.

The show\_badge option only applies to the first notification created under a
new notification channel. The show\_badge option will be ignored if the channel
already exists. Clearing or dismissing all notification under a notification channel will not delete
it, and the show\_badge option will remain at the same value. However, clearing
the channel or clearing all channels will reset this option, and will allow to create the same
channel with a different show\_badge value.

If the user selects a notification on the mobile device, the predefined
`"notificationselected"` action will be sent to the current dialog instruction. For
more details, see [mobile.getLastNotificationInteractions](3450-mobile-getlastnotificationinteractions.md "Get the last user interactions on mobile app notifications.").

## Example

```
DEFINE options RECORD
    id INTEGER,
    number INTEGER,
    channel_name STRING,
    title STRING,
    content STRING,
    icon STRING,
    show_badge BOOLEAN
END RECORD
LET options.id = NULL
LET options.title = "Arrival"
LET options.content = "Shipment arrived"
CALL ui.Interface.frontCall("mobile","createNotification",[options],[options.id])
DISPLAY "Returned notification id: ", options.id
```

## Related links

**Related concepts**  

[mobile.getLastNotificationInteractions](3450-mobile-getlastnotificationinteractions.md "Get the last user interactions on mobile app notifications.")

[mobile.clearNotifications](3448-mobile-clearnotifications.md "Drops notifications displayed on the mobile device.")

[ios.getBadgeNumber](3470-ios-getbadgenumber.md "Returns the current badge number associated to the app.")

[ios.setBadgeNumber](3471-ios-setbadgenumber.md "Sets the current badge number associated to the app.")
