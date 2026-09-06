---
title: "Genero Mobile common front calls"
source: "fgl-topics/c_fgl_frontcalls_mobile.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls"
type: "concept"
---

# Genero Mobile common front calls

> This section describes common front calls provided by all mobile front-ends.

This table shows the functions implemented by all mobile front-ends in the
"`mobile`" module.

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("mobile", "chooseContact", [], [result]) | Lets the user choose a contact from the mobile device contact list and returns the vCard. |
| ui.Interface.frontCall("mobile", "choosePhoto", [], [path]) | Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier. |
| ui.Interface.frontCall("mobile", "chooseVideo", [], [path]) | Lets the user select a video from the mobile device's video gallery and returns a video identifier. |
| ui.Interface.frontCall("mobile", "createNotification", [options], [id] ) | Creates or updates a local notification to be displayed on the mobile device. |
| ui.Interface.frontCall("mobile", "clearNotifications", [options], [status] ) | Drops notifications displayed on the mobile device. |
| ui.Interface.frontCall("mobile", "composeMail", [to, subject, content, cc, bcc, attachments ...], [result]) | Invokes the user's default mail application for a new mail to send. |
| ui.Interface.frontCall("mobile", "composeSMS", [ recipients, content ], [ result ] ) | Sends an SMS text to one or more phone numbers. |
| ui.Interface.frontCall("mobile", "connectivity", [], [result] ) | Returns the type of network available for the mobile device. |
| ui.Interface.frontCall("mobile", "getGeolocation", [], [status, latitude, longitude] ) | Returns the Global Positioning System (GPS) location of a mobile device. |
| ui.Interface.frontCall("mobile", "getLastNotificationInteractions", [], [notifications] ) | Get the last user interactions on mobile app notifications. |
| ui.Interface.frontCall("mobile","getRemoteNotifications", [], [data] ) | This front call retrieves push notification messages. |
| ui.Interface.frontCall("mobile", "importContact", [vcard], [result] ) | Creates a new contact, or merges to an existing entry, the contact details passed in a vCard string. |
| ui.Interface.frontCall("mobile", "isEmulator", [], [result] ) | Indicates if the mobile app runs or displays forms on an emulator/simulator. |
| ui.Interface.frontCall("mobile", "isForeground", [], [result] ) | Indicates if the mobile app is in foreground mode. |
| ui.Interface.frontCall("mobile", "newContact", [defaults],[vcard]) | Opens contact input form to create a new entry in the contact database. |
| ui.Interface.frontCall("mobile","registerForRemoteNotifications", [], [registration-token] ) | This front call registers a mobile device for push notifications. |
| ui.Interface.frontCall("mobile", "runOnServer", [ appurl, timeout ], [] ) | Run an application from the Genero Application Server using the specified URL. |
| ui.Interface.frontCall("mobile", "scanBarCode", [options], [code, type] ) | Allow the user to scan a barcode with a mobile device |
| ui.Interface.frontCall("mobile", "takePhoto", [], [path] ) | Lets the user take a picture with the mobile device and returns the corresponding picture identifier. |
| ui.Interface.frontCall("mobile", "takeVideo", [], [path]) | Lets the user take a video with the mobile device and returns the corresponding video identifier. |
| ui.Interface.frontCall("mobile","unregisterFromRemoteNotifications", [], [] ) | This front call unregisters the mobile device from push notifications. |

## Related links

**Related concepts**  

[Genero Mobile Android front calls](3463-genero-mobile-android-front-calls.md "This section describes front calls specific to the Android platform.")

[Genero Mobile iOS front calls](3469-genero-mobile-ios-front-calls.md "This section describes front calls specific to the iOS platform.")
