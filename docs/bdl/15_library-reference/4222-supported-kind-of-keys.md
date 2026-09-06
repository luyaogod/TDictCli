---
title: "Supported kind of keys"
source: "fgl-topics/r_gws_XmlCryptoKey_supported_keys.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The CryptoKey class > Supported kind of keys"
type: "reference"
---

# Supported kind of keys

> Types of keys supported by the xml.CryptoKey class.

The `xml.CryptoKey` class supports a wide range of key types for digital
signatures, encryption, and key agreement, including Digital Signature Algorithm (DSA), RSA, HMAC,
AES, TripleDES, Diffie-Hellman, and ECDSA with various hash algorithms.

| Identifier | Description | Usage | Type |
| --- | --- | --- | --- |
| http://www.w3.org/2000/09/xmldsig#dsa-sha1 | Asymmetric DSA key with SHA1 for signature purposes.Uses a private DSA key for signature and needs an associated public DSA key or X509 certificate containing it, to verify it.See [specification](http://www.w3.org/TR/xmldsig-core/#sec-SignatureAlg) for details.**Important:**SHA-1 is deprecated. Use the `dsa-sha256` key instead. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2009/xmldsig11#dsa-sha256 | Asymmetric DSA key with SHA-256 for signature purposes.Uses a private DSA key for signature and needs an associated public DSA key or X509 certificate containing it, to verify it.See [specification](https://www.w3.org/TR/xmldsig-core1/#sec-DSA) for details. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2000/09/xmldsig#rsa-sha1 | Asymmetric RSA key with SHA1 for signature purposes.Uses a private RSA key for signature and needs an associated public RSA key or X509 certificate containing it, to verify it.See [specification](http://www.w3.org/TR/xmldsig-core/#sec-SignatureAlg) for details. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmldsig-more#rsa-sha256 | Asymmetric RSA key with SHA256 for signature purposes.Uses a private RSA key for signature and needs an associated public RSA key or X509 certificate containing it, to verify it.See [specification](http://www.w3.org/TR/xmldsig-core2/#sec-PKCS1) for details. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmldsig-more#rsa-sha384 | Asymmetric RSA key with SHA384 for signature purposes.Uses a private RSA key for signature and needs an associated public RSA key or X509 certificate containing it, to verify it.See [specification](http://www.w3.org/TR/xmldsig-core2/#sec-PKCS1) for details. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmldsig-more#rsa-sha512 | Asymmetric RSA key with SHA512 for signature purposes.Uses a private RSA key for signature and needs an associated public RSA key or X509 certificate containing it, to verify it.See [specification](http://www.w3.org/TR/xmldsig-core2/#sec-PKCS1) for details. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2007/05/xmldsig-more#sha256-rsa-MGF1 | Use this key type for RSA signatures requiring SHA-256 and RSA-PSS padding, equivalent to the PS256 algorithm in JWT and OpenID Connect. Use a private RSA key to sign and an associated public RSA key or X509 certificate to verify.See [specification](https://datatracker.ietf.org/doc/html/rfc6931#section-2.3.10) for details. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2000/09/xmldsig#hmac-sha1 | Message Authentication Code key with SHA1 for signature purposes.Uses the same password for signature and to verify it, and key size is free.See [specification](http://www.w3.org/TR/xmldsig-core/#sec-SignatureAlg) for details. | SIGNATURE | HMAC |
| http://www.w3.org/2001/04/xmldsig-more#hmac-sha256 | Message Authentication Code key with SHA256 for signature purposes.Uses the same password for signature and to verify it, and key size is free.See [specification](http://www.w3.org/TR/xmldsig-core2/#sec-HMAC) for details. | SIGNATURE | HMAC |
| http://www.w3.org/2001/04/xmldsig-more#hmac-sha384 | Message Authentication Code key with SHA384 for signature purposes.Uses the same password for signature and to verify it, and key size is free.See [specification](http://www.w3.org/TR/xmldsig-core2/#sec-HMAC) for details. | SIGNATURE | HMAC |
| http://www.w3.org/2001/04/xmldsig-more#hmac-sha512 | Message Authentication Code key with SHA512 for signature purposes.Uses the same password for signature and to verify it, and key size is free.See [specification](http://www.w3.org/TR/xmldsig-core2/#sec-HMAC) for details. | SIGNATURE | HMAC |
| http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1 | ECDSA key with SHA-1 for signature purposes. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha224 | ECDSA key with SHA-224 for signature purposes. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256 | ECDSA key with SHA-256 for signature purposes. Typically used with SHA-256 curves: secp256k1 or secp256r1 (also known as prime256v1) | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384 | ECDSA key with SHA-384 for signature purposes. Typically used with curve secp384r1 | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512 | ECDSA key with SHA-512 for signature purposes. Typically used with curve secp521r1 | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2021/04/xmldsig-more#ecdsa-sha3-224 | ECDSA key with SHA3-224 for signature purposes. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2021/04/xmldsig-more#ecdsa-sha3-256 | ECDSA key with SHA3-256 for signature purposes. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2021/04/xmldsig-more#ecdsa-sha3-384 | ECDSA key with SHA3-384 for signature purposes. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2021/04/xmldsig-more#ecdsa-sha3-512 | ECDSA key with SHA3-512 for signature purposes. | SIGNATURE | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmlenc#aes128-cbc | Symmetric AES128 key for encryption purposes.Uses a common key of 128bits for encrypting and decrypting XML documents.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-Block) for details. | ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#aes192-cbc | Symmetric AES192 key for encryption purposes.Uses a common key of 192bits for encrypting and decrypting XML documents.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-Block) for details. | ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#aes256-cbc | Symmetric AES256 key for encryption purposes.Uses a common key of 256bits for encrypting and decrypting XML documents.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-Block) for details. | ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#tripledes-cbc | Symmetric TripleDes key for encryption purposes.Uses a common key of 192bits for encrypting and decrypting XML documents.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-Block) for details. | ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#kw-aes128 | Symmetric AES128 key wrap for key encryption purposes.Uses a common key of 128bits for encrypting and decrypting a symmetric key.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-SymmetricKeyWrap) for details. | KEY ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#kw-aes192 | Symmetric AES192 key wrap for key encryption purposes.Uses a common key of 192bits for encrypting and decrypting a symmetric key.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-SymmetricKeyWrap) for details. | KEY ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#kw-aes256 | Symmetric AES256 key wrap for key encryption purposes.Uses a common key of 256bits for encrypting and decrypting a symmetric key.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-SymmetricKeyWrap) for details. | KEY ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#kw-tripledes | Symmetric TripleDes key wrap for key encryption purposes.Uses a common key of 192bits for encrypting and decrypting a symmetric key.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-SymmetricKeyWrap) for details. | KEY ENCRYPTION | SYMMETRIC |
| http://www.w3.org/2001/04/xmlenc#rsa-1\_5 | Asymmetric RSA key for key encryption purposes.Uses a public RSA key or a X509 certificate containing it to encrypt a symmetric key, and needs the associated private RSA key to decrypt it.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-KeyTransport) for details. | KEY ENCRYPTION | PUBLIC or PRIVATE |
| http://www.w3.org/2001/04/xmlenc#rsa-oaep-mgf1p | Asymmetric RSA key for key encryption purposes.Uses a public RSA key or a X509 certificate containing it to encrypt a symmetric key, and needs the associated private RSA key to decrypt it.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-KeyTransport) for details. | KEY ENCRYPTION | PUBLIC or PRIVATE |
| Diffie-Hellman identifier: http://www.w3.org/2001/04/xmlenc#DHKeyValue | Diffie-Hellman key agreement algorithm. Derives a shared secret. The resulting shared secret is a HMAC or symmetric key for encryption purposes.See [specification](http://www.w3.org/TR/xmlenc-core/#sec-Alg-KeyAgreement) for details. | KEY AGREEMENT | PUBLIC or PRIVATE |

## ECDSA signature algorithms

The Elliptic Curve Digital Signature Algorithm (ECDSA) is widely used for creating digital
signatures due to its strong security and efficiency, as it achieves high levels of security with
relatively small key sizes. You might consider using ECDSA keys when less computational power is
available, making it a useful option for your mobile apps secure communications.

ECDSA utilizes a variety of named elliptic curves, which are standardized curves rigorously
analyzed for security and performance. An elliptic curve is defined by a specific set of
mathematical parameters that influence the cryptographic operations. The choice of curve
significantly impacts both the security and performance of ECDSA operations.

Refer to Table 2 for the officially supported named curves.
Understanding these curves and their uses may be helpful for making informed decisions about which
curve to use.

For optimal security, it is essential to use the appropriate key size with the corresponding
algorithm listed in Supported kind of keys: using a key generated
with secp256k1 (256 bits) curve with the ecdsa-sha512 algorithm (512 bits)
will not enhance security compared to using the same key with the ecdsa-sha256
algorithm. Always ensure that the key and algorithm are aligned to best practices for effective
cryptographic security.

There are specific methods for working with ECDSA—[xml.CryptoKey.generateEllipticCurveKey](4199-xml-cryptokey-generateellipticcurvekey.md "Generates a new ECDSA key for the specified named elliptic curve."), [xml.CryptoKey.loadEllipticCurve](4209-xml-cryptokey-loadellipticcurve.md "Loads an ECDSA key using a specific public point."), and [xml.CryptoKey.getEllipticCurveName](4201-xml-cryptokey-getellipticcurvename.md "Returns the name of the elliptic curve for an ECDSA key, if available.") —while all other methods in the [xml.CryptoKey](4191-the-cryptokey-class.md "The xml.CryptoKey class provides methods to manipulate HMAC, symmetric and asymmetric keys needed for signing, verifying, encrypting and decrypting XML documents or document fragments.") class can handle ECDSA keys similarly to DSA
keys. Similarly, in the [xml.Signature](4255-the-signature-class.md "The xml.Signature class provides methods to create detached, enveloped or enveloping XML signatures of one or more references of XML documents or document fragments, and to determine whether a signed referenced document has been modified afterwards.") class, ECDSA keys
are used just like DSA keys.

| Elliptic curve name | Description of use |
| --- | --- |
| secp256k1 | Commonly used in cryptocurrencies like Bitcoin. It is known for its efficiency and is not a standard from the National Institute of Standards and Technology (NIST). |
| secp256r1 | This NIST standardized curve is widely used in various applications, including TLS and ECDSA.For `secp256r1`, the resulting key will be seen as `prime256v1` by [`getEllipticCurveName()`](4201-xml-cryptokey-getellipticcurvename.md "Returns the name of the elliptic curve for an ECDSA key, if available."). |
| prime256v1 | This curve is another name for secp256r1 above. It is used in similar applications as a NIST standard. |
| secp384r1 | This NIST standardized curve is used for applications requiring higher security, such as digital signatures and key exchange. |
| secp521r1 | This NIST standardized curve provides even higher security and is used in applications that require strong cryptographic assurances. |

## Related links

**Related concepts**  

[Create an enveloping signature using a DSA key](4300-create-an-enveloping-signature-using-a-dsa-key.md "Use this example to digitally sign an XML document with a DSA key, producing an enveloping signature that guarantees the document has not been tampered with.")
