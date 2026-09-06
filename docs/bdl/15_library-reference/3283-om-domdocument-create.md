---
title: "om.DomDocument.create"
source: "fgl-topics/c_fgl_ClassDomDocument_create.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.create"
type: "concept"
---

# om.DomDocument.create

> Create a new empty om.DomDocument object.

## Syntax

```
om.DomDocument.create(
   tagName STRING )
  RETURNS om.DomDocument
```

1. tagName defines the tag name of the root element.

## Usage

Use the class method `om.DomDocument.create()` to instantiate a new, empty DOM
document object.

To hold the reference to a DOM document object, define a variable with the
`om.DomDocument` type.

## Example

```
DEFINE d om.DomDocument
LET d = om.DomDocument.create("Vehicles")
...
```

## Related links

**Related concepts**  

[om.DomDocument.createFromXmlFile](3285-om-domdocument-createfromxmlfile.md "Create a new om.DomDocument object from a XML file.")

[om.DomDocument.createFromString](3284-om-domdocument-createfromstring.md "Create a new om.DomDocument object from an XML string.")
