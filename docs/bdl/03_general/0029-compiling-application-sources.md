---
title: "Compiling application sources"
source: "fgl-topics/c_fgl_intro_BDL_018.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > Compiling application sources"
type: "concept"
---

# Compiling application sources

> You need to compile the source files in order to run the application.

A program can consist of a single source code module, but generally it will be organized in
multiple modules, will involve form specification files and perhaps localized string files.

Database schema files are required when you define program data types and variables in terms of
an existing database column or table, by using the `DEFINE ... LIKE` statement.

Before running your application with the runtime system, you need to use compilation tools in
order to build the various runtime files.

![Genero compilation tools diagram](../_images/TUT101-2.jpg)

*Genero compilation tools*

Program module dependency is defined with the `IMPORT FGL` instruction, and
modules can be grouped into packages.

## Related links

**Related concepts**  

[Program structure](../09_advanced-features/0785-program-structure.md "Explains the organization of a BDL program.")

[Form definitions](../11_user-interface/1537-form-definitions.md "This section describes how to define application forms and program resources related to the presentation layer.")

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")
