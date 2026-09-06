---
title: "xml.Encryption.decryptElementContent"
source: "fgl-topics/c_gws_XmlEncryption_decryptElementContent.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.decryptElementContent"
type: "concept"
---

# xml.Encryption.decryptElementContent

> Decrypts the encrypted data DomNode using the symmetric key.

## Syntax

```
decryptElementContent(
   node xml.DomNode )
```

1. node defines the encrypted [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

## Usage

The [EncryptedData](http://www.w3.org/TR/xmlenc-core/#sec-EncryptedData) DomNode node is replaced at the same place
in the XML document with the resulting child nodes.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
