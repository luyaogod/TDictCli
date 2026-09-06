---
title: "Generating a symmetric AES256 key"
source: "fgl-topics/c_gws_XmlCryptoKey_example_generate_AES256.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Examples > Generating a symmetric AES256 key"
type: "concept"
---

# Generating a symmetric AES256 key

> This example demonstrates how to generate a symmetric AES256 key using the xml.CryptoKey class. It shows key creation, generation, and displays key properties such as size, type, and usage.

```
IMPORT xml

MAIN
  DEFINE key xml.CryptoKey
  LET key = xml.CryptoKey.Create("http://www.w3.org/2001/04/xmlenc#aes256-cbc")
  TRY
    CALL key.generateKey(NULL)
    DISPLAY "Key size (in bits) : ",key.getSize() # displays 256 (bits)
    DISPLAY "Key type : ",key.getType() # displays SYMMETRIC
    DISPLAY "Key usage : ",key.getUsage() # displays ENCRYPTION
  CATCH
    DISPLAY "Unable to generate key :",status
  END TRY
END MAIN
```

> **Note:**
>
> All keys in PEM or DER format were created with the OpenSSL tool.
