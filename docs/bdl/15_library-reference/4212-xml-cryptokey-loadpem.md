---
title: "xml.CryptoKey.loadPEM"
source: "fgl-topics/c_gws_XmlCryptoKey_loadPEM.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.loadPEM"
type: "concept"
---

# xml.CryptoKey.loadPEM

> Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format.

## Syntax

```
loadPEM(
   filename STRING )
```

1. filename defines the filename or an entry in the FGLPROFILE file.

   If your FGLPROFILE file contains an entry for an encryption key
   like `xml.myRsa.key="/opt/fourjs/crt/myRsa.pem"` (where "myRsa" is the key
   identifier), pass "myRsa" to the method to retrieve that value. For an example of loading an
   encryption key, refer to [Loading an asymmetric RSA key](4226-loading-an-asymmetric-rsa-key.md "In this example, RSA keys are loaded from PEM files. Code samples show how to load the key file from the PEM file or from FGLPROFILE entries."). For more on
   security-related FGLPROFILE entries, go to [XML configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md) and [FGLPROFILE: XML cryptography](../16_web-services/4921-fglprofile-xml-cryptography.md "Use FGLPROFILE file entries to define XML cryptography and use the fglpass agent to get the private key passwords.")

## Usage

If the DSA, RSA or ECDSA private key or Diffie-Hellman parameters are protected with a password,
the recommended way is to unprotect it with the openssl tool and to put the key
file on a restricted file system. However, you can use a [script](../16_web-services/4916-fglprofile-entries-for-web-services.md) or the
fglpass
[agent](../16_web-services/4916-fglprofile-entries-for-web-services.md) to provide
the password to the application.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
