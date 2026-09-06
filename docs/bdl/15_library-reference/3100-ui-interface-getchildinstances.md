---
title: "ui.Interface.getChildInstances"
source: "fgl-topics/c_fgl_ClassInterface_getChildInstances.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.getChildInstances"
type: "concept"
---

# ui.Interface.getChildInstances

> Get the number of child instances for a given program name.

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

## Syntax

```
ui.Interface.getChildInstances(
   name STRING )
  RETURNS INTEGER
```

1. name is the name of a child program attached to the container of the current
   program.

## Usage

The `ui.Interface.getChildInstances()` class method returns the number of child
instances of a program attached to the current container (parent) program, based on the name of the
child program passed as parameter.

Child programs are attached to a given container by using the [`ui.Interface.setContainer()`](3118-ui-interface-setcontainer.md "Define the parent container for the current program.")
method.

The name of a child program is defined by the [`ui.Interface.setName()`](3120-ui-interface-setname.md "Define the name of the current program for the front-end.") method.

The `getChildInstances()` method is typically used to check if a given
child program is already started, to avoid multiple instances of the same program in a given
container.
