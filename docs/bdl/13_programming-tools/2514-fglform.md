---
title: "fglform"
source: "fgl-topics/c_fgl_tools_fglform.html"
breadcrumb: "Programming tools > Command reference > fglform"
type: "concept"
---

# fglform

> The fglform tool compiles form specification files into XML formatted files used by programs.

## Syntax 1: Compiling forms

```
fglform [comp-options] [prepro-options] file-list
```

1. In this form, fglform [compiles the form specification files](2531-compiling-form-specification-files-per.md "The .per form definition files must be compiled to .42f XML files, in order to be loaded by the runtime system.") to a .42f file.
2. comp-options are described in compilation options.
3. prepro-options are described in preprocessor options.

where file-list
is:

```
{ form-name[.per]
| pattern
| @argfile
} [...]
```

1. form-name.per is the form
   specification file. The .per extension is optional.
2. pattern is a `MATCHES`-style pattern to find files, like
   `'[a-z]*.per'`.
3. argfile defines a file that contains a list of .per
   sources to be compiled. Each line must specify a filename or a pattern.

## Syntax 2: Extracting localized strings

```
fglform -m [prepro-options] form-name[.per]
```

1. In this form, fglform extracts [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") from the source.
2. prepro-options are described in preprocessor options.
3. form-name.per is the form
   specification file. The .per extension is optional

## Syntax 3: Information options

```
fglform info-option
```

1. info-option can be any of the informational options.

## Options:

| Option | Description |
| --- | --- |
| `--make` | Do not recompile .42f form files up-to-date with provided .per source. See [Compiling in make mode](2531-compiling-form-specification-files-per.md). |
| `-M` | Write error messages to standard output instead of creating a .err error file. |
| `-o` or `--output-dir` | Specify the output directory where .42f files must be created. fglform automatically creates the target and intermediate directories if they do not exist. See also [Output directory for .42f form files](2531-compiling-form-specification-files-per.md). |
| `-k` or `--keep-going` | When compiling a set of .per files passed as arguments, if a form file produces an error, the remaining files are processed. See [Compiling several .per sources in a single command](2531-compiling-form-specification-files-per.md). |
| `--verbose` | Print detailed compilation information. |
| `-W { all }` | Produce warning messages. Only `-W all` option is supported for now. |

| Option | Description |
| --- | --- |
| `-E` | Preprocess only. See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |
| `-p option` | Preprocessing control, where option can be one of:`nopp`: Disable preprocessing.`noli`: No line number information (only with `-E` option).`fglpp`: Use # syntax instead of & syntax.`auto`: Detect # syntax or & syntax automatically. |
| `-I path` | Provides a single path to search for include files. See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |
| `-D ident[=value]` | Defines the macro 'ident' with an optional value (default is 1). See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |
| `-U ident` | Undefines the macro 'ident'. See [Source preprocessor](2562-source-preprocessor.md "A typical preprocessor like in the C language.") for more details. |

| Option | Description |
| --- | --- |
| `-V` or `--version` | Displays version information. |
| `-h` or `--help` | Displays options for the tool. |
| `-i [ mbcs ]` | Displays information about the current locale / character set settings. See [Application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."). |

## Usage

The fglform command line tool compiles a .per
form specification file into a .42f compiled
version:

```
fglform custform.per
```

The .per extension is optional, if not used, fglform will
automatically search for files with this extension.

The .42f compiled version is an XML formatted file used by
programs when a form definition is loaded with the [`OPEN FORM`](../11_user-interface/1578-open-form.md "Declares a compiled form in the program.")
or [`OPEN WINDOW WITH
FORM`](../11_user-interface/1572-open-window.md "Creates and displays a new window.") instructions.

If case of error, the fglform command execution status is different from zero.
Consequently, it is possible to detect compilation errors in scripts and makefiles.

## Related links

**Related concepts**  

[Compiling form specification files (.per)](2531-compiling-form-specification-files-per.md "The .per form definition files must be compiled to .42f XML files, in order to be loaded by the runtime system.")

[Form specification files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")

[fglcomp](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.")
