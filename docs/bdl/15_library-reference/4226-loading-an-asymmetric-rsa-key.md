---
title: "Loading an asymmetric RSA key"
source: "fgl-topics/c_gws_XmlCryptoKey_example_load_RSA.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Examples > Loading an asymmetric RSA key"
type: "concept"
---

# Loading an asymmetric RSA key

> In this example, RSA keys are loaded from PEM files. Code samples show how to load the key file from the PEM file or from FGLPROFILE entries.

## Loading an RSA key from a PEM file

This example demonstrates loading an RSA key from a PEM file using [xml.CryptoKey.loadPEM](4212-xml-cryptokey-loadpem.md "Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format."). The key must be loaded before it can be used for
encryption or decryption operations.

```
IMPORT xml

MAIN
  DEFINE key xml.CryptoKey
  LET key = xml.CryptoKey.Create("http://www.w3.org/2001/04/xmlenc#rsa-1_5")
  TRY
    CALL key.loadPEM("RSA1024Key.pem")
    CALL key.setFeature("KeyName","MyRsaKey")
    DISPLAY "Key size (in bits) : ",key.getSize() # displays 1024 (bits)
    DISPLAY "Key type : ",key.getType() # displays PRIVATE or PUBLIC
    DISPLAY "Key usage : ",key.getUsage() # displays KEYENCRYPTION
  CATCH
    DISPLAY "Unable to load key :",status
  END TRY
END MAIN
```

> **Note:**
>
> All keys in PEM or DER format were created with the OpenSSL tool.

## Loading an RSA key from an FGLPROFILE entry

You can also load a PEM key by reference using an FGLPROFILE entry instead of a direct file path.
This allows you to manage key file locations centrally in your application configuration.

Example FGLPROFILE entry:

```
# fglprofile
xml.myRSA.key = "/opt/local/cert-key.pem"
```

Ensure your FGLPROFILE environment variable points to the correct fglprofile
file.

Example code:

```
IMPORT xml

MAIN
  DEFINE key xml.CryptoKey
  LET key = xml.CryptoKey.Create("http://www.w3.org/2001/04/xmlenc#rsa-1_5")
  TRY
    CALL key.loadPEM("myRSA") -- loads the key file referenced by the FGLPROFILE entry
    CALL key.setFeature("KeyName","MyRsaKey")
    DISPLAY "Key size (in bits) : ",key.getSize()
    DISPLAY "Key type : ",key.getType()
    DISPLAY "Key usage : ",key.getUsage()
  CATCH
    DISPLAY "Unable to load key :",status
  END TRY
END MAIN
```

> **Note:**
>
> The string passed to `xml.CryptoKey.loadPEM` is the logical key name as defined in
> the FGLPROFILE, not a file path.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
