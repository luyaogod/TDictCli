---
title: "Genero programs"
source: "fgl-topics/c_fgl_intro_BDL_029.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > Genero programs"
type: "concept"
---

# Genero programs

> Genero Business Development Language (BDL) is a programming language based on simple and readable syntax.

The program logic is written in text files with the .4gl file extension,
called program source modules. Module sources are compiled (fglcomp) into p-code
modules with the .42m file extension, that can be executed by the
runtime system (fglrun). Application programs are build with a group of
.42m modules.

While Genero BDL supports legacy linking to create .42r program files, it is
highly recommended to use the `IMPORT FGL` solution. This approach leverages the full
capabilities of the fglcomp compiler and runtime system. You may also want to use
`PACKAGE` structures to streamline module organization and deployment.

## Related links

**Related concepts**  

[Program structure](../09_advanced-features/0785-program-structure.md "Explains the organization of a BDL program.")
