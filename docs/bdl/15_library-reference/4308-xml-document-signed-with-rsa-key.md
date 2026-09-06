---
title: "XML document (signed with RSA key)"
source: "fgl-topics/c_gws_XmlSignature_example_signed_rsa_xml_document.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Sample: XML and signed documents > XML document (signed with RSA key)"
type: "concept"
---

# XML document (signed with RSA key)

> The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a RSA key) might look like.

This is the output produced by running the [Create an enveloped signature using a RSA key](4302-create-an-enveloped-signature-using-a-rsa-key.md "In this code sample, a XML document (\"MyDocument.xml\") is loaded and signed with a RSA key.") example with the [XML document (unsigned)](4305-xml-document-unsigned.md "Sample content provided for the purpose of testing examples.") document as input.

```
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<?target data?>
<MyElement>
  <MySecret xmlns:dec="http://tempuri.org">This is my first signed XML node...</MySecret>
  <MyLogin xml:id="log">This is my second signed XML node.</MyLogin>
  <MyCode xml:id="code">1234</MyCode>
  <dsig:Signature xmlns:dsig="http://www.w3.org/2000/09/xmldsig#">
    <dsig:SignedInfo>
      <dsig:CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315"/>
      <dsig:SignatureMethod Algorithm="http://www.w3.org/2000/09/xmldsig#rsa-sha1"/>
      <dsig:Reference URI="#code">
        <dsig:Transforms>
          <dsig:Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/>
          <dsig:Transform Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/>
        </dsig:Transforms>
        <dsig:DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1"/>
        <dsig:DigestValue>DCDNxibEA3AHpFMtzvj6hxd7p5A=</dsig:DigestValue>
      </dsig:Reference>
    </dsig:SignedInfo>
    <dsig:SignatureValue>jocgUrZPKR8jvery4gG4V34qx7/yxODSPJq//iS3Q5Ps7lPADNBEVK4Y50HIdrkodcYLZjBkvuGMT89nTeT24W/Dw/XEeMWXRmy/Mj1/rza8JMaP46F+2MZ6tlGWlyA2tRZNExe5TPA8Wo6jTSN3KX3aLoLkwRsLBt50Zr8zz8xFt9dZciNWnsD6y/UgQzNYfLovMw54AHGk+5FzRWMgwtTseISWxSF+9zsgiQStrrXzy1SaRycQTAjz4PF6HebGWJcECLa+r/iLtigbTmgL3Mj7mkmw90M3mNncqZKBFmjNxTZCPiMQHbSvTgOBe8REwCrclHJkyYP14NsxEg6LZQ==</dsig:SignatureValue>
  </dsig:Signature>
</MyElement>
```

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")

**Related reference**  

[Supported kind of keys](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
