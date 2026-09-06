---
title: "Deploying application files"
source: "fgl-topics/c_fgl_intro_BDL_024.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > Deploying application files"
type: "concept"
---

# Deploying application files

> To deploy an application, you must deploy all of the required runtime and resource files. Many (but not all) of these files are compiled from the source files.

![Deployment files diagram](../_images/TUT102-2.jpg)

*Deployment files*

These program files must be deployed at the user site:

- .42m - PCode modules
- .42f - Runtime form files
- .42s - Compiled localized string files
- .4ad, .4st, etc - XML resource files

## Runtime environment settings

The .42m modules are found from the directory set provided in the FGLLDPATH
environment variable.

The resource files (.42f, .42s, etc) are found at
runtime from the directory set provided in the FGLRESOURCEPATH environment variable.

The fglprofile configuration file and environment variables can be used to
change the behavior of programs.

## Directory structure for program files

When deploying your program files, you have the following options:

1. Copy all program files (.42m, .42f, etc) into the same
   distribution directory. The runtime files will be found by fglrun, when executing
   the main module from that directory as current working directory. This solution is suitable for
   small applications with a few program files.
2. Group .42m program files and (.42f, etc) resource
   files in distinct directories, to distinguish application domains, and easily disable/remove some
   parts of your software from the production site. The directory structure of the production
   environment can be the same than the source directory tree (without the sources of course!). Define
   FGLLDPATH (to all dir paths to .42m files) and FGLRESOURCEPATH for resource
   files.
3. Use the package concept to structure your modules, and clone the source directory structure to
   the production environment (without sources). Define FGLLDPATH (to top-dir) and
   FGLRESOURCEPATH for resource files.

## Related links

**Related concepts**  

[Program execution](../09_advanced-features/0828-program-execution.md "This section describes program execution and language instructions related to program execution.")

[Configuration](../07_configuration/0482-configuration.md "These topics cover configuration options of the Genero Business Development Language.")

[Organizing modules in packages](../09_advanced-features/0818-organizing-modules-in-packages.md "Modules to be imported can be grouped in packages.")

[Packaging web applications](../13_programming-tools/2645-packaging-web-applications.md "Describes methods of packaging the runtime files and resources of your web applications and services using the fglgar tool.")
