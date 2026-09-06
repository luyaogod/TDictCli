---
title: "Firebase Cloud Messaging (FCM)"
source: "fgl-topics/c_fgl_mobile_push_notif_gcm.html"
breadcrumb: "Mobile applications > Push notifications > Firebase Cloud Messaging (FCM)"
type: "concept"
---

# Firebase Cloud Messaging (FCM)

> Follow this procedure to implement push notification with FCM.

## Introduction to FCM push notification

The push notification solution described in this section is based on the Firebase
Cloud Messaging service. Familiarize yourself with FCM by visiting the [Firebase Cloud Messaging](http://firebase.google.com) web
site.
> **Important:**
>
> Firebase Cloud Messaging comes with implicit features that you need to carefully consider. For
> example, user data collection for analytics can be done by FCM when doing push notifications; This
> feature, disabled by default when building Genero apps for Android, can be activated with [the
> `-bfac`
> gmabuildtool option](5106-gmabuildtool.md). Take the time to read the documentation available at
> [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging).

Firebase Cloud Messaging services allow push servers to send notification message data to registered
Android™ or iOS devices.

The system involves the following actors:

- The Firebase Cloud Messaging (FCM) service:

  FCM provides push server and client
  identification. It also handles all aspects of queuing of messages and delivery to the target
  application running on registered devices.
- The database containing device tokens:

  The database used to store device tokens must be a
  multi-user database, because two distinct programs will use the database.
- The registration tokens maintainer:

  A Web services server program maintaining the database of
  registration tokens with application user information. This program must listen to new device
  registration events and store them in a database. The push server program can then query this database
  to build the list of registration tokens to identify the devices to be notified.
- The push server program:

  Implemented by a third-party service or as a Genero BDL program using
  the Web services API. This push server program will send notification messages to FCM with two
  connection servers (HTTP and XMPP).
- Devices running the Genero app registered to the push notification server:

  Registered devices use the push notification client API to [register](../15_library-reference/3457-mobile-registerforremotenotifications.md "This front call registers a mobile device for push notifications."), [get notifications](../15_library-reference/3452-mobile-getremotenotifications.md "This front call retrieves push notification messages.") data and [unregister](../15_library-reference/3462-mobile-unregisterfromremotenotifications.md "This front call unregisters the mobile device from push notifications.") from the
  service.

> **Note:**
>
> In order to send push notifications to a set of devices in one send operation, FCM supports the
> creation of device groups, that will be referenced with a unique token. This solution
> is not detailed here. If you want to use FCM device grouping, refere to the [Firebase documentation](https://firebase.google.com/docs/cloud-messaging/android/device-group).

## Prerequisites

The following elements are needed to setup a push notification system using FCM:

- [GMA requirements](../04_installation/0050-install-genero-mobile-for-android.md "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA.") (JDK, Android SDK, etc)
- A Google developer account.
- An FCM project created in
  the Google FCM console.
- The google-auth-library-oauth2-http-\*.jar java archive, used by the push server.
- [Apache
  Maven](https://maven.apache.org/index.html), to get the JAR deps for
  google-auth-library-oauth2-http-\*.jar

## Get the Java Google authentication JAR

For authentication, the push
server needs google-auth-library-oauth2-http-\*.jar, and related
dependency jar files.

Steps to get the required JARs:

1. Make sure Maven is properly installed and configured
   with:

   ```
   mvn -version
   ```
2. Create and go to a new directory to hold the Maven project
3. Create a file name pom.xml with the following
   content:

   ```
   <?xml version="1.0" encoding="UTF-8"?>
   <project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
     <modelVersion>4.0.0</modelVersion>
     <groupId>com.dummy.app</groupId>
     <artifactId>dummy-app</artifactId>
     <version>1.0-SNAPSHOT</version>
     <name>dummy-app</name>
     <url>http://www.dummy.com</url>
     <properties>
       <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
       <maven.compiler.source>1.7</maven.compiler.source>
       <maven.compiler.target>1.7</maven.compiler.target>
     </properties>
     <dependencies>
       <!-- https://mvnrepository.com/artifact/com.google.auth/google-auth-library-oauth2-http -->
       <dependency>
         <groupId>com.google.auth</groupId>
         <artifactId>google-auth-library-oauth2-http</artifactId>
         <version>1.20.0</version>
       </dependency>
     </dependencies>
     <build>
     </build>
   </project>
   ```
4. Execute the following command:

   ```
   mvn dependency:copy-dependencies
   ```
5. Add all JAR files found in $PWD/target/dependency into the CLASSPATH
   environment variable.

## Creating a FCM project

To initiate a push notification service dedicated to your applications, you must first create a
Firebase Cloud Messaging project in the FCM console. Creating an FCM project will provide you:

- The Project ID (only for info, it is not used in the code)
- The Project Number (FCM\_PROJECT\_NUM)
- The google-services.json configuration file (to be bundled with your
  app.)
- The JSON file containing the private key from Google service accounts (FCM\_APP\_CRED)

Steps to setup your FCM project:

1. Go to the [FCM
   console](https://console.firebase.google.com) and log in with your Google developer account.
2. Add a new FCM project.
3. Write down the FCM Project Number (`FCM_PROJECT_NUM`). This id will be used in
   the URI to send push notifications from the push server program.
4. Add an (Android) app to your FCM project: Specify the same package name specified when building
   your app with the `--build-app-package-name` option of
   gmabuildtool.
5. Download the google-services.json configuration file. This file needs to be
   added to the [appdir](5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA).") with the
   other program files.
6. Skip the other FCM project creation steps.
7. In the project overview page, go to the settings of the newly created project.
8. Select the "Cloud Messaging" panel.
9. Generate a new private key from the Service Accounts tab. Download and securely store this JSON
   file, referenced here as "service-account-file.json". This private key will be
   used to build the access token, used in the `"Authorization"` HTTP
   header as `"Bearer access-token"`, when posting a message to the
   FCM server endpoint.

For more details about FCM project creation, visit the [Firebase Cloud Messaging](http://firebase.google.com/) web site.

## Implementing the registration tokens maintainer

To handle device registrations on the server side of your application, the same code base can be
used for FCM and APNs token-based frameworks.

For more details, see [Implementing a token maintainer](5115-implementing-a-token-maintainer.md "The token maintainer is a BDL Web Services server program that handles push token registration from mobile apps.").

## Implementing the push server

The push server will produce application notification messages that will be transmitted to the
FCM service. The FCM service will then spread them to all mobile devices registered to the service
with the Sender ID.

> **Important:**
>
> The size of an FCM notification content cannot exceed 4 Kilobytes. If more information needs to
> be passed, apps must contact the server part to query for more information after receiving the push
> message; however, this is only possible when a network is available.

The push server will use RESTful HTTP POST requests to send notifications through the FCM service
to the following URL:

`"https://fcm.googleapis.com/v1/projects/fcm-project-num/messages"`

In this URL template, fcm-project-num needs to be replaced by the your FCM
Project Number (`FCN_PROJECT_NUM`).

For more details about the JSON request structure in a FCM HTTP POST, visit the [Firebase Cloud Messaging web
site](https://firebase.google.com/).

The HTTP POST header must contain the following
attributes:

```
Content-Type: application/json
Authorization: Bearer access-token
```

where access-token is build from the FCM project private key. Use Java to
produce the access token:

```
IMPORT JAVA com.google.auth.oauth2.GoogleCredentials
IMPORT JAVA java.io.FileInputStream
IMPORT JAVA java.lang.String
IMPORT JAVA java.util.Arrays

FUNCTION get_access_token() RETURNS STRING
    TYPE string_array_type ARRAY[] OF java.lang.String
    DEFINE credentials_file STRING
    DEFINE google_credentials GoogleCredentials
    DEFINE file_input_stream FileInputStream
    DEFINE scopes string_array_type
    DEFINE access_token STRING
    LET credentials_file = fgl_getenv("FCM_APP_CRED")
    IF length(credentials_file) == 0 THEN
       LET credentials_file = FCM_APP_CRED
       IF credentials_file == "..." THEN
          DISPLAY "ERROR: FCM_APP_CRED is not defined."
          EXIT PROGRAM 1
       END IF
    END IF
    LET scopes = string_array_type.create(1)
    LET scopes[1] = FCM_SCOPE_MESSAGING
    LET file_input_stream = FileInputStream.create(credentials_file)
    LET google_credentials = GoogleCredentials
          .fromStream(file_input_stream)
          .createScoped(Arrays.asList(scopes))
    CALL google_credentials.refresh()
    LET access_token = google_credentials.getAccessToken().getTokenValue()
    RETURN access_token
END FUNCTION
```

The push server program can be implemented with the Web Services API to make RESTful requests as
follows:

```
IMPORT com
IMPORT util

CONSTANT FCM_PUSH_BASE_URI = "https://fcm.googleapis.com/v1/projects/%1/messages:%2"
CONSTANT FCM_HTTP_CONTENT_TYPE = "application/json; charset=utf-8"

FUNCTION fcm_send_notif_http(
    project_num STRING,
    access_token STRING,
    req_msg STRING -- JSON
) RETURNS STRING
    DEFINE req com.HttpRequest,
           resp com.HttpResponse,
           res STRING
    IF req_msg.getLength() > 4096 THEN
       LET res = "ERROR : FCM message cannot exceed 4 kilobytes"
       RETURN res
    END IF
    LET req = com.HttpRequest.Create(SFMT(FCM_PUSH_BASE_URI,project_num,"send"))
    CALL req.setHeader("Content-Type", FCM_HTTP_CONTENT_TYPE)
    CALL req.setHeader("Authorization", SFMT("Bearer %1", access_token))
    CALL req.setMethod("POST")
    TRY
        CALL req.doTextRequest(req_msg)
        LET resp = req.getResponse()
        CASE resp.getStatusCode()
        WHEN 400 -- Can happen when a device id is invalid (has unregistered)
            LET res = "400 error (maybe some device ids are invalid?)"
        WHEN 200
            LET res = "OK"
        OTHERWISE
            LET res = SFMT("HTTP Error (%1) %2",
                           resp.getStatusCode(),
                           resp.getStatusDescription())
        END CASE
    CATCH
        LET res = SFMT("ERROR : %1 (%2)", status, sqlca.sqlerrm)
    END TRY
    RETURN res
END FUNCTION
```

The body of the HTTP POST request must be a JSON formatted record using a structure similar to
the following example:

```
{
   "message":{
      "token": "registraton-id",
      "data":{
          "title": "Hello world!",
          "content": "This is a push notification demo",
          "icon": "genero_notification"
      }
   }
}
```

The `"registration-id"` is a device token obtained
from the token maintainer database. This `"token"` field can also contain a
"notification-key" which identifies a group of devices, obtained via the [HTTP v1 API to manage device groups](https://firebase.google.com/docs/cloud-messaging/android/device-group).

The recommendation for GMA is that the `"icon"` resource is packaged in the APK
and is accessible by name (as gma\_ic\_genero.png in the drawable folders).

The next code example implements a function collects device registration ids
(`fcm_collect_reg_tokens`), and for each registered device, constructs a message in
JSON (`msgrec`) and sends a push notification
(`fcm_send_notif_http`):

```
TYPE t_fcm_message RECORD
       message RECORD
         token STRING,
         data RECORD
           title STRING,
           content STRING,
           icon STRING,
           extra_data STRING
         END RECORD
       END RECORD
     END RECORD

FUNCTION fcm_send_notification(rec t_push_rec) RETURNS STRING
    DEFINE reg_ids DYNAMIC ARRAY OF STRING,
           x INTEGER,
           msgrec t_fcm_message,
           res STRING,
           res_list DYNAMIC ARRAY OF STRING
    CALL fcm_collect_reg_tokens(reg_ids)
    IF reg_ids.getLength() == 0 THEN
       RETURN "No registered devices..."
    END IF
    LET msgrec.message.data.title = rec.msg_title
    LET msgrec.message.data.content = rec.user_data
    LET msgrec.message.data.icon = "genero_notification"
    FOR x=1 TO reg_ids.getLength()
        LET msgrec.message.token = reg_ids[x]
        LET res = fcm_send_notif_http( rec.project_num, rec.access_token,
                                       util.JSON.stringify(msgrec) )
        LET res_list[x] = SFMT("Send to token %1 : %2",
                               reg_ids[x].subString(1,10)||"...",res)
    END FOR
    RETURN util.JSON.format(util.JSON.stringify(res_list))
END FUNCTION
```

To use the tokens database maintained by a [token maintainer](5115-implementing-a-token-maintainer.md "The token maintainer is a BDL Web Services server program that handles push token registration from mobile apps.") program, your FCM push server can collect registration tokens as shown in
the following
example:

```
FUNCTION fcm_collect_tokens(reg_ids DYNAMIC ARRAY OF STRING) RETURNS ()
    DEFINE rec RECORD
               id INTEGER,
               notification_type VARCHAR(10),
               registration_token VARCHAR(250),
               badge_number INTEGER,
               app_user VARCHAR(50),
               reg_date DATETIME YEAR TO FRACTION(3)
           END RECORD
    DECLARE c1 CURSOR FOR
      SELECT * FROM tokens
       WHERE notification_type = "FCM"
    CALL reg_ids.clear()
    FOREACH c1 INTO rec.*
        CALL reg_ids.appendElement()
        LET reg_ids[reg_ids.getLength()] = rec.registration_token
    END FOREACH
END FUNCTION
```

## Handle push notifications in mobile apps

To handle push notifications in mobile apps, the same code base can be used for FCM and APNs
token-based frameworks.

For more details see [Handling notifications in the mobile app](5116-handling-notifications-in-the-mobile-app.md "This topic describes how to handle push notification in the app running on mobile devices.").

## Related links

**Related concepts**  

[Apple Push Notification Service (APNs)](5114-apple-push-notification-service-apns.md "Follow this procedure to implement push notification with APNs.")
