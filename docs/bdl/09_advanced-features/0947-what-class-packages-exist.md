---
title: "What class packages exist?"
source: "fgl-topics/c_fgl_oop_007.html"
breadcrumb: "Advanced features > OOP support > What class packages exist?"
type: "concept"
---

# What class packages exist?

> A set of utility packages including useful classes are part of the distribution.

[Built-in packages](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.") such as
`ui`, `om` and `base`, are part of the runtime system
and can be referenced directly.

[Extension packages](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.") such as
`util`, `os`, `com` and `xml` need to be
loaded explicitly with the `IMPORT` instruction, at the beginning of program
modules.

Genero supports usage of [Java classes and objects](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.").
Note that using Java from a Genero program will create a Java Virtual Machine (JVM) that will be
part of the runtime system process.
