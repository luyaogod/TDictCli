---
title: "BDL 1.10 new features"
source: "fgl-topics/fgl_whatsnew_110.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 1.10 new features"
type: "topic"
---

# BDL 1.10 new features

> Features added in 1.10 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 1.10 upgrade guide](0340-bdl-1-10-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 1.10.").

| Overview | Reference |
| --- | --- |
| The language supports now built-in classes, a new object-oriented way to program in BDL. | See [OOP support](../09_advanced-features/0942-oop-support.md "Describes Object Oriented Programming basics in the language."). |
| `CONSTANT` keyword to define constants in your programs. | See [Constants](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values."). |
| The language now supports dynamic arrays with automatic memory allocation. | See [Dynamic arrays](../08_language-basics/0734-dynamic-arrays.md). |
| A set of XML Utilities are provided in the runtime library as built-in classes. | See [The om package](../15_library-reference/3280-the-om-package.md "These topics cover the built-in classes of the om package"). |
| The `STRING` data type can be used to manipulate character strings without a length limit as with `CHAR`/`VARCHAR`. | See [STRING](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation."). |

| Overview | Reference |
| --- | --- |
| The Dynamic User Interface is the major new concept in Genero. It is the basement for the new graphical user interface. | See [User interface basics](../11_user-interface/1508-user-interface-basics.md "This section introduces to the foundation of the Genero user interface."). |
| Compared to classic IBM Informix 4gl, interactive instructions such as `INPUT`, `DISPLAY ARRAY`, have been extended with new control blocks and control instructions. | See [Dialog instructions](../11_user-interface/1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation."). |
| Form specification files (.per) support now extended layout definition with the `LAYOUT` section. | See [Form definitions](../11_user-interface/1537-form-definitions.md "This section describes how to define application forms and program resources related to the presentation layer."). |
| Defining Window Containers (a.k.a. MDI) is a simple way to group programs. | Desupported in V4. See [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine."). |

| Overview | Reference |
| --- | --- |
| The new `SCHEMA` instruction allows you to specific a database schema, without having an implicit connection, when the program executes. | See [Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions."). |
