---
title: "Setting a HMAC key"
source: "fgl-topics/c_gws_XmlCryptoKey_example_set_HMAC.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Examples > Setting a HMAC key"
type: "concept"
---

# Setting a HMAC key

> This example demonstrates how to create and set an HMAC key using the xml.CryptoKey class. It shows key initialization, setting the key value, and retrieving key properties such as size, type, and usage.

```
IMPORT xml

MAIN
  DEFINE key xml.CryptoKey
  LET key = xml.CryptoKey.Create("http://www.w3.org/2000/09/xmldsig#hmac-sha1")
  TRY
    CALL key.setKey("secretpassword")
    # displays 112 (size of secretpassword in bits)
    DISPLAY "Key size (in bits) : ",key.getSize()
    DISPLAY "Key type : ",key.getType() # displays HMAC
    DISPLAY "Key usage : ",key.getUsage() # displays SIGNATURE
  CATCH
    DISPLAY "Unable to set key :",status
  END TRY
END MAIN
```

> **Note:**
>
> All keys in PEM or DER format were created with the OpenSSL tool.
