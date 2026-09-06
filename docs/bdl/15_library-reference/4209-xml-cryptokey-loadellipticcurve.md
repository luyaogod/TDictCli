---
title: "xml.CryptoKey.loadEllipticCurve"
source: "fgl-topics/c_gws_XmlCryptoKey_loadEllipticCurve.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadEllipticCurve"
type: "concept"
---

# xml.CryptoKey.loadEllipticCurve

> Loads an ECDSA key using a specific public point.

## Syntax

```
loadEllipticCurve(
   curveName STRING,
   x STRING,
   y STRING )
```

1. curveName: Name of the elliptic curve, as described in [ECDSA signature algorithms](4222-supported-kind-of-keys.md) (for example, secp256r1, secp256k1).
2. x: X coordinate of the public point, base64 encoded.
3. y: Y coordinate of the public point, base64 encoded.

## Usage

Creates an ECDSA `xml.CryptoKey` from `curveName`,
`x`, and `y` so a JSON Web Key (JWK) entry in a JSON Web Key Set
(JWKS) can be converted to the corresponding `xml.CryptoKey` for token verification
and interoperability.

Use this method when Genero Identity Provider receives a JSON-formatted ECDSA public key and
needs to construct the corresponding public key for JWT verification.

The method raises an exception if the curve name or coordinates are invalid.

For an example of use, go to [Create an ECDSA public key from coordinates](4233-create-an-ecdsa-public-key-from-coordinates.md).

## Related links

**Related concepts**  

[Create an ECDSA public key from coordinates](4233-create-an-ecdsa-public-key-from-coordinates.md "Create an ECDSA public key from coordinates")

**Related information**  

[ECDSA signature algorithms](4222-supported-kind-of-keys.md "ECDSA signature algorithms")
