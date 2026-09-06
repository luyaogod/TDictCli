---
title: "xml.Serializer.OptimizedStaxToVariable"
source: "fgl-topics/c_gws_XmlSerializer_OptimizedStaxToVariable.html"
breadcrumb: "Library reference > Extension packages > The xml package > XML serialization classes > The Serializer class > xml.Serializer methods > xml.Serializer.OptimizedStaxToVariable"
type: "concept"
---

# xml.Serializer.OptimizedStaxToVariable

> Serializes an XML element node into a BDL variable using a StaxReader object.

This API implements the XML-binary Optimized packaging specification.
See [https://www.w3.org/TR/xop10/](http://www.w3.org/TR/xop10/).

## Syntax

```
xml.Serializer.OptimizedStaxToVariable(
   stax xml.StaxReader,
   var fgl-type,
   xopTable RECORD )
```

1. stax is a [StaxReader](4121-the-staxreader-class.md "The StaxReader class provides methods compatible with Streaming API for XML(StAX) for reading XML documents.") object where the cursor points
   to an XML Element node.
2. var is any Genero BDL [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."). Or it can be a structured type, like a
   `RECORD` or `ARRAY`, including elements with these data types.
   Optional [XML mapping attributes](../16_web-services/4958-xml-serialization-rules-and-customization.md) can be added to the definition of variables for XML serialization.
3. xopTable is a dynamic array, defined as
   follows:

   ```
   DEFINE XOPTable DYNAMIC ARRAY OF RECORD
     cid STRING, # Content-ID to identify the part in a XML Optimized document
     data BYTE,  # Blob handled as part in a XML Optimized document
     file STRING # Name of the file handled as part in an XML Optimized document
   END RECORD
   ```

   The
   `XOPtable` dynamic array is necessary to keep the relation between the data to be
   handled as separate part in an XML Optimized document via an href attribute containing the
   Content-ID value. This parameter can be NULL.

## Usage

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Optimized APIs

Optimized APIs work in the same method as the non-Optimized APIs, with the addition that the
optimized API supports XML-binary Optimized format, and return in the XOPTable
(if not NULL) the BYTE or the file on disk to handle as a separate part based on the Content-ID.

For instance, if a BYTE has to be sent as an attachment via MTOM, the API will create an XML
Optimized node with an href containing the Content-ID returned in the XOP table. This node will
include a reference to that BYTE.

For example, given this example of an XML-optimized
document:

```
<m:data xmlns:m='http://example.org/stuff'>
  <m:photo>
  <xop:Include xmlns:xop='http://www.w3.org/2004/08/xop/include' 
  href='cid:myref@tempuri.org'/>
  </m:photo>
</m:data>
```

If you have NOT used XMLOptimizedContent, the `XOPTable` will contain one element
where:

- `cid` contains "`myref@tempuri.org`".
- `data` contains the BYTE to be sent or to be received as an
  attachment.
- `file` contains NULL.

If you have used XMLOptimizedContent, the `XOPtable` contains one element
where:

- `cid` contains "`myref@tempuri.org`".
- `data` contains NULL.
- `file` contains the filename to be sent or to be received as
  an attachment..
