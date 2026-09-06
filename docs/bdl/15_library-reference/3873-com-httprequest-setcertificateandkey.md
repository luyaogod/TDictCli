---
title: "com.HttpRequest.setCertificateAndKey"
source: "fgl-topics/c_gws_ComHTTPRequest_setCertificateAndKey.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setCertificateAndKey"
type: "concept"
---

# com.HttpRequest.setCertificateAndKey

> Specifies the certificate and key to use for the HttpRequest request.

## Syntax

```
setCertificateAndKey(
   certificate STRING,
   privateKey STRING)
```

1. certificate defines the absolute path of the certificate file
   (.crt).
2. privateKey defines the absolute path of the private key file
   (.pem).

## Usage

The `setCertificateAndKey()` method allows you to define the certificate and key
to use for authentication for the current `HttpRequest`.

For example, you can set the client certificate to use for an HTTPS connection at runtime instead
of setting entries in the FGLPROFILE file before running the application.

When you use `setCertificateAndKey()`, you automatically override [FGLPROFILE](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.") entries for
`security.global.certificate` and `security.idsec.certificate` and the
`security.global.privatekey` and `security.idsec.privatekey`. To
return to using settings in the FGLPROFILE entries, you must use the corresponding [clearCertificateAndKey()](3853-com-httprequest-clearcertificateandkey.md "Removes the client certificate and key set by setCertificateAndKey().") method.

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
