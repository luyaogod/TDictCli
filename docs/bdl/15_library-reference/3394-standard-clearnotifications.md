---
title: "standard.clearNotifications"
source: "fgl-topics/c_fgl_frontcall_standard_clearnotifications.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.clearNotifications"
type: "concept"
---

# standard.clearNotifications

> Drops notifications displayed on the host system of the front end.

## Syntax

```
ui.Interface.frontCall("standard", "clearNotifications",
   [options], [status] )
```

1. options - Is a record defined with the following members:
   - id\_list - A dynamic array of integers containing the list of notification ids
     to be removed.
   - channel\_name - The name of a group of notifications to be removed.
2. status - Holds the status of the front call execution.

## Usage

The "`clearNotifications`" front call deletes all notifications identified by the
ids provided in the options record.

Notifications cleared with this front call will not appear in the next call to [`getLastNotificationInteractions`](3402-standard-getlastnotificationinteractions.md "Report all user interactions with notifications since last call to this function.").

When passing `NULL` as option parameter, all notifications
(remote and local) will be deleted.

## Example

```
DEFINE ret STRING
DEFINE options RECORD 
    id_list DYNAMIC ARRAY OF INTEGER,
    channel_name STRING
END RECORD
LET options.id_list[1] = 1768
LET options.id_list[2] = 2234
LET options.id_list[3] = 3239
LET options.channel_name = "Shipping"
CALL ui.Interface.frontCall("standard","clearNotifications",[options],[ret])
```

## Related links

**Related concepts**  

[standard.createNotification](3397-standard-createnotification.md "Creates a new local notification to be displayed on the host system of the front end.")
