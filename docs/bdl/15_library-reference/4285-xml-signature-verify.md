---
title: "xml.Signature.verify"
source: "fgl-topics/c_gws_XmlSignature_verify.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.verify"
type: "concept"
---

# xml.Signature.verify

> Verifies that all references in this signature object have not changed.

## Syntax

```
verify(
   doc xml.DomDocument )
  RETURNS INTEGER
```

1. doc defines the [XML document](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards.").

## Usage

Returns TRUE if valid, FALSE otherwise.

If the signature type is:

- Enveloping: then *doc* must be NULL because all document fragment references are inside the
  Signature itself
- Enveloped: then *doc* must be the XML document where the signature was enveloped
- Detached: then *doc* can be null if all references are absolute, otherwise it can be the
  XML document the fragment references are referencing

See [XML Signature
concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.") for more details.

By default, the validation process uses the CryptoKey set with `xml.Signature.setKey()` to
verify the signature. However, if the signature contains a X509 certificate or a X509 retrieval
method, it uses the list of trusted certificates, or if the signature contains a RSA or DSA
retrieval method, it uses the RSA or DSA public key automatically loaded.

> **Note:**
>
> See Windows® .NET special [recommendation](4278-xml-signature-setcanonicalization.md "Sets the canonicalization method to use for the signature.").

Before loading the XML document to verify the signature, you might need to set some options to
retrieve the "id" nodes with the [`xml.DomDocument.setFeature()`](4006-xml-domdocument-setfeature.md "Sets a feature for this xml.DomDocument object.") method:

```
DEFINE doc xml.DomDocument
...
CALL doc.setFeature(feature, TRUE)
...
```

Here feature must be `"auto-id-attribute"` if the
`"id"` attribute has no namespace, or `"auto-id-qualified-attribute"`,
when `"id"` has a namespace.

This is especially needed when you encounter error messages such
as:

```
Xml security operation failed : libxml2 library function failed : expr=xpointer(id('id-1436767651')).
```

Meaning that the parser was unable to find the `"id"` attribute in the XML
document.

Note that the `"auto-id-*"` features will declare all XML attributes where the
name is `"id"`, `"ID"`, `"Id"` or `"iD"`
to be of type `ID`, and thus be usable via [`xml.DomDocument.getElementById()`](3982-xml-domdocument-getelementbyid.md "Returns the xml.DomNode element that has an attribute of type ID with the given value.") method used during signature
validation.

If needed, you can also set features for a specific attribute with the [`xml.DomNode.setIdAttribute()`](4077-xml-domnode-setidattribute.md "Set the XML Attribute of given name to be of type ID. Declare (or undeclare) the ID as user-determined.")
method, or with the [`xml.DomNode.setIdAttributeNS()`](4078-xml-domnode-setidattributens.md "Set the namespace-qualified XML Attribute of given name and namespace to be of type ID. Declare (or undeclare) the ID as user-determined.") method.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
