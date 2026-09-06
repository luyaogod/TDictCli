---
title: "xml.Encryption.RSAEncrypt"
source: "fgl-topics/c_gws_XmlEncryption_RSAEncrypt.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.RSAEncrypt"
type: "concept"
---

# xml.Encryption.RSAEncrypt

> Encrypts the specified string using the RSA key and returns it encoded in BASE64.

## Syntax

```
xml.Encryption.RSAEncrypt(
   filename STRING,
   str STRING )
  RETURNS STRING
```

1. filename defines the filename of an RSA public or private key in PEM format
   or an encryption key entry in the FGLPROFILE file.

   If your FGLPROFILE file contains an entry for an encryption key
   like `xml.myRsa.key="/opt/fourjs/crt/myRsa.pem"` (where "myRsa" is the key
   identifier), pass "myRsa" to the method to retrieve that value. For an example of loading an
   encryption key, refer to [Loading an asymmetric RSA key](4226-loading-an-asymmetric-rsa-key.md "In this example, RSA keys are loaded from PEM files. Code samples show how to load the key file from the PEM file or from FGLPROFILE entries."). For more on
   security-related FGLPROFILE entries, go to [XML configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md) and [FGLPROFILE: XML cryptography](../16_web-services/4921-fglprofile-xml-cryptography.md "Use FGLPROFILE file entries to define XML cryptography and use the fglpass agent to get the private key passwords.")
2. str defines the string to be encrypted.

## Usage

RSA encryption is only intended for short strings that cannot exceed the size of the RSA key
minus 12 bytes. For instance, if you have a RSA key of 512 bits, you password cannot exceed 512/8-12
= 52 bytes. If you need to handle big strings, you must use symmetric keys and the [`xml.Encryption.EncryptString`](4323-xml-encryption-encryptstring.md "Encrypts the specified string using the symmetric key, and returns the encrypted string encoded in BASE64.")
method. However, you can use RSA keys to encrypt symmetric key values.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Loading an asymmetric RSA key](4226-loading-an-asymmetric-rsa-key.md "In this example, RSA keys are loaded from PEM files. Code samples show how to load the key file from the PEM file or from FGLPROFILE entries.")

[Encrypting with an RSA key (xml.Encryption.RSAEncrypt)](4227-encrypting-with-an-rsa-key.md "When encrypting data with an RSA key, this example shows how to use a PEM key file or a key referenced from FGLPROFILE entries to perform the encryption.")
