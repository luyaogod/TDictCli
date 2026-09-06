---
title: "Understanding source code documentation"
source: "fgl-topics/c_fgl_AutoDoc_002.html"
breadcrumb: "Programming tools > Source documentation > Understanding source code documentation"
type: "concept"
---

# Understanding source code documentation

> This is an introduction to source code documentation.

Documenting sources is an important task in software development, to share the code among
applications and achieve better re-usability.

Source documentation must be concise, clear, and complete. However, documenting sources can be
boring and subject to mistakes if large repetitive documentation sections have to be written by
hand.

Source documentation can be produced automatically with the [fglcomp](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler. The compiler can generate source documentation from the
.4gl files of your project with minimum effort. The resulting source
documentation is generated in simple HTML format and can be published on a web server.

Source documentation is generated with the `--build-doc` option of
fglcomp. To extract documentation from a .4gl source:

```
fglcomp --build-doc filename.4gl
```

You can generate default documentation from the existing sources. For a better description of
the code, add special `#+` comments in your sources to describe code elements such
as functions, function parameters, and return values.

By default, only `PUBLIC` symbols are documented. If you want to include
`PRIVATE` symbols, use the `--doc-private` option:

```
fglcomp --build-doc --doc-private filename.4gl
```

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.
