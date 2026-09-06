---
title: "Distinguish class and object methods"
source: "fgl-topics/c_fgl_oop_005.html"
breadcrumb: "Advanced features > OOP support > Distinguish class and object methods"
type: "concept"
---

# Distinguish class and object methods

> Class methods can be invoked from the class, while object methods can only be invoked from the variable referencing the object.

Methods can be invoked like regular functions, by passing parameters and/or returning values, and
can be used in expressions when they return a scalar value.

## Class methods

Class methods are called
by using the class identifier as prefix, with the period as
separator. The class identifier includes the package name and class
name.

```
package.classname.method( parameter [,...] )
```

For example, to call the `refresh()` method of
the `Interface` class, which is part of the `ui` package:

```
CALL ui.Interface.refresh()
```

## Object methods

Object methods are
called through the variable referencing the object. To use object
methods, the object must exist. Call the object methods by
using the object variable as a prefix, with a period as the
separator.

```
object.method( parameter [,...] )
```

For example, to call the `setFieldActive()` method
of an object of the `Dialog` class,
which is part of the `ui` package:

```
DEFINE d ui.Dialog
LET d = ui.Dialog.getCurrent()
CALL d.setFieldActive("cust_addr", FALSE)
```
