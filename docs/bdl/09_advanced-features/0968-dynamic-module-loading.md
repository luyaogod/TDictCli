---
title: "Dynamic module loading"
source: "fgl-topics/c_fgl_optimization_003.html"
breadcrumb: "Advanced features > Optimization > Runtime system basics > Dynamic module loading"
type: "concept"
description: "Type of Genero programs A Genero Business Development Language program is made of several .42m modules. A program can be created by linking modules together (to get a .42r file), or by defining the ..."
---

# Dynamic module loading

## Type of Genero programs

A Genero Business Development Language program is made of several .42m
modules.

A program can be created by [linking](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.") modules
together (to get a .42r file), or by defining the dependency between modules
with [`IMPORT FGL`](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") (and have only
.42m program files); in this case no linking is required.

## When are .42m modules loaded?

As a general rule, .42m program modules are loaded at runtime on demand,
when a module element (.i.e symbol) is required by the caller.

For example, when executing a [`CALL`](../08_language-basics/0675-call.md "The CALL instruction invokes a specified function or method.") instruction, the runtime system checks if the module implementing the
function is already in memory. If not, the module is first loaded, then module variables are
initialized, and then the function is called.

When using the [debugger](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs."), all modules are loaded
implicitly when starting the debugger, for both linked programs and programs using imported
modules.

## Overwriting .42m modules in production environments

Modules loaded by running programs are not affected by file replacements: Programs continue to
run with an image of the module file that was originally loaded.

However, replacing program modules during execution should be used with care: Since
.42m modules are loaded on demand (when a symbol of the module is referenced),
some modules of a running program may not yet be loaded.

When replacing a module while programs are running, invalid symbol errors can occur if the module
to be loaded does not correspond to the rest of the program modules that were loaded before the file
replacements.

See following scenario:

1. Program starts with V1 of main.42m, needing V1 of module
   libutil.42m (loaded later on demand).
2. Administrator upgrades application and installs main.42m and
   libutil.42m version V2.
3. Program running with V1 copy of main.42m calls a function
   from libutil.42m: runtime loads V2 of that module, while V1
   is expected.

When live application updates are mandatory, consider installing new program and resource files
(V2) in a different directory as the currently running version (V1), and use the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") and [FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.") environment variables to point
to the new files when starting a new (V2) program instance.

Note that on Windows™ platforms, program files currently
in use cannot be overwritten, because of Windows OS memory
mapping limitations. You need to turn off memory mapping with the [FGLPROFILE entry `fglrun.mmapDisable`](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.").
