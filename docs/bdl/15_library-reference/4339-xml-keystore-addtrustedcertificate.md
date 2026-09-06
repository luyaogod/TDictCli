---
title: "xml.KeyStore.AddTrustedCertificate"
source: "fgl-topics/c_gws_XmlKeyStore_AddTrustedCertificate.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The KeyStore class > KeyStore methods > xml.KeyStore.AddTrustedCertificate"
type: "concept"
---

# xml.KeyStore.AddTrustedCertificate

> Registers in the keystore the given X509 certificate as a trusted certificate for the application.

## Syntax

```
xml.KeyStore.AddTrustedCertificate(
   global xml.CryptoX509 )
```

1. global defines the X509 certificate to register.

## Usage

This method registers the given [X509
certificate](4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.") as a trusted certificate for the application. It will be used for signature
verification if no other certificate was set for that purpose.

The method has the same effect as the FGLPROFILE entry `xml.keystore.calist`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
