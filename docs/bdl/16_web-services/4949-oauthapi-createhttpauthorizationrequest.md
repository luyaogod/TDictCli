---
title: "OAuthAPI.CreateHTTPAuthorizationRequest()"
source: "fgl-topics/c_gws_oauthapi_CreateHTTPAuthorizationRequest.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.CreateHTTPAuthorizationRequest()"
type: "concept"
---

# OAuthAPI.CreateHTTPAuthorizationRequest()

> Create an HttpRequest with OAuth access token.

## Syntax

```
FUNCTION CreateHTTPAuthorizationRequest(
   url STRING )
RETURNS com.HttpRequest
```

1. url is a `STRING` with the URL to connect to.

Returns a `com.HttpRequest` object. `NULL` may be
returned if the access token can not be renewed.

## Usage

Use this function to create an `HttpRequest` object with OAuth access token for
access to a secure RESTful web service.

In case of error, a `NULL` value will be returned.

## OAuthAPI.CreateHTTPAuthorizationRequest function

```
IMPORT com
IMPORT FGL OAuthAPI

MAIN

DEFINE req com.HttpRequest
DEFINE resp com.HttpResponse
DEFINE url STRING

TRY
    # call the init function first

    # get request path
    LET url = fgl_getenv("OIDC_USERINFO_ENDPOINT")
 
    # Create oauth request
    LET req = OAuthAPI.CreateHTTPAuthorizationRequest(url)

    # Perform request
    CALL req.setMethod("GET")
    CALL req.setHeader("Accept", "application/json")
    CALL req.DoRequest()
        
    # Retrieve response
    LET resp = req.getResponse()

    # Process response
    CASE resp.getStatusCode()
        WHEN 200 #Success
        # ...
        OTHERWISE
            DISPLAY " Error code is: ",  resp.getStatusCode()
            EXIT PROGRAM 1
    END CASE
CATCH
    DISPLAY "ERROR : ",status,sqlca.sqlerrm
    EXIT PROGRAM 1
END TRY

END MAIN
```
