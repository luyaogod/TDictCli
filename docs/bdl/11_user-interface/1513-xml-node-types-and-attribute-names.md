---
title: "XML node types and attribute names"
source: "fgl-topics/c_fgl_DynamicUI_022.html"
breadcrumb: "User interface > User interface basics > The abstract user interface tree > XML node types and attribute names"
type: "concept"
description: "XML is case-sensitive. By convention, node types in the AUI tree use uppercase/lowercase combinations to indicate word boundaries. Therefore, the nodes and attributes of an abstract user interface ..."
---

# XML node types and attribute names

XML is case-sensitive. By convention, node types in the AUI tree use uppercase/lowercase
combinations to indicate word boundaries. Therefore, the nodes and attributes of an abstract user
interface tree are handled as follows:

- Node types - the first letter of the node type is always capitalized. Subsequent letters are
  lowercase, unless the type consists of multiple words joined together. In that case, the first
  letter of each of the multiple words is capitalized (the camel-case convention). Examples:
  `Label`, `FormField`, `DateEdit`,
  `Edit`.
- Attribute names - the first letter of the name is always lowercase; subsequent letters are also
  lowercase, unless the name consists of multiple words joined together. In that case, the first
  letter of each subsequent word is capitalized (the Lower camel-case convention). Examples:
  `text`, `colName`, `width`,
  `tabIndex`
- Attribute values - the values are enclosed in quotes, and the runtime system does not convert
  them.

If you reference AUI tree XML nodes or attributes in your code, you must always respect the
naming conventions.
