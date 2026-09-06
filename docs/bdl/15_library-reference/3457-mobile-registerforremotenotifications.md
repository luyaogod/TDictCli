---
title: "mobile.registerForRemoteNotifications"
source: "fgl-topics/c_fgl_frontcall_mobile_registerforremotenotifications.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.registerForRemoteNotifications"
type: "concept"
---

# mobile.registerForRemoteNotifications

> This front call registers a mobile device for push notifications.

## Syntax

```
ui.Interface.frontCall("mobile","registerForRemoteNotifications",
   [], [registration-token] )
```

1. registration-token - Registration token to be sent to the push notification
   provider. For GMA/Android, this is the "registration token" obtained from Firebase Cloud Messaging
   (FCM).

## Usage

The "`registerForRemoteNotifications`" front call registers the mobile device for
push notifications. Once the registration procedure is done (see below for platform specifics), it
is possible to get notification events through the `notificationpushed` predefined
action, and retrieve notification data with the [mobile.getRemoteNotifications](3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") front call.

> **Note:**
>
> The app does not need to register for
> notification each time it is restarted. Even if the app is closed, the registration is still active
> until the `unregisterFromRemoteNotifications` front call is performed. At first
> execution, an app will typically ask if the user wants to get push notifications and register to the
> push service if needed. To disable push notification, apps usually implement an option that can be
> disabled (to unregister) and re-enabled (to register again) by the user. On Android™, the app must register for
> notification each time it is upgraded.

## Registering with FCM on Android

On Android the
registration-token is the registration token returned by FCM. Once registered
with the FCM service, the app must also send this registration token to the FCM application server.
Registration tokens are typically sent to the FCM application server using a RESTFul HTTP POST.

> **Note:**
>
> Android apps using push notification
> services need specific permissions to be defined in the manifest, such as
> `android.permission.POST_NOTIFICATIONS`. These Android permissions will be
> automatically set by the [gmabuildtool](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.").

For more details, see Firebase Could Messaging (FCM) documentation at <https://firebase.google.com>.

## Example

The following code example registers with Firebase Cloud Messaging or Apple Push
Notification service. It then sends the registration token to the push notification provider:

```
IMPORT com  -- For RESTful post
IMPORT util -- JSON API

DEFINE registration_token STRING
DEFINE req com.HttpRequest,
       obj util.JSONObject,
       resp com.HttpResponse

-- First get the registration token
CALL ui.Interface.frontCall(
        "mobile", "registerForRemoteNotifications", 
        [ ], [ registration_token ] )

-- Then send registration token to push notification provider
TRY
    LET req = com.HttpRequest.create("http://SERVER_IP:4930")
    CALL req.setHeader("Content-Type", "application/json")
    CALL req.setMethod("POST")
    CALL req.setTimeOut(5)
    LET obj = util.JSONObject.create()
    CALL obj.put("registration_token", registration_token)
    CALL req.doTextRequest(obj.toString())
    LET resp = req.getResponse()
    IF resp.getStatusCode() != 200 THEN
       MESSAGE SFMT("HTTP Error (%1) %2",
                    resp.getStatusCode(),
                    resp.getStatusDescription())
    ELSE
       MESSAGE "Registration token sent."
    END IF
CATCH
    MESSAGE SFMT("Could not post registration token to server: %1", status)
END TRY
...
```

## Related links

**Related concepts**  

[mobile.unregisterFromRemoteNotifications](3462-mobile-unregisterfromremotenotifications.md "This front call unregisters the mobile device from push notifications.")
