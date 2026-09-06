---
title: "com.HttpRequest.clearCipher"
source: "fgl-topics/c_gws_ComHTTPRequest_clearCipher.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.clearCipher"
type: "concept"
---

# com.HttpRequest.clearCipher

> Removes the cipher set by setCipher().

## Syntax

```
clearCipher()
```

## Usage

Removes the cipher set by [setCipher()](3875-com-httprequest-setcipher.md "Defines the type of cipher to use for encryption and decryption.").

Once the value is removed, [FGLPROFILE](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
entries for `security.global.cipher` and `ws.idws.cipher` that exist
will be used for secured communication instead.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Examples using com.HttpRequest methods](3887-examples-httprequest.md "These examples use methods of the com.HttpRequest class.")

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
