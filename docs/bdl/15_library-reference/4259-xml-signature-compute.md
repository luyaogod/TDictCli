---
title: "xml.Signature.compute"
source: "fgl-topics/c_gws_XmlSignature_compute.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.compute"
type: "concept"
---

# xml.Signature.compute

> Computes the signature of all references set in this Signature object.

## Syntax

```
compute(
   doc xml.DomDocument )
```

1. doc defines the [XML document](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.").

## Usage

If the signature type is:

- Enveloping: then *doc* must be NULL because all document fragment references are inside the
  Signature itself
- Enveloped: then *doc* must be the XML document where the signature must be added afterwards
  to get a valid enveloped signature
- Detached: then *doc* can be NULL if all references are absolute, otherwise it can be the
  XML document fragment references that are referencing

See [XML Signature
concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.") for more details.

Also, see Windows® .NET special [recommendation](4278-xml-signature-setcanonicalization.md "Sets the canonicalization method to use for the signature.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
