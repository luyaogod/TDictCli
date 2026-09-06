---
title: "xml.CryptoKey.loadDER"
source: "fgl-topics/c_gws_XmlCryptoKey_loadDER.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadDER"
type: "concept"
---

# xml.CryptoKey.loadDER

> Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in DER format.

## Syntax

```
loadDER(
   filename STRING )
```

1. filename defines the filename or an [entry](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") in the FGLPROFILE file.

## Usage

If the DSA, RSA or ECDSA private key or Diffie-Hellman parameters is protected with a password,
the recommended way is to unprotect it with the openssl tool and to put the key
file on a restricted file system. However, you can use a [script](../16_web-services/4916-fglprofile-entries-for-web-services.md) or the
fglpass
[agent](../16_web-services/4916-fglprofile-entries-for-web-services.md) to provide
the password to the application.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
