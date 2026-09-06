---
title: "Digest identifier"
source: "fgl-topics/r_gws_XmlSignature_digest_identifier.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Digest identifier"
type: "reference"
description: "Table 1. Digest identifiers Identifier Description http://www.w3.org/2000/09/xmldsig#sha1 See specification for details. Computes the digest of the reference set with createReference() , by applying a ..."
---

# Digest identifier

| Identifier | Description |
| --- | --- |
| http://www.w3.org/2000/09/xmldsig#sha1See [specification](http://www.w3.org/TR/xmldsig-core/#sec-MessageDigests) for details. | Computes the digest of the reference set with `createReference()`, by applying a hash operation using a SHA algorithm of 160 bits.**Note:**It is the only digest algorithm recommended by the W3C. |
| http://www.w3.org/2001/04/xmlenc#sha512See [specification](http://www.w3.org/TR/xmldsig-core/#sec-MessageDigests) for details. | Computes the digest of the reference set with `createReference()`, by applying a hash operation using a SHA algorithm of 512 bits. |
| http://www.w3.org/2001/04/xmldsig-more#sha384See [specification](http://www.w3.org/TR/xmldsig-core/#sec-MessageDigests) for details. | Computes the digest of the reference set with `createReference()`, by applying a hash operation using a SHA algorithm of 384 bits. |
| http://www.w3.org/2001/04/xmlenc#sha256See [specification](http://www.w3.org/TR/xmldsig-core/#sec-MessageDigests) for details. | Computes the digest of the reference set with `createReference()`, by applying a hash operation using a SHA algorithm of 256 bits. |
| http://www.w3.org/2001/04/xmldsig-more#sha224See [specification](http://www.w3.org/TR/xmldsig-core/#sec-MessageDigests) for details. | Computes the digest of the reference set with `createReference()`, by applying a hash operation using a SHA algorithm of 224 bits. |
| http://www.w3.org/2001/04/xmldsig-more#md5See [specification](http://www.w3.org/TR/xmldsig-core/#sec-MessageDigests) for details. | Computes the digest of the reference set with `createReference()`, by applying a hash operation using a MD5 algorithm. |
| http://www.w3.org/2001/04/xmlenc#ripemd160See [specification](http://www.w3.org/TR/xmldsig-core/#sec-MessageDigests) for details. | Computes the digest of the reference set with `createReference()`, by applying a hash operation using a RIPEMD algorithm. |

## Related links

**Related concepts**  

[xml.Signature.createReference](4263-xml-signature-createreference.md "Creates a new reference that will be signed with the compute() method")
