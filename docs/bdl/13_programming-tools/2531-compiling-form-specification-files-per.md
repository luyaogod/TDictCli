---
title: "Compiling form specification files (.per)"
source: "fgl-topics/c_fgl_CompilingPrograms_forms.html"
breadcrumb: "Programming tools > Compiling source files > Compiling form specification files (.per)"
type: "concept"
---

# Compiling form specification files (.per)

> The .per form definition files must be compiled to .42f XML files, in order to be loaded by the runtime system.

## Understanding .per source compilation

Form specification files (with .per file extension) must be compiled to
runtime form files (with .42f file extension) by using the [fglform](2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.") tool.

Compiled form files are XML independent from the platform and processor architecture.

The following lines show a compilation in a UNIX™ shell
session:

```
$ cat form.per
LAYOUT
GRID
{
[f01   ]
}
END
END
ATTRIBUTES
f01 = FORMONLY.field1;
END

$ fglform form.per

$ ls -s form.42f
   4 form.42f
```

## Output directory for .42f form files

By default, fglform creates the .42f files in the current
working directory, for the source files specified in the command line.

The output directory for .42f files can be specified with the [`--output-dir` option](2514-fglform.md) of fglform. If the directory does
not exist, it will be automatically created.

```
$ fglform --output-dir ./bin pm_form.per
$ ls ./bin
pm_form.42f
```

## Compiling in make mode

The fglform compiler supports the `--make` option, to compile
all provided sources, in the way the make utility does: If the
.42f file is older than the .per file, it will be
recompiled. If the .42f is up-to-date with the .per, it is
not recompiled.

```
$ rm -f *.42f
$ fglform --make pm_form.per
[fglform pm_form.per]
$ ls *.42f
pm_form.42f
$ fglform --make pm_form.per
[fglform: nothing to do]
$ touch pm_form.per
$ fglform --make pm_form.per
[fglform pm_form.per]
```

## Verbose compilation

Consider using the `--verbose` option of the compiler to get detailed information
about the source compilation:

```
$ fglform --verbose *.per
[parsing main_form.per]
[wrote main_form.42f]
[parsing form_101.per]
[wrote form_101.42f]
```

## Automatic compilation of imported forms

When compiling a .per module that includes other forms with the
`FORM` instruction, fglform will automatically compile the included forms, if the
.per source is more recent as the .42f file. The included
forms can be located in a different directory as the main form.

For more details, see [FORM clause](../11_user-interface/1716-form-clause.md "Reuse the definition of a form in the current form.").

## Compiling several .per sources in a single command

Several .per source files can be provided to fglform.

When compiling multiple .per source files with a single
fglorm command, the process stops, if an error occurs during the compilation of a
source file.

For example:

```
$ rm *.42f
$ fglform --verbose form1.per form2.per form3.per
[parsing form1.per]
[wrote form1.42f]
[parsing form2.per]
[wrote form2.42f]
[parsing form3.per]
[wrote form3.42f]
```

The [fglcomp](2534-compiling-program-code-files-4gl.md "The .4gl source files must be compiled to .42m p-code files, in order to be loaded by the runtime system.") and [fglform](2531-compiling-form-specification-files-per.md "The .per form definition files must be compiled to .42f XML files, in order to be loaded by the runtime system.") compilers support `MATCHES`-style pattern on the
command line, to achieve pathname expansion:

```
fglcomp [a-z]*.4gl
```

This feature
exists on any supported platform, including Microsoft™ Windows®.
> **Note:**
>
> On UNIX plarforms, pathname expansion is part of the shell
> interpreter. File names are resolved before Genero commands get the specified argument(s). In order
> to use the Genero build-in pathname expansion on a Unix platform, specify the search pattern in
> single quotes:
>
> ```
> fglcomp '[a-z]*.4gl'
> ```

For more details about supported patterns and wildcards, see the [os.Path.glob](../15_library-reference/3721-os-path-glob.md "Returns a list of files matching the specified pattern.") utility method.

Use the `--keep-going` option, to continue with the compilation of the remaining
form files, in case one of the form files passed as argument fails to
compile:

```
$ fglform --verbose --keep-going "form?.per"
[parsing form0.per]
[building cancelled form0]
fglcomp form0.per: *** Error
The compilation was not successful.  Errors found: 18.
The file 'form0.err' has been written.
[parsing form1.per]
[wrote form1.42f]
[parsing form2.per]
[wrote form2.42f]
[fglform: *** Errors in 1 of 3 files]
```

## Providing the source files in an arguments file

The fglform compiler supports the `@argfile`
argument, to provide a file that contains the list of source files to be compiled. This can be used
when it is not possible to pass all source files in the command line.

Only source files and patterns must be specified in the arguments file. Compiler options must be
provided in the command line.

The argument file must contain one source file per line:

```
$ cat mysources.txt
form1.per
form2.per
form3.per
$ fglform -M -W all @mysources.txt
```

Note that the arguments file can contain expansion patterns as in the fglcomp
command line:

```
$ cat mysources.txt
form[123]*.per
$ fglform -M -W all @mysources.txt
```

Multiple `@argfile` files can be provided, and mixed with
filenames and patterns in the command line. The fglform compiler will build a
list of sources to be compiled.

## Handling fglform compiler errors

If an error occurs, the compiler writes an error file with the .err
extension.

```
$ cat form.per
LAYOUT
GRID
{
}

$ fglform form.per
The compilation was not successful.  Errors found: 1.
 The file 'form.err' has been written.

$ cat form.err
LAYOUT
GRID
{
}
# A grammatical error has been found at '}', expecting SCR_TEXT.
# See error number -6803.
```

With the `-M` option, you can force the compiler to display an error message
instead of generating an .err error file (line break added for documentation
readability):

```
$ fglform -M form.per
form.per:4:1:4:1:error:(-6803)
 A grammatical error has been found at '}', expecting SCR_TEXT.
```

## Produce compiler warnings with -W

By default, the compiler does not raise any warnings. You can turn on warnings with the
`-W` option:

```
$ cat form.per
LAYOUT
GRID
{
[f01     ]
}
END
END
ATTRIBUTES
f01 = FORMONLY.field1, QUERYCLEAR;
END

$ fglform -W all form.per
form.per:9: warning (-8005) Deprecated feature: The QUERYCLEAR attribute is ignored
```
