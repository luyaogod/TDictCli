---
title: "The CryptoX509 class"
source: "fgl-topics/c_gws_XmlCryptoX509.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class"
type: "concept"
---

# The CryptoX509 class

> The xml.CryptoX509 class provides methods to manipulate X509 certificates needed for identification of individual persons, groups or any entities during XML encryption or signature process.

This class is provided in the `xml` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`xml` package with:

```
IMPORT xml
```

It also provides additional load and save functions to interact with other applications
in XML or in BASE64, such as in WS-Security compliant applications. It follows the [XML-Signature](http://www.w3.org/TR/xmldsig-core/)
and [XML-Encryption](http://www.w3.org/TR/xmlenc-core/)
specifications.

The `status` variable is set to zero after a successful method call.

## Child topics

- [xml.CryptoX509 methods](4235-cryptox509-methods.md): Methods for the xml.CryptoX509 class.
- [CryptoX509 Features](4250-cryptox509-features.md): Features of the xml.CryptoX509 class.
- [Examples](4251-examples.md): xml.CryptoX509 usage examples.
