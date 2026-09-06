---
title: "xml.CryptoKey.generateEllipticCurveKey"
source: "fgl-topics/c_gws_XmlCryptoKey_generateEllipticCurveKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.generateEllipticCurveKey"
type: "concept"
---

# xml.CryptoKey.generateEllipticCurveKey

> Generates a new ECDSA key for the specified named elliptic curve.

## Syntax

```
generateEllipticCurveKey(
  curveName STRING )
```

## Usage

Generates a new ECDSA key using the specified curve name. The `curveName`
parameter must be one of the following strings:

- secp256k1
- prime256v1 (also known as secp256r1)
- secp384r1
- secp521r1

For more information about why you might choose one elliptic curve over another, refer to [Elliptic curves](4222-supported-kind-of-keys.md)

The `xml.CryptoKey.generateEllipticCurveKey` method raises an exception if called
on a non-ECDSA `xml.CryptoKey` object or if the `curveName` is
invalid.

**Example:**

```
IMPORT xml
DEFINE key xml.CryptoKey
LET key = xml.CryptoKey.Create(key_url)
CALL key.generateEllipticCurveKey("curveName")
-- The key now contains a new ECDSA key for the specified curve
```

Where:

- key\_url is the URL for the [supported
  key algorithm](4222-supported-kind-of-keys.md) of one of the supported curves listed above.
- curveName is one of the supported curves shown above.

For optimal security, it is essential to use the appropriate key size with the corresponding
algorithm: using a key generated with secp256k1 (256 bits) with the
ecdsa-sha512 algorithm (512 bits) will not enhance security compared to using the same
key with the ecdsa-sha256 algorithm. Always ensure that the key and algorithm are
aligned to best practices for effective cryptographic security.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related information**  

[ECDSA signature algorithms](4222-supported-kind-of-keys.md "ECDSA signature algorithms")
