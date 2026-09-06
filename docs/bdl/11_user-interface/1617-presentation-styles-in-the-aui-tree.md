---
title: "Presentation styles in the AUI tree"
source: "fgl-topics/c_fgl_presentation_styles_010.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Presentation styles in the AUI tree"
type: "concept"
---

# Presentation styles in the AUI tree

> Where to find presentation styles definitions in the AUI Tree?

Presentation styles are loaded in the abstract user interface tree, under the
`UserInterface` node, in a `StyleList` node following the presentation
style syntax.

The `StyleList` node holds a list of `Style` nodes that
define a set of attribute values. Attribute values are defined in
`StyleAttribute` nodes, with a `name` and a
`value` attribute.

## Related links

**Related concepts**  

[The abstract user interface tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")

[Syntax of presentation styles file (.4st)](1609-syntax-of-presentation-styles-file-4st.md "A .4st presentation styles file is an XML file defining style attributes to be applied by front-ends.")
