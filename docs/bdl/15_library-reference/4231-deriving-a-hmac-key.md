---
title: "Deriving a HMAC key"
source: "fgl-topics/c_gws_XmlCryptoKey_example_derive_HMAC.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Examples > Deriving a HMAC key"
type: "concept"
---

# Deriving a HMAC key

> This example demonstrates how to derive an HMAC key using the xml.CryptoKey class. It shows generating a random seed, setting a base password, and deriving a key suitable for encryption or signature functions.

```
IMPORT xml
IMPORT security

MAIN
  DEFINE key xml.CryptoKey
  # will contain a random binary data encoded in Base64
  DEFINE seedBase64 STRING
  LET key = xml.CryptoKey.CreateDerivedKey(
    "http://www.w3.org/2000/09/xmldsig#hmac-sha1")
  TRY
    # Creates a random 24 bytes long binary data encoded into a Base64 form string
    CALL key.setKey("secretpassword")
    # Derives the 14 bytes long "secretpassword" into a 64 bytes long key
    # from a random 24 bytes long seed value and shifting the resulting key
    # from 255 bytes
    LET seedBase64 = security.RandomGenerator.CreateRandomString(24)
    CALL key.deriveKey(
      "http://schemas.xmlsoap.org/ws/2005/02/sc/dk/p_sha1",
      NULL,seedBase64,NULL,255,64)
    # Displays 512 (size of 'secretpassword' derivation in bits)
    DISPLAY "Key size (in bits) : ",key.getSize()
    # Note: Key is derived and can be used in 
    # any encryption or signature function
  CATCH
    DISPLAY "Unable to derive key :",status
  END TRY
END MAIN
```

> **Note:**
>
> All keys in PEM or DER format were created with the OpenSSL tool.
