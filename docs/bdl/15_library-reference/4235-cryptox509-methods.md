---
title: "xml.CryptoX509 methods"
source: "fgl-topics/c_gws_XmlCryptoX509_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods"
type: "concept"
---

# xml.CryptoX509 methods

> Methods for the xml.CryptoX509 class.

| Name | Description |
| --- | --- |
| xml.CryptoX509.Create() RETURNS xml.CryptoX509 | Constructor of an empty CryptoX509 object. |
| xml.CryptoX509.CreateFromNode( node xml.DomNode ) RETURNS xml.CryptoX509 | Constructor of a new CryptoX509 object from a XML X509 certificate node. |

| Name | Description |
| --- | --- |
| getCount() RETURNS INTEGER | Returns the number of X509 certificates contained in one CryptoX509 object. |
| getIdentifier() RETURNS STRING | Gets the identification part of a X509 certificate |
| getThumbprintSHA1() RETURNS STRING | Gets the SHA1 encoded thumbprint identifying the X509 certificate. |

| Name | Description |
| --- | --- |
| createPublicKey( url STRING ) RETURNS xml.CryptoKey | Creates a new public CryptoKey object for the given URL. |

| Name | Description |
| --- | --- |
| load( doc xml.DomDocument ) | Loads the given XML document with ds:X509Data as root node in a CryptoX509 object. |
| loadDER( filename STRING ) | Loads a X509 certificate from a file in DER format. |
| loadFromString( str STRING ) | Loads the given X509 certificate in BASE64 string format into this CryptoX509 object. |
| loadPEM( filename STRING ) | Loads a X509 certificate from a file in PEM format. |
| save() RETURNS xml.DomDocument | Saves the CryptoX509 certificate into a XML document with ds:X509Data element as root node. |
| saveToString() RETURNS STRING | Saves the CryptoX509 certificate into a BASE64 string format. |

| Name | Description |
| --- | --- |
| getFeature( feature STRING ) RETURNS STRING | Get the value of a given feature of a CryptoX509 object. |
| setFeature( feature STRING, value STRING ) | Sets or resets the given feature for this CryptoX509 object. |

## Child topics

- [xml.CryptoX509.getCount](4239-xml-cryptox509-getcount.md): Returns the number of X509 certificates contained in one CryptoX509 object.
