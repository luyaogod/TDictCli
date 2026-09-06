---
title: "CryptoKey Features"
source: "fgl-topics/r_gws_XmlCryptoKey_features.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey Features"
type: "reference"
---

# CryptoKey Features

> Features of the xml.CryptoKey class.

| Name | Description |
| --- | --- |
| `KeyName`See [W3C KeyName](http://www.w3.org/TR/xmldsig-core/#sec-KeyName) specification for details. | Defines or returns whether a user-defined key name is added during a XML signature or encryption in order to identify it to other applications, or by the [xml.key store](4335-the-keystore-class.md "The xml.KeyStore class handles a key store for your application.").The default value is NULL, meaning that no key name is used. |
| `KeyValue`See [W3C KeyValue](http://www.w3.org/TR/xmldsig-core/#sec-KeyValue) specification for details. | Defines or returns whether the public part of the asymmetric key is added during a XML signature or encryption.Only for RSA and DSA keys.The default value is FALSE, meaning that no key value is used. |
| `RetrievalMethod`See [W3C RetrievalMethod](http://www.w3.org/TR/xmldsig-core/#sec-RetrievalMethod) specification for details. | Defines or returns the URL where the XML form of:a DSA or RSA public key will be set during a XML signature, and loaded during a XML verification process.a RSA public key will be set and used to encrypt a XML node during XML encryptiona symmetric key with encryption usage will be used to encrypt a XML node or decrypt it backThe default value is NULL, meaning that no retrieval method is used.The XML form of a DSA or RSA public key can be obtained by the [`xml.CryptoKey.savePublic`](4217-xml-cryptokey-savepublic.md "Saves the public part of an asymmetric RSA, DSA or ECDSA CryptoKey object, or the parameters and the public key of the Diffie-Hellman object into a XML document.") method.The XML form of a symmetric key can be obtained by the [`xml.Encryption.encryptKey`](4322-xml-encryption-encryptkey.md "Encrypts the given symmetric or HMAC key as an encrypted-key node and returns it as root node of a new XML document.") method. |
