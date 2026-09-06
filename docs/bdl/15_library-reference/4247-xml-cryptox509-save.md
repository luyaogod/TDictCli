---
title: "xml.CryptoX509.save"
source: "fgl-topics/c_gws_XmlCryptoX509_save.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.save"
type: "concept"
---

# xml.CryptoX509.save

> Saves the CryptoX509 certificate into a XML document with ds:X509Data element as root node.

## Syntax

```
save()
  RETURNS xml.DomDocument
```

## Usage

This method saves a CryptoX509 certificate in an XML document with the ds:X509Data element as
root node. See the [w3.org site](http://www.w3.org/TR/xmldsig-core/#sec-X509Data) for more
information on the XML-Signature specification for ds:X509Data as root node.

(See also the [RetrievalMethod](4250-cryptox509-features.md "Features of the xml.CryptoX509 class.") feature)

Returns a [xml.DomDocument](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")
object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
