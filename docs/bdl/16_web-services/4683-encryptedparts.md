---
title: "EncryptedParts"
source: "fgl-topics/c_gws_ssl_security_how_to_009.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to handle WS security > SOAP security standards > EncryptedParts"
type: "concept"
---

# EncryptedParts

> The EncryptedParts section of the policy specifies which part of the message should be encrypted.

```
<sp:EncryptedParts xmlns:sp="http://schemas.xmlsoap.org/ws/2005/07/securitypolicy">
  <sp:Body />
```

- `sp:Body` indicates the body message needs to be encrypted

Encrypt the body using the algorithm referenced in assertion [AlgorithmSuite](4680-security-bindings.md "Understand the mechanism of bindings that allows the secure exchange of SOAP messages over HTTP. The demo application security policy is referenced to illustrate this."):

- Create an encryption key using TripleDesRsa15 algorithm (it generates a TripleDES symmetric key
  and then encrypts it with a RSA1.5 public key), like in [example2](../15_library-reference/4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") that uses AES256 in the CryptoKey
  chapter.
- Encrypt the body with the created key.

To find the exact syntax of security message read the specifications
"Web Services Security: SOAP Message Security 1.0".
