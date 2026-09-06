---
title: "The NodeList class"
source: "fgl-topics/c_fgl_ClassNodeList.html"
breadcrumb: "Library reference > Built-in packages > The om package > The NodeList class"
type: "concept"
---

# The NodeList class

> A om.NodeList object hold a list of DOM nodes.

The list is created from an [`om.DomNode.selectByTagName()`](3321-om-domnode-selectbytagname.md "Finds descendant DOM nodes based on a tag name.") or [`om.DomNode.selectByPath()`](3320-om-domnode-selectbypath.md "Finds descendant DOM nodes from an XPath-like pattern.")
method.

After creating the node list, you can process the nodes with the `getLength()`
and `item()` methods of the `om.NodeList` object.

> **Important:**
>
> Build-in XML support classes `om.*` have known limitations. Refere to the [Limitations of XML built-in classes](../09_advanced-features/0951-limitations-of-xml-built-in-classes.md "Built-in XML classes have some limitations you must be aware of.") page for more details.

## Child topics

- [om.NodeList methods](3331-om-nodelist-methods.md): Methods of the om.NodeList class.
- [Examples](3334-examples.md): om.NodeList usage examples.
