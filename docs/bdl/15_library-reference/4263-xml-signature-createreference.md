---
title: "xml.Signature.createReference"
source: "fgl-topics/c_gws_XmlSignature_createReference.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Signature methods > xml.Signature.createReference"
type: "concept"
---

# xml.Signature.createReference

> Creates a new reference that will be signed with the compute() method

## Syntax

```
createReference(
   uri STRING,
   digest STRING )
  RETURNS INTEGER
```

1. uri represents the data to be
   signed.
2. digest defines a URL as [identifier](4295-digest-identifier.md) for the hash algorithm.

## Usage

The returned value represents the index for any further manipulation of this reference.

The uri can be:

- An absolute URL such as http://, https://, tcp://, tcps://, file:/// and
  alias:// (see [FGLPROFILE Configuration](../16_web-services/4915-web-services-fglprofile-configuration.md "The configuration for the Genero Web Services is defined from entries in the FGLPROFILE file.")
  for more details about URL mapping with aliases), and where the data can be a XML document or any
  kind of data such as images or HTML pages.
- NULL to sign the whole document, but only one NULL is allowed in the entire signature.
- An empty URI (`uri=""`). In this case, the uri attribute
  identifies the root element of the XML document containing the reference.
- A fragment like `#tobesigned`. Note that a DOM node fragment is identified via the
  value of an attribute of type ID such as `xml:id` or any attribute whose type was
  changed to ID with [`xml.DomNode.setIdAttribute()`](4077-xml-domnode-setidattribute.md "Set the XML Attribute of given name to be of type ID. Declare (or undeclare) the ID as user-determined.") or [`xml.DomNode.setIdAttributeNS()`](4078-xml-domnode-setidattributens.md "Set the namespace-qualified XML Attribute of given name and namespace to be of type ID. Declare (or undeclare) the ID as user-determined.").

If the uri attribute is omitted altogether, the receiving application is
expected to know the identity of the object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[xml.Signature.compute](4259-xml-signature-compute.md "Computes the signature of all references set in this Signature object.")
