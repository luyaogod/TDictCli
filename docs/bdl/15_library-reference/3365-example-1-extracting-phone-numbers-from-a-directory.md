---
title: "Example 1: Extracting phone numbers from a directory"
source: "fgl-topics/c_fgl_ClassSaxDocumentHandler_example_1.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxDocumentHandler class > Examples > Example 1: Extracting phone numbers from a directory"
type: "concept"
description: "This example shows how to write a SAX filter to extract phone numbers from a directory file written in XML. MAIN DEFINE f om.SaxDocumentHandler LET f = om.SaxDocumentHandler.createForName(\"module1\") ..."
---

# Example 1: Extracting phone numbers from a directory

This example shows how to write a SAX filter to extract phone numbers
from a directory file written in XML.

```
MAIN
  DEFINE f om.SaxDocumentHandler
  LET f = om.SaxDocumentHandler.createForName("module1")
  CALL f.readXmlFile("customers.xml")
END MAIN
```

> **Note:**
>
> The parameter of the `createForName()` method specifies the name of
> a source file that has been compiled into a .42m file
> ("module1.42m" in our example).

The module module1.4gl:

```
FUNCTION startDocument()
END FUNCTION

FUNCTION processingInstruction(name,data)
  DEFINE name,data STRING
END FUNCTION

FUNCTION startElement(name,attr)
  DEFINE name STRING
  DEFINE attr om.SaxAttributes 
  DEFINE i INTEGER
  CASE name
    WHEN "Customer"
      DISPLAY attr.getValue("lname")," ", attr.getValue("fname")
    WHEN "CellPhone"
      DISPLAY "   Cell phone: ", attr.getValue("number")
    WHEN "WorkPhone"
      DISPLAY "   Work phone: ", attr.getValue("number")
  END CASE
END FUNCTION

FUNCTION endElement(name)
  DEFINE name STRING
END FUNCTION

FUNCTION endDocument()
END FUNCTION

FUNCTION characters(chars)
  DEFINE chars STRING
END FUNCTION

FUNCTION skippedEntity(chars)
  DEFINE chars STRING
END FUNCTION
```

The XML file customers.xml:

```
<Customers>
  <Customer customer_num="101" fname="Ludwig" lname="Pauli"
    company="All Sports Supplies" address1="213 Erstwild Court"
    address2="" city="Sunnyvale" state="CA" zip-code="94086">
    <CellPhone number="408-789-8075" />
    <WorkPhone number="873-123-4543" />
  </Customer>
  <Customer customer_num="102" fname="Carole" lname="Sadler"
    company="Sports Spot" address1="785 Geary St"
    address2="" city="San Francisco" state="CA" zip-code="94117">
    <CellPhone number="415-822-1289" />
    <WorkPhone number="834-842-8373" />
  </Customer>
  <Customer customer_num="103" fname="Philip" lname="Currie"
    company="Phil's Sports" address1="654 Poplar"
    address2="P. O. Box 3498" city="Palo Alto" state="CA"
    zip-code="94303">
    <CellPhone number="415-328-4543" />
    <WorkPhone number="932-118-4824" />
  </Customer>
</Customers>
```

Output:

```
Pauli Ludwig
   Cell phone: 408-789-8075
   Work phone: 873-123-4543
Sadler Carole
   Cell phone: 415-822-1289
   Work phone: 834-842-8373
Currie Philip
   Cell phone: 415-328-4543
   Work phone: 932-118-4824
```
