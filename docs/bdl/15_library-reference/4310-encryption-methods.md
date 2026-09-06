---
title: "xml.Encryption methods"
source: "fgl-topics/c_gws_XmlEncryption_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods"
type: "concept"
---

# xml.Encryption methods

> Methods for the xml.Encryption class.

| Name | Description |
| --- | --- |
| xml.Encryption.Create() RETURNS xml.Encryption | Constructor of an Encryption object. |

| Name | Description |
| --- | --- |
| xml.Encryption.DecryptString( key xml.CryptoKey , str STRING ) RETURNS STRING | Decrypts an encrypted string encoded in BASE64, using the specified symmetric key, and returns the string in clear text. |
| xml.Encryption.EncryptString( key xml.CryptoKey, str STRING ) RETURNS STRING | Encrypts the specified string using the symmetric key, and returns the encrypted string encoded in BASE64. |
| xml.Encryption.RSADecrypt( filename STRING, str STRING ) RETURNS STRING | Decrypts the BASE64 encrypted string using the RSA key and returns it in clear text |
| xml.Encryption.RSAEncrypt( filename STRING, str STRING ) RETURNS STRING | Encrypts the specified string using the RSA key and returns it encoded in BASE64. |

The methods listed in Table 2 do not
belong to the XML encryption specification, but are helper functions to allow BDL applications to
encrypt and decrypt short passwords with RSA keys, or big strings by using symmetric keys. Notice
that a common way to encrypt data is to use symmetric keys, and to use RSA keys to encrypt the
symmetric key value.

| Name | Description |
| --- | --- |
| getEmbeddedKey() RETURNS xml.CryptoKey | Get a copy of the embedded symmetric key that was used in the last decryption operation. |
| setCertificate( cert xml.CryptoX509 ) | Assigns a copy of the X509 certificate to this encryption object. |
| setKey( key xml.CryptoKey ) | Assigns a copy of the symmetric key to this encryption object. |
| setKeyEncryptionKey( key xml.CryptoKey ) | Assigns a copy of the key-encryption key to this encryption object. |

| Name | Description |
| --- | --- |
| decryptElement( node xml.DomNode ) | Decrypts the encrypted data DomNode using the symmetric key. |
| decryptElementContent( node xml.DomNode ) | Decrypts the encrypted data DomNode using the symmetric key. |
| encryptElement( node xml.DomNode ) | Encrypts the element DomNode and all its children using the symmetric key. |
| encryptElementContent( node xml.DomNode ) | Encrypts all child nodes of the element DomNode using the symmetric key. |

| Name | Description |
| --- | --- |
| decryptElementDetached( node xml.DomNode ) RETURNS xml.DomNode | Decrypts the specified encrypted data DomNode using the symmetric key, and returns it in a new element `xml.DomNode`. |
| decryptElementContentDetached( node xml.DomNode ) RETURNS xml.DomNode | Decrypts the encrypted data DomNode using the symmetric key, and returns all its children in one new document fragment type `xml.DomNode`. |
| encryptElementDetached( node xml.DomNode ) RETURNS xml.DomNode | Encrypts the element DomNode and all its children using the symmetric key, and returns them as one new encrypted-data node. |
| encryptElementContentDetached( node xml.DomNode ) RETURNS xml.DomNode | Encrypts all child nodes of the element DomNode using the symmetric key, and returns them as one new encrypted-data node. |

| Name | Description |
| --- | --- |
| decryptKey( doc xml.DomDocument, url STRING ) RETURNS xml.CryptoKeY | Decrypts the [EncryptedKey](http://www.w3.org/TR/xmlenc-core/#sec-EncryptedKey) as root in the given XML document, and returns a new CryptoKey of the given kind. |
| encryptKey( symkey xml.CryptoKey ) RETURNS xml.DomDocument | Encrypts the given symmetric or HMAC key as an encrypted-key node and returns it as root node of a new XML document. |
