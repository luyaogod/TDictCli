---
title: "xml.Encryption.encryptElement"
source: "fgl-topics/c_gws_XmlEncryption_encryptElement.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.encryptElement"
type: "concept"
---

# xml.Encryption.encryptElement

> Encrypts the element DomNode and all its children using the symmetric key.

## Syntax

```
encryptElement(
   node xml.DomNode )
```

1. node defines the ELEMENT
   [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") to
   encrypt.

## Usage

The element [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") specified
in *node* and all its children are replaced at the same place in the XML document with the
resulting [EncryptedData](http://www.w3.org/TR/xmlenc-core/#sec-EncryptedData) node.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
