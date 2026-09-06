---
title: "xml.DomDocument.CreateDocumentNS"
source: "fgl-topics/c_gws_XmlDomDocument_createDocumentNS.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.CreateDocumentNS"
type: "concept"
---

# xml.DomDocument.CreateDocumentNS

> Constructor of a xml.DomDocument with a root namespace-qualified XML root element

## Syntax

```
xml.DomDocument.CreateDocumentNS(
   prefix STRING,
   name STRING,
   ns STRING )
  RETURNS xml.DomDocument
```

1. prefix defines the prefix of the XML Element or NULL.
2. name defines the XML element.
3. ns is the namespace of the XML Element.

Returns an `xml.DomDocument` object.

## Usage

Constructor of a `xml.DomDocument` with a root namespace-qualified XML root
element where prefix is the prefix of the XML element or `NULL`,
name is the name of the XML element, and ns is the namespace
of the XML element. Returns an `xml.DomDocument` object.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

Create a `xml.DomDocument` with an initial root node named `"List"`
with `abc` as the prefix and `http://www.mysite.com/xmlapi` as the
namespace:

```
xml.domdocument.createDocumentNS("abc","List","http://www.mysite.com/xmlapi")
```

Produces:

```
<abc:List xmlns:abc="http://www.mysite.com/xmlapi">
[...]
</abc:List>
```

## Related links

**Related concepts**  

[Example 1 : Create a namespace qualified document with processing instructions](4018-example-1-create-a-namespace-qualified-document-with-process.md "Example 1 : Create a namespace qualified document with processing instructions")
