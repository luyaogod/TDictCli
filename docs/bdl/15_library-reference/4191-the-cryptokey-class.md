---
title: "The CryptoKey class"
source: "fgl-topics/c_gws_XmlCryptoKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class"
type: "concept"
---

# The CryptoKey class

> The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.

This class is provided in the `xml` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`xml` package with:

```
IMPORT xml
```

It follows the [XML-Signature](http://www.w3.org/TR/xmldsig-core/) and [XML-Encryption](http://www.w3.org/TR/xmlenc-core/)
specifications.

The `status` variable is set to zero after a successful method call.

## Child topics

- [xml.CryptoKey methods](4192-cryptokey-methods.md): Methods for the xml.CryptoKey class.
- [Supported kind of keys](4222-supported-kind-of-keys.md): Types of keys supported by the xml.CryptoKey class.
- [Derived keys](4223-derived-keys.md)
- [CryptoKey Features](4224-cryptokey-features.md): Features of the xml.CryptoKey class.
- [Examples](4225-examples.md): xml.CryptoKey usage examples.
