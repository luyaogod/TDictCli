---
title: "Load and save methods usage examples"
source: "fgl-topics/c_gws_XML_DomDocument_class_008.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > Load and save methods usage examples"
type: "concept"
---

# Load and save methods usage examples

> Load and save method usage examples for the xml.DomDocument class.

You can load an existing XML document. Before loading an XML document you need to [create the DomDocument](3963-xml-domdocument-create.md "Constructor of an empty xml.DomDocument object.") object.

An `xml.DomDocument` object can load files using different URI:
http://, https://, tcp://,
tcps://, file:// and alias://.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

```
load("data.xml")
load("http://www.w3schools.com/xml/cd_catalog.xml")
load("https://localhost:6394/ws/r/calculator?WSDL")
load("file:///data/cd_catalog.xml")
load("tcp://localhost:4242/")
load("tcps://localhost:4243/")
load("alias://demo")
```

In
the example the `demo` alias is defined in fglprofile as
`ws.demo.url = "http://www.w3schools.com/xml/cd_catalog.xml"`

```
loadfromstring("<List> <elt>First element</elt>
 <elt>Second element</elt> <elt>Third element</elt> </List>")
```

The
example, produces a subtree with a root node `List` and three nodes
`elt` and three `textnode`.

A `xml.DomDocument` can be saved at different URI beginning with:
http://, https://, tcp://,
tcps://, file:// and alias://.

```
save("myfile.xml")
save("http://myserver:8080/data/save1.xml")
save("file:///data/save.xml")
save("tcp://localhost:4242/")
save("alias://test")
```

In the example the `test` alias is defined in
fglprofile as `ws.test.url =
"http://localhost:8080/data/save3.xml"`

The `saveToString`
method saves the `xml.DomDocument` in a string.

See [`getErrorsCount()`](3986-xml-domdocument-geterrorscount.md "Returns the number of errors encountered during the loading, saving or validation of a XML document.") and [`getErrorDescription()`](3985-xml-domdocument-geterrordescription.md "Returns the error description at the given position.")
to retrieve error messages related to XML document.

The `normalize` method
emulates a `xml.DomDocument` save and load. It can be called at any stage of the
`xml.DomDocument` building. This removes empty text nodes and sets namespace
declarations as if the document had been saved.
