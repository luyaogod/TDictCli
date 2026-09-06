---
title: "xml.CryptoKey methods"
source: "fgl-topics/c_gws_XmlCryptoKey_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods"
type: "concept"
---

# xml.CryptoKey methods

> Methods for the xml.CryptoKey class.

| Name | Description |
| --- | --- |
| xml.CryptoKey.Create( url STRING ) RETURNS xml.CryptoKey | Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL. |
| xml.CryptoKey.CreateDerivedKey( url STRING ) RETURNS xml.CryptoKey | Constructor of an empty CryptoKey object based on a URL. The crypto key must be derived before use. |
| xml.CryptoKey.CreateFromNode( url STRING, node xml.DomNode ) RETURNS xml.CryptoKey | Constructor of a new CryptoKey object based on a URL, from a XML node based on the XML-Signature and XML-Encryption specification. |

| Name | Description |
| --- | --- |
| compareTo( toCompare xml.CryptoKey ) RETURNS INTEGER | Compares a CryptoKey object to a second key. |
| getEllipticCurveName() RETURNS STRING | Returns the name of the elliptic curve for an ECDSA key, if available. |
| getSHA1() RETURNS STRING | Returns the SHA1 encoded key identifier in a base64 encoded STRING. |
| getSize() RETURNS INTEGER | Returns the size of the key in bits. |
| getType() RETURNS STRING | Returns the type of key. |
| getUsage() RETURNS STRING | Returns the usage of the key. |
| getURL() RETURNS STRING | Returns the key identifier as a URL. |

See also [The Diffie-Hellman key agreement algorithm](../16_web-services/4592-the-diffie-hellman-key-agreement-algorithm.md "Understand how Genero Web Services supports different key encryption methods for shared secret communication using the Diffie-Hellman key-agreement algorithm and the GWS XML security classes."). For more about ECDSA elliptic
curve keys, see [Supported kind of keys](4222-supported-kind-of-keys.md).

| Name | Description |
| --- | --- |
| deriveKey( url STRING, label STRING, seed STRING, created STRING, offset INTEGER, bytes INTEGER ) | Derives the symmetric or HMAC CryptoKey object using the given method identifier and concatenating the optional label, the mandatory seed value and the optional created date as initial random value. |
| generateKey( keySize INTEGER ) | Generates a random key of given size (in bits). |
| generateEllipticCurveKey( curveName STRING ) | Generates a new ECDSA key for the specified named elliptic curve. |
| setKey( key STRING ) | Defines the value of a HMAC or Symmetric key. |

For more about ECDSA elliptic curve keys, see [Supported kind of keys](4222-supported-kind-of-keys.md).

| Name | Description |
| --- | --- |
| computeKey( pub xml.CryptoKey, url STRING ) RETURNS xml.CryptoKey | Computes the shared secret based on the given modulus, generator, the private key, and the other peer's public key. The returned key can be any symmetric/HMAC or symmetric/encryption key type. It can be used for symmetric signature or symmetric encryption. |
| loadBin( filename STRING ) | Loads a symmetric or HMAC key from a file in raw format. |
| loadEllipticCurve( curveName STRING, x STRING, y STRING ) | Loads an ECDSA key using a specific public point. |
| loadDER( filename STRING ) | Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in DER format. |
| loadFromString( str STRING ) | Loads the given key in BASE64 string format into a CryptoKey object. |
| loadPEM( filename STRING ) | Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format. |
| loadPrivate( doc xml.DomDocument ) | Loads the private asymmetric RSA key from the given XML document. |
| loadPublic( doc xml.DomDocument ) | Loads the public part of an asymmetric RSA, ECDSA or DSA CryptoKey object, or the parameters and the public key of the Diffie-Hellman object from a XML document. |
| loadPublicFromString( str STRING ) | Populate the current CryptoKey object with the passed public key. |
| savePrivate() RETURNS xml.DomDocument | Saves the private key part of an asymmetric RSA CryptoKey object into a XML document according to the [XKMS2.0 specification](http://www.w3.org/TR/xkms2/#XKMS_2_0_Section_8_2). |
| savePublic() RETURNS xml.DomDocument | Saves the public part of an asymmetric RSA, DSA or ECDSA CryptoKey object, or the parameters and the public key of the Diffie-Hellman object into a XML document. |
| savePublicToString() RETURNS STRING | Save the current `xml.CryptoKey`'s public part in the returned base64 string. |
| saveToString() RETURNS STRING | Saves the CryptoKey object into a BASE64 string format. |

| Name | Description |
| --- | --- |
| getFeature( feature STRING ) RETURNS STRING | Returns the value of the given feature for this CryptoKey object, or NULL. |
| setFeature( feature STRING, value STRING ) | Sets or resets the value of a feature for a CryptoKey object. |

## Child topics

- [xml.CryptoKey.loadEllipticCurve](4209-xml-cryptokey-loadellipticcurve.md): Loads an ECDSA key using a specific public point.
