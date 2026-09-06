---
title: "Sign with the originator private RSA or DSA key, and verify with a RSA or DSA retrieval method"
source: "fgl-topics/c_gws_XmlSignature_concepts_3.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > XML Signature concepts > Sign with the originator private RSA or DSA key, and verify with a RSA or DSA retrieval method"
type: "concept"
---

# Sign with the originator private RSA or DSA key, and verify with a RSA or DSA retrieval method

> Use if the sender of the XML document provides the public RSA or DSA key in XML form (and via http, tcp or a file protocol).

Only the originator can sign a message with this specific pair of keys.
Any other peer needs the corresponding public key and does not have access to the private key.

## How to sign

1. Create a RSA or DSA key with the [constructor](4195-xml-cryptokey-create.md "Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.") of the CryptoKey
   class.
2. [Load](4212-xml-cryptokey-loadpem.md "Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format.") the RSA or
   DSA private key into the CryptoKey object.
3. Set the RetrievalMethod [feature](4220-xml-cryptokey-setfeature.md "Sets or resets the value of a feature for a CryptoKey object.") on the
   CryptoKey object with the URL where the XML form of the public RSA or DSA key is available.
4. Create a blank signature with the [constructor](4260-xml-signature-create.md "Constructor of a blank Signature object.") of the Signature
   class.
5. [Assign](4281-xml-signature-setkey.md "Defines the key used for signing or validation.") the
   CryptoKey object to the Signature object.
6. [Create](4263-xml-signature-createreference.md "Creates a new reference that will be signed with the compute() method") one or more references to be signed.
7. [Compute](4259-xml-signature-compute.md "Computes the signature of all references set in this Signature object.") the
   signature.
8. [Retrieve](4265-xml-signature-getdocument.md "Returns a new DomDocument object representing the signature in XML.")
   the XML signature document from the Signature object.

## How to verify

1. Create a signature with the [constructor](4261-xml-signature-createfromnode.md "Constructor of a new Signature object from a XML Signature node, based on the XML-Signature specification.")
   of the Signature class and from a XML signature node obtained after the above compute
   operation.
2. [Verify](4285-xml-signature-verify.md "Verifies that all references in this signature object have not changed.") the
   signature validity.
   > **Note:**
   >
   > There is no key nor certificate to set in the Signature object during
   > validation.

## Related links

**Related concepts**  

[Examples](4297-examples.md "xml.Signature usage examples.")
