---
title: "Sign with the originator private RSA or DSA key, and verify with the originator X509 certificate associated to the private RSA or DSA key"
source: "fgl-topics/c_gws_XmlSignature_concepts_4.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > XML Signature concepts > Sign with the originator private RSA or DSA key, and verify with the originator X509 certificate associated to the private RSA or DSA key"
type: "concept"
---

# Sign with the originator private RSA or DSA key, and verify with the originator X509 certificate associated to the private RSA or DSA key

> Use if the receiver of the XML document has the X509 certificate associated to the RSA or DSA private key.

Only the originator can sign a message with this specific pair of keys.
Any other peer needs the corresponding public key and does not have access to the private key.

## How to sign

1. Create a RSA or DSA key with the [constructor](4195-xml-cryptokey-create.md "Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.") of the CryptoKey
   class.
2. [Load](4212-xml-cryptokey-loadpem.md "Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format.") the RSA or
   DSA private key into the CryptoKey object.
3. Create a blank signature with the [constructor](4260-xml-signature-create.md "Constructor of a blank Signature object.") of the Signature
   class.
4. [Assign](4281-xml-signature-setkey.md "Defines the key used for signing or validation.") the
   CryptoKey object to the Signature object.
5. [Create](4260-xml-signature-create.md "Constructor of a blank Signature object.") one or more
   references to be signed.
6. [Compute](4259-xml-signature-compute.md "Computes the signature of all references set in this Signature object.") the
   signature.
7. [Retrieve](4265-xml-signature-getdocument.md "Returns a new DomDocument object representing the signature in XML.")
   the XML signature document from the Signature object.

## How to verify

1. Create a X509 certificate with the [constructor](4236-xml-cryptox509-create.md "Constructor of an empty CryptoX509 object.") of the
   CryptoX509 class.
2. [Load](4246-xml-cryptox509-loadpem.md "Loads a X509 certificate from a file in PEM format.") the X509
   certificate into the CryptoKey object.
3. Create the RSA or DSA [public
   key](4238-xml-cryptox509-createpublickey.md "Creates a new public CryptoKey object for the given URL.") from the X509 certificate of the CryptoX509 object.
4. Create a signature with the [constructor](4261-xml-signature-createfromnode.md "Constructor of a new Signature object from a XML Signature node, based on the XML-Signature specification.")
   of the Signature class and from a XML signature node obtained after the above compute
   operation.
5. [Assign](4281-xml-signature-setkey.md "Defines the key used for signing or validation.") the
   CryptoKey object containing the public key to the Signature object.
6. [Verify](4285-xml-signature-verify.md "Verifies that all references in this signature object have not changed.") the
   signature validity.

## Related links

**Related concepts**  

[Examples](4297-examples.md "xml.Signature usage examples.")
