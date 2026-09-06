---
title: "Node creation methods usage examples"
source: "fgl-topics/c_gws_XML_DomDocument_class_007.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > Node creation methods usage examples"
type: "concept"
---

# Node creation methods usage examples

> Node creation methods usage examples for the xml.DomDocument class.

Creating a node for the `xml.DomDocument` is done in two steps:

- Create the node.
- Add the node to the `xml.DomDocument`.

Each time you create a node, you need to append it at the right place in the
`xml.DomDocument`. To add a node to the document use the
`xml.DomDocument`
[management
methods](3960-xml-domdocument-methods.md "Methods for the xml.DomDocument class.") or the `xml.DomNode`
[manipulation methods](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.").

```
createNode("<LastName>PATTERSON</LastName><FirstName>Andrew</FirstName>")
```

Creates
a structure of nodes.

```
createElement("CardList")
```

Produces

```
<CardList>
```

```
createElementNS("cny", "Company", "http://www.mysite.com/")
```

Produces `<cny:Company
xmlns:cny="http://www.mysite.com/"/>` or `<cny:Company
/>`. See [Cautions](4015-cautions.md "Some things you need to be aware of when working with the xml.DomDocument class.") for
more details.

```
createAttribute("Country")
```

Creates a Country attribute
node.

- To set a value to the attribute, use the method [setNodeValue](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") of the
  `xml.DomNode` class.
- To add the attribute to an element node, use the method [setAttributeNode](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") of the
  `xml.DomNode` class.

```
createAttributeNS("tw","Town","http://www.mysite.com/cities")
```

Produces `xmlns:tw="http://www.mysite.com/cities"
tw:Town=""`

- To set a value to the attribute use the method [setNodeValue](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") of the
  `xml.DomNode` class.
- To add the attribute to an element node use the method [setAttributeNodeNS](4020-the-domnode-class.md "The xml.DomNode class provides methods to manipulate a node of a DomDocument object.") of the
  `xml.DomNode` class.
- For optimization reasons, the namespace is not written beside the attribute until the saving of
  the `xml.DomDocument`.
- When accessing the element node, the namespace is not listed in
  the list of children. In the example above, `tw:Town=""` is
  in the list of children, not `xmlns:tw="http://www.mysite.com/cities"`.
- To access the namespace during the DomDocument building, use the method [normalize](4014-load-and-save-methods-usage-examples.md "Load and save method usage examples for the xml.DomDocument class.") first.
  Normalize writes the namespace declaration at the appropriate place. If there is no previous
  declaration, it will be accessible as an attribute of this element, otherwise it will be an
  attribute of one of the ancestors of the element.

```
createTextNode("My Company")
```

Creates a text
node.

```
createComment("End of the card")
```

Produces `<!--End of the
card-->`

```
createCDATASection("<website><a href=\"www.mysite.com\">My 
Company</a></website>")
```

Produces
`<![CDATA[<website><a href="www.mysite.com">My
Company</a></website>]]>`

```
createEntityReference("title")
```

Creates the
entity reference `&title`.

```
createProcessingInstruction("xml-stylesheet", "type=\"text/xsl\"
href=\"card.xsl\"")
```

Produces
`<?xml-stylesheet type="text/xsl"href="card.xsl"?>`

```
createDocumentType("Card", NULL, NULL,"<!ELEMENT 
Card (lastname, firstname, company, location)>")
```

Produces `<!DOCTYPE
Card [ <!ELEMENT Card (lastname , firstname , company , location)>]>`

- Only inline DTD are supported. The DTD has to been inserted in
  the DomDocument at an appropriate place.

```
createDocumentFragment
```

Is a method that creates a lightweight
`xml.DomDocument`. It represents a subtree of nodes that do not need to conform to
well-formed XML rules. This makes DocumentFragment easier to manipulate than a
`xml.DomDocument`.

```
for i=1 to 5
  let node = doc.createElement("Card")
  call root.appendchild(node) end for
```

This produces a subtree with 5 Card nodes that do
not have any root node. Once the subtree is completed, it can be added to the
`xml.DomDocument` object like any other node.
