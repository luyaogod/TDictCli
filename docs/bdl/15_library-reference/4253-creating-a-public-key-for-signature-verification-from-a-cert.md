---
title: "Creating a public key for signature verification from a certificate"
source: "fgl-topics/c_gws_XmlCryptoX509_example_create.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > Examples > Creating a public key for signature verification from a certificate"
type: "concept"
description: "IMPORT xml MAIN DEFINE x509 xml.CryptoX509 DEFINE key xml.CryptoKey LET x509 = xml.CryptoX509.Create() TRY CALL x509.loadPEM(\"RSA1024Certificate.crt\"); CATCH DISPLAY \"Unable to load certificate ..."
---

# Creating a public key for signature verification from a certificate

```
IMPORT xml

MAIN
  DEFINE x509 xml.CryptoX509
  DEFINE key xml.CryptoKey
  LET x509 = xml.CryptoX509.Create()
  TRY
    CALL x509.loadPEM("RSA1024Certificate.crt");
  CATCH
    DISPLAY "Unable to load certificate :",status
    EXIT PROGRAM
  END TRY
  TRY
    LET key = x509.createPublicKey("http://www.w3.org/2000/09/xmldsig#rsa-sha1")
    DISPLAY "Key size (in bytes) : ",key.getSize() # displays 1024 (bits)
    DISPLAY "Key type : ",key.getType() # displays PUBLIC
    DISPLAY "Key usage : ",key.getUsage() # displays SIGNATURE
  CATCH
    DISPLAY "Unable to create public key :",status
  END TRY
END MAIN
```

> **Note:**
>
> All certificates in PEM format were created with the OpenSSL tool.
