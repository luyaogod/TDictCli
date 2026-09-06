---
title: "com.HttpRequest.setVerifyServer"
source: "fgl-topics/c_gws_ComHTTPRequest_setVerifyServer.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setVerifyServer"
type: "concept"
---

# com.HttpRequest.setVerifyServer

> Defines if certificates for applications or services are validated on each request.

## Syntax

```
setVerifyServer(
   val BOOLEAN )
```

1. val specifies TRUE or FALSE.

## Usage

The `setVerifyServer()` method allows you to specify if certificates for
applications or services run by the server are validated or not at each request.

For example, if set to `FALSE`, once the server has been validated against the
local certificate authority, no additional request is performed to validate certificates for
applications or services run by the server. Default value is `TRUE` (certificate
validation is done for all requests for HTTPS applications or services).
> **Warning:**
>
> **Using unsafe SSL connections**
>
> Setting `setVerifyServer(FALSE)` has security
> implications. It can be used in development but should not be used in production.

When you
use `setVerifyServer()`, you automatically override [FGLPROFILE](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.") entries for
`security.global.verifyserver` and `ws.idws.verifyserver`. To return
to using settings in the FGLPROFILE entries, you must use the corresponding [clearVerifyServer()](3856-com-httprequest-clearverifyserver.md "Removes the value set by setVerifyServer.") method.

## Example

```
IMPORT com

MAIN
    DEFINE req com.HttpRequest
    DEFINE resp com.HttpResponse

    LET req = com.HttpRequest.Create("https://myserver/")
    CALL req.setMethod("GET")
    CALL req.setCertificateAndKey("client.crt","client.key")
    CALL req.setCipher("AES128-SHA256")
    CALL req.setVerifyServer(FALSE)
    TRY
        CALL req.doRequest()
        LET resp = req.getResponse()
        DISPLAY resp.getStatusCode()
        DISPLAY resp.getTextResponse()
    CATCH
        DISPLAY "ERROR :",status||" ("||sqlca.sqlerrm||")"
        EXIT PROGRAM -1
    END TRY
END MAIN
```

FGLPROFILE must have the Certificate Authority setting:

```
security.global.ca ="ca.crt"
```

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
