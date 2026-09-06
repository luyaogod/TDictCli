---
title: "Example 1 : Create a namespace qualified document with processing instructions"
source: "fgl-topics/c_gws_XML_DomDocument_class_015.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > Examples > Example 1 : Create a namespace qualified document with processing instructions"
type: "concept"
description: "To create the following XML document on disk: <?Target1 This is my first PI ?> <MyPre:RootNode xmlns:MyPre=\"http://www.tempuri.org\" > <MyPre:Element /> </MyPre:RootNode> <?Target2 This is my last PI ..."
---

# Example 1 : Create a namespace qualified document with processing instructions

To create the following XML document on
disk:

```
<?Target1 This is my first PI ?>
<MyPre:RootNode xmlns:MyPre="http://www.tempuri.org" >
<MyPre:Element />
</MyPre:RootNode>
<?Target2 This is my last PI ?>
```

Write the following
code:

```
IMPORT xml

MAIN
  DEFINE doc xml.DomDocument
  DEFINE pi xml.DomNode
  DEFINE node xml.DomNode
  DEFINE elt xml.DomNode

  # Create a document with an initial namespace qualified root node
  LET doc = xml.DomDocument.CreateDocumentNS("MyPre", "RootNode", "http://www.tempuri.org")
  # Create a Processing instruction
  LET pi = doc.createProcessingInstruction("Target1", "This is my first PI")
  # And add it at the begining of the document
  CALL doc.prependDocumentNode(pi)
  # Create another Processing instruction
  LET pi = doc.createProcessingInstruction("Target2", "This is my last PI")
  # And add it at the end of the document
  CALL doc.appendDocumentNode(pi)
  # Retrieve initial root node of the document
  LET elt = doc.getDocumentElement()
  # Create a new Element node
  LET node = doc.createElementNS("MyPre", "Element", "http://www.tempuri.org")
  # And add it as child of the RootNode 
  CALL elt.appendChild(node)
  # Then save the document on disk
  CALL doc.save("MyFile.xml")
END MAIN
```
