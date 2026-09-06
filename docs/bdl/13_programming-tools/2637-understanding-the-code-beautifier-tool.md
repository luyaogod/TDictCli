---
title: "Understanding the code beautifier tool"
source: "fgl-topics/c_fgl_beautifier_basics.html"
breadcrumb: "Programming tools > Source code beautifier > Understanding the code beautifier tool"
type: "concept"
---

# Understanding the code beautifier tool

> This is an introduction to the source code beautifier tool.

The code beautifier tool built in the [fglcomp compiler](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") can be used to reformat .4gl source
modules and get a readable source code, using consistent indentation and source layout rules across
modules.

The code beautifier supports uppercase or lowercase keyword coding conventions. See the [`--fo-lowercase-keywords`](2639-code-beautifier-tool-usage.md) option usage for more details.

Formatting options can be specified on [the
fglcomp command line](2638-syntax-of-the-code-beautifier-tool.md), or in a configuration file (it is recommended to use
the [configuration
file](2638-syntax-of-the-code-beautifier-tool.md) instead of commande line options).

The beautifier tool:

- Indents source lines for nested code blocks
- Converts language keywords to uppercase or lowercase
- Adds or removes space characters between tokens
- Breaks long lines
- Can use space or tab as indentiation characters
- Can indent parts of the source file

By default, the new reformatted source code is written to the standard output stream (stdout).
Use the [`--fo-inplace`](2638-syntax-of-the-code-beautifier-tool.md "The source code beautifier option of fglcomp is --format.") option to
replace directly the original source file.

> **Important:**
>
> The input source file must compile. Indenting a source file with errors will
> produce an unpredictable result.

In addition to the fglcomp --format tool, other command line tools are
provided to integrate with version control systems such as GIT and SVN:

- [fglformatdiff](2639-code-beautifier-tool-usage.md) can reformat code line changes referenced by a
  diff command output.
- [fglgitformat](2639-code-beautifier-tool-usage.md) can reformat code line changes referenced by a GIT commit.
  Use this tool instead of fglformatdiff, if GIT is your version control
  system.
