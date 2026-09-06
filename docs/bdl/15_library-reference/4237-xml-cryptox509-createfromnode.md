---
title: "xml.CryptoX509.CreateFromNode"
source: "fgl-topics/c_gws_XmlCryptoX509_CreateFromNode.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.CreateFromNode"
type: "concept"
---

# xml.CryptoX509.CreateFromNode

> Constructor of a new CryptoX509 object from a XML X509 certificate node.

## Syntax

```
xml.CryptoX509.CreateFromNode(
   node xml.DomNode )
  RETURNS xml.CryptoX509
```

1. node defines an element in [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") node with [X509Data](http://www.w3.org/TR/xmldsig-core/#sec-X509Data) as local
   name, which is based on the XML-Signature specification namespace
   "http://www.w3.org/2000/09/xmldsig#".

## Usage

Returns a [xml.CryptoX509](4234-the-cryptox509-class.md "The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.") object or NULL.

If the X509 certificate is incomplete, the certificate will be
created from the application global certificate list if one of `SubjectName` or
`Issuer` matches. (See [xml.KeyStore.addCertificate](4336-keystore-methods.md "Methods for the xml.KeyStore class.")
for more details.)

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
