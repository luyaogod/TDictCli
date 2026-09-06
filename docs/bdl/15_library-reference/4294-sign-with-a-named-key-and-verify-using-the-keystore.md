---
title: "Sign with a named key and verify using the keystore"
source: "fgl-topics/c_gws_XmlSignature_concepts_7.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > XML Signature concepts > Sign with a named key and verify using the keystore"
type: "concept"
---

# Sign with a named key and verify using the keystore

> Use if the sender and the receiver exchange multiple XML documents signed with different keys.

## How to sign

1. Create a HMAC, RSA or DSA key with the [constructor](4195-xml-cryptokey-create.md "Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.") of the CryptoKey
   class.
2. [Set](4221-xml-cryptokey-setkey.md "Defines the value of a HMAC or Symmetric key.") the HMAC key or
   [load](4212-xml-cryptokey-loadpem.md "Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format.") the RSA or DSA
   key in the CryptoKey object.
3. Set the [KeyName](4224-cryptokey-features.md "Features of the xml.CryptoKey class.")
   [feature](4220-xml-cryptokey-setfeature.md "Sets or resets the value of a feature for a CryptoKey object.") with
   the name identifying the key.
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

1. Create a HMAC, RSA or DSA key with the [constructor](4195-xml-cryptokey-create.md "Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.") of the
   CryptoKey.
2. [Set](4221-xml-cryptokey-setkey.md "Defines the value of a HMAC or Symmetric key.") the HMAC key or
   [load](4212-xml-cryptokey-loadpem.md "Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format.") the RSA or DSA
   key in the CryptoKey object.
3. Set the [KeyName](4224-cryptokey-features.md "Features of the xml.CryptoKey class.")
   [feature](4220-xml-cryptokey-setfeature.md "Sets or resets the value of a feature for a CryptoKey object.") with
   the name identifying the key.
4. [Register](4338-xml-keystore-addkey.md "Registers in the keystore the given key by name for the application.") the key to
   be used by key name for any signature verification.
5. Create a signature with the [constructor](4261-xml-signature-createfromnode.md "Constructor of a new Signature object from a XML Signature node, based on the XML-Signature specification.")
   of the Signature class and from a XML signature node obtained after the above compute
   operation.
6. [Verify](4285-xml-signature-verify.md "Verifies that all references in this signature object have not changed.") the
   signature validity.

> **Note:**
>
> It is recommended that steps 1 to 4 are done once at application start-up for each
> key used in the application. Steps 5 - 6 can then be executed quickly for any XML signature to be
> checked.

## Related links

**Related concepts**  

[Examples](4297-examples.md "xml.Signature usage examples.")
