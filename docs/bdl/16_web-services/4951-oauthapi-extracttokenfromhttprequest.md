---
title: "OAuthAPI.ExtractTokenFromHTTPRequest()"
source: "fgl-topics/c_gws_oauthapi_ExtractTokenFromHTTPRequest.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > OAuthAPI.ExtractTokenFromHTTPRequest()"
type: "concept"
---

# OAuthAPI.ExtractTokenFromHTTPRequest()

> Return the OAuth access token from a HTTP request service object.

## Syntax

```
FUNCTION ExtractTokenFromHTTPRequest(
   req com.HttpServiceRequest)
RETURNS STRING
```

1. req is a `com.HttpServiceRequest` object.

Returns the access token. `NULL` may be returned if the access token is not
found.

## Usage

Use the `ExtractTokenFromHTTPRequest()` function to extract the access token from
an `com.HttpServiceRequest` object accessing a RESTful Web service.

In case of error, a `NULL` value will be returned.

## OAuthAPI.ExtractTokenFromHTTPRequest function

```
IMPORT com
IMPORT FGL OAuthAPI

DEFINE req com.HttpServiceRequest 
DEFINE access_token STRING

MAIN

   TRY
      # Start server for all registered Web Services
      CALL com.WebServiceEngine.Start()
      
      WHILE TRUE
         LET req = com.WebServiceEngine.GetHttpServiceRequest(-1)
         IF req IS NULL THEN
            DISPLAY "HTTP request timeout...: ", CURRENT YEAR TO FRACTION
            EXIT PROGRAM 1
         ELSE 
            # Retrieve access token
            LET access_token = OAuthAPI.ExtractTokenFromHTTPRequest(req)
            IF access_token IS NOT NULL THEN
               DISPLAY "token is: ", access_token
            END IF 
         END IF
      END WHILE
   CATCH
      DISPLAY "ERROR : ",status,sqlca.sqlerrm
      EXIT PROGRAM 1
   END TRY

END MAIN
```

## Related links

**Related concepts**  

[OAuth access without GAS](4957-get-oauth-access-without-gas.md "This example illustrates how to obtain OAuth access to a web service when the application is not operating behind a Genero Application Server (GAS), specifically for applications that include mobile, desktop, Genero Web Applications, and Text User Interface (TUI) apps.")
