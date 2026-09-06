---
title: "XML Signature concepts"
source: "fgl-topics/c_gws_XmlSignature_concepts.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > XML Signature concepts"
type: "concept"
---

# XML Signature concepts

> The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.


## Child topics

- [Sign and verify with a common shared HMAC key](4288-sign-and-verify-with-a-common-shared-hmac-key.md): Use if the sender of the XML document and the receiver share a common secret key.
- [Sign with the originator private RSA or DSA key, and verify with the originator public RSA or DSA key](4289-sign-with-the-originator-private-rsa-or-dsa-key-and-verify-w.md): Use if the receiver of the XML document has the RSA or DSA public key of the sender.
- [Sign with the originator private RSA or DSA key, and verify with a RSA or DSA retrieval method](4290-sign-with-the-originator-private-rsa-or-dsa-key-and-verify-w.md): Use if the sender of the XML document provides the public RSA or DSA key in XML form (and via http, tcp or a file protocol).
- [Sign with the originator private RSA or DSA key, and verify with the originator X509 certificate associated to the private RSA or DSA key](4291-sign-with-the-originator-private-rsa-or-dsa-key-and-verify-w.md): Use if the receiver of the XML document has the X509 certificate associated to the RSA or DSA private key.
- [Sign with the originator private RSA or DSA key, and verify with trusted X509 certificates](4292-sign-with-the-originator-private-rsa-or-dsa-key-and-verify-w.md): Use if the sender of the XML document adds a X509 certificate that was signed by another trusted X509 certificate.
- [Sign with the originator private RSA or DSA key, and verify with a X509 certificate retrieval method and trusted X509 certificates](4293-sign-with-the-originator-private-rsa-or-dsa-key-and-verify-w.md): Use if the sender of the XML document adds a X509 retrieval method that was signed by another trusted X509 certificate.
- [Sign with a named key and verify using the keystore](4294-sign-with-a-named-key-and-verify-using-the-keystore.md): Use if the sender and the receiver exchange multiple XML documents signed with different keys.
