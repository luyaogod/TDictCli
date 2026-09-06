---
title: "xml.Encryption.RSADecrypt"
source: "fgl-topics/c_gws_XmlEncryption_RSADecrypt.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.RSADecrypt"
type: "concept"
---

# xml.Encryption.RSADecrypt

> Decrypts the BASE64 encrypted string using the RSA key and returns it in clear text

## Syntax

```
xml.Encryption.RSADecrypt(
   filename STRING,
   str STRING )
  RETURNS STRING
```

1. filename defines the filename of an RSA private key in PEM format or an
   encryption key entry in the FGLPROFILE file.

   If your FGLPROFILE file contains an entry for an encryption key
   like `xml.myRsa.key="/opt/fourjs/crt/myRsa.pem"` (where "myRsa" is the key
   identifier), pass "myRsa" to the method to retrieve that value. For an example of loading an
   encryption key, refer to [Loading an asymmetric RSA key](4226-loading-an-asymmetric-rsa-key.md "In this example, RSA keys are loaded from PEM files. Code samples show how to load the key file from the PEM file or from FGLPROFILE entries."). For more on
   security-related FGLPROFILE entries, go to [XML configuration](../16_web-services/4916-fglprofile-entries-for-web-services.md) and [FGLPROFILE: XML cryptography](../16_web-services/4921-fglprofile-xml-cryptography.md "Use FGLPROFILE file entries to define XML cryptography and use the fglpass agent to get the private key passwords.")
2. str defines a string that was encrypted with the [fglpass](../16_web-services/4557-encryption-base64-and-password-agent-with-fglpass-tool.md "Genero Web Services supports password encryption with fglpass as password agent.") tool or with the [`xml.Encryption.RSAEncrypt`](4326-xml-encryption-rsaencrypt.md "Encrypts the specified string using the RSA key and returns it encoded in BASE64.") method.

## Usage

RSA decryption is only intended for short strings that cannot exceed the size of the RSA key
minus 12 bytes. For instance, if you have a RSA key of 512 bits, your password cannot exceed
512/8-12 = 52 bytes. If you need to handle big strings, you must use symmetric keys and the [`DecryptString`](4317-xml-encryption-decryptstring.md "Decrypts an encrypted string encoded in BASE64, using the specified symmetric key, and returns the string in clear text.") method. However,
you can use RSA keys to decrypt symmetric key values.

> **Important:**
>
> You must ensure that access to the RSA private key file is restricted only to
> the authorized person or group of persons.

If the RSA private key is protected with a password, the recommended way is to unprotect it with
the [openssl](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") tool
and to put the key file on a restricted file system. But you can also use a [script](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") or the
fglpass
[agent](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.") to provide
the password to the application.

For example, you can encrypt a database password with the fglpass tool and
store it in the FGLPROFILE file, then you can decrypt it with the [`base.Application.getResourceEntry`](2975-base-application-getresourceentry.md "Returns the value of a FGLPROFILE entry.") and the [`xml.Encryption.RSADecrypt`](4325-xml-encryption-rsadecrypt.md "Decrypts the BASE64 encrypted string using the RSA key and returns it in clear text") method
to connect to the database.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Decrypting with an RSA key (xml.Encryption.RSADecrypt)](4228-decrypting-with-an-rsa-key.md "When decrypting data with an RSA key, this example shows how to use a PEM key file or a key referenced from FGLPROFILE entries to perform the decryption.")

[Loading an asymmetric RSA key](4226-loading-an-asymmetric-rsa-key.md "In this example, RSA keys are loaded from PEM files. Code samples show how to load the key file from the PEM file or from FGLPROFILE entries.")
