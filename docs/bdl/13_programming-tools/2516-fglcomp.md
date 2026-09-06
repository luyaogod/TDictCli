---
title: "fglcomp"
source: "fgl-topics/c_fgl_tools_fglcomp.html"
breadcrumb: "Programming tools > Command reference > fglcomp"
type: "concept"
---

# fglcomp

> The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.

## Syntax 1: Compiling .4gl sources

```
fglcomp [comp-options] [prepro-options] file-list
```

1. In this form, fglcomp [compiles
   the source files(s)](2530-compiling-source-files.md "Describes how to build the runtime files from source files.") to .42m p-code modules.
2. comp-options are described in compilation options.
3. prepro-options are described in preprocessor options.
4. In this form, fglcomp can handle several .4gl sources
   files as arguments, see file-list.

where file-list
is:

```
{ module[.4gl]
| pattern
| @argfile
} [...]
```

1. module.4gl is a source file to
   compile. The .4gl extension is optional.
2. pattern is a `MATCHES`-style pattern to find files, like
   `'[a-z]*.4gl'`.
3. argfile defines a file that contains a list of .4gl
   sources to be compiled. Each line must specify a filename or a pattern.

## Syntax 2: Produce module dependency rules

```
fglcomp --dependencies [prepro-options] file-list
```

1. In this form, `fglcomp` produces [imported
   module dependency rules](2534-compiling-program-code-files-4gl.md).
2. prepro-options are described in preprocessor options.
3. In this form, fglcomp can handle several
   .4gl sources files as arguments, see file-list.

## Syntax 3: Building documentation

```
fglcomp --build-doc [doc-options] [prepro-options] module[.4gl]
```

1. In this form, fglcomp produces [source
   documentation](2546-source-documentation.md "Explains how to automatically generate documentation from your sources.").
2. doc-options are described in documentation options.
3. prepro-options are described in preprocessor options.
4. module.4gl is a source file to
   process. The .4gl extension is optional.

## Syntax 4: Extracting localized strings

```
fglcomp -m [prepro-options] module[.4gl]
```

1. In this form, fglcomp extracts [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") from the source.
2. prepro-options are described in preprocessor options.
3. module.4gl is a source file to
   process. The .4gl extension is optional.

## Syntax 5: Source code formatting

```
fglcomp --format [format-options] module[.4gl]
```

1. In this form, fglcomp [reformats the source
   file](2636-source-code-beautifier.md "Reformat the source code for better readability.").
2. format-options can be any of the formatting options.
3. module.4gl is a source file to
   process. The .4gl extension is optional.

## Syntax 6: Add imported module prefixes

```
fglcomp --qualify-imports [refact-options]
     [prepro-options] module[.4gl]
```

1. In this form, fglcomp adds [imported
   module prefixes](2642-qualifying-imported-symbols.md "The fglcomp compiler can automatically qualify imported symbols.") to symbols.
2. refact-options are described in refactoring options.
3. prepro-options are described in preprocessor options.
4. module.4gl is a source file to
   process. The .4gl extension is optional.

## Syntax 7: Mark SQL host variables with $

```
fglcomp --mark-host-variables [refact-options]
     [prepro-options] module[.4gl]
```

1. In this form, fglcomp adds a [`$` dollar sign before SQL host variables](2644-mark-sql-host-variables-with.md "The fglcomp compiler can mark SQL host variables with a $ dollar sign.").
2. refact-options are described in refactoring options.
3. prepro-options are described in preprocessor options.

## Syntax 8: Fix symbols character case

```
fglcomp --fix-case [refact-options]
     [prepro-options] module[.4gl]
```

1. In this form, fglcomp makes [symbol
   match their definition exactly](2643-make-symbols-case-match-definition.md "The fglcomp compiler can make all symbol names match their definition exactly.").
2. refact-options are described in refactoring options.
3. prepro-options are described in preprocessor options.

## Syntax 9: Information options

```
fglcomp info-option
```

1. info-option can be any of the informational options.

## Options

| Option | Description |
| --- | --- |
| `-M` | Write error messages to standard output instead of creating a .err error file. |
| `-k` or `--keep-going` | When compiling a set of .4gl files passed as arguments, if a source file produces an error, the remaining files are processed. See [Compiling several .4gl sources in a single command](2534-compiling-program-code-files-4gl.md). |
| `--make` | Do not recompile .42m modules up-to-date with provided .4gl source. See [Compiling in make mode](2534-compiling-program-code-files-4gl.md). |
| `--simulate` | Print the list of .4gl sources that would be compiled (can be used with or without `--make`). See [Compiling in make mode](2534-compiling-program-code-files-4gl.md). |
| `-W warning-argument` | Produce warning messages.The possible warning arguments are listed in the table Table 2.The `-W` option also supports the negative form of arguments by using the `no-` prefix as in: `no-return`, `no-unused`, `no-stdsql`. You might need to use these negative forms in order to disable some warning when using the `-W all` option:fglcomp -W all -W no-stdsql customers.4glSwitches will be enabled/disabled in the order of appearance in the command line. |
| `--timestamp` | Add compilation timestamp to build information in 42m header. See [42m module information](2539-42m-module-information.md "Describes how to handle module information in .42m p-code files."). |
| `--omit-source-name` | Omit the source filename in the build information of the 42m header. See [42m module information](2539-42m-module-information.md "Describes how to handle module information in .42m p-code files."). |
| `-o` or `--output-dir` | Specify the output directory where .42m files must be created. When using [packages](../09_advanced-features/0816-package.md "Defines the package the module belongs to."), the directory structure of sources will be mirrored to the output directory. fglcomp automatically creates the target and intermediate directories if they do not exist. See also [Output directory for .42m pcode files](2534-compiling-program-code-files-4gl.md). |
| `--tag=string` | Write a custom string in the build information of the 42m header. See [42m module information](2539-42m-module-information.md "Describes how to handle module information in .42m p-code files."). |
| `--build-rdd` | While compiling, generate the module.rdd Report Data Definition file (of `REPORT` routines). |
| `--verbose` | Print detailed compilation information. |
| `--implicit=type` | Specify whether or not to compile imported modules, if the .42m does not exist, or if the .4gl source is more recent as the .42m.Here type can be one of:`none`: Disable [auto-compilation of imported modules](../09_advanced-features/0819-auto-compilation-of-imported-modules.md "Imported local and package modules are compiled automatically if needed.").`42m` (default): Compile [imported modules](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") if needed. |
| `-r` or `--resolve-calls` | Throw an error on references to undeclared functions. Each external function must be made known to the compiler by `IMPORT FGL`. When using this option, the linking phase is no longer needed; a source (.4gl) file compiled with this option must not be linked. See [IMPORT FGL](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") for more details. |
| `--java-option=option` | Passes Java runtime options when initializing the JNI interface.See [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") for more details. |
| `-S` | With this option, fglcomp extracts all [Static SQL](../10_sql-support/1115-static-sql-statements.md "Describes static SQL statements supported in the language.") commands from the source while compiling the source code. |

| -W argument | Included in `-W all` | Description |
| --- | --- | --- |
| `all` | N/A | `-W all` enables all warning flags. |
| `error` | N/A | `-W error` makes the compiler stop, if any warning is raised, as if an error occurred. |
| `to-err-file` | N/A | `-W to-err-file` writes warnings to the .err file when this file is produced. By default warnings go to the stderr stream. |
| `form-field-name` | Yes | `-W form-field-name` instructs the compiler to produce warnings for invalid field names used in `INPUT BY NAME` and `CONSTRUCT BY NAME` dialogs, for example with `NEXT FIELD`. A runtime error will be raised whenever fglrun executes this code. |
| `fragile-cursor` | Yes | `-W fragile-cursor` detects risky SQL cursor usage that must be reviewed:`DECLARE` statements using local function host/into variables, with `OPEN`, `FETCH` or `FOREACH` wihout `USING/INTO` clause, in another function scope.`OPEN` with `USING` clause, followed by `OPEN` or `FOREACH` without `USING` clause, that can lead to unexpected results with Informix OPTOFC option. |
| `unused` | Yes | `-W unused` displays a message for all declared symbols not used. Includes function parameters, defined variables, as well as unused [`DECLARE`](../10_sql-support/1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.") and [`PREPARE`](../10_sql-support/1142-prepare-sql-statement.md "Prepares an SQL statement for execution.") statements and unused [imported modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module."). |
| `unused-parameter` | Yes | `-W unused-parameter` displays a message for all unused function parameters. |
| `return` | Yes | `-W return` displays a warning if the same function returns different number of values with several [`RETURN`](../08_language-basics/0768-returning-values.md "A function can return values with the RETURN instruction.") statements. |
| `stdsql` | Yes | `-W stdsql` displays a message for all non-portable SQL statements or language instructions. |
| `colname` | Yes | `-W colname` displays a message for invalid column names in static SQL statements, using the current database schema and table names of the statement. See [`-W colname` usage](../10_sql-support/1117-using-program-variables-in-static-sql.md). |
| `case` | Yes | `-W case` displays a message when a symbol does not match the character case of its definition. See [Make symbols case match definition](2643-make-symbols-case-match-definition.md "The fglcomp compiler can make all symbol names match their definition exactly."). |
| `print` | Yes | `-W print` displays a message when the [`PRINT`](../12_reports/2491-print.md "Formats and prints a row of data in a report routine.") instruction is used outside a `REPORT`. |
| `implicit` | Yes | `-W implicit` warns on references to undeclared functions. A function is undeclared if not defined in the current module or in any [imported module](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module."). |
| `api-doc` | Yes | `-W apidoc` prints a warning for invalid source documentation tags when using the [`--build-doc` option](2546-source-documentation.md "Explains how to automatically generate documentation from your sources."). |
| `unqualified-imports` | No | `-W unqualified-imports` prints warnings for any imported symbol that is not prefixed with its module. This warning is not included in `-W all`. See [Qualifying imported symbols](2642-qualifying-imported-symbols.md "The fglcomp compiler can automatically qualify imported symbols."). |
| `keywords` | Yes | `-W keywords` prints a warning when a symbol uses a name that matches a keyword and can lead to code ambiguity.This warning is enabled by default, and can be disabled with `-W no-keywords` . |
| `shadow` | No | `-W shadow` detects symbols hiding symbols defined at a higher scope. |

| Option | Description |
| --- | --- |
| `--doc-private` | When using the `--build-doc` option, include `PRIVATE` symbols to the documentation. |

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

| Option | Description |
| --- | --- |
| `-E` | Preprocess only. See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |
| `-p option` | Preprocessing control, where option can be one of:`nopp`: Disable preprocessing.`noli`: No line number information (only with `-E` option).`fglpp`: Use # syntax instead of & syntax.`auto`: Detect # syntax or & syntax automatically. |
| `-I path` | Provides a single path to search for include files. See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |
| `-D ident[=value]` | Defines the macro 'ident' with an optional value (default is 1). See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |
| `-U ident` | Undefines the macro 'ident'. See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |

| Option | Description |
| --- | --- |
| `--inplace` | Instead of writing the refactoring result into stdout, replace the original source file with the refactored code. When using this option, fglcomp makes a copy of the original source, using the .4gl~ file extension. |

| Option | Description |
| --- | --- |
| `-V` or `--version` | Displays version information. |
| `-h` or `--help` | Displays options for the tool. |
| `-i [ mbcs ]` | Displays information about the current locale / character set settings. See [Application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."). |

## Usage

The fglcomp command line tool compiles a .4gl into a
.42m p-code module:

```
fglcomp customers.4gl
```

If a compilation error occurs, the compiler generates an error file with an
.err extension. The error file contains the original source code with
error messages. Use the option `-M` to display the error messages to standard
error instead of producing the .err file.

If case of error, the fglcomp command execution status is different from zero.
Consequently, it is possible to detect compilation errors in scripts and makefiles.

## Related links

**Related concepts**  

[Compiling program code files (.4gl)](2534-compiling-program-code-files-4gl.md "The .4gl source files must be compiled to .42m p-code files, in order to be loaded by the runtime system.")

[fglform](2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.")
