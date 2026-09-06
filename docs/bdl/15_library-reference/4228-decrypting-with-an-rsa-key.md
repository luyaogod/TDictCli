---
title: "Decrypting with an RSA key (xml.Encryption.RSADecrypt)"
source: "fgl-topics/c_gws_XmlCryptoKey_example_decrypt_rsa.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Examples > Decrypting with an RSA key"
type: "concept"
---

# Decrypting with an RSA key (xml.Encryption.RSADecrypt)

> When decrypting data with an RSA key, this example shows how to use a PEM key file or a key referenced from FGLPROFILE entries to perform the decryption.

The following examples show how to use the [xml.Encryption.RSADecrypt](4325-xml-encryption-rsadecrypt.md "Decrypts the BASE64 encrypted string using the RSA key and returns it in clear text") method to decrypt a string using an RSA key.

**Example 1: Using a PEM file path**

```
IMPORT xml

DEFINE encryptedStr, decryptedStr STRING

MAIN
  LET encryptedStr = xml.Encryption.RSAEncrypt("/opt/local/cert-key.pem", "Mary had a little lamb")
  LET decryptedStr = xml.Encryption.RSADecrypt("/opt/local/cert-key.pem", encryptedStr)
  DISPLAY "Encrypted string: ", encryptedStr
  DISPLAY "Decrypted string: ", decryptedStr
END MAIN
```

**Example 2: Using an FGLPROFILE entry**

Assuming your FGLPROFILE contains:

```
# fglprofile
xml.myRSA.key = "/opt/local/cert-key.pem"
```

Ensure your FGLPROFILE environment variable
points to the correct fglprofile file.

The code:

```
IMPORT xml

DEFINE encryptedStr, decryptedStr STRING
MAIN
  LET encryptedStr = xml.Encryption.RSAEncrypt("myRSA", "Mary had a little lamb")
  LET decryptedStr = xml.Encryption.RSADecrypt("myRSA", encryptedStr)
  DISPLAY "Encrypted string: ", encryptedStr
  DISPLAY "Decrypted string: ", decryptedStr
END MAIN
```

In both cases, the first argument to `xml.Encryption.RSAEncrypt` and
`xml.Encryption.RSADecrypt` can be a file path or a logical key name defined in the
FGLPROFILE.

> **Tip:**
>
> You can also use [The fglpass tool](../16_web-services/4558-the-fglpass-tool.md "Use fglpass to encrypt passwords (output as BASE64), run a protected password agent that supplies private-key passphrases to applications, and encode or decode files in BASE64.") (`fglpass -e`) to
> encrypt a plaintext password to RSA/BASE64 outside the application (for example, when preparing
> FGLPROFILE entries). For usage and examples, go to [Decrypt a password from BASE64 (fglpass)](../16_web-services/4563-decrypt-a-password-from-base64.md "This task shows how to decrypt an RSA-encrypted password using the fglpass tool.")

## Related links

**Related concepts**  

[xml.Encryption.RSAEncrypt](4326-xml-encryption-rsaencrypt.md "Encrypts the specified string using the RSA key and returns it encoded in BASE64.")

[xml.Encryption.RSADecrypt](4325-xml-encryption-rsadecrypt.md "Decrypts the BASE64 encrypted string using the RSA key and returns it in clear text")

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
