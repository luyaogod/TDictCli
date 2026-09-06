---
title: "CryptoX509 Features"
source: "fgl-topics/r_gws_XmlCryptoX509_features.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 Features"
type: "reference"
---

# CryptoX509 Features

> Features of the xml.CryptoX509 class.

| Feature | Description |
| --- | --- |
| `X509Certificate`See [specification](http://www.w3.org/TR/xmldsig-core/#sec-X509Data) for details. | Defines or returns whether the complete X509 certificate is added during XML signature or encryption.Default value is FALSE. |
| `X509SubjectName`See [specification](http://www.w3.org/TR/xmldsig-core/#sec-X509Data) for details. | Defines or returns whether the subject name of the X509 certificate is added during XML signature or encryption.Default value is FALSE. |
| `X509IssuerSerial`See [specification](http://www.w3.org/TR/xmldsig-core/#sec-X509Data) for details. | Defines or returns whether the issuer name and serial number of the X509 certificate is added during XML signature or encryption.Default value is FALSE. |
| `RetrievalMethod`See [specification](http://www.w3.org/TR/xmldsig-core/#sec-X509Data) for details. | Defines or returns the URL where the XML form of the X509 certificate will be set during a XML signature, and loaded during a XML verification process, and based on that CryptoX509 object.Default value is NULL, meaning that no retrieval method is used.**Note:**The XML form of a X509 certificate can be obtain by the [`xml.CryptoX509.save()`](4247-xml-cryptox509-save.md "Saves the CryptoX509 certificate into a XML document with ds:X509Data element as root node.") method. |
