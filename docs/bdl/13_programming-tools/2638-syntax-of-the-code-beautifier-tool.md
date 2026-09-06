---
title: "Syntax of the code beautifier tool"
source: "fgl-topics/c_fgl_beautifier_syntax.html"
breadcrumb: "Programming tools > Source code beautifier > Syntax of the code beautifier tool"
type: "concept"
---

# Syntax of the code beautifier tool

> The source code beautifier option of fglcomp is --format.

## Syntax

To beautify a source file, use the `--format` option of
fglcomp:

```
fglcomp --format [ format-option [...] ]
```

1. format-option is a formatting option.

## Options

| Option | Description |
| --- | --- |
| `--fo-align-consecutive-assignments={0\|1}` | Aligns equal signs of consecutive variable assignments. Default is 0 (no alignment) |
| `--fo-align-consecutive-types={0\|1}` | Aligns type specification of consecutive variable definitions. Default is 0 (no alignment) |
| `--fo-align-trailing-comments={0\|1}` | Aligns consecutive comments. Default is 0 (no alignment) |
| `--fo-inplace` | Write formatted output back to the provided file, instead of stdout. Creates a copy of the original file in filename.4gl~ |
| `--fo-fallback-style=filename` | Specify the configuration filename to be used if no .fgl-format file is found. |
| `--fo-column-limit=integer` | Define the source line width limit. Default is 80. |
| `--fo-indent-width=integer` | Number of columns to use for indentation. Default is 4. |
| `--fo-continuation-indent-width=integer` | Indent width for line continuations. Default is 4. |
| `--fo-label-indent={0\|1}` | When 1, indent instruction clauses such as `WHEN` in a `CASE` instruction. Default is 1 (enabled). |
| `--fo-pack={0\|1}` | When 1, try to put as many items on the same line as possible. When 0, use one line for each item. Default is 0 (do not pack). |
| `--fo-lowercase-keywords={0\|1}` | When 1, produce lowercase keywords. When 0, produce uppercase keywords. Default is 0 (uppercase). |
| `--fo-use-tab={0\|1}` | When 1, instead of using space characters all over for column indentation (the default), use tabs to fill whitespace that spans at least from one tab stop to the next one, based on `--fo-tab-width` and `--fo_indent-width`. |
| `--fo-tab-width=integer` | Defines the number of columns used for tab stop, when `--fo-use-tab` is specified. The default is 8 columns. |
| `--fo-lines=start-line:end-line` | Formats only the range of lines specified. This option can be used multiple times to specify several pieces of code to be reformatted. |

## Configuration file

Formatting options can be provided in a configuration file. It is strongly recommended to use a
formatting configuration file instead of the command line options.

By convention, the name of the formatting configuration file is
".fgl-format".

The fglcomp --format tool will search for the .fgl-format
file in the directory containing the source file passed as an argument, then in the parent directory
of the source directory, and then its parent directory, and so on. The formatter will use the first
.fgl-format file found. If no .fgl-format is found,
fglcomp --format will read the file specified with the
`--fo-fallback-style=filename` option.
> **Tip:**
>
> Place the
> .fgl-format configuration file in the top directory of your project, to use the
> same formatting options for all sources located under this directory.

When using command line options and a configuration file, the compiler will combine formatting
options from the configuation file and from the command line. When the same formatting option is
specified in the configuration file and on the command line, the option of the command line takes
precedence.

The formatting configuration file is a simple text file with lines defining formatting options.
Lines starting with a hash character (#) are ignored.

The `--fo-inplace` option must be specified as a command line parameter. It will
not be read from the configuration file.

Example:

```
# My code formatting options
--fo-column-limit=80
--fo-indent-width=4
--fo-continuation-indent-width=4
--fo-label-indent=1
```

## Plugging the beautifier into VIM

For vim users, it is possible to define function keys and commands to invoke
the source code beautifier, and indent the current source file with a simple key stroke.

For more details about VIM configuration, see [Configure VIM for Genero BDL](2544-configure-vim-for-genero-bdl.md).
