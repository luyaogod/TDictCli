---
title: "P-Code compatibility"
source: "fgl-topics/c_fgl_Migrate_to_all_pcode.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > General BDL upgrade guide > P-Code compatibility"
type: "concept"
---

# P-Code compatibility

> P-Code incompatibility (within .42m files) may be introduced from version to version.

Recompilation is only needed when the p-code becomes incompatible. When executing a program with
an older p-code version as expected, fglrun will raise the error [-6201](../15_library-reference/4483-genero-bdl-errors.md).

The [product version number](0086-version-number-meaning.md "A product version number identifies a specific release of the software product.") has the form
`X.YY.ZZ` and identifies a given Genero BDL release.

If the new version is different from the older version by its first digit (`X`),
you must recompile the program code sources and form files. For example, when upgrading from
`5.01.02` to `6.00.12`, recompilation is required.

Recompilation is not required when upgrading to a version with a different
`YY.ZZ` part. For example, there is no need to recompile when upgrading from
`5.00.03` to `5.01.17`, typically in a production enviroment.

However, there is potentially some benefits to be gained from recompiling your sources when
upgrading any Genero version: The fglcomp compiler can be improved, to produce
more optimized p-code, detect source code mistakes with new warning options, and speed up
compilation time. Consequently, it's always a good idea to recompile your sources with the latest
compiler version.

When upgrading, you may need to recreate the [C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") libraries: C extension libraries must be provided as dynamically loadable
modules and a rebuild is generally not needed. However, if the C-Extension API header files have
changed, you must recompile your C sources. Check $FGLDIR/include/f2c for C
Extension API header file changes.

## Related links

**Related concepts**  

[Compiling source files](../13_programming-tools/2530-compiling-source-files.md "Describes how to build the runtime files from source files.")
