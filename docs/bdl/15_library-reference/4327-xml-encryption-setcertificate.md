---
title: "xml.Encryption.setCertificate"
source: "fgl-topics/c_gws_XmlEncryption_setCertificate.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.setCertificate"
type: "concept"
---

# xml.Encryption.setCertificate

> Assigns a copy of the X509 certificate to this encryption object.

## Syntax

```
setCertificate(
   cert xml.CryptoX509 )
```

1. cert defines the copy of the [X509](4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.") certificate.

## Usage

The certificate will then be added to any further XML document or node encryption.

- NULL is allowed to avoid the certificate being added.
- To encrypt using a certificate, you must use the [`createPublicKey`](4238-xml-cryptox509-createpublickey.md "Creates a new public CryptoKey object for the given URL.") method of
  the X509 class to obtain the public key embedded in the certificate, and then provide it to the
  encryption object with [`setKeyEncryptionKey`](4329-xml-encryption-setkeyencryptionkey.md "Assigns a copy of the key-encryption key to this encryption object.") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
