---
title: "Push notifications"
source: "fgl-topics/c_fgl_mobile_push_notifications.html"
breadcrumb: "Mobile applications > Push notifications"
type: "concept"
---

# Push notifications

> This section describes how to implement push notification with Genero.

Enterprise mobile applications can use push notifications to produce urgent and important updates
for users.

A push notification is a short message sent by a central server entity to an app
installed on a mobile device. In order to be notified, the app must register itself to a push
service (a global service such as Firebase Cloud Messaging), and register also to a
token maintainer (part of the custom application). To indicate that fresh information
is available, notifications are sent by the push provider to the push
service, which broadcasts notifications to registered devices. The apps can then get details
about the notification and display a little hint to the end user.

![Push notification workflow diagram](../_images/push_notif_1.jpg)

*This figure describes the workflow for a push notification (items in yellow are the components that can be implemented with Genero BDL)*

Workflow:

1. The app registers to the push service.
2. The push service generates a unique token to identify the device+app and returns
   this token to the app.
3. The app transmits the token to the token maintainer.
4. The token maintainer stores the new token in a database.
5. Some event occurs in the global application workflow that requires a push
   notification to warn all registered devices/apps.
6. The push provider reads the database for registered tokens.
7. The push provider sends push notification requests to the push
   service.
8. The push service broadcasts the notification messages to all registered
   devices.

There are several push notification mechanisms available. This chapter covers the
Firebase Cloud Messaging (FCM) and Apple Push Notification services (APNs).

> **Important:**
>
> On March 31 2021 Apple has discontinued the APNs legacy binary protocol. It is now mandatory to
> use APNs services over HTTP/2. However, Genero Web Services API does not support HTTP/2 yet. Until
> GWS supports HTTP/2, you can use external tools such as CURL to send APNs requests from the push
> provider server program.

Common components can be implemented on the same code base for both FCM and APNs push
notification mechanisms: The mobile app and the token maintainer.

The complete source code is available of the FourJSGenero GitHub: <https://github.com/FourjsGenero/ex_push_notification>

## Child topics

- [Firebase Cloud Messaging (FCM)](5113-firebase-cloud-messaging-fcm.md): Follow this procedure to implement push notification with FCM.
- [Apple Push Notification Service (APNs)](5114-apple-push-notification-service-apns.md): Follow this procedure to implement push notification with APNs.
- [Implementing a token maintainer](5115-implementing-a-token-maintainer.md): The token maintainer is a BDL Web Services server program that handles push token registration from mobile apps.
- [Handling notifications in the mobile app](5116-handling-notifications-in-the-mobile-app.md): This topic describes how to handle push notification in the app running on mobile devices.
