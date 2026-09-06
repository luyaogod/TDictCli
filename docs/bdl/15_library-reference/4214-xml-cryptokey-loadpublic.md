---
title: "xml.CryptoKey.loadPublic"
source: "fgl-topics/c_gws_XmlCryptoKey_loadPublic.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadPublic"
type: "concept"
---

# xml.CryptoKey.loadPublic

> Loads the public part of an asymmetric RSA, ECDSA or DSA CryptoKey object, or the parameters and the public key of the Diffie-Hellman object from a XML document.

## Syntax

```
loadPublic(
   doc xml.DomDocument )
```

1. doc defines a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") object.

## Usage

This method populates the RSA, DSA ECDSA CryptoKey object with the public key parts contained in
the given `xml.DomDocument`.

The asymmetric RSA CryptoKey object public
key part is based on the XML-Signature specification for [RSA](http://www.w3.org/TR/xmldsig-core/#sec-RSAKeyValue). The DSA
CryptoKey object public key part is based on the [DSA](http://www.w3.org/TR/xmldsig-core/#sec-DSAKeyValue) specification.

The ECDSA CryptoKey object public key is based on the [XML
Signature Syntax and Processing Version 1.1](https://www.w3.org/TR/xmldsig-core/#sec-ECKeyValue) specification.

For Diffie-Hellman, the input parameter is a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.") object containing a
representation of the Diffie-Hellman key based on the XML-Signature specification for the
Diffie-Hellman key values. This method populates the Diffie-Hellman object with the parameters and
the public key contained in the given `xml.DomDocument`.

If the public key node exists in the `xml.DomDocument` but is empty, it won't be
possible to use the key unless the document contains valid modulus and generator parameters and you
call [`xml.CryptoKey.generateKey`](4200-xml-cryptokey-generatekey.md "Generates a random key of given size (in bits).") with a size of zero (0). In this case, you won't
be in possession of the other peer's public key.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
