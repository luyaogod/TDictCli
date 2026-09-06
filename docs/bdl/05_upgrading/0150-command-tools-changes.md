---
title: "Command tools changes"
source: "fgl-topics/c_fgl_Migrate_to_400_fgl_tools.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Command tools changes"
type: "concept"
---

# Command tools changes

> Modifications to consider regarding command line tools.

This topic describes changes that may need code review.

See also [new 4.00 features of
commands](0060-bdl-4-00-new-features.md).

## Auto-compilation with fglcomp

When using [`IMPORT FGL`](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") for a
module that is found in another directory through FGLLDPATH, fglcomp will not
auto-compile that module, if the .42m is not up to date with the
.4gl source. This does not change with version 4.00.

However, with the introduction of [packages](../09_advanced-features/0816-package.md "Defines the package the module belongs to.") in
version 4.00, .4gl source modules that can be found by
fglcomp from package paths, will be automatically compiled, if the
.42m is out of date.

The `--output-dir` option of fglcomp can be used, to specify
the destination directory for .42m files.

For more details, see [Automatic compilation of imported modules](../13_programming-tools/2534-compiling-program-code-files-4gl.md).

## Deprecated command tools

The following command line tools are deprecated, and only supported for backward compatibility:

- [fgl2p](../13_programming-tools/2515-fgl2p.md "The fgl2p tool compiles source files and assembles p-code modules into a .42r program or a .42x library."): use
  fglcomp or fgllink .

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Command
tools changes in 3.20](0177-command-tools-changes.md "Modifications to consider regarding command line tools.").

Notable changes introduced in maintenance releases:

- The [fglrun -m/-M options](0177-command-tools-changes.md) have been deprecated in 3.20.06

## Related links

**Related concepts**  

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
