---
title: "xml.CryptoKey.generateKey"
source: "fgl-topics/c_gws_XmlCryptoKey_generateKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.generateKey"
type: "concept"
---

# xml.CryptoKey.generateKey

> Generates a random key of given size (in bits).

## Syntax

```
generateKey(
   keySize INTEGER )
```

1. keySize defines the size of the key
   to generate.

## Usage

For symmetric keys, the size is fixed by the key [identifier](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
and cannot be changed. The only authorized values are the real key size or NULL.

For Diffie-Hellman, the input parameter (size `INTEGER`) is the size of the
Diffie-Hellman modulus. If the given size is greater than zero (0), it populates the Diffie-Hellman
object by randomly generating a modulus of the given size and a private key, and computes the public
key. The used generator is two (2). If the given size is zero (0), it completes the Diffie-Hellman
object by choosing a private key and computing the public key based on the previously loaded
parameters. For more details on loading parameters, see [Table 4](4235-cryptox509-methods.md).

For ECDSA, use [xml.CryptoKey.generateEllipticCurveKey](4199-xml-cryptokey-generateellipticcurvekey.md "Generates a new ECDSA key for the specified named elliptic curve.").

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
