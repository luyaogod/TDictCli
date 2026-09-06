---
title: "Create an ECDSA public key from coordinates"
source: "fgl-topics/c_gws_XmlCryptoKey_example_loadEllipticCurve.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Examples > Create an ECDSA public key from coordinates"
type: "concept"
description: "This example creates an ECDSA public key from raw x / y point coordinates, validates its serialized representation, and shows how to dynamically load a public key when a PEM file is not available. ..."
---

# Create an ECDSA public key from coordinates

This example creates an ECDSA public key from raw `x`/`y` point
coordinates, validates its serialized representation, and shows how to dynamically load a public key
when a PEM file is not available.

```
IMPORT Xml
FUNCTION TEST1()
  DEFINE key Xml.CryptoKey
  DEFINE str_tmp STRING
  DEFINE pubkey STRING
  LET pubkey = "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEno/am4KD9n9sf6kEPSde11E9Hr+W2PFfoSXthTkdeyEFs+Obrnt5OcPAb2x27P0m7cwuf/51eyO4cJ7G1+MSQQ=="
  TRY
    LET key = Xml.CryptoKey.Create("http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256")
    CALL key.loadEllipticCurve("secp256r1", "no/am4KD9n9sf6kEPSde11E9Hr+W2PFfoSXthTkdeyE=", "BbPjm657eTnDwG9sduz9Ju3MLn/+dXsjuHCextfjEkE=")
    LET str_tmp = key.saveToString()
    IF NOT (str_tmp.equals(pubkey)) THEN
      DISPLAY "TEST 1 Expected '"||pubkey||"' as key value but found '"||str_tmp||"'"
    END IF
    DISPLAY "TEST 1 OK"
  CATCH
    DISPLAY "TEST 1 Error "|| sqlca.sqlerrm
  END TRY
END FUNCTION
```

## Related links

**Related concepts**  

[xml.CryptoKey.loadEllipticCurve](4209-xml-cryptokey-loadellipticcurve.md "Loads an ECDSA key using a specific public point.")
