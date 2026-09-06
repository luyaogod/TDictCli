---
title: "OAuth access without GAS"
source: "fgl-topics/c_gws_oauthapi_helper_example.html"
breadcrumb: "Web services > Reference > OAuthAPI library > OAuthAPI: library > Examples > Get OAuth access without GAS"
type: "concept"
---

# OAuth access without GAS

> This example illustrates how to obtain OAuth access to a web service when the application is not operating behind a Genero Application Server (GAS), specifically for applications that include mobile, desktop, Genero Web Applications, and Text User Interface (TUI) apps.

## OAuthAPI helper functions

```
IMPORT com
IMPORT FGL OAuthAPI
IMPORT util

PUBLIC TYPE QueryUserInfoResponseBodyType DYNAMIC ARRAY OF RECORD
    user_id INTEGER,
    given_name STRING,
    family_name STRING,
    email STRING
END RECORD

# Error codes
PUBLIC CONSTANT C_SUCCESS = 0

PUBLIC FUNCTION QueryUserInfo(idp_url STRING, user_login STRING, user_password STRING) RETURNS(INTEGER, QueryUserInfoResponseBodyType)
    DEFINE fullpath base.StringBuffer
    DEFINE contentType STRING
    DEFINE req com.HttpRequest
    DEFINE resp com.HttpResponse
    DEFINE resp_body QueryUserInfoResponseBodyType
    DEFINE json_body STRING
    DEFINE expire INTEGER
    DEFINE meta OAuthAPI.OpenIDMetadataType
    DEFINE access_token  STRING
    DEFINE expires_in    INTEGER

    CALL OAuthAPI.FetchOpenIDMetadata(5, idp_url) RETURNING meta.*
    CALL OAuthAPI.RetrievePasswordToken(5, meta.token_endpoint, user_login, user_password, "profile") RETURNING access_token, expire 
    IF NOT OAuthAPI.initService(5, access_token) THEN
      DISPLAY "Error: unable to initiate OAuthAPI"
      EXIT PROGRAM 1
    END IF

    TRY

        # Prepare request path
        LET fullpath = base.StringBuffer.create()
        CALL fullpath.append("/userinfo/v1/query")

        WHILE TRUE
            # Create oauth request and configure it
            LET req =
                OAuthAPI.CreateHTTPAuthorizationRequest(
                    SFMT("%1%2", meta.userinfo_endpoint, fullpath.toString()))

            # Perform request
            CALL req.setMethod("GET")
            CALL req.setHeader("Accept", "application/json")
            CALL req.doRequest()

            # Retrieve response
            LET resp = req.getResponse()
            # Retry if access token has expired
            IF NOT OAuthAPI.RetryHTTPRequest(resp) THEN
                EXIT WHILE
            END IF
        END WHILE
        # Process response
        INITIALIZE resp_body TO NULL
        LET contentType = resp.getHeader("Content-Type")
        CASE resp.getStatusCode()

            WHEN 200 #Success
                IF contentType MATCHES "*application/json*" THEN
                    # Parse JSON response
                    LET json_body = resp.getTextResponse()
                    CALL util.JSON.parse(json_body, resp_body)
                    RETURN C_SUCCESS, resp_body
                END IF
                RETURN -1, resp_body

            OTHERWISE
                RETURN resp.getStatusCode(), resp_body
        END CASE
    CATCH
        RETURN -1, resp_body
    END TRY
END FUNCTION
```
