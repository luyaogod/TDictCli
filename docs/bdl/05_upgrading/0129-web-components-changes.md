---
title: "Web components changes"
source: "fgl-topics/c_fgl_Migrate_to_401_web_components.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.01 upgrade guide > Web components changes"
type: "concept"
---

# Web components changes

> Modifications to consider when using web components.

## fglsvgcanvas: STRING-typed parameters

Starting with BDL 4.01.03 / WCG 4.01.02, the [fglsvgcanvas](../11_user-interface/2421-the-fglsvgcanvas-web-component.md "The fglsvgcanvas built-in web component implements a drawing canvas for Scalable Vector Graphics content.") API functions to create SVG
elements take now `STRING`-typed parameters for coordinates and sizes. Before this
change, the functions were defined with `DECIMAL` parameters, and could not take
arguments as percentage, with units or as list of values. With `STRING` parameters it
is now possible to do the
following:

```
LET t1 = fglsvgcanvas.text("50%","70%","abcd",NULL)
LET t1 = fglsvgcanvas.text("10,20,30,40","20,22,24,26","abcd",NULL)
```

For a complete definition of fglsvgcanvas functions, see [fglsvgcanvas: SVG drawing module](../15_library-reference/2841-fglsvgcanvas-svg-drawing-module.md).

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Web component changes in BDL 4.00](0145-web-components-changes.md "Modifications to consider when using web components.").

Notable changes introduced in maintenance releases:

- No particular change to consider.

## Related links

**Related concepts**  

[Web components](../11_user-interface/2378-web-components.md "This section describes how to use web components in your application.")
