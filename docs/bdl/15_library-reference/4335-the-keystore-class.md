---
title: "The KeyStore class"
source: "fgl-topics/c_gws_XmlKeyStore.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The KeyStore class"
type: "concept"
---

# The KeyStore class

> The xml.KeyStore class handles a key store for your application.

This class is provided in the `xml` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`xml` package with:

```
IMPORT xml
```

The `xml.KeyStore` class provides static methods to handle a key store that is
global for the entire application.

It enables you to register X509 and trusted certificates, and any kind of key by name for
automatic XML signature validation or XML decryption.

The `status` variable is set to zero after a successful method call.

## Child topics

- [xml.KeyStore methods](4336-keystore-methods.md): Methods for the xml.KeyStore class.
