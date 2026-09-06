---
title: "XML document (signed with HMAC key)"
source: "fgl-topics/c_gws_XmlSignature_example_signed_hmac_xml_document.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Sample: XML and signed documents > XML document (signed with HMAC key)"
type: "concept"
---

# XML document (signed with HMAC key)

> The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a HMAC key) might look like.

This is the output produced by running the [Create a detached signature using a HMAC key](4298-create-a-detached-signature-using-a-hmac-key.md "In the example, an XML document (\"MyDocument.xml\") is loaded and signed with a HMAC key.") example with the [XML document (unsigned)](4305-xml-document-unsigned.md "Sample content provided for the purpose of testing examples.") document as input.

```
<?xml version="1.0" encoding="UTF-8"?>
<dsig:Signature xmlns:dsig="http://www.w3.org/2000/09/xmldsig#">
  <dsig:SignedInfo>
    <dsig:CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315"/>
    <dsig:SignatureMethod Algorithm="http://www.w3.org/2000/09/xmldsig#hmac-sha1"/>
    <dsig:Reference URI="#code">
      <dsig:Transforms>
        <dsig:Transform Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/>
      </dsig:Transforms>
      <dsig:DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1"/>
      <dsig:DigestValue>DCDNxibEA3AHpFMtzvj6hxd7p5A=</dsig:DigestValue>
    </dsig:Reference>
  </dsig:SignedInfo>
  <dsig:SignatureValue>ZMU4rBeDJt/nHZmDglm4wlUroJM=</dsig:SignatureValue>
</dsig:Signature>
```

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")

**Related reference**  

[Supported kind of keys](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
