---
title: "com.HttpRequest.clearVerifyServer"
source: "fgl-topics/c_gws_ComHTTPRequest_clearVerifyServer.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.clearVerifyServer"
type: "concept"
---

# com.HttpRequest.clearVerifyServer

> Removes the value set by setVerifyServer.

## Syntax

```
clearVerifyServer()
```

## Usage

Removes the value set by [setVerifyServer](3885-com-httprequest-setverifyserver.md "Defines if certificates for applications or services are validated on each request.") for the validation, or not, of certificates at each request for application
or service.

Once the value is removed, [FGLPROFILE](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
entries for `security.global.verifyserver` and `ws.idws.verifyserver`
that exist will be used for secured communication instead.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
