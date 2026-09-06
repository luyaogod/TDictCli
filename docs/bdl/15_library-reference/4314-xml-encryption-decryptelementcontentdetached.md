---
title: "xml.Encryption.decryptElementContentDetached"
source: "fgl-topics/c_gws_XmlEncryption_decryptElementContentDetached.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.decryptElementContentDetached"
type: "concept"
---

# xml.Encryption.decryptElementContentDetached

> Decrypts the encrypted data DomNode using the symmetric key, and returns all its children in one new document fragment type xml.DomNode.

## Syntax

```
decryptElementContentDetached(
   node xml.DomNode )
  RETURNS xml.DomNode
```

1. node defines the encrypted [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

## Usage

This method decrypts the [EncryptedData](http://www.w3.org/TR/xmlenc-core/#sec-EncryptedData)
DomNode referenced by node and returns all its children in one new [DOCUMENT\_FRAGMENT\_NODE](4082-domnode-types.md "List of types for the xml.DomNode class.")
[node](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

The resulting child nodes are not added at any place in the XML document. It's up to the user to
insert them in the right place, and to remove the encrypted node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
