---
title: "Implementing a token maintainer"
source: "fgl-topics/c_fgl_mobile_push_notif_token_maint.html"
breadcrumb: "Mobile applications > Push notifications > Implementing a token maintainer"
type: "concept"
---

# Implementing a token maintainer

> The token maintainer is a BDL Web Services server program that handles push token registration from mobile apps.

## Basics

In order to implement a push notification mechanism, you need to set up a server part (token
maintainer and push notification server), based on a push notification framework such as Firebase
Cloud Messaging (FCM) or Apple Push Notification service (APNs). In addition, you need to handle
notification events in your mobile app. This section describes how to implement the token
maintainer, the server program that maintains the list of registered devices (i.e. registration
tokens for FCM or device tokens for APNs).

The max length of a push client token can vary depending on the push framework provider. If you
need to store registration tokens in a database, check the max size for a token and consider using a
large column type such as `VARCHAR(250)`.

The same code base can be used for Android™
(using FCM) and iOS (using APNs) applications: The token maintainer will basically handle RESTful
HTTP requests coming from the internet for token registration and token unregistration. For each
of these requests, the program will insert a new record or delete an existing record in a dedicated
database table.

The database used to store tokens must be created before starting the token maintainer program. By
default, the push demo program uses SQLite (dbmsqt) and the name of the database is "tokendb". To
create this SQLite database, simply create an empty file with this name.

The push provider/server program can then query the tokens table to build the list of target
devices for push notifications.

In the context of APNs, the token maintainer must also handle badge numbers for each registered
device: When consuming notification messages, the iOS app must inform the token maintainer that
the badge number has changed. This function is implemented with the "`badge_number`"
command.

The token maintainer is a Web Services server program which must be deployed behind a GAS to
handle load balancing. You can, however, write code to test your program in development without a GAS.

The act of registering/unregistering push tokens is application specific: When registering tokens,
you typically want to add application user information. Genero BDL allows you to implement a token
maintainer in a simple way.

## MAIN block and database creation

Start with the MAIN block, and the connection to a database. In this tutorial, we use SQLite as
the database. The program will automatically create the database file and the tokens table if it
does not yet exist.

```
IMPORT util
IMPORT com
IMPORT os

CONSTANT DEFAULT_PORT = 9999

MAIN
    CALL open_create_db() 
    CALL handle_registrations()
END MAIN

FUNCTION open_create_db() RETURNS ()
    DEFINE dbsrc VARCHAR(100),
           x INTEGER
    IF NOT os.Path.exists("tokendb") THEN
       CALL create_empty_file("tokendb")
    END IF
    LET dbsrc = "tokendb+driver='dbmsqt'"
    CONNECT TO dbsrc
    WHENEVER ERROR CONTINUE
    SELECT COUNT(*) INTO x FROM tokens
    WHENEVER ERROR STOP
    IF sqlca.sqlcode<0 THEN
       CREATE TABLE tokens (
              id INTEGER NOT NULL PRIMARY KEY,
              notification_type VARCHAR(10) NOT NULL,
              registration_token VARCHAR(250) NOT NULL UNIQUE,
              badge_number INTEGER NOT NULL,
              app_user VARCHAR(50) NOT NULL, -- UNIQUE
              reg_date DATETIME YEAR TO FRACTION(3) NOT NULL
       )
    END IF
END FUNCTION

FUNCTION create_empty_file(fn STRING) RETURNS ()
    DEFINE c base.Channel
    LET c = base.Channel.create()
    CALL c.openFile(fn, "w")
    CALL c.close()
END FUNCTION
```

## Handling registration and unregistration requests

The next function is typical Web Service server code using the Web Services API to handle RESTful
requests. Note that the TCP port is defined as a constant that is used to set FGLAPPSERVER
automatically when not running behind the GAS:

```
IMPORT util
IMPORT com

CONSTANT DEFAULT_PORT = 9999

MAIN
    ...
    CALL handle_registrations()
END MAIN

FUNCTION handle_registrations() RETURNS ()
    DEFINE req com.HttpServiceRequest,
           url, method, version, content_type STRING,
           reg_data, reg_result STRING
    IF length(fgl_getenv("FGLAPPSERVER"))==0 THEN
       -- Normally, FGLAPPSERVER is set by the GAS
       DISPLAY SFMT("Setting FGLAPPSERVER to %1", DEFAULT_PORT)
       CALL fgl_setenv("FGLAPPSERVER", DEFAULT_PORT)
    END IF
    CALL com.WebServiceEngine.Start()
    WHILE TRUE
       TRY
          LET req = com.WebServiceEngine.GetHTTPServiceRequest(20)
       CATCH
          IF status==-15565 THEN
             DISPLAY "TCP socket probably closed by GAS, stopping process..."
             EXIT PROGRAM 0
          ELSE
             DISPLAY "Unexpected getHttpServiceRequest() exception: ", status
             DISPLAY "Reason: ", sqlca.sqlerrm
             EXIT PROGRAM 1
          END IF
       END TRY
       IF req IS NULL THEN -- timeout
          DISPLAY SFMT("HTTP request timeout...: %1", CURRENT YEAR TO FRACTION)
          CALL show_tokens()
          CONTINUE WHILE
       END IF
       LET url = req.getUrl()
       LET method = req.getMethod()
       IF method IS NULL OR method != "POST" THEN
          IF method == "GET" THEN
             CALL req.sendTextResponse(200,NULL,"Hello from token maintainer...")
          ELSE
             DISPLAY SFMT("Unexpected HTTP request: %1", method)
             CALL req.sendTextResponse(400,NULL,"Only POST requests supported")
          END IF
          CONTINUE WHILE
       END IF
       LET version = req.getRequestVersion()
       IF version IS NULL OR version != "1.1" THEN
          DISPLAY SFMT("Unexpected HTTP request version: %1", version)
          CONTINUE WHILE
       END IF
       LET content_type = req.getRequestHeader("Content-Type")
       IF content_type IS NULL
          OR content_type NOT MATCHES "application/json*" -- ;Charset=UTF-8
       THEN
          DISPLAY SFMT("Unexpected HTTP request header Content-Type: %1", content_type)
          CALL req.sendTextResponse(400,NULL,"Bad request")
          CONTINUE WHILE
       END IF
       TRY
          CALL req.readTextRequest() RETURNING reg_data
       CATCH
          DISPLAY SFMT("Unexpected HTTP request read exception: %1", status)
       END TRY
       LET reg_result = process_command(url, reg_data)
       CALL req.setResponseCharset("UTF-8")
       CALL req.setResponseHeader("Content-Type","application/json")
       CALL req.sendTextResponse(200,NULL,reg_result)
    END WHILE
END FUNCTION
```

## Processing registration and unregistration commands

The next function is called when a RESTful request is to be processed. The URL will define the
type of command to be executed by the server:

- If the URL contains "/token\_maintainer/register", a new token must be inserted in the
  database.
- If the URL contains "/token\_maintainer/unregister", an existing token must be deleted from the
  database.

```
IMPORT util

FUNCTION process_command(url STRING, data STRING) RETURNS STRING
    DEFINE data_rec RECORD
               notification_type VARCHAR(10),
               registration_token VARCHAR(250),
               badge_number INTEGER,
               app_user VARCHAR(50)
           END RECORD,
           p_id INTEGER,
           p_ts DATETIME YEAR TO FRACTION(3),
           result_rec RECORD
               status INTEGER,
               message STRING
           END RECORD,
           result STRING
    LET result_rec.status = 0
    TRY
       CASE
         WHEN url MATCHES "*token_maintainer/register"
           CALL util.JSON.parse( data, data_rec )
           SELECT id INTO p_id FROM tokens
                  WHERE registration_token = data_rec.registration_token
           IF p_id > 0 THEN
              LET result_rec.status = 1
              LET result_rec.message = SFMT("Token already registered:\n [%1]",
                                            data_rec.registration_token)
              GOTO pc_end
           END IF
           SELECT MAX(id) + 1 INTO p_id FROM tokens
           IF p_id IS NULL THEN LET p_id=1 END IF
           LET p_ts = util.Datetime.toUTC(CURRENT YEAR TO FRACTION(3))
           WHENEVER ERROR CONTINUE
           INSERT INTO tokens
               VALUES( p_id, data_rec.notification_type,
                       data_rec.registration_token, 0, data_rec.app_user, p_ts )
           WHENEVER ERROR STOP
           IF sqlca.sqlcode==0 THEN
              LET result_rec.message = SFMT("Token is now registered:\n [%1]",
                                            data_rec.registration_token)
           ELSE
              LET result_rec.status = -2
              LET result_rec.message = SFMT("Could not insert token in database:\n [%1]",
                                            data_rec.registration_token)
           END IF
         WHEN url MATCHES "*token_maintainer/unregister"
           CALL util.JSON.parse( data, data_rec )
           DELETE FROM tokens
                  WHERE registration_token = data_rec.registration_token
           IF sqlca.sqlerrd[3]==1 THEN
              LET result_rec.message = SFMT("Token unregistered:\n [%1]",
                                            data_rec.registration_token)
           ELSE
              LET result_rec.status = -3
              LET result_rec.message = SFMT("Could not find token in database:\n [%1]",
                                            data_rec.registration_token)
           END IF
         WHEN url MATCHES "*token_maintainer/badge_number"
            CALL util.JSON.parse( data, data_rec )
            WHENEVER ERROR CONTINUE
              UPDATE tokens
                 SET badge_number = data_rec.badge_number
               WHERE registration_token = data_rec.registration_token
            WHENEVER ERROR STOP
            IF sqlca.sqlcode==0 THEN
               LET result_rec.message =
                   SFMT("Badge number updated for Token:\n [%1]\n New value:[%2]\n",
                        data_rec.registration_token, data_rec.badge_number)
            ELSE
               LET result_rec.status = -4
               LET result_rec.message = SFMT("Badge update failed for token:\n [%1]",
                                             data_rec.registration_token)
            END IF
       END CASE
    CATCH
       LET result_rec.status = -1
       LET result_rec.message = SFMT("Failed to register token:\n [%1]",
                                     data_rec.registration_token)
    END TRY
LABEL pc_end:
    DISPLAY result_rec.message
    LET result = util.JSON.stringify(result_rec)
    RETURN result
END FUNCTION
```

## Showing the current registered tokens

The following function is called after a WebServiceEngine timeout, when no request is to be
processed. Its purpose is just to show the current list of registered tokens in a server log
(stdout):

```
FUNCTION show_tokens() RETURNS ()
    DEFINE rec RECORD -- Use CHAR to format
               id INTEGER,
               notification_type CHAR(10),
               registration_token CHAR(250),
               badge_number INTEGER,
               app_user CHAR(50),
               reg_date DATETIME YEAR TO FRACTION(3)
           END RECORD
    DECLARE c1 CURSOR FOR SELECT * FROM tokens ORDER BY id
    FOREACH c1 INTO rec.*
        DISPLAY "   ", rec.id, ": ",
                       rec.notification_type,": ",
                       rec.app_user[1,10], " / ",
                       "(",rec.badge_number USING "<<<<&", ") ",
                       rec.registration_token[1,20],"..."
    END FOREACH
    IF rec.id == 0 THEN
       DISPLAY "No tokens registered yet..."
    END IF
END FUNCTION
```

## Related links

**Related concepts**  

[Firebase Cloud Messaging (FCM)](5113-firebase-cloud-messaging-fcm.md "Follow this procedure to implement push notification with FCM.")

[Apple Push Notification Service (APNs)](5114-apple-push-notification-service-apns.md "Follow this procedure to implement push notification with APNs.")
