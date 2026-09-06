---
title: "Saving the subjectName of a certificate in XML"
source: "fgl-topics/c_gws_XmlCryptoX509_example_save.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > Examples > Saving the subjectName of a certificate in XML"
type: "concept"
description: "IMPORT xml MAIN DEFINE x509 xml.CryptoX509 DEFINE doc xml.DomDocument LET x509 = xml.CryptoX509.Create() TRY CALL x509.loadPEM(\"RSA1024Certificate.crt\"); CATCH DISPLAY \"Unable to load certificate ..."
---

# Saving the subjectName of a certificate in XML

```
IMPORT xml

MAIN
  DEFINE x509 xml.CryptoX509
  DEFINE doc xml.DomDocument
  LET x509 = xml.CryptoX509.Create()
  TRY
    CALL x509.loadPEM("RSA1024Certificate.crt");
  CATCH
    DISPLAY "Unable to load certificate :",status
    EXIT PROGRAM
  END TRY
  TRY
    CALL x509.setFeature("X509SubjectName",TRUE)
    LET doc = x509.save()
    CALL doc.setFeature("format-pretty-print",TRUE)
    CALL doc.save("RSAX509SubjectName.xml")
  CATCH
    DISPLAY "Unable to save certificate :",status
  END TRY
END MAIN
```

> **Note:**
>
> All certificates in PEM format were created with the OpenSSL tool.
