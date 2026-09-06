---
title: "Understanding interfaces"
source: "fgl-topics/c_fgl_Interface_intro.html"
breadcrumb: "Language basics > Interfaces > Understanding interfaces"
type: "concept"
---

# Understanding interfaces

> This is an introduction to interfaces.

Interfaces are a way to achieve polymorphism in Genero BDL.

While [user-defined types](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") describe the data
structure and storage, interfaces describe the behavior of an object (type), by
declaring the list of operations ([methods](0772-methods.md "A function declared with a receiver type defines a method for this type.")) that can
be done on associated types.

The code using a variable declared as an interface can then act on different typed variables for
which these methods exist.

The typical and best practice with interfaces is to write functions that take interfaces as
parameters, where generic code manipulates the objects referenced by these interfaces, using common
methods associating the interface to the object types. The parameters of these generic functions can
also be collections of interfaces such as [dynamic arrays](0734-dynamic-arrays.md)
or [dictionaries](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") of interfaces.

For example, a `"Shape"` `INTERFACE` can define methods for shapes
such as `"area()"`, that will be implemented by the real specialized object types
such as `"Rectangle"` or `"Circle"`, as `RECORD` with
methods. Functions manipulating `"Shape"` interfaces don't need to know about the
implementation of the real types, and can be reused if new types like `"Triangle"` or
`"Ellipse"` are added.
