---
title: "Sign with the originator private RSA or DSA key, and verify with trusted X509 certificates"
source: "fgl-topics/c_gws_XmlSignature_concepts_5.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > XML Signature concepts > Sign with the originator private RSA or DSA key, and verify with trusted X509 certificates"
type: "concept"
---

# Sign with the originator private RSA or DSA key, and verify with trusted X509 certificates

> Use if the sender of the XML document adds a X509 certificate that was signed by another trusted X509 certificate.

Only the originator can sign a message with this specific pair of keys.
Any other peer needs the corresponding public key and does not have access to the private key.

## How to sign

1. Create a RSA or DSA key with the [constructor](4195-xml-cryptokey-create.md "Initializes a xml.CryptoKey object. Constructor of an empty CryptoKey object based on a URL.") of the CryptoKey
   class.
2. [Load](4212-xml-cryptokey-loadpem.md "Loads an asymmetric DSA key, an asymmetric RSA key, an asymmetric ECDSA key or Diffie-Hellman parameters from a file in PEM format.") the RSA or
   DSA private key into the CryptoKey object.
3. Create a X509 certificate with the [constructor](4236-xml-cryptox509-create.md "Constructor of an empty CryptoX509 object.") of the
   CryptoX509 class.
4. [Load](4246-xml-cryptox509-loadpem.md "Loads a X509 certificate from a file in PEM format.") the X509
   certificate associated to the RSA or DSA private key into the CryptoKey object.
5. Create a blank signature with the [constructor](4260-xml-signature-create.md "Constructor of a blank Signature object.") of the Signature
   class.
6. [Assign](4281-xml-signature-setkey.md "Defines the key used for signing or validation.") the
   CryptoKey object to the Signature object.
7. [Assign](4279-xml-signature-setcertificate.md "Defines the X509 certificate to be added to the signature object when signing a document.") the CryptoX509 object to the Signature object.
8. [Create](4263-xml-signature-createreference.md "Creates a new reference that will be signed with the compute() method") one or more references to be signed.
9. [Compute](4259-xml-signature-compute.md "Computes the signature of all references set in this Signature object.") the
   signature.
10. [Retrieve](4265-xml-signature-getdocument.md "Returns a new DomDocument object representing the signature in XML.")
    the XML signature document from the Signature object.

## How to verify

1. Create a X509 certificate with the [constructor](4236-xml-cryptox509-create.md "Constructor of an empty CryptoX509 object.") of the
   CryptoX509 class.
2. [Load](4246-xml-cryptox509-loadpem.md "Loads a X509 certificate from a file in PEM format.") the X509
   certificate that was used to sign the originator X509 certificate into the CryptoX509 object.
3. [Add](4339-xml-keystore-addtrustedcertificate.md "Registers in the keystore the given X509 certificate as a trusted certificate for the application.") the X509 certificate as trusted certificate to the application.
4. Create a signature with the [constructor](4261-xml-signature-createfromnode.md "Constructor of a new Signature object from a XML Signature node, based on the XML-Signature specification.")
   of the Signature class and from a XML signature node obtained after the above compute
   operation.
5. [Verify](4285-xml-signature-verify.md "Verifies that all references in this signature object have not changed.") the
   signature validity.

> **Note:**
>
> Point 1 to 3 can be omitted if entry [`xml.application.calist`](../16_web-services/4916-fglprofile-entries-for-web-services.md) has been set in FGLPROFILE file with the trusted
> certificate.

> **Note:**
>
> There is no key nor certificate to set in the Signature object during
> validation.

## Related links

**Related concepts**  

[Examples](4297-examples.md "xml.Signature usage examples.")
