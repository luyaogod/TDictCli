---
title: "Example 1: Search nodes by tag name"
source: "fgl-topics/c_fgl_ClassNodeList_example_1.html"
breadcrumb: "Library reference > Built-in packages > The om package > The NodeList class > Examples > Example 1: Search nodes by tag name"
type: "concept"
description: "XML file: Vehicles.xml <?xml version='1.0' encoding='ASCII'?> <Vehicles> <Vehicle type=\"car\" name=\"DeLorean DMC-12\" year=\"1983\" color=\"Gray\" weight=\"1230\"> <Engine energy=\"petrol\" size=\"2849\" ..."
---

# Example 1: Search nodes by tag name

XML file: Vehicles.xml

```
<?xml version='1.0' encoding='ASCII'?>
<Vehicles> 
   <Vehicle type="car" name="DeLorean DMC-12" year="1983" color="Gray" weight="1230">
      <Engine energy="petrol" size="2849" power="185" />
      <Wheels count="4" width="280" diameter="550" />
   </Vehicle>
   <Vehicle type="car" name="Corolla" year="2003" color="Blue" weight="1146">
      <Engine energy="gasoline" size="1200" power="75" />
      <Wheels count="4" width="220" diameter="525" />
   </Vehicle> 
   <Vehicle type="bus" name="Maxibus" year="1998" color="Yellow" weight="5278">
      <Engine energy="diesel" size="4100" power="445" />
      <Wheels count="6" width="315" diameter="925" />
   </Vehicle>
</Vehicles>
```

Program file:

```
MAIN
    DEFINE d om.DomDocument
    DEFINE nl om.NodeList
    DEFINE r, n om.DomNode
    DEFINE i INTEGER

    LET d = om.DomDocument.createFromXmlFile("Vehicles.xml")
    LET r = d.getDocumentElement()
    LET nl = r.selectByTagName("Vehicle")
    FOR i = 1 TO nl.getLength()
        LET n = nl.item(i)
        DISPLAY n.getAttribute("name")
    END FOR

END MAIN
```

Output:

```
DeLorean DMC-12
Corolla
Maxibus
```

## Related links

**Related concepts**  

[om.DomNode.selectByTagName](3321-om-domnode-selectbytagname.md "Finds descendant DOM nodes based on a tag name.")
