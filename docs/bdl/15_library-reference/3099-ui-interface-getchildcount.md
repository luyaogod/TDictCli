---
title: "ui.Interface.getChildCount"
source: "fgl-topics/c_fgl_ClassInterface_getChildCount.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.getChildCount"
type: "concept"
---

# ui.Interface.getChildCount

> Get the number of children in a parent container.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

## Syntax

```
ui.Interface.getChildCount()
  RETURNS INTEGER
```

## Usage

The `ui.Interface.getChildCount()` class method returns the number of child
programs attached to the current container (parent) program.

Child programs are attached to a given container by using the [`ui.Interface.setContainer()`](3118-ui-interface-setcontainer.md "Define the parent container for the current program.")
method.

Container and child program identifiers/names must be defined by the [`ui.Interface.setName()`](3120-ui-interface-setname.md "Define the name of the current program for the front-end.") method.
