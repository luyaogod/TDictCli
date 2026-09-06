---
title: "xml.Signature methods"
source: "fgl-topics/c_gws_XmlSignature_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods"
type: "concept"
---

# xml.Signature methods

> Methods for the xml.Signature class.

| Name | Description |
| --- | --- |
| xml.Signature.Create() RETURNS xml.Signature | Constructor of a blank Signature object. |
| xml.Signature.CreateFromNode( node xml.DomNode ) RETURNS xml.Signature | Constructor of a new Signature object from a XML Signature node, based on the XML-Signature specification. |

| Name | Description |
| --- | --- |
| RetrieveObjectDataListFromSignatureNode( signNode xml.DomNode, index INTEGER ) RETURNS xml.DomNodeList | Returns a DomNodeList containing all embedded XML nodes related to the signature object |

> **Note:**
>
> In addition to this class method categorized under Object Access, there are also object
> methods. These are listed in Table 9.

| Name | Description |
| --- | --- |
| setCertificate( cert xml.CryptoX509 ) | Defines the X509 certificate to be added to the signature object when signing a document. |
| setKey( key xml.CryptoKey ) | Defines the key used for signing or validation. |

| Name | Description |
| --- | --- |
| setCanonicalization( url STRING ) | Sets the canonicalization method to use for the signature. |
| setID( id STRING ) | Sets an ID value for the signature. |

| Name | Description |
| --- | --- |
| getCanonicalization() RETURNS STRING | Returns the canonicalization identifier of the signature. |
| getDocument() RETURNS xml.DomDocument | Returns a new DomDocument object representing the signature in XML. |
| getID() RETURNS STRING | Returns the ID value of the signature. |
| getSignatureMethod() RETURNS STRING | Returns the algorithm method of the signature. |
| getType() RETURNS STRING | Returns a string with the type of the signature object. |

| Name | Description |
| --- | --- |
| appendReferenceTransformation( referenceIndex INTEGER, method STRING [, args ] ) | Appends transformations related to the specified reference index. |
| createReference( uri STRING, digest STRING ) RETURNS INTEGER | Creates a new reference that will be signed with the `compute()` method |
| setReferenceID( index INTEGER, id STRING ) | Sets an ID for the signature reference in the specified signature object. |

| Name | Description |
| --- | --- |
| getReferenceCount() RETURNS INTEGER | Returns the number of references in this Signature object. |
| getReferenceDigest( index INTEGER ) RETURNS STRING | Returns the digest algorithm identifier of the reference. |
| getReferenceURI( index INTEGER ) RETURNS STRING | Returns the URI of the reference in this signature object. |
| getReferenceID( index INTEGER ) RETURNS STRING | Returns the ID value of the reference in this signature object. |
| getReferenceTransformation( referenceIndex INTEGER, index INTEGER ) RETURNS STRING | Gets the transformation identifier related to the reference of index *referenceIndex*. |
| getReferenceTransformationCount( referenceIndex INTEGER ) RETURNS INTEGER | Returns the number of transformations referenced in this signature object. |

| Name | Description |
| --- | --- |
| appendObjectData( index INTEGER, node xml.DomNode ) | Appends a copy of a XML DomNode to the signature object index. |
| createObject() RETURNS INTEGER | Creates a new object that will embed additional XML nodes. |
| setObjectID( index INTEGER, id STRING ) | Sets an ID for the signature object. |

| Name | Description |
| --- | --- |
| getObjectCount() RETURNS INTEGER | Returns the number of objects in this Signature object. |
| getObjectId( index INTEGER ) RETURNS STRING | Returns the ID value of the signature object. |

> **Note:**
>
> In addition to these object methods categorized under Object Access, there is also a class
> method. It is listed in Table 2.

| Name | Description |
| --- | --- |
| compute( doc xml.DomDocument ) | Computes the signature of all references set in this Signature object. |
| SignString( key xml.CryptoKey, str STRING ) RETURNS STRING | Sign the passed string according to the specified key. |
| verify( doc xml.DomDocument ) RETURNS INTEGER | Verifies that all references in this signature object have not changed. |
| VerifyString( key xml.CryptoKey, originalStr STRING, signature STRING ) RETURNS INTEGER | Verify the signature is consistent with the given key and the original message. |
