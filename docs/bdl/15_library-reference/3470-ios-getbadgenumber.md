---
title: "ios.getBadgeNumber"
source: "fgl-topics/c_fgl_frontcall_ios_getbadgenumber.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile iOS front calls > ios.getBadgeNumber"
type: "concept"
---

# ios.getBadgeNumber

> Returns the current badge number associated to the app.

## Syntax

```
ui.Interface.frontCall("ios", "getBadgeNumber",
  [],[value])
```

1. value - Holds the current badge number.

## Usage

The iOS "`getBadgeNumber`" front call returns the current badge
number associated to the app.

> **Important:**
>
> This front call is only available for an
> application running on an iOS device.

The badge number appears on the app icon and is typically used for [Push notifications](../17_mobile-applications/5112-push-notifications.md "This section describes how to implement push notification with Genero.").

> **Important:**
>
> In order to query or set the badge number, the app program must have
> executed a [`registerForRemoteNotifications`](3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.")
> front call before (in the current or prior execution instance). This registration is required
> in order to set the appropriate app permissions to access badge number data.

## Example

```
DEFINE value INTEGER
CALL ui.interface.frontcall("ios","getBadgeNumber",[],[value])
```

## Related links

**Related concepts**  

[Deploying mobile apps on iOS devices](../17_mobile-applications/5107-deploying-mobile-apps-on-ios-devices.md "This section contains information to create a mobile application to be deployed on iOS devices.")

[ios.setBadgeNumber](3471-ios-setbadgenumber.md "Sets the current badge number associated to the app.")
