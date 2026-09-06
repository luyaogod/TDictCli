---
title: "Serialization option flags"
source: "fgl-topics/r_gws_XmlSerializer_option_flags.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > Serialization option flags"
type: "reference"
---

# Serialization option flags

> Serialization option flags for the xml.Serializer class.

| Flag | Description |
| --- | --- |
| `xml_ignoretimezone` | Defines that the serializer ignores the time zone information, during the marshalling and un-marshalling process of a BDL DATETIME data type.A value of zero means `FALSE`. The default is `FALSE`.Throws an exception in case of errors, and updates status with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number."). |
| `xml_ignoreunknownelements` | Force the XML serializer to ignore unexpected elements that are not defined in the XML schema. See [XML-to-BDL conversion options](../16_web-services/5047-xml-to-bdl-conversion-options.md "This topic describes the options that relax XML-to-BDL deserialization by allowing the serializer to ignore unexpected XML attributes or elements.").A value of zero means FALSE. The default is FALSE.Throws an exception in case of errors, and updates status with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number."). |
| `xml_ignoreunknownattributes` | Force the XML serializer to ignore unexpected attributes that are not defined in the XML schema. See [XML-to-BDL conversion options](../16_web-services/5047-xml-to-bdl-conversion-options.md "This topic describes the options that relax XML-to-BDL deserialization by allowing the serializer to ignore unexpected XML attributes or elements.").A value of zero means `FALSE`. The default is `FALSE`.Throws an exception in case of errors, and updates status with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number."). |
| `xml_usetypedefinition` | Defines whether the serializer must specify the type of data during serialization. This will add an "`xsi:type`" attribute to each XML data type.A value of zero means `FALSE`. The default is `FALSE`.Throws an exception in case of errors, and updates status with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number."). |
| `xml_useutctime` | Defines that BDL DATETIME data type is converted to UTC time during the serializer marshalling process. Be aware that this conversion does not take daylight saving time in winter and summer months into account, and therefore local time in some timezones may differ from UTC by one hour.A value of zero means `FALSE`. The default is `FALSE`.Throws an exception in case of errors, and updates status with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number."). |
| `xop_threshold` | When using the optimized serializer APIs, you can set a size that determines whether the `BYTE` is handled as a separate part (with an XML-Binary Optimized document created for it) or whether it is transmitted inline and encoded in base64 format.CALL xml.serializer.setOption("xop_threshold",5000)In this example, `BYTE` variables whose size is greater than 5000 bytes are handled as a separate part, otherwise they are handled inline and encoded in base64 format.By default, the size is zero (0), and all `BYTE` variables are handled as separate parts. |
| `xs_processcontents` | Defines the way to generate wildcard elements and attributes in XML schemas via the XML schema `processContents` tag. See Table 2Throws an exception in case of errors, and updates status with an [error code](4483-genero-bdl-errors.md "System error messages sorted by error number."). |

| Value | Description |
| --- | --- |
| 0 | No processContents tag will be generated (default) |
| 1 | Generation of `processContents="skip"`. |
| 2 | Generation of `processContents="lax"`. |
| 3 | Generation of `processContents="strict"`. |
