---
title: "OAuthAPI.init()"
source: "fgl-topics/c_gws_oauthapi_init.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.init()"
type: "concept"
---

# OAuthAPI.init()

> To be called in a Genero application accessing a secure RESTful web service started behind a Genero Application Server.

## Syntax

```
FUNCTION Init( 
   cnx_timeout INTEGER, 
   client_id STRING,
   client_secret STRING )
  RETURNS BOOLEAN
```

1. cnx\_timeout is a connection timeout to the REST service
   with value in seconds.
2. client\_id is the application ID assigned to the app when
   registered.
3. client\_secret is the application secret created for the
   app.

Returns `FALSE` if the mandatory access token is null.

## Usage

Use the `OAuthAPI.init()` function to initialize the information provided by the
Identity Provider for accessing a secure RESTful web service. This function retrieves environment
variables containing the OAuth information obtained from the Genero Identity Provider (GIP) or a
third-party Identity Provider.

The `OAuthAPI.init()` function registers the access token and relevant access
information, such as the subject, scopes, and endpoints, with the Genero Web Services (GWS) engine,
ensuring that the application can securely interact with the web service.

In case of error, a `NULL` value will be returned.

## OAuthAPI.init function

```
IMPORT FGL OAuthAPI

DEFINE my_user_id STRING
MAIN
  # ...
   
  # Init OAuthAPI
  IF NOT OAuthAPI.init(5, "AF350CBC-8801-4DFB-9A78-A95B25BB32AF", "8JEq3HBfxrmj/8vMP66iaRQnGrWVyjqr") THEN
    DISPLAY "Error: unable to initialize OAuth"
    EXIT PROGRAM 1
  ELSE
    LET my_user_id = OAuthAPI.getIDSubject()
  END IF

  # ... 
  
END MAIN
```
