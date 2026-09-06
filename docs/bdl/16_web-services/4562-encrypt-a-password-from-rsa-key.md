---
title: "Encrypt a password from RSA key (fglpass)"
source: "fgl-topics/c_gws_ssl_fglpass_007.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > Encrypt a password from RSA key"
type: "concept"
---

# Encrypt a password from RSA key (fglpass)

> This task shows how to encrypt a password with an RSA key using the fglpass tool.

The [fglpass](4558-the-fglpass-tool.md "Use fglpass to encrypt passwords (output as BASE64), run a protected password agent that supplies private-key passphrases to applications, and encode or decode files in BASE64.") tool can encrypt a
password using an RSA key or certificate, and then encode it in BASE64 form. This allows you to add
a protected password in the FGLPROFILE file for future use by any BDL application.
> **Tip:**
>
> You can also use [xml.Encryption.RSAEncrypt](../15_library-reference/4326-xml-encryption-rsaencrypt.md "Encrypts the specified string using the RSA key and returns it encoded in BASE64.") to encrypt a plaintext
> password to BASE64 from within your Genero program (so you do not need to run `fglpass
> -e` or other external tools); go to [Encrypting with an RSA key (xml.Encryption.RSAEncrypt)](../15_library-reference/4227-encrypting-with-an-rsa-key.md "When encrypting data with an RSA key, this example shows how to use a PEM key file or a key referenced from FGLPROFILE entries to perform the encryption.") for usage and examples.

1. To encrypt a password from a RSA key and encoded in BASE64,
   enter:

   ```
   fglpass -e -k myprivatekey.pem
   ```

   When encrypting a password, you can either supply the
   certificate with the `-c` option (`fglpass -e -c`) or provide the
   private key using the `-k` option.
   > **Note:**
   >
   > The private key file also contains (or allows derivation of) the corresponding public key, so
   > when you supply the private key the public portion is extracted and used to encrypt the password;
   > the private key is required later to decrypt it.
2. You are prompted to enter the password you want to encrypt.

   ```
   Enter password :hello
   ```

   The fglpass tool outputs the BASE64 form of the encrypted password on the
   console.

   ```
   BASE64 BEGIN
   Pzk/fNRhetdJDZz5kjNg7P0XET4XsW6bys/fi0DvugxRPh9d/s41oAws65
   JY0EPb2zytQjxZ/dwaaRzJPYoQmA==
   BASE64 END
   ```

   The BASE64 encrypted password is the string between the `BASE64
   BEGIN` and `BASE64 END`.
3. For details on storing the encrypted password in fglprofile, go to [Encrypt a HTTP authenticate password for FGLPROFILE](4559-encrypt-password-in-fglprofile-file.md "Use the fglpass tool to encrypt a password for storing in the FGLPROFILE file.").

## Related links

**Related concepts**  

[The fglpass tool](4558-the-fglpass-tool.md "Use fglpass to encrypt passwords (output as BASE64), run a protected password agent that supplies private-key passphrases to applications, and encode or decode files in BASE64.")

[Decrypt a password from BASE64 (fglpass)](4563-decrypt-a-password-from-base64.md "This task shows how to decrypt an RSA-encrypted password using the fglpass tool.")

[xml.Encryption.RSAEncrypt](../15_library-reference/4326-xml-encryption-rsaencrypt.md "Encrypts the specified string using the RSA key and returns it encoded in BASE64.")

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")

[Using the password agent](4561-using-the-password-agent.md "Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand.")
