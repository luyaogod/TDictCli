---
title: "com.HttpRequest.setCipher"
source: "fgl-topics/c_gws_ComHTTPRequest_setCipher.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > HttpRequest methods > com.HttpRequest.setCipher"
type: "concept"
---

# com.HttpRequest.setCipher

> Defines the type of cipher to use for encryption and decryption.

## Syntax

```
setCipher(
   cipher STRING )
```

1. cipher defines the algorithm to use during secured communication.

## Usage

The `setCipher()` allows you to define the type of cipher to use for encryption
and decryption. For more details about using ciphers, refer to [www.openssl.org](http://www.openssl.org) (external link) and search for
"cipher".

For example, you can set the cipher to use for an HTTPS communication at runtime instead of
setting entries in the FGLPROFILE file before running the application.

When you use `setCipher()`, you automatically override [FGLPROFILE](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.") cipher entries for
`security.global.cipher` and `ws.idws.cipher`. To return to using
settings in the FGLPROFILE entries, you must use the corresponding [clearCipher()](3854-com-httprequest-clearcipher.md "Removes the cipher set by setCipher().") method.

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
