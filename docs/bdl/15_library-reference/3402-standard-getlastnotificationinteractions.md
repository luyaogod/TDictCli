---
title: "standard.getLastNotificationInteractions"
source: "fgl-topics/c_fgl_frontcall_standard_getlastnotificationinteractions.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.getLastNotificationInteractions"
type: "concept"
---

# standard.getLastNotificationInteractions

> Report all user interactions with notifications since last call to this function.

## Syntax

```
ui.Interface.frontCall("standard", "getLastNotificationInteractions",
   [], [notifications] )
```

1. notifications - Is a `DYNAMIC ARRAY OF RECORD` defined as follows:
   - id `STRING` - Identifies a notification created by a [`createNotification`](3397-standard-createnotification.md "Creates a new local notification to be displayed on the host system of the front end.")
     frontcall.
   - type `STRING` - Indicates the type of action taken by the end
     user. Values can be: `"selected"` or `"deleted"`.

## Usage

The "`getLastNotificationInteractions`" front call returns a list of notification
ids, with the corresponding user interaction that has been done (selected or deleted)

If no notification has been selected or deleted by the user since the last call,
`NULL` is returned.

This front call can be used after receiving a [`notificationselected`](../11_user-interface/2257-list-of-predefined-actions.md) predefined action, to detect if notifications have
been deleted or selected by the end user.

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
     CALL ui.Interface.frontCall("standard",
             "getLastNotificationInteractions",[],[nl])
     ...
```

## Related links

**Related concepts**  

[standard.createNotification](3397-standard-createnotification.md "Creates a new local notification to be displayed on the host system of the front end.")

[standard.clearNotifications](3394-standard-clearnotifications.md "Drops notifications displayed on the host system of the front end.")
