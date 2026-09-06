---
title: "OAuthAPI.RetryHTTPRequest()"
source: "fgl-topics/c_gws_oauthapi_RetryHTTPRequest.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.RetryHTTPRequest()"
type: "concept"
---

# OAuthAPI.RetryHTTPRequest()

> Retry an HttpRequest with OAuth access token to check if the access token has expired.

## Syntax

```
FUNCTION RetryHTTPRequest(
   resp com.HttpResponse )
  RETURNS BOOLEAN
```

1. resp is a `com.HttpResponse` object.

Returns `TRUE` if the access token has expired and the request must be sent again,
`FALSE` otherwise.

## Usage

Use this function to check whether the access token to a secure RESTful Web service has expired.
If expired, it will perform the refresh and then do the request.

In case of error, a `NULL` value will be returned.

## OAuthAPI.RetryHTTPRequest function

```
IMPORT com
IMPORT FGL OAuthAPI

DEFINE req com.HttpRequest 
DEFINE resp com.HttpResponse
DEFINE url STRING

MAIN

   TRY

      # Get path
      LET url = fgl_getenv("OIDC_USERINFO_ENDPOINT")

      WHILE TRUE
         # Create oauth request
         LET req = OAuthAPI.CreateHTTPAuthorizationRequest(url)

         # Perform request
         CALL req.setMethod("GET")
         CALL req.setHeader("Accept", "application/json")
         CALL req.DoRequest()
         
         # Retrieve response
         LET resp = req.getResponse()
         # Retry if access token has expired
         IF NOT OAuthAPI.RetryHTTPRequest(resp) THEN
            EXIT WHILE
         END IF   
      END WHILE

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

## Related links

**Related concepts**  

[OAuth access without GAS](4957-get-oauth-access-without-gas.md "This example illustrates how to obtain OAuth access to a web service when the application is not operating behind a Genero Application Server (GAS), specifically for applications that include mobile, desktop, Genero Web Applications, and Text User Interface (TUI) apps.")
