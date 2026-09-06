---
title: "Loading a certificate from a PEM file"
source: "fgl-topics/c_gws_XmlCryptoX509_example_load.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > Examples > Loading a certificate from a PEM file"
type: "concept"
---

# Loading a certificate from a PEM file

> Examples of how to load X.509 certificates from PEM files or FGLPROFILE entries in Genero.

## Loading a certificate from a PEM file

```
IMPORT xml

MAIN
  DEFINE x509 xml.CryptoX509
  LET x509 = xml.CryptoX509.Create()
  TRY
  CALL x509.loadPEM("certificate.crt");
    DISPLAY "Id : ",x509.getIdentifier()
  CATCH
    DISPLAY "Unable to load certificate :",status
  END TRY
END MAIN
```

> **Note:**
>
> All certificates in PEM format were created with the OpenSSL tool.

## Loading a certificate from an FGLPROFILE entry

You can also load a PEM certificate by reference using an FGLPROFILE entry instead of a direct file path. This allows you to manage certificate file locations centrally in your application configuration.

Example FGLPROFILE entry:

```
# fglprofile
xml.myCert.cert = "/opt/local/certificate.crt"
```

Example code:

```
IMPORT xml

MAIN
  DEFINE x509 xml.CryptoX509
  LET x509 = xml.CryptoX509.Create()
  TRY
  CALL x509.loadPEM("myCert") -- loads the certificate file referenced by the FGLPROFILE entry
    DISPLAY "Id : ",x509.getIdentifier()
  CATCH
    DISPLAY "Unable to load certificate :",status
  END TRY
END MAIN
```

> **Note:**
>
> The string passed to `xml.CryptoX509.loadPEM` is the logical certificate
> name as defined in the FGLPROFILE, not a file path.

For details about the `xml.CryptoX509.loadPEM` method, see [`xml.CryptoX509.loadPEM`](4246-xml-cryptox509-loadpem.md "Loads a X509 certificate from a file in PEM format.").
