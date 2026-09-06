---
title: "xml.CryptoKey.getEllipticCurveName"
source: "fgl-topics/c_gws_XmlCryptoKey_getEllipticCurveName.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.getEllipticCurveName"
type: "concept"
---

# xml.CryptoKey.getEllipticCurveName

> Returns the name of the elliptic curve for an ECDSA key, if available.

## Syntax

```
getEllipticCurveName()
  RETURNS STRING
```

## Usage

Returns the name of the elliptic curve associated with the loaded ECDSA key. The returned value will be one of the following strings, depending on the curve used:

- secp256k1
- secp256r1
- prime256v1
- secp384r1
- secp521r1

> **Important:**
>
> **secp256r1**
>
> For `secp256r1`, the returned name will be
> `prime256v1`.

This method does not raise exceptions. If the key is not an ECDSA key or if no curve is
associated, an empty string will be returned.

**Example:**

```
IMPORT xml
DEFINE key xml.CryptoKey
CALL key.loadPEM("ecdsa_private.pem")
DISPLAY key.getEllipticCurveName() -- e.g., "prime256v1"
```

## Related links

**Related information**  

[ECDSA signature algorithms](4222-supported-kind-of-keys.md "ECDSA signature algorithms")
