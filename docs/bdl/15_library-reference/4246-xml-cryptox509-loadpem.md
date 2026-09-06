---
title: "xml.CryptoX509.loadPEM"
source: "fgl-topics/c_gws_XmlCryptoX509_loadPEM.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoX509 class > CryptoX509 methods > xml.CryptoX509.loadPEM"
type: "concept"
---

# xml.CryptoX509.loadPEM

> Loads a X509 certificate from a file in PEM format.

## Syntax

```
loadPEM(
   filename STRING )
```

1. *filename* defines the filename or an [entry](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") in the FGLPROFILE
   file.

   If your FGLPROFILE contains an entry like
   `xml.myRsa.key="/opt/fourjs/crt/myRsa.pem"` (where "myRsa" is the key identifier),
   pass "myRsa" to the method to retrieve that value. For an example of loading an encryption key,
   refer to [Loading a certificate from a PEM file](4252-loading-a-certificate-from-a-pem-file.md "Examples of how to load X.509 certificates from PEM files or FGLPROFILE entries in Genero."). For more on security-related
   FGLPROFILE entries, go to [XML configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md) and [FGLPROFILE: XML cryptography](../16_web-services/4921-fglprofile-xml-cryptography.md "Use FGLPROFILE file entries to define XML cryptography and use the fglpass agent to get the private key passwords.").

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Loading a certificate from a PEM file](4252-loading-a-certificate-from-a-pem-file.md "Examples of how to load X.509 certificates from PEM files or FGLPROFILE entries in Genero.")
