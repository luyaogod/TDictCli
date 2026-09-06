---
title: "om.DomNode.selectByPath"
source: "fgl-topics/c_fgl_ClassDomNode_selectByPath.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.selectByPath"
type: "concept"
---

# om.DomNode.selectByPath

> Finds descendant DOM nodes from an XPath-like pattern.

## Syntax

```
selectByPath(
   path STRING )
  RETURNS om.NodeList
```

1. path is an XPath-like pattern.

## Usage

The `selectByPath()` method scans the DOM tree for descendant nodes
from the specified XPath-like pattern.

> **Important:**
>
> The `selectByPath()` method supports a limited XPath
> syntax.

The search pattern must always start with `/` or `//`, must contain
at least on tag name, and the attributes expression allows only equality comparison.

The pattern supported by `selectByPath()` is limited to the following syntax:

```
{ / | // } TagName [ [@AttributeName="Value"] ] [...]
```

DOM node tag names and attributes names are case-sensitive.

The method creates a list of nodes as an `om.NodeList` object. This list object is
then used to process the nodes found.

| XPath expression | Description |
| --- | --- |
| /Vehicle | The element `<Vehicle/>`. |
| //Vehicle | All `<Vehicle/>` elements. |
| //Transport/Vehicle | All `<Vehicle/>` elements that are children of `<Transport/>` elements. |
| //Transport//Engine | All `<Engine/>` elements that are arbitrary descendants of `<Transport/>` elements. |
| //Transport/*/Engine | All `<Engine/>` elements that are grandchildren of `<Transport/>` elements. |
| //Transport/* | All elements that are children of `<Transport/>` elements. |
| //Engine[@power="185"] | All `<Engine/>` elements where the `power` attribute equals `185`. |
| //Vehicle[@type="car"]/Engine[@power="185"] | All `<Engine/>` elements where the `power` attribute equals `185`, children of `<Vehicle/>` elements where the `type` attribute is `"car"`. |

## Example

```
DEFINE node om.DomNode,
       nodelist om.NodeList
...
LET nodelist = node.selectByPath("//Grid/Table[@tabName=\"t1\"]")
```

For a complete example, see [Example 2: Search nodes by XPath](3336-example-2-search-nodes-by-xpath.md).

## Related links

**Related concepts**  

[The NodeList class](3330-the-nodelist-class.md "A om.NodeList object hold a list of DOM nodes.")
