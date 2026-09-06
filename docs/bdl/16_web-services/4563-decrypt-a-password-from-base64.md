---
title: "Decrypt a password from BASE64 (fglpass)"
source: "fgl-topics/c_gws_ssl_fglpass_008.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > Decrypt a password from BASE64"
type: "concept"
---

# Decrypt a password from BASE64 (fglpass)

> This task shows how to decrypt an RSA-encrypted password using the fglpass tool.

The [fglpass](4558-the-fglpass-tool.md "Use fglpass to encrypt passwords (output as BASE64), run a protected password agent that supplies private-key passphrases to applications, and encode or decode files in BASE64.") tool uses the RSA private key that was
used to encrypt it or that is associated to a certificate containing the public part of that private
key.

> **Tip:**
>
> You can also use [xml.Encryption.RSADecrypt()](../15_library-reference/4325-xml-encryption-rsadecrypt.md "Decrypts the BASE64 encrypted string using the RSA key and returns it in clear text") to decrypt an RSA-encrypted password from within your Genero
> program (so you do not need to run `fglpass -d` or other external tools); go to [Decrypting with an RSA key (xml.Encryption.RSADecrypt)](../15_library-reference/4228-decrypting-with-an-rsa-key.md "When decrypting data with an RSA key, this example shows how to use a PEM key file or a key referenced from FGLPROFILE entries to perform the decryption.") for usage and examples.

- To decrypt, run:

  ```
  fglpass -d -k RSAPriv.pem
  ```
- If RSAPriv.pem is protected, you are prompted for the passphrase:

  ```
  Enter  pass phrase for RSAPriv.pem:
  ```
- Then paste the BASE64-encoded encrypted password when prompted:

  ```
  Enter password :Pzk/fNRhetdJDZz5kjNg7P0XET4XsW6bys/fi0
  DvugxRPh9d/s41oAws65JY0EPb2zytQjxZ/dwaaRzJPYoQmA==
  ```
- `fglpass` prints the decrypted password in clear text to the console (for example
  `hello`).

## Related links

**Related concepts**  

[The fglpass tool](4558-the-fglpass-tool.md "Use fglpass to encrypt passwords (output as BASE64), run a protected password agent that supplies private-key passphrases to applications, and encode or decode files in BASE64.")

[Encrypt a password from RSA key (fglpass)](4562-encrypt-a-password-from-rsa-key.md "This task shows how to encrypt a password with an RSA key using the fglpass tool.")

[xml.Encryption.RSADecrypt](../15_library-reference/4325-xml-encryption-rsadecrypt.md "Decrypts the BASE64 encrypted string using the RSA key and returns it in clear text")

[FGLPROFILE entries for web services](4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")

[Using the password agent](4561-using-the-password-agent.md "Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand.")
