---
title: "Review user-made C routines"
source: "fgl-topics/c_fgl_MigI4GL_028.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Review user-made C routines"
type: "concept"
---

# Review user-made C routines

> Genero BDL provides libraries which may replace some of the C routines required by I4GL applications.

IBM®
Informix® 4GL (I4GL) applications often need additional
utility C routines implemented in C-Extensions. For example, to access the file system and read the
content of a directory with the [`os.Path`](../15_library-reference/3697-the-os-path-class.md "The os.Path class provides functions to manipulate files and directories on the machine where the program executes.") class. Writing C-Extensions is an important cost in cross-platform
portability and maintenance.

Genero Business Development Language (BDL) provides a set of [libraries](../15_library-reference/2723-library-reference.md "Reference for classes and functions provided as built-in or extension packages.") that include functions and classes which can
probably replace some of the routines written for I4GL application. For example, BDL implements
typical file management functions to search directories and files.

If portability is a concern (for example if you want to move from
a UNIX™ platform to a
Microsoft™
Windows™ or
Mac OS-X™ platform),
review your C routines and check whether there is a replacement
built into the language or in one of the libraries provided.

Genero BDL even allows use of the huge Java class library with the [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.").
