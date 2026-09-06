---
title: "Handling notifications in the mobile app"
source: "fgl-topics/c_fgl_mobile_push_notif_app_side.html"
breadcrumb: "Mobile applications > Push notifications > Handling notifications in the mobile app"
type: "concept"
---

# Handling notifications in the mobile app

> This topic describes how to handle push notification in the app running on mobile devices.

## Basics

In order to implement a push notification mechanism, you need to set up a server part (token
maintainer and push notification server), based on a push notification framework such as Firebase
Cloud Messaging (FCM) or Apple® Push
Notification service (APNs). In Addition, you need to handle notification events in your mobile app.
This section describes how to implement push notification in the app with the push notification API
available in Genero BDL.

The same code base can be used to handle push notifications for Android™ (using FCM) and iOS (using APNs) devices. Only the content of the
notification message will have to be processed with specific code, as the structure of the message
differs depending on standards defined by the push notification framework.

## Genero API for push notifications

Genero BDL provides an API to handle push notification on mobile apps. Dedicated front calls are
available to register to a push server, fetch push notification data, check notifications user
interactions, clear notifications, and eventually unregister:

- [mobile.registerForRemoteNotifications](../15_library-reference/3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.")
- [mobile.getRemoteNotifications](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.")
- [mobile.getLastNotificationInteractions](../15_library-reference/3450-mobile-getlastnotificationinteractions.md "Get the last user interactions on mobile app notifications.")
- [mobile.clearNotifications](../15_library-reference/3448-mobile-clearnotifications.md "Drops notifications displayed on the mobile device.")
- [mobile.unregisterFromRemoteNotifications](../15_library-reference/3462-mobile-unregisterfromremotenotifications.md "This front call unregisters the mobile device from push notifications.")

To detect when a notification message arrives from the push server, a specific action called
`notificationpushed` must be used by app code on a `ON ACTION` handler.
This special action is referenced as a [predefined
action](../11_user-interface/2257-list-of-predefined-actions.md).

## Android app permissions for FCM push notifications

Android apps using push notification services need specific permissions (Android manifest), such
as `android.permission.POST_NOTIFICATIONS`.

Permissions are automatically set for Android APK packages by the [GMA
buildtool](5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.").

See the FCM documentation for more details about required permissions for push notifications.

## iOS app certificates for APNs push notifications

iOS apps must be created with an Apple
certificate for development or distribution, linked to an App ID (or Bundle ID) with push notification
enabled. The provisioning profile used when building the IPA must be linked to the App ID with push
enabled. Certificate, provisioning and bundle id must be specified to the [GMI buildtool](5110-gmibuildtool.md "The gmibuildtool is a utility to create and test applications for an iOS devices.").

## Handling push notification in the app

To handle push notifications in your mobile app, perform the following steps:

1. Register to the push service and
   get the registration token
2. Send the push notification
   token to your token maintainer
3. Handle notification
   arrival
4. Handle notification
   interactions
5. Clear received
   notification
6. Un-register from the push
   servers

## 1 - Registering to the push service and to the push provider

Register the app to the push notification service with the [`"registerForRemoteNotifications"`](../15_library-reference/3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.") front call.

- When using FCM, you must provide Sender ID to identify the FCM project.
- When using APNs, you can set the Sender ID to NULL.

> **Note:**
>
> The app does not need to register for
> notification each time it is restarted. Even if the app is closed, the registration is still active
> until the `unregisterFromRemoteNotifications` front call is performed. At first
> execution, an app will typically ask if the user wants to get push notifications and register to the
> push service if needed. To disable push notification, apps usually implement an option that can be
> disabled (to unregister) and re-enabled (to register again) by the user. On Android, the app must register for
> notification each time it is upgraded.

> **Important:**
>
> When an app restarts, if notifications are pending and the app has already registered for push
> notification in a previous execution, the `notificationpushed` action will be raised
> as soon as a dialog with the corresponding `ON ACTION` handler activates. The app
> then performs a [mobile.getRemoteNotifications](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") front call as in
> the usual way, to get the pending notifications pushed to the device while the app was off.
>
> However, special consideration needs to be given to iOS devices. When push notification arrives
> for an iOS app that has not started, there is no mechanism to wake up the app and get the push data.
> Therefore, when the user starts the app from the springboard, there will never be any push data
> available. Depending on the context, implement the following programming patterns to solve this
> problem:
>
> 1. If the push notification contains a badge number, the app can verify if the badge is greater
>    than 0 (with the [`getBadgeNumber`](../15_library-reference/3470-ios-getbadgenumber.md "Returns the current badge number associated to the app.") front call) in order to perform a
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

The `registerForRemoteNotifications` front call will return a registration
token for the app which will be used by the push server (a.k.a push provider).

- When using FCM, the returned identifier is the FCM "registration token".
- When using APNs, the returned identifier is the APNs "device token".

```
...
    LET rec.tm_host = "https://pushreg.example.orion"
    LET rec.tm_port = 4930
    LET rec.user_name = "mike"
    LET rec.notification_type = "FCM"
...
    DIALOG ATTRIBUTES(UNBUFFERED)
      ...
      ON ACTION register
         LET rec.registration_token = register(rec.notification_type, rec.user_name)
...

FUNCTION register(
    notification_type STRING,
    app_user STRING
) RETURNS STRING
    DEFINE registration_token STRING
    TRY
        CALL ui.Interface.frontCall(
                "mobile", "registerForRemoteNotifications",
                [ ], [ registration_token ] )
        IF tm_command( "register", notification_type,
                       registration_token, app_user, 0 ) < 0 THEN
           RETURN NULL
        END IF
    CATCH
        ERROR "Registration failed!"
        RETURN NULL
    END TRY
    MESSAGE SFMT("Registration succeeded (token=%1)", registration_token)
    RETURN registration_token
END FUNCTION
```

## 2 - Sending a push notification token to your token maintainer

Once registered to the FCM or APNs service, the app must also register to the push server or push
provider by sending the token obtained in step 1.

This is typically done by using a RESTFul HTTP POST, sending the token (along with additional
application user information) to a dedicated server program that maintains the list of registered
devices/tokens.

The device token maintainer can be implemented in BDL as a Web Service program, as described in
[Implementing a token maintainer](5115-implementing-a-token-maintainer.md "The token maintainer is a BDL Web Services server program that handles push token registration from mobile apps.").

In this tutorial, the `tm_command()` function implements token registration (as well
as badge number handling for APNs):

```
IMPORT com
IMPORT util

...
    LET rec.tm_host = "https://pushreg.example.orion"
    LET rec.tm_port = 4930
...

FUNCTION tm_command(
    command STRING,
    notification_type STRING,
    registration_token STRING,
    app_user STRING,
    badge_number INTEGER
) RETURNS INTEGER
    DEFINE url STRING,
           json_obj util.JSONObject,
           req com.HttpRequest,
           resp com.HttpResponse,
           json_result STRING,
           result_rec RECORD
                          status INTEGER,
                          message STRING
                      END RECORD
    TRY
        LET url = SFMT( "http://%1:%2/token_maintainer/%3",
                        rec.tm_host, rec.tm_port, command )
        LET req = com.HttpRequest.Create(url)
        CALL req.setHeader("Content-Type", "application/json")
        CALL req.setMethod("POST")
        CALL req.setConnectionTimeOut(5)
        CALL req.setTimeOut(5)
        LET json_obj = util.JSONObject.create()
        CALL json_obj.put("notification_type", notification_type)
        CALL json_obj.put("registration_token", registration_token)
        CALL json_obj.put("app_user", app_user)
        CALL json_obj.put("badge_number", badge_number)
        CALL req.doTextRequest(json_obj.toString())
        LET resp = req.getResponse()
        IF resp.getStatusCode() != 200 THEN
           ERROR SFMT("HTTP Error (%1) %2",
                      resp.getStatusCode(),
                      resp.getStatusDescription())
           RETURN -2
        ELSE
           LET json_result = resp.getTextResponse()
           CALL util.JSON.parse(json_result, result_rec)
           IF result_rec.status >= 0 THEN
              RETURN 0
           ELSE
              MESSAGE SFMT("Notification maintainer message:\n %1", result_rec.message)
              RETURN -3
           END IF
        END IF
    CATCH
        ERROR SFMT("Failed to post token registration command: %1", status)
        RETURN -1
    END TRY
END FUNCTION
```

When the app is declared as push notification client to the push server, continue with the normal
program flow.

## 3 - Handling push notification events

To get and handle notification events, the current active dialog must implement the
`notificationpushed` special action.

To control action view rendering defaults and current field validation behavior when the
`notificationpushed` action is used, consider setting action default attributes for
this action in your [.4ad file](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") as follows
(like done in FGLDIR/lib/default.4ad):

```
<ActionDefaultList>
  ...
  <ActionDefault name="notificationpushed" validate="no" defaultView="no" contextMenu="no"/>
  ...
</ActionDefaultList>
```

Another option is to define these action defaults attributes in the `ON ACTION`
handler:

```
ON ACTION notificationpushed ATTRIBUTES(VALIDATE=NO, DEFAULTVIEW=NO)
   ...
```

In the `ON ACTION` block for this action, query for notification messages by using
the [`"getRemoteNotifications"`](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") front call, (passing the Sender ID as parameter
when using FCM, for APNs the Sender ID must be NULL). This front call returns a JSON string
containing a list of notification messages to be processed:

```
...
      ON ACTION notificationpushed
         CALL handle_notification()
...

FUNCTION handle_notification() RETURNS ()
    DEFINE notif_list STRING,
           notif_array util.JSONArray,
           notif_item util.JSONObject,
           notif_data util.JSONObject,
           aps_record util.JSONObject,
           id, info, other_info STRING,
           i, x INTEGER
    CALL ui.Interface.frontCall(
              "mobile", "getRemoteNotifications",
              [ ], [ notif_list ] )
    TRY
        IF notif_list.trimLeft() NOT LIKE "[%" THEN -- GBC-3043 Workaround
           LET notif_list = "[ ", notif_list, " ]"
        END IF
        LET notif_array = util.JSONArray.parse(notif_list)
        IF notif_array.getLength() > 0 THEN
           CALL setup_badge_number(notif_array.getLength())
        END IF
        FOR i=1 TO notif_array.getLength()
            LET info = NULL
            LET other_info = NULL
            LET notif_item = notif_array.get(i)
            -- Try APNs msg format
            LET aps_record = notif_item.get("aps")
            IF aps_record IS NOT NULL THEN
               LET id = notif_item.get("id")
               LET info = aps_record.get("alert")
               LET notif_data = notif_item.get("custom_data")
               IF notif_data IS NOT NULL THEN
                  LET other_info = notif_data.get("other_info")
               END IF
            ELSE
               -- Try FCM msg format
               LET notif_data = notif_item.get("data")
               LET id = notif_data.get("id")
               LET info = notif_data.get("content")
               LET other_info = notif_data.get("other_info")
            END IF
            IF info IS NULL THEN
               LET info = "Unexpected message format"
            END IF
            LET x = x + 1
            LET rec.notifications = rec.notifications, "\n",
                    SFMT("%1(%2): Notifiation received: ID=%3\n %4[%5]",
                         x, CURRENT HOUR TO SECOND, id, info, other_info)
        END FOR
    CATCH
        ERROR "Could not extract notification info"
    END TRY
END FUNCTION
```

With APNs, the app must handle the badge numbers attached to the device token. The app must:

1. Query the current badge number with the `getBadgeNumber` front call.
2. Compute the new badge number based on the number of notifications consumed.
3. Reset the badge number with the setBadgeNumber front call.
4. Inform the token maintainer to sync the badge number in the central database.

The following function handles badge numbers for the app:

```
FUNCTION setup_badge_number(consumed INTEGER) RETURNS ()
    DEFINE badge_number INTEGER
    TRY -- If the front call fails, we are not on iOS...
        CALL ui.Interface.frontCall("ios", "getBadgeNumber", [], [badge_number])
    CATCH
        RETURN
    END TRY
    IF badge_number>0 THEN
       LET badge_number = badge_number - consumed
    END IF
    CALL ui.Interface.frontCall("ios", "setBadgeNumber", [badge_number], [])
    IF tm_command( "badge_number", "APNS", rec.registration_token,
                   rec.user_name, badge_number) < 0 THEN
       ERROR "Could not send new badge number to token maintainer."
       RETURN
    END IF
END FUNCTION
```

## 4 - Handle notification interactions

To detect user taps on notification popups, use the `notificationselected` action
and the [`getLastNotificationInterfactions`](../15_library-reference/3450-mobile-getlastnotificationinteractions.md "Get the last user interactions on mobile app notifications.") front call:

```
...
      ON ACTION notificationselected
         CALL handle_notification_selection()
...

FUNCTION handle_notification_selection() RETURNS ()
    DEFINE notif_array DYNAMIC ARRAY OF RECORD
               id STRING,
               type STRING
           END RECORD,
           x INTEGER
    TRY
        CALL ui.Interface.frontCall("mobile", "getLastNotificationInteractions",
                                    [], [notif_array] )
        FOR x=1 TO notif_array.getLength()
            LET rec.notifications = rec.notifications, "\n",
                    SFMT("%1(%2): Notification selected: ID=%3 TYPE=%4",
                         x, CURRENT HOUR TO SECOND,
                         notif_array[x].id, notif_array[x].type)
        END FOR
    CATCH
        ERROR "Could not get selected notifications info"
    END TRY
END FUNCTION
```

## 5 - Clear notifications

To clear notifications on the device, use the [`clearNotifications`](../15_library-reference/3448-mobile-clearnotifications.md "Drops notifications displayed on the mobile device.")
front call:

```
...
      ON ACTION clear
         CALL clear_notifications()
...

FUNCTION clear_notifications() RETURNS ()
    DEFINE res STRING
    TRY
        CALL ui.Interface.frontCall(
                "mobile", "clearNotifications",
                [ ], [ res ] )
        LET rec.notifications = res
    CATCH
        ERROR "Clear notifications failed!"
    END TRY
END FUNCTION
```

## 6 - Unregistering the app from push notification

If the app no longer wants to get push notifications, unregister from the push provider (using
a RESTful POST, in the `regunreg_token()` function), and unregister from the push
service by using the [`"unregisterFromRemoteNotifications"`](../15_library-reference/3462-mobile-unregisterfromremotenotifications.md "This front call unregisters the mobile device from push notifications.") front call.

- When using FCM, you must pass the FCM Sender ID as parameter.
- When using APNs, the parameter must be NULL.

```
...
    LET rec.tm_host = "https://pushreg.example.orion"
    LET rec.tm_port = 4930
    CALL unregister(rec.registration_token, rec.app_user)
...

FUNCTION unregister(
    notification_type STRING,
    registration_token STRING,
    app_user STRING
) RETURNS ()
    IF tm_command( "unregister", notification_type,
                   registration_token, app_user, 0 ) < 0 THEN
       RETURN
    END IF
    TRY
        CALL ui.Interface.frontCall(
                "mobile", "unregisterFromRemoteNotifications",
                [ ], [ ] )
    CATCH
        ERROR "Un-registration failed!"
        RETURN
    END TRY
    MESSAGE "Un-registration succeeded!"
END FUNCTION
```

## Related links

**Related concepts**  

[Firebase Cloud Messaging (FCM)](5113-firebase-cloud-messaging-fcm.md "Follow this procedure to implement push notification with FCM.")

[Apple Push Notification Service (APNs)](5114-apple-push-notification-service-apns.md "Follow this procedure to implement push notification with APNs.")
