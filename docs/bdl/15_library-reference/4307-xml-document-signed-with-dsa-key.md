---
title: "XML document (signed with DSA key)"
source: "fgl-topics/c_gws_XmlSignature_example_signed_dsa_xml_document.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Sample: XML and signed documents > XML document (signed with DSA key)"
type: "concept"
---

# XML document (signed with DSA key)

> The sample content provided here is for the purpose of demonstrating what a signed document (a document signed with a DSA key) might look like.

This is the output produced by running the [Create an enveloping signature using a DSA key](4300-create-an-enveloping-signature-using-a-dsa-key.md "Use this example to digitally sign an XML document with a DSA key, producing an enveloping signature that guarantees the document has not been tampered with.") example with the [XML document (unsigned)](4305-xml-document-unsigned.md "Sample content provided for the purpose of testing examples.") document as input.

```
<?xml version="1.0" encoding="UTF-8"?>
<dsig:Signature xmlns:dsig="http://www.w3.org/2000/09/xmldsig#">
	<dsig:SignedInfo>
		<dsig:CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315"/>
		<dsig:SignatureMethod Algorithm="http://www.w3.org/2009/xmldsig11#dsa-sha256"/>
		<dsig:Reference URI="#data" Type="http://www.w3.org/2000/09/xmldsig#Object">
			<dsig:Transforms>
				<dsig:Transform Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/>
			</dsig:Transforms>
			<dsig:DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256"/>
			<dsig:DigestValue>nXJaZMTnZYcbMp+pfFjZWRyKk+EkSte8YBty8+K9hwQ=</dsig:DigestValue>
		</dsig:Reference>
	</dsig:SignedInfo>
	<dsig:SignatureValue>z0F4fi3Iebot1hTHeZnQWmr3X5g+z3Tvgdb+D6QsOilUZE5B4Csw1r7YidpA9S8cIitLb9Y+eHhTb7S44t9nNw==</dsig:SignatureValue>
	<dsig:Object Id="data">
		<MyElement>
			<MySecret xmlns:dec="http://tempuri.org">This is my first signed XML node...</MySecret>
			<MyLogin xml:id="log">This is my second signed XML node.</MyLogin>
			<MyCode xml:id="code">1234</MyCode>
		</MyElement>
	</dsig:Object>
</dsig:Signature>
```

## Related links

**Related concepts**  

[XML Signature concepts](4287-xml-signature-concepts.md "The purpose of a signature is to guarantee the integrity of a XML document, that it was not altered, and that it still contains the same data as when it was created. An additional purpose of a signature is to authenticate the author of the document. There are different ways to achieve this guarantee.")

**Related reference**  

[Supported kind of keys](4222-supported-kind-of-keys.md "Types of keys supported by the xml.CryptoKey class.")
