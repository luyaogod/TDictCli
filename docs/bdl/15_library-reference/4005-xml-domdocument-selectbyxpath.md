---
title: "xml.DomDocument.selectByXPath"
source: "fgl-topics/c_gws_XmlDomDocument_selectByXPath.html"
breadcrumb: "Library reference > Extension packages > The xml package > The Document Object Modeling (DOM) classes > The DomDocument class > xml.DomDocument methods > xml.DomDocument.selectByXPath"
type: "concept"
---

# xml.DomDocument.selectByXPath

> Returns a xml.DomNodeList object containing all xml.DomNode objects matching an XPath 1.0 expression.

## Syntax

```
selectByXPath(
   expr STRING
   [, args ] )
  RETURNS xml.DomNodeList
```

1. expr defines the XPath 1.0 expression.
2. args is a variable list of parameters that define a list of
   prefixes bound to namespaces to resolve qualified names in the XPath expression.

## Usage

This method returns a [DomNodeList](4085-the-domnodelist-class.md "The xml.DomNodeList class provides methods to manipulate a list of DomNode objects.") object containing all `xml.DomNode` objects matching an XPath 1.0
expression.

> **Important:**
>
> This method is not part of W3C standard
> API.

args is a variable list of parameters that define a list of prefix/namespace
pairs used to resolve qualified names in the XPath expression. You must provide an even number of
parameters, where each set of two parameters represent a prefix (alias) and its corresponding
namespace:

```
selectByXPath( expr,
  "prefix-1", "namespace-1",
  "prefix-2", "namespace-2",
  "prefix-3", "namespace-3",
  ... )
```

A namespace must be an absolute URI; for example, 'http://' or
'file://'.

If args is `NULL`, the prefixes and namespaces defined in the
document itself are used when
available:

```
selectByXPath( expr, NULL )
```

If the XML document uses a default namespace without a prefix, you must provide a prefix in the
`selectByXPath()` expression. In this example, the `xmlns` attribute
in the XML document HTML tag does not have a
prefix.

```
<html lang="en" xml:lang="en" xmlns="http://www.w3.org/1999/xhtml">
```

In
this case, the search for the tag name `FORM` with the attribute
`NAME` equal to "login" with args
`NULL` in the example will not be found because the namespace does not have a prefix.

```
selectByXPath('//FORM[@NAME="login"]',NULL)
```

Instead you must provide a prefix
for the namespace; it can be any name, it does not have to be "alias" as in the example. You specify
the prefix and namespace in the args parameter of the
`selectByXPath()` similar to
this:

```
selectByXPath('//alias:FORM[@NAME="login"]',"alias","http://www.w3.org/1999/xhtml")
```

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Example

```
selectByXPath( "../../d:Record/*[last()]", 
  "d", "http://defaultnamespace" )

selectByXPath( "ns:Record", NULL )

selectByXPath( "ns1:Records/ns2:Record", 
  "ns1", "http://namespace1",
  "ns2", "http://namespace2" )
```

Invalid example:

```
selectByXPath( "ns1:Record", "ns1" )
```

## Related links

**Related concepts**  

[Navigation methods usage examples](4011-navigation-methods-usage-examples.md "Examples using the navigation methods of the xml.DomDocument class.")
