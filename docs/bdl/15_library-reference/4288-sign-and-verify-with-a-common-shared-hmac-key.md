---
title: "Sign and verify with a common shared HMAC key"
source: "fgl-topics/c_gws_XmlSignature_concepts_HMAC.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > XML Signature concepts > Sign and verify with a common shared HMAC key"
type: "concept"
---

# Sign and verify with a common shared HMAC key

> Use if the sender of the XML document and the receiver share a common secret key.

## How to sign

1. Create a HMAC key with the [constructor](4195-xml-cryptokey-create.md "Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.") of the CryptoKey class.
2. [Set](4221-xml-cryptokey-setkey.md "Defines the value of a HMAC or Symmetric key.") or [load](4208-xml-cryptokey-loadbin.md "Loads a symmetric or HMAC key from a file in raw format.") the common shared key
   value in the CryptoKey object.
3. Create a blank signature with the [constructor](4260-xml-signature-create.md "Constructor of a blank Signature object.") of the Signature
   class.
4. [Assign](4281-xml-signature-setkey.md "Defines the key used for signing or validation.") the
   CryptoKey object to the Signature object.
5. [Create](4263-xml-signature-createreference.md "Creates a new reference that will be signed with the compute() method") one or more references to be signed.
6. [Compute](4259-xml-signature-compute.md "Computes the signature of all references set in this Signature object.") the
   signature.
7. [Retrieve](4265-xml-signature-getdocument.md "Returns a new DomDocument object representing the signature in XML.")
   the XML signature document from the Signature object.

## How to verify

1. Create a HMAC key with the [constructor](4195-xml-cryptokey-create.md "Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.") of the CryptoKey.
2. [Set](4221-xml-cryptokey-setkey.md "Defines the value of a HMAC or Symmetric key.") or [load](4208-xml-cryptokey-loadbin.md "Loads a symmetric or HMAC key from a file in raw format.") the common shared key
   value in the CryptoKey object.
3. Create a signature with the [constructor](4261-xml-signature-createfromnode.md "Constructor of a new Signature object from a XML Signature node, based on the XML-Signature specification.")
   of the Signature class and from a XML signature node obtained after the above compute
   operation.
4. [Assign](4281-xml-signature-setkey.md "Defines the key used for signing or validation.") the
   CryptoKey object to the Signature object.
5. [Verify](4285-xml-signature-verify.md "Verifies that all references in this signature object have not changed.") the
   signature validity.

## Related links

**Related concepts**  

[Examples](4297-examples.md "xml.Signature usage examples.")
