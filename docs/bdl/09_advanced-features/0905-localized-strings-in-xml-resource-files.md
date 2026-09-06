---
title: "Localized strings in XML resource files"
source: "fgl-topics/c_fgl_localized_strings_004.html"
breadcrumb: "Advanced features > Localization > Localized strings > Localized strings in XML resource files"
type: "concept"
---

# Localized strings in XML resource files

> In XML resource files, localized string specification must follow the XML syntax and therefore must be defined as an XML node.

## Syntax: Localized string in XML files

```
  <ParentNode attribute = "default" [...] >
     <LStr attribute = "sid" [...] />
  </ParentNode>
```

1. ParentNode is the node type of the parent where the localized strings must be
   applied.
2. attribute is the attribute in the parent node that will get the localized
   string identified by *sid*.
3. default is the default text of an attribute, if not localized string is found
   for *sid*.
4. sid is a character string literal that defines both the string identifier and
   the default text.

## Description

In .42m p-code modules, the localized strings are coded in a proprietary
binary format. But, for XML files such as action defaults files (.4ad),
the localized strings must be written with a specific node, following the XML standards. To
support localized strings in XML files, any file loaded into the Abstract User Interface
tree is parsed to search for `<LStr>` nodes. The
`<LStr>` nodes define the same attributes as in the parent node with
localized string identifiers, for example:

```
<Label text="Hello!" >
  <LStr text="label01" />
</Label>
```

The runtime system automatically replaces corresponding attributes in the parent node
(`text="Hello!"`), with the localized text for the string identifier
(`label01`) found in the compiled string files. After interpretation, the
`<LStr>` nodes are removed from the XML data.

To take effect, a localized attribute in the `<LStr>` node must have a
corresponding attribute in the parent node.
