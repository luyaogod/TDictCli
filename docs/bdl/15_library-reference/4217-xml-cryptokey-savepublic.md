---
title: "xml.CryptoKey.savePublic"
source: "fgl-topics/c_gws_XmlCryptoKey_savePublic.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.savePublic"
type: "concept"
---

# xml.CryptoKey.savePublic

> Saves the public part of an asymmetric RSA, DSA or ECDSA CryptoKey object, or the parameters and the public key of the Diffie-Hellman object into a XML document.

## Syntax

```
savePublic()
  RETURNS xml.DomDocument
```

## Usage

This method saves the public key parts of an RSA, DSA or ECDSA CryptoKey object in an
`xml.DomDocument`.

The asymmetric RSA CryptoKey object public
key part is based on the XML-Signature specification for [RSA](http://www.w3.org/TR/xmldsig-core/#sec-RSAKeyValue). The DSA
CryptoKey object public key part is based on the [DSA](http://www.w3.org/TR/xmldsig-core/#sec-DSAKeyValue) specification.

For Diffie-Hellman, the method is used for the public key exchanged between the two peers.

See also the CryptoKey [RetrievalMethod](4224-cryptokey-features.md "Features of the xml.CryptoKey class.") feature.

Returns a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")
object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
