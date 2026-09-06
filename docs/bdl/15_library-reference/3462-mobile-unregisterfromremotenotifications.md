---
title: "mobile.unregisterFromRemoteNotifications"
source: "fgl-topics/c_fgl_frontcall_mobile_unregisterfromremotenotifications.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.unregisterFromRemoteNotifications"
type: "concept"
---

# mobile.unregisterFromRemoteNotifications

> This front call unregisters the mobile device from push notifications.

## Syntax

```
ui.Interface.frontCall("mobile","unregisterFromRemoteNotifications",
   [], [] )
```

## Usage

The "`unregisterFromRemoteNotifications`" front call unregisters the
device from push notifications after it has been registered with the [mobile.registerForRemoteNotifications](3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.") front
call.

## Example

```
CALL ui.Interface.frontCall(
        "mobile", "unregisterFromRemoteNotifications", 
        [ ], [ ] )
...
```

## Related links

**Related concepts**  

[mobile.registerForRemoteNotifications](3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.")
