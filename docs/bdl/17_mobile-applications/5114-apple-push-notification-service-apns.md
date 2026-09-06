---
title: "Apple Push Notification Service (APNs)"
source: "fgl-topics/c_fgl_mobile_push_notif_apns.html"
breadcrumb: "Mobile applications > Push notifications > Apple Push Notification Service (APNs)"
type: "concept"
---

# Apple Push Notification Service (APNs)

> Follow this procedure to implement push notification with APNs.

> **Important:**
>
> On March 31 2021 Apple has discontinued the APNs legacy binary protocol. It is now mandatory to
> use APNs services over HTTP/2. However, Genero Web Services API does not support HTTP/2 yet. Until
> GWS supports HTTP/2, you can use external tools such as CURL to send APNs requests from the push
> provider server program.

## Introduction to APNs push notification

The push notification solution described in this section is based on the Apple Push Notification
Service.

Apple Push Notification service allows push servers to send notification message data to registered
iOS (and macOS) devices.

For more details, see the Apple Push Notification service (APNs) documentation on
<https://developer.apple.com>.

The APNs service transports and routes a remote notification from a given provider to a given
device. A notification is a short message built from two pieces of data: the device token and the
payload.

Each device needs to be identified by its device token, and the provider must send individual
notification messages for each registered device.

The system involves the following components:

- The Apple Push Notification Service (APNs):

  APNs provides push server and client
  identification. It also handles all aspects of message queuing and delivery to the target
  applications running on registered devices. The APNs system includes a feedback service that can be
  queried to check for devices that have unregistered and no longer need to be notified.
- The database containing device tokens:

  The database used to store device tokens must be a
  multi-user database, because two distinct programs will use the database.
- The device tokens maintainer:

  A Web Services server program maintaining the database of device
  tokens, with application user information. This program must listen to new device registration
  events, store them in a database, and from time to time query the APNs feedback service to check for
  unregistrations.
- The push provider:

  This program will send notification messages to the APNs server. The push
  provider program will query the device token database to know which devices need to be
  notified.
- Devices running the Genero app registered to the push notification server:

  Registered devices
  use the push notification client API to [register](../15_library-reference/3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications."), [get notifications](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") data and [unregister](../15_library-reference/3462-mobile-unregisterfromremotenotifications.md "This front call unregisters the mobile device from push notifications.") from the
  service.

## APNs push notification security

iOS apps must be created with an Apple certificate
for development or distribution, linked to an App ID (or Bundle ID) with push notification enabled.
The provisioning profile used when building the IPA must be linked to the App ID with push enabled.
Certificate, provisioning and bundle id must be specified to the [GMI buildtool](5110-gmibuildtool.md "The gmibuildtool is a utility to create and test applications for an iOS devices.").

Check Apple Push Notification documentation for more details about certificate requirements for
push notifications.

The Apple® Push Notification Certificate
identifies the push notification service for a given mobile app. This certificate is created from
an App ID (or Bundle ID) and is used by the APNs system to dispatch the notification message to the
registered devices.

You can create two types of APN certificates for a given App ID:

- Sandbox (for development and test purpose)
- Production (for deployment)

An APNS push notification provider program needs to establish a secure connection to Apple's
APNs server.

To create an Apple Push Notification
Certificate:

1. Log in to [Apple's Member Center](http://developer.apple.com/membercenter/index.action) with you iOS developer or enterprise account.
2. Select Certificates, Identifiers & Profiles.
3. Under App IDs, make sure that you have created an App ID with the Push Notification service
   enabled, for development and/or distribution.
4. Under Certificates, select the + symbol.
5. Select Apple Push Notification service
   SSL (Sandbox) for development, or Apple Push
   Notification service SSL (Sandbox & Production) for production.
6. Choose the App ID with push notifications service enabled.
7. Follow the instructions to create a Certificate Signing Request (CSR) file from your Mac, then
   click Continue.
8. Back in the web browser and IOS Certificate page, upload the CSR file you have generated.
9. Generate the certificate.
10. Once the certificate is generated, download it. The certificate will be downloaded into your
    Downloads folder, as "aps\_development.cer" file in case the sandbox certificate
    is chosen, or as "aps.cer" in case the production certificate is chosen. In the
    next sections, this certificate file will be referenced as "myapp.cer".
11. Double-click this file to import the certificate into the Mac® Keychain®.
12. The new certificate is now listed in the Certificates list.
13. Open your Keychain app and locate
    the certificate you created, export the private key in p12 format (for example,
    myapp-key.p12). Note that you will be asked for a password to encode the
    .p12 file, and for your session password, to exported Keychain files.

On the Genero push provider server, you will need the public certificate
(myapp.crt file) and the private key (myapp-key.pem file)
for you app.

In order to authenticate the APNs server, you will also need the root certificate authority
(apple\_entrust\_root\_certification\_authority.pem), that can be downloaded from
Apple's web site.

Create the myapp.crt file (public certificate) from the
myapp.cer file, with the openssl x509 command:

```
$ openssl x509 -in myapp.cer -inform der -out myapp.crt
```

Convert
the myapp-key.p12 file (containing the private key) to a
myapp-key.pem format, with the openssl pkcs12
command:

```
$ openssl pkcs12 -in myapp-key.p12 -out myapp-key.pem -nodes -clcerts
```

You need to enter the passphrase for the .p12 file so that
openssl can read it. Then you need to enter a new passphrase that will be used to
encrypt the .pem file.

> **Note:**
>
> Alternatively, if you do not have a Mac, you can generate a Certificate Signing Request directly
> from openssl:
>
> ```
> openssl req -nodes -newkey rsa:2048 -keyout myapp.key \
>     -out CertificateSigningRequest.certSigningRequest
> ```
>
> Generate the Push Notification Certificate (myapp.cer) from the Apple
> developer site, by uploading the CertificateSigningRequest.certSigningRequest
> file generated with openssl.
>
> Download the myapp.cer file and convert this file to a
> .pem file:
>
> ```
> openssl x509 -in myapp.cer -inform der -out aps_cert.pem
> ```
>
> Concat the generated .pem with the private key to produce the final .pem file:
>
> ```
> cat aps_cert.pem myapp.key > aps.pem
> ```

## Identifying target devices

Each APNs client device is identified by a device token. A device token is an opaque
identifier of a device that APNs gives to the device when an app registers itself for push notification.
It enables APNs to locate in a unique manner the device on which the client app is installed. The device
shares the device token with the push provider. The push provider must produce notification messages for
each device by including the device token in the message structure.

> **Important:**
>
> The mobile app obtains its device token by registering to the APNs service
> with the [mobile.registerForRemoteNotifications](../15_library-reference/3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications.") front call. It
> is then in charge of sending its device token to the push provider; typically through a RESTFul request.
> The push provider must collect and store the device tokens, as they need to be specified in a push
> notification message send by the push provided.

## Notification content (payload)

In a notification message, the payload is a JSON-defined property list that specifies
how the user of an app on a device is to be alerted.

The payload must contain a list of "`aps`" records. Each "`aps`" record
represents a notification message to be displayed as a hint on the device (for example, by adding a badge
number to the app icon). The "`aps`" records can also contain custom data in a separate
set of JSON attributes.

In the Genero mobile app, the notification messages are obtained by using the [mobile.getRemoteNotifications](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") front call, after a
`notificationpushed` action was detected with an `ON ACTION`
handler.
> **Important:**
>
> When an iOS app is
> in background, silent push notifications can occur, but notification message data (i.e. the payload)
> may not be available. In such case, GMI is able to detect that a notification arrived (i.e. when
> the app badge number is greater than zero) and raise the `notificationpushed` action,
> but the `getRemoteNotifications` front call will return no message data
> (data return param is `NULL`). If such case, implement a fallback
> mechanism (based on RESTFul web services for example), to contact the push notification provider
> and retrieve the message information.

Example of notification record list (JSON array) returned by the `getRemoteNotifications`
front call:

```
[
 {
  "aps" :
  {
    "alert" : "My first push",
    "badge" : 1,
    "sound" : "default",
    "content-available" : 1
  }
 },
 {
  "aps" :
  {
    "alert" :
    {
      "title" : "Push",
      "body" : "My second push"
    }
    "badge" : 2,
    "sound" : "default",
    "content-available" : 1
  },
  "new_ids" : [ "XV234", "ZF452", "RT563" ],
  "updated_ids" : [ "AC634", "HJ153" ]
 }
]
```

## Badge number handling

With APNs, badge number handling is in charge of the application code: The push provider sends a
badge number in the payload records, the app can check the message content, and must communicate with
a server component, to indicate that the notification message has been consumed. The server program
can then maintain a badge number for each registered device, decrementing the badge number.

In order to set or query the badge number for your app, use the following front calls:

- [ios.setBadgeNumber](../15_library-reference/3471-ios-setbadgenumber.md "Sets the current badge number associated to the app.")
- [ios.getBadgeNumber](../15_library-reference/3470-ios-getbadgenumber.md "Returns the current badge number associated to the app.")

In this tutorial, badge numbers are stored on the server database. The token maintainer handlers
requests from apps to sync the badge number for a given device token, and the push provider program
reads the database to set the badge number in the notification payload. When the app consumes
messages, it queries and resets the app badge number with the `getBadgeNumber` /
`setBadgeNumber` front calls, and informs the token maintainer to sync the badge
number in the central database.

## Implementing the registration tokens maintainer

To handle device registrations on the server side of your application, the same code base can be
used for FCM and APNs token-based frameworks.

For more details, see [Implementing a token maintainer](5115-implementing-a-token-maintainer.md "The token maintainer is a BDL Web Services server program that handles push token registration from mobile apps.").

## Implementing the APNS push provider

The push provider will produce application notification messages that will be transmitted to the
APNs service. The APNs service will then spread them to all registered mobile devices, identified by
their device token.

To send notification messages, the push provider must send HTTP/2 POST requests to the APNS service.

To send a notification message, the push provider must know the device tokens of the registered devices
/ applications. A distinct notification message must be sent for each registered device.

- For APNs
  development:

  ```
  https://api.sandbox.push.apple.com:443/3/device/device-token
  ```
- For APNs
  production:

  ```
  https://api.push.apple.com:443/3/device/device-token
  ```

For more details, see [Sending Notification Requests to APNs](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns/) documentation.

Sending APNS POST requests must be done over HTTP/2. Since GWS does not yet support HTTP/2 protocol,
you need to use an external tool such as CURL.

The following example demonstrates how to implement
a function to send an APNs notification message. The function takes a device token and a JSON object as
parameters.

```
IMPORT util

CONSTANT APNS_APPID = "app-bundle-id"
CONSTANT APNS_CERTIF = "path-to-certif-file"
CONSTANT APNS_DEVICE_URL = "https://api.sandbox.push.apple.com:443/3/device/%1"

FUNCTION apns_send_notif_http(
    deviceTokenHexa STRING,
    push_type STRING,
    priority INTEGER,
    notif_obj util.JSONObject
) RETURNS STRING
    DEFINE dt DATETIME YEAR TO SECOND,
           exp INTEGER,
           data STRING,
           cmd STRING,
           s INTEGER

    LET dt = CURRENT + INTERVAL(10) MINUTE TO MINUTE
    LET exp = util.Datetime.toSecondsSinceEpoch(dt)

    IF length(push_type) == 0 THEN
        LET push_type = "alert"
    END IF

    LET data = notif_obj.toString()

    LET cmd = "curl -vs --http2 "
              || SFMT('--header "apns-topic: %1" ',APNS_APPID)
              || SFMT('--header "apns-push-type: %1" ',push_type)
              || IIF(priority IS NOT NULL, SFMT('--header "apns-priority: %1" ',priority), "")
              || SFMT('--cert "%1" ',APNS_CERTIF) -- *.pem
              || SFMT("--data '%1' ",data) -- JSON string!
              || SFMT(APNS_DEVICE_URL,deviceTokenHexa)

    RUN cmd RETURNING s
    IF s != 0 THEN
        RETURN "ERROR : Failed to execute CURL command"
    END IF

    RETURN NULL

END FUNCTION
```

The next code example implements a function that creates the JSON object defining notification
content (payload). That object can be passed to the `apns_send_notif_http()`
function described above:

```
IMPORT util

FUNCTION apns_simple_popup_notif(
    notif_obj util.JSONObject,
    msg_title STRING,
    user_data STRING,
    badge_number INTEGER
) RETURNS ()
    DEFINE aps_obj, data_obj util.JSONObject

    LET aps_obj = util.JSONObject.create()
    CALL aps_obj.put("alert", msg_title)
    CALL aps_obj.put("sound", "default")
    CALL aps_obj.put("badge", badge_number)
    CALL aps_obj.put("content-available", 1)
    CALL notif_obj.put("aps", aps_obj)

    LET data_obj = util.JSONObject.create()
    CALL data_obj.put("other_info", user_data)

    CALL notif_obj.put("custom_data", data_obj)

END FUNCTION
```

In order to use the tokens database maintained by a [token maintainer program](5115-implementing-a-token-maintainer.md "The token maintainer is a BDL Web Services server program that handles push token registration from mobile apps."), your APNs push
provider can collect device tokens as shown in the example below. Note that the dynamic array
contains token ids and badge numbers:

```
FUNCTION apns_collect_tokens(
    reg_ids DYNAMIC ARRAY OF RECORD
                       token STRING,
                       badge INTEGER
                   END RECORD
) RETURNS ()
    DEFINE rec RECORD
               id INTEGER,
               notification_type VARCHAR(10),
               registration_token VARCHAR(250),
               badge_number INTEGER,
               app_user VARCHAR(50),
               reg_date DATETIME YEAR TO FRACTION(3)
           END RECORD,
           x INTEGER
    DECLARE c1 CURSOR FOR
      SELECT * FROM tokens
       WHERE notification_type = "APNS"
    CALL reg_ids.clear()
    FOREACH c1 INTO rec.*
        LET x = reg_ids.getLength() + 1
        LET reg_ids[x].token = rec.registration_token
        LET reg_ids[x].badge = rec.badge_number
    END FOREACH
END FUNCTION
```

In order to handle badge numbers for each registered device, implement a function to update
badge numbers in database:

```
FUNCTION save_badge_number(token STRING, badge INTEGER) RETURNS ()
    UPDATE tokens SET
        badge_number = badge
    WHERE registration_token = token
END FUNCTION
```

The above functions can then be used to send a push message to all registered devices:

```
IMPORT util

FUNCTION apns_send_message(
    msg_title STRING,
    user_data STRING
) RETURNS STRING
    DEFINE reg_ids DYNAMIC ARRAY OF RECORD
                       token STRING,
                       badge INTEGER
                   END RECORD,
           notif_obj util.JSONObject,
           info_msg STRING,
           new_badge, i INTEGER
    CALL apns_collect_tokens(reg_ids)
    IF reg_ids.getLength() == 0 THEN
       RETURN "No registered devices..."
    END IF
    LET info_msg = "Send:"
    FOR i=1 TO reg_ids.getLength()
        LET new_badge = reg_ids[i].badge + 1
        CALL save_badge_number(reg_ids[i].token, new_badge)
        LET notif_obj = util.JSONObject.create()
        CALL apns_simple_popup_notif(notif_obj, msg_title, user_data, new_badge)
        LET info_msg = info_msg, "\n",
            apns_send_notif_http(reg_ids[i].token, NULL, NULL, notif_obj)
    END FOR
    RETURN info_msg
END FUNCTION
```

## Handle push notifications in mobile apps

To handle push notifications in mobile apps, the same code base can be used for FCM and APNs
token-based frameworks.

For more details see [Handling notifications in the mobile app](5116-handling-notifications-in-the-mobile-app.md "This topic describes how to handle push notification in the app running on mobile devices.").

## Related links

**Related concepts**  

[Firebase Cloud Messaging (FCM)](5113-firebase-cloud-messaging-fcm.md "Follow this procedure to implement push notification with FCM.")
