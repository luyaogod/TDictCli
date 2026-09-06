---
title: "HTML document usage example"
source: "fgl-topics/c_gws_XML_DomDocument_usage_html_doc.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > HTML document usage example"
type: "concept"
---

# HTML document usage example

> The HTML language provides tags that allow the user to provide an embedded style sheet (the "style" tag) and to write embedded client side script (the "script" tag). According to the HTML 4.0 specification, the content of these tags must be managed as CDATA section.

> **Note:**
>
> For more information, see the [HTML 4.0 specification](http://www.w3.org/TR/1998/REC-html40-19980424/).

Because HTML document management via the `xml.DomDocument` object provides HTML
compliancy only (and not strict HTML management), there is a specific way to add these nodes inside
a loaded HTML document:

1. [Create an element node](3972-xml-domdocument-createelement.md "Creates a XML Element xml.DomNode object for an xml.DomDocument object") with the name
   of the tag to be created.
2. [Append the element node](4024-xml-domnode-appendchild.md "Adds a child DomNode object to the end of the child list for a DomNode object") to its
   parent.
3. [Create a CDATASection node](3966-xml-domdocument-createcdatasection.md "Creates an XML CData xml.DomNode object for an xml.DomDocument object.") with
   the required embedded piece of style sheet or piece of script content.
4. [Append the CDATASection node](4024-xml-domnode-appendchild.md "Adds a child DomNode object to the end of the child list for a DomNode object")  to the
   previously-created element node.

By following this procedure, the "`script`" and "`style`" tags
content are recognized as `CDATA` section content and not `TEXT`
section content and will be preserved. Other methods for adding nodes to the document manage text
and therefore will not treat these types of content properly, resulting in invalid HTML code.

## Example

```
IMPORT xml
MAIN
  DEFINE myDoc xml.DomDocument
  DEFINE myEltNode, myAttrNode, bodyNode, myCdataNode xml.DomNode
  DEFINE nodeLst xml.DomNodeList
  TRY
    LET myDoc = xml.DomDocument.Create()
    CALL myDoc.setFeature("enable-html-compliancy", 1)
    CALL myDoc.load("testHtml.html")
    LET myEltNode = myDoc.createElement("script")
    LET myCdataNode = myDoc.createCDATASection("document.write(\"CDATA\");")
    LET myAttrNode = myDoc.createAttribute("type")
    CALL myAttrNode.setNodeValue("text/javascript")
    LET nodeLst = myDoc.getElementsByTagName("body")
    LET bodyNode = nodeLst.getItem(1)
    CALL bodyNode.appendChild(myEltNode)
    CALL myEltNode.setAttributeNode(myAttrNode)
    CALL myEltNode.appendChild(myCdataNode)
  CATCH
    DISPLAY "ERROR : ", status, " - ", sqlca.sqlerrm
    EXIT PROGRAM(-1)
  END TRY
END MAIN
```
