---
title: "xml.Encryption.Create"
source: "fgl-topics/c_gws_XmlEncryption_Create.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.Create"
type: "concept"
---

# xml.Encryption.Create

> Constructor of an Encryption object.

## Syntax

```
xml.Encryption.Create()
  RETURNS xml.Encryption
```

## Usage

Returns an [Encryption](4309-the-encryption-class.md "The xml.Encryption class provides methods to encrypt and decrypt XML documents, nodes or symmetric keys.") object or NULL.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
