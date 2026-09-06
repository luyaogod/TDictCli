---
title: "xml.Encryption.encryptElementDetached"
source: "fgl-topics/c_gws_XmlEncryption_encryptElementDetached.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.encryptElementDetached"
type: "concept"
---

# xml.Encryption.encryptElementDetached

> Encrypts the element DomNode and all its children using the symmetric key, and returns them as one new encrypted-data node.

## Syntax

```
encryptElementDetached(
   node xml.DomNode )
  RETURNS xml.DomNode
```

1. node defines the ELEMENT [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") to encrypt.

## Usage

Encrypts the element DomNode specified in *node* and all its children using the symmetric
key, and returns them as one new [EncryptedData](http://www.w3.org/TR/xmlenc-core/#sec-EncryptedData)
[node](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

The resulting DomNode is not added at any place in the XML document. It's up to the user to
insert it at the right place, and to remove the nodes in clear form.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
