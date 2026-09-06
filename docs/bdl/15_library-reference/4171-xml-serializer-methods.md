---
title: "xml.Serializer methods"
source: "fgl-topics/c_gws_XmlSerializer_methods.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods"
type: "concept"
---

# xml.Serializer methods

> Methods for the xml.Serializer class.

| Name | Description |
| --- | --- |
| xml.Serializer.CreateXmlSchemas( var fgl-type, schemas RECORD ) | Creates XML schemas corresponding to the given variable, and fills a dynamic array with xml.DomDocument objects each representing a XML schema. |
| xml.Serializer.DomToStax( node xml.DomNode, stax xml.StaxWriter ) | Serializes a XML DomNode object to a StaxWriter object. |
| xml.Serializer.DomToVariable( node xml.DomNode, var fgl-type ) | Serializes a XML element node into a BDL variable using a DomNode object. |
| xml.Serializer.GetOption( str STRING ) RETURNS INTEGER | Gets a global option value from the serializer engine. |
| xml.Serializer.OptimizedDomToVariable( node xml.DomNode, var fgl-type, xopTable RECORD ) | Serializes a XML element node into a BDL variable using a DomNode object. |
| xml.Serializer.OptimizedSoapSection5ToVariable( node xml.DomNode, var fgl-type, xopTable RECORD) | Serializes an XML element node into a BDL variable in Soap Section 5 encoding. |
| xml.Serializer.OptimizedStaxToVariable( stax xml.StaxReader, var fgl-type, xopTable RECORD ) | Serializes an XML element node into a BDL variable using a StaxReader object. |
| xml.Serializer.OptimizedVariableToDom( var fgl-type, node xml.DomNode, xopTable RECORD ) | Serializes a BDL variable into a XML element node using a DomNode object. |
| xml.Serializer.OptimizedVariableToSoapSection5( var fgl-type, node xml.DomNode, xopTable RECORD) | Serializes a BDL variable into a XML element node in Soap Section 5 encoding. |
| xml.Serializer.OptimizedVariableToStax( var fgl-type, stax xml.StaxWriter, xopTable RECORD ) | Serializes a BDL variable into a XML element node using a StaxWriter object. |
| xml.Serializer.SetOption( optionName STRING, optionValue INTEGER ) | Sets a global option value for the serializer engine |
| xml.Serializer.SoapSection5ToVariable( node xml.DomNode, var fgl-type ) | Serializes an XML element node into a BDL variable in Soap Section 5 encoding. |
| xml.Serializer.StaxToVariable( stax xml.StaxReader, var fgl-type ) | Serializes an XML element node into a BDL variable using a StaxReader object. |
| xml.Serializer.StaxToDom( stax xml.StaxReader, node xml.DomNode ) | Serializes an XML element node into a DomNode object using a StaxReader object. |
| xml.Serializer.VariableToDom( var fgl-type, node xml.DomNode ) | Serializes a BDL variable into a XML element node using a DomNode object. |
| xml.Serializer.VariableToSoapSection5( var fgl-type, node xml.DomNode ) | Serializes a BDL variable into a XML element node in Soap Section 5 encoding. |
| xml.Serializer.VariableToStax( var fgl-type, stax xml.StaxWriter ) | Serializes a BDL variable into a XML element node using a StaxWriter object. |

## Child topics

- [xml.Serializer.OptimizedDomToVariable](4176-xml-serializer-optimizeddomtovariable.md): Serializes a XML element node into a BDL variable using a DomNode object.
- [xml.Serializer.OptimizedSoapSection5ToVariable](4177-xml-serializer-optimizedsoapsection5tovariable.md): Serializes an XML element node into a BDL variable in Soap Section 5 encoding.
- [xml.Serializer.OptimizedStaxToVariable](4178-xml-serializer-optimizedstaxtovariable.md): Serializes an XML element node into a BDL variable using a StaxReader object.
- [xml.Serializer.OptimizedVariableToDom](4179-xml-serializer-optimizedvariabletodom.md): Serializes a BDL variable into a XML element node using a DomNode object.
- [xml.Serializer.OptimizedVariableToSoapSection5](4180-xml-serializer-optimizedvariabletosoapsection5.md): Serializes a BDL variable into a XML element node in Soap Section 5 encoding.
- [xml.Serializer.OptimizedVariableToStax](4181-xml-serializer-optimizedvariabletostax.md): Serializes a BDL variable into a XML element node using a StaxWriter object.
