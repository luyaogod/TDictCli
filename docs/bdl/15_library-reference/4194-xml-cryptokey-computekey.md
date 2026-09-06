---
title: "xml.CryptoKey.computeKey"
source: "fgl-topics/c_gws_XmlCryptoKey_computeKey.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > CryptoKey methods > xml.CryptoKey.computeKey"
type: "concept"
---

# xml.CryptoKey.computeKey

> Computes the shared secret based on the given modulus, generator, the private key, and the other peer's public key. The returned key can be any symmetric/HMAC or symmetric/encryption key type. It can be used for symmetric signature or symmetric encryption.

## Syntax

```
computeKey(
   pub xml.CryptoKey,
   url STRING )
  RETURNS xml.CryptoKey
```

1. pub defines the other peer's public
   key ([`xml.CryptoKey`](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.")).
2. url defines the shared secret key type as a URL identifier
   (`STRING`).

## Usage

> **Important:**
>
> This method is for Diffie-Hellman key-agreement algorithm only.

Returns a `xml.CryptoKey`
`sharedSecret`: a `xml.CryptoKey` object of the specified type.

In the 3DES case, no key weakness test is done. If the compound shared secret is week, the other
peer involved in the communication may raise an error. It depends on the language used on the other
side.

In order to be able to compute an AES256 shared secret of the Java side, you need to add or
replace the files local\_policy.jar and
US\_export\_policy.jar located in $JDK\_HOME/jre/lib/security
by the Java Cryptographic Extension corresponding to your JDK version. You can find this extension
at http://www.oracle.com/technetwork/java/javase/downloads/index.html.

If the shared secret key length is less than the Diffie-Hellman key length, only the first needed
bytes will be taken. For example, if the Diffie-Hellman is 512 bits length and the shared secret is
a 3DES key, then only the first 192 bits will be used by the computation. In a 3DES shared secret
case, `xml.CryptoKey.computeKey()` is calculated, whereas in AES shared secret case,
the Diffie-Hellman key is truncated.

If the shared secret key length is bigger than the Diffie-Hellman key length, an error is
raised.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
