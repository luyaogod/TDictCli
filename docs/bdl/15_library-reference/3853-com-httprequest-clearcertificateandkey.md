---
title: "com.HttpRequest.clearCertificateAndKey"
source: "fgl-topics/c_gws_ComHTTPRequest_clearCertificateAndKey.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.clearCertificateAndKey"
type: "concept"
---

# com.HttpRequest.clearCertificateAndKey

> Removes the client certificate and key set by setCertificateAndKey().

## Syntax

```
clearCertificateAndKey()
```

## Usage

Removes the certificate and key set by [setCertificateAndKey()](3873-com-httprequest-setcertificateandkey.md "Specifies the certificate and key to use for the HttpRequest request.").

Once the certificate and key are removed, [FGLPROFILE](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.") entries for `security.global.certificate`,
`security.idsec.certificate`, `security.global.privatekey`, and
`security.idsec.privatekey` that exist will be used for authentication instead.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
