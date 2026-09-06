---
title: "The fglpass tool"
source: "fgl-topics/c_gws_ssl_fglpass_003.html"
breadcrumb: "Web services > Security > Encryption, BASE64 and password agent with fglpass tool > The fglpass tool"
type: "concept"
---

# The fglpass tool

> Use fglpass to encrypt passwords (output as BASE64), run a protected password agent that supplies private-key passphrases to applications, and encode or decode files in BASE64.

The Genero Web Services package provides a command line tool called fglpass,
which can be used to encrypt a password from an X.509 certificate or a RSA private key. The
encrypted password is displayed on the console in Base64 form, composed only of alphanumeric
characters, and therefore easily usable in any text file.

See [fglpass](../13_programming-tools/2525-fglpass.md "The fglpass tool allows you to encrypt passwords.") for
more details.

## Related links

**Related concepts**  

[Using the password agent](4561-using-the-password-agent.md "Run fglpass in agent mode to securely hold private-key passphrases entered at startup and provide them to BDL applications on demand.")
