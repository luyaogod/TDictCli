---
title: "xml.CryptoKey.CreateFromNode"
source: "fgl-topics/c_gws_XmlCryptoKey_CreateFromNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.CreateFromNode"
type: "concept"
---

# xml.CryptoKey.CreateFromNode

> Constructor of a new CryptoKey object based on a URL, from a XML node based on the XML-Signature and XML-Encryption specification.

## Syntax

```
xml.CryptoKey.CreateFromNode(
   url STRING,
   node xml.DomNode )
  RETURNS xml.CryptoKey
```

1. [url](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.") defines a key
   identifier restricted to PUBLIC/PRIVATE keys.
2. node defines an [`ELEMENT` node](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") whose
   local name is either:
   - [DSAKeyValue](http://www.w3.org/TR/xmldsig-core/#sec-DSAKeyValue) or [RSAKeyValue](http://www.w3.org/TR/xmldsig-core/#sec-RSAKeyValue)
     belonging to the XML-Signature namespace http://www.w3.org/2000/09/xmldsig#
   - [RSAKeyPair](http://www.w3.org/TR/xkms2/#XKMS_2_0_Section_8_2) belonging to the XKMS 2.0 namespace
     http://www.w3.org/2002/03/xkms#

## Usage

Returns a [CryptoKey](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") object or
NULL.

If the local name is RSAKeyValue or RSAKeyPair, the URL must be a RSA key. If the local name is
DSAKeyValue or ECDSAKeyValue, the URL must be a DSA key.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
