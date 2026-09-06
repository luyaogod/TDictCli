---
title: "xml.Signature.VerifyString"
source: "fgl-topics/c_gws_XmlSignature_verifyString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.VerifyString"
type: "concept"
---

# xml.Signature.VerifyString

> Verify the signature is consistent with the given key and the original message.

## Syntax

```
VerifyString(
   key xml.CryptoKey,
   originalStr STRING,
   signature STRING )
  RETURNS INTEGER
```

1. key defines the [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") to use for
   verification.
2. originalStr defines the signed string
   in its clear form.
3. toBeVerifiedStr defines the signature
   to be verified.

## Usage

The key can be a HMAC key, a RSA private key or a DSA private key. The HMAC key must be the same
as the one used for signing. The public RSA and DSA key must be the public key corresponding to the
private key used for signing.

Returns 1 when verification is successful; 0 (zero) is returned if verification fails.

This method does not belong to the XML encryption specification.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
