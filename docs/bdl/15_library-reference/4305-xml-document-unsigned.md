---
title: "XML document (unsigned)"
source: "fgl-topics/c_gws_XmlSignature_example_mydocument_xml.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML security classes > The Signature class > Examples > Sample: XML and signed documents > XML document (unsigned)"
type: "concept"
---

# XML document (unsigned)

> Sample content provided for the purpose of testing examples.

Use the sample content provided here for the purpose of testing code in [Examples](4297-examples.md "xml.Signature usage examples.") for signing documents with different signature keys.

Copy the contents to a file named "MyDocument.xml" in a directory where you need to sign a sample
XML document.

```
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<?target data?>
<MyElement>
  <MySecret xmlns:dec="http://tempuri.org">This is my first signed XML node...</MySecret>
  <MyLogin xml:id="log">This is my second signed XML node.</MyLogin>
  <MyCode xml:id="code">1234</MyCode>  
</MyElement>
```
