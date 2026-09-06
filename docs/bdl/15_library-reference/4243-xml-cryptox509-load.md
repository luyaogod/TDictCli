---
title: "xml.CryptoX509.load"
source: "fgl-topics/c_gws_XmlCryptoX509_load.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.load"
type: "concept"
---

# xml.CryptoX509.load

> Loads the given XML document with ds:X509Data as root node in a CryptoX509 object.

## Syntax

```
load(
   doc xml.DomDocument )
```

1. doc defines a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")
   object.

## Usage

This method loads an XML document with ds:X509Data based on the XML-Signature specification as
root node in a CryptoX509 object. See the [w3.org site](http://www.w3.org/TR/xmldsig-core/#sec-X509Data) for more
information on the XML-Signature specification for ds:X509Data as root node.

If the X509 certificate is incomplete, the certificate will be
created from the application global certificate list if one of `SubjectName` or
`Issuer` matches. (See [xml.KeyStore.addCertificate](4336-keystore-methods.md "Methods for the xml.KeyStore class.")
for more details.)

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
