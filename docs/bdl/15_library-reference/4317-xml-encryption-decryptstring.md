---
title: "xml.Encryption.DecryptString"
source: "fgl-topics/c_gws_XmlEncryption_DecryptString.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Encryption class > Encryption methods > xml.Encryption.DecryptString"
type: "concept"
---

# xml.Encryption.DecryptString

> Decrypts an encrypted string encoded in BASE64, using the specified symmetric key, and returns the string in clear text.

## Syntax

```
xml.Encryption.DecryptString(
   key xml.CryptoKey ,
   str STRING )
  RETURNS STRING
```

1. key defines the symmetric [key](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") to use for decryption.
2. str defines the encrypted string for
   decryption.

## Usage

This method decrypts the encrypted string str encoded in BASE64, using the
symmetric key specified in key, and returns the string in clear text.

The key's usage must be for `encryption`, see [Supported kind of keys](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.").

There are some considerations to take into account when
encrypting / decrypting a string in languages other than Genero BDL. The encryption and
decryption of a string follows XML-Encryption specification. It defines two rules you must
adapt when coding in another language:

1. The initialization vector (IV) is concatenated to the beginning of the encrypted
   string.
2. A padding value defined as ISO 10126 may be appended to the end of the encrypted
   string.

For more details see <https://www.w3.org/TR/xmlenc-core/#sec-Alg-Block>.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").
