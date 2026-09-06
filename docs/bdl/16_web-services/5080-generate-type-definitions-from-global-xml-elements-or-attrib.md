---
title: "Generate TYPE definitions from global XML elements or attributes"
source: "fgl-topics/c_gws_fglwsdl_finlinetypes.html"
breadcrumb: "Web services > Reference > Using fglwsdl to generate code from WSDL or XSD schemas > Generate TYPE definitions from global XML elements or attributes"
type: "concept"
description: "If a WSDL or a XSD has global XML elements or attributes defined with an inlined type, the -fInlineTypes option of fglwsdl generates a TYPE definition representing that inline type, using the original ..."
---

# Generate TYPE definitions from global XML elements or attributes

If a WSDL or a XSD has global XML elements or attributes defined with an inlined type, the
-fInlineTypes option of fglwsdl generates a
`TYPE` definition representing that inline type, using the original WSDL/XSD name of
the element or attribute, concatenated with the string 'GlobalAttributeType' or
'GlobalElementType'.

For example, when using fglwsdl -fInlineTypes, the following schema:

```
<xs:element name="getAlertListRequestFlow">
  <xs:complexType>
    <xs:sequence>
      <xs:element name="getAlertListRequest" type="amp:getAlertListRequest" />
    </xs:sequence>
  </xs:complexType>
</xs:element>
```

will produce:

```
TYPE tgetAlertListRequestFlowGlobalElementType RECORD
         ATTRIBUTES(XMLSequence)
   getAlertListRequest tgetAlertListRequest
      ATTRIBUTES(XMLName="getAlertListRequest")
END RECORD
DEFINE getAlertListRequestFlow tgetAlertListRequestFlowGlobalElementType
          ATTRIBUTES(XMLName="getAlertListRequestFlow")
```

Instead of:

```
DEFINE getAlertListRequestFlow RECORD
          ATTRIBUTES(XMLName="getAlertListRequestFlow",XMLSequence)
   getAlertListRequest tgetAlertListRequest
      ATTRIBUTES(XMLName="getAlertListRequest")
END RECORD
```
