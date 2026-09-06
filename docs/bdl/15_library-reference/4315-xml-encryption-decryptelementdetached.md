---
title: "xml.Encryption.decryptElementDetached"
source: "fgl-topics/c_gws_XmlEncryption_decryptElementDetached.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.decryptElementDetached"
type: "concept"
---

# xml.Encryption.decryptElementDetached

> Decrypts the specified encrypted data DomNode using the symmetric key, and returns it in a new element xml.DomNode.

## Syntax

```
decryptElementDetached(
   node xml.DomNode )
  RETURNS xml.DomNode
```

1. node defines the encrypted [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

## Usage

This method decrypts the [EncryptedData](http://www.w3.org/TR/xmlenc-core/#sec-EncryptedData)
DomNode node. The resulting element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") and its children are not added at any
place in the XML document. It's up to the user to insert it at the right place, and to remove the
encrypted node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
