---
title: "XMLTypenamespace"
source: "fgl-topics/c_gws_XML_attribute_XMLTypenamespace.html"
breadcrumb: "Web services > Reference > XML serialization rules and customization > XML serialization attributes > XMLTypenamespace"
type: "concept"
description: "Force the XML type namespace of a variable by adding xsi:type at serialization or by checking xsi:type at de-serialization. Important: The attribute must be used with the XMLType attribute; otherwise, ..."
---

# XMLTypenamespace

Force the XML type namespace of a variable by adding `xsi:type` at serialization
or by checking `xsi:type` at de-serialization.

> **Important:**
>
> The attribute must be used with the [XMLType](5039-xmltype.md)
> attribute; otherwise, the generated source code will not compile..

## Example

```
DEFINE myVar RECORD ATTRIBUTES(XMLName="Root",
                               XMLNamespace="http://tempuri.org")
  val1  FLOAT   ATTRIBUTES(XMLName="Val1"),
  val2  INTEGER ATTRIBUTES(XMLName="Val2",
                           XMLType="MyRecord",
                           XMLTypenamespace="http://mynamespace.org")
END RECORD
```

```
<fjs1:Root xmlns:fjs1="http://tempuri.org">
  <fjs1:Val1>0.5</fjs1:Val1>
  <fjs1:Val2 xmlns:fjs2="http://mynamespace.org" 
   xsi:type="fjs2:MyRecord">-18547</fjs1:Val2>
</fjs1:Root>
```
