---
title: "Encryption, BASE64 and password agent with fglpass tool"
source: "fgl-topics/c_gws_ssl_fglpass_001.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool"
type: "concept"
---

# Encryption, BASE64 and password agent with fglpass tool

> Genero Web Services supports password encryption with fglpass as password agent.

For security reasons, it is recommended that you avoid storing
clear passwords in a file. The Genero Web Services enables the password
encryption of a HTTP Authenticate entry in the FGLPROFILE file. The
encrypted password is decrypted by the Genero Web Services engine
when required.

## Child topics

- [The fglpass tool](4558-the-fglpass-tool.md): Use fglpass to encrypt passwords (output as BASE64), run a protected password agent that supplies private-key passphrases to applications, and encode or decode files in BASE64.
- [Encrypt a HTTP authenticate password for FGLPROFILE](4559-encrypt-password-in-fglprofile-file.md): Use the fglpass tool to encrypt a password for storing in the FGLPROFILE file.
- [Encrypt a HTTP authenticate password using a certificate in the Windows key store](4560-encrypt-password-for-windows-key-store.md): Use the fglpass tool to encrypt a password to store in the Windows® key store.
- [Using the password agent](4561-using-the-password-agent.md): Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand.
- [Encrypt a password from RSA key (fglpass)](4562-encrypt-a-password-from-rsa-key.md): This task shows how to encrypt a password with an RSA key using the fglpass tool.
- [Decrypt a password from BASE64 (fglpass)](4563-decrypt-a-password-from-base64.md): This task shows how to decrypt an RSA-encrypted password using the fglpass tool.
- [Encode a file in BASE64 form](4564-encode-a-file-in-base64-form.md): The fglpass tool can encode a file in BASE64 form.
- [Decode a file encoded in BASE64 form](4565-decode-a-file-encoded-in-base64-form.md): The fglpass tool can decode a BASE64 encoded file.
