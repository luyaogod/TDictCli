---
title: "The Java interface"
source: "fgl-topics/c_fgl_JavaBridge_001.html"
breadcrumb: "Extending the language > The Java interface"
type: "concept"
---

# The Java interface

> The Java interface allows you to import Java classes and instantiate Java objects in your programs.

The Java interface gives access to the huge standard Java libraries, as well as custom
libraries for specific purposes.

The methods of Java objects can be called with other Java objects referenced in a program, as
well as with native language data types such as `INTEGER`, `DECIMAL`,
`CHAR`.

The Java interface of Genero has the following limitations:

1. It is not possible to use Java generic types such as `java.util.Vector<E>`,
   with a type parameter (for example, in pure Java: `Vector<MyClass> v = new
   Vector<MyClass>()` ).
2. Database connections cannot be shared between Java and Genero programs.
3. Java graphical objects cannot be used in Genero forms.

## Child topics

- [Prerequisites and installation](2656-prerequisites-and-installation.md)
- [Getting started with the Java interface](2661-getting-started-with-the-java-interface.md)
- [Advanced programming](2667-advanced-programming.md)
- [Examples](2699-examples.md): Java interface usage examples.
