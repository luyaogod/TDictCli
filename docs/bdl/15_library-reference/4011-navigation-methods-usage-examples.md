---
title: "Navigation methods usage examples"
source: "fgl-topics/c_gws_XML_DomDocument_class_005.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > Navigation methods usage examples"
type: "concept"
---

# Navigation methods usage examples

> Examples using the navigation methods of the xml.DomDocument class.

`xml.DomDocument` navigation functions deal with nodes immediately under the
`xml.DomDocument` object, except for search features. To navigate through all the
nodes, you can refer to the navigation functions of the class [xml.DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

```
<?xml version="1.0" encoding="ISO-8859-1"?>
<?xml-stylesheet type="text/xsl" href="card.xsl"?>
<!-- demo card -->
<CardList xml:id="1" >[...]
</CardList>
```

In the example the first node of the document is `xml-stylesheet`. Use
`getFirstDocumentNode` to get the node. The element at position 2 is the comment
`<!-- demo card -->`. Use `getDocumentNodeItem` function to get
the node.

The last node of the document is `CardList`. Use `getLastDocumentNode` to get the node.

The number of nodes in the document is 3. This is the result of the function `getDocumentNodesCount`. This
function only counts the number of children immediately under the `xml.DomDocument`.

> **Note:**
>
> The first line of the example, `<?xml version="1.0"
> encoding="ISO-8859-1"?>`, is not considered as a node. To access the information of the
> first line, use [`getXmlVersion()`](3991-xml-domdocument-getxmlversion.md "Returns the document version as defined in the XML document declaration.") and [`getXmlEncoding`](3990-xml-domdocument-getxmlencoding.md "Returns the document encoding as defined in the XML document declaration.") functions.

Caution, if the example is in pretty printed format, the results are not the same. There are
additional text nodes representing the carriage
returns.

```
<?xml version="1.0" encoding="ISO-8859-1"?>
<?xml-stylesheet type="text/xsl" href="card.xsl"?>
<!-- demo card -->
<CardList xml:id="1" > [...]
</CardList>
```

See
[Cautions](4015-cautions.md "Some things you need to be aware of when working with the xml.DomDocument class.")
section for more details.

You can select nodes using their tag names, by XPath, or by their attributes value (if of type
ID, `xml:id` for example). The [`getElementsbyTagName`](3983-xml-domdocument-getelementsbytagname.md "Returns a xml.DomNodeList object containing all XML Element xml.DomNode objects with the same tag name in the document.") and [`getElementsbyTagNameNS`](3984-xml-domdocument-getelementsbytagnamens.md "Returns an xml.DomNodeList object containing all namespace qualified XML Element xml.DomNode objects with the same tag name and namespace in the entire document") methods return a [DomNodeList](4085-the-domnodelist-class.md "The xml.DomNodeList class provides methods to manipulate a list of DomNode objects.") object, unlike the other
methods that return a [DomNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") object. The
`xml.DomNodeList` is restricted to containing objects with the same tag name and/or
namespace. The `selectByXPath` method also returns a `xml.DomNodeList` object, but
each node can have a different name.

```
getElementsByTagNameNS("message","http://schemas.xmlsoap.org/wsdl/")
```

Get
the message nodes that have http://schemas.xmlsoap.org/wsdl/
as the namespace.

```
getElementsByTagNameNS("message","*")
```

Get all the
message nodes, regardless of the namespace they have.

```
getElementsByTagName("message")
```

Get all the message
nodes that do not have any namespace.

```
selectByXPath("//xs:element",NULL)
```

Get all the `xs:element`
nodes that have a namespace corresponding to the prefix xs .

```
selectByXPath("//Card",NULL)
```

Get all the Card nodes that do not have any
namespace.

```
getElementById("1")
```

Get the unique node whose attribute of type ID has a
value of "1".
