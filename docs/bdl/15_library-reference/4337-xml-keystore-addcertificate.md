---
title: "xml.KeyStore.AddCertificate"
source: "fgl-topics/c_gws_XmlKeyStore_AddCertificate.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The KeyStore class > KeyStore methods > xml.KeyStore.AddCertificate"
type: "concept"
---

# xml.KeyStore.AddCertificate

> Registers in the keystore the given X509 certificate as a certificate for the application.

## Syntax

```
xml.KeyStore.AddCertificate( 
   global xml.CryptoX509 )
```

1. global defines the X509 certificate to register.

## Usage

This method registers a given X509 certificate to be used when an incomplete X509 certificate is
detected during signature or encryption. To complete the process the certificate's issuer name and
serial number are checked.

The method has the same effect as the FGLPROFILE entry `xml.keystore.x509list`.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
