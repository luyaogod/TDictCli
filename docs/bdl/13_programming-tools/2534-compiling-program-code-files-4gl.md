---
title: "Compiling program code files (.4gl)"
source: "fgl-topics/c_fgl_CompilingPrograms_002.html"
breadcrumb: "Programming tools > Compiling source files > Compiling program code files (.4gl)"
type: "concept"
---

# Compiling program code files (.4gl)

> The .4gl source files must be compiled to .42m p-code files, in order to be loaded by the runtime system.

## Understanding .4gl source compilation

Genero BDL source code modules (with .4gl file extension) must be compiled
to p-code modules (with .42m file extension) by using the [fglcomp](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") tool.

Compiled p-code modules are independent of the platform and processor architecture. They are
interpreted by the Genero runtime system ([fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.")).

The following lines show the compilation of the prog.4gl source, in a UNIX™ shell session:

```
$ cat prog.4gl
MAIN
  DISPLAY "hello"
END MAIN

$ fglcomp prog.4gl

$ ls -s prog.42m
   4 prog.42m
```

## Finding modules with FGLLDPATH

The fglcomp compiler uses the list of directories defined in the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable, in order to find the
.4gl sources and .42m compiled modules imported by the
.4gl source file(s) passed as argument to fglcomp.

## Output directory for .42m pcode files

By default, fglcomp creates the .42m files in the current
working directory, for the source files specified in the command line, and for the modules imported
with `IMPORT FGL module-name`, when these are located in the
current working directory.

The output directory for .42m files can be specified with the [`--output-dir` option](2516-fglcomp.md) of fglcomp. If the directory does
not exist, it will be automatically created.

When importing modules via [packages](../09_advanced-features/0816-package.md "Defines the package the module belongs to."), the
.42m files of the modules that belong to a package are by default created in
the directory of the .4gl source files. This is mandatory since the same
modules names can be reused in different packages.

When using packages, the output directory specified with the `--output-dir` option
defines the top-dir to create a tree of sub-directories containing the
.42m files of modules imported via packages. This structure can directly be
used as runtime distribution directory tree, since modules imported via packages are found
automatically by fglrun, from the package-path reflecting the
directory path. The FGLLDPATH environment variable can be used to specify the
top-dir.

## Automatic compilation of imported modules

When compiling a .4gl module that imports other modules with the
`IMPORT FGL` instruction, if the .4gl source is more recent as
the .42m file or the .42m file does not exist,
fglcomp will automatically compile the imported modules, if they are located in
the same directory as the compiled module, or when the module is imported via a package and located
in a different directory.

For more details about auto-compilation of imported modules, see [Auto-compilation of imported modules](../09_advanced-features/0819-auto-compilation-of-imported-modules.md "Imported local and package modules are compiled automatically if needed.").

## Verbose compilation

Consider using the `--verbose` option of the compiler to get detailed information
about the source compilation:

```
$ fglcomp --verbose main.4gl
[parsing main.4gl]
[parsing mod1.4gl]
[parsing mod2.4gl]
[building mod2]
[writing mod2.42m]
[building mod1]
[writing mod1.42m]
[building main]
[writing main.42m]
...
```

## Compiling several .4gl sources in a single command

Several .4gl source files can be provided to fglcomp: The
compiler builds a dependency tree of imported modules, and compiles the files in the calculated
order.

When compiling multiple .4gl source files with a single
fglcomp command, the process stops, if an error occurs during the compilation of
a source file.

For example (assuming no .42m file exists before executing this command), if
the main.4gl module imports module1.4gl, but does not
import module2.4gl, when passing main.4gl module1.4gl
module2.4gl as arguments to fglcomp, the compiler will first compile
the module2.4gl because it is independent. Then
module1.4gl is compiled,because it is imported by
main.4gl, and finaly main.4gl:

```
$ rm *.42m
$ fglcomp --verbose main.4gl module1.4gl module2.4gl
[parsing module2.4gl]
[building module2]
[writing module2.42m]
[parsing module1.4gl]
[building module1]
[writing module1.42m]
[parsing main.4gl]
[building main]
[writing main.42m]
...
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
source files, in case one of the source files passed as argument fails to
compile:

```
$ fglcomp -M --verbose --keep-going "module?.4gl"
[parsing module1.4gl]
[building module1]
[writing module1.42m]
[parsing module2.4gl]
module2.4gl:2:13:2:30:error:(-6609) A grammatical error has been found ...
[parsing module2 cancelled: 1 errors]
[parsing module3.4gl]
[building module3]
[writing module3.42m]
[fglcomp: *** Errors in 1 of 3 files]
```

> **Tip:**
>
> fglcomp provides the `--make` option, to optimize the
> compilation of a large number of modules passed as arguments to the compiler. See Compiling in make mode.

## Stricter compilation options

By default, fglcomp allows you to compile legacy application sources in a
relaxed manner. For example, it is possible to compile a module calling a function which is not
defined in that module, and resolved during the link process.

In recent developments based on features such as `IMPORT FGL`, consider using
fglcomp options that make the compilation stricter:

- `--resolve-calls`: To check that all functions used by a module are defined
  locally, or by imported modules.
- `-W all`: To print all type of warnings.
- `-W error`: To stop compilation if a warning occurs.
- `-M`: To write error messages to standard output instead of producing a
  .err file.

For more details about fglcomp options, see [the fglcomp
command reference topic](2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.").

Additionally, the [linker](2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.") provides options
such as `--print-imports` and `--print-missing-imports`, that produce
useful information to use new language features.

## Providing the source files in an arguments file

The fglcomp compiler supports the `@argfile`
argument, to provide a file that contains the list of source files to be compiled. This can be used
when it is not possible to pass all source files in the command line.

Only source files and patterns must be specified in the arguments file. Compiler options must be
provided in the command line.

The argument file must contain one source file per line:

```
$ cat mysources.txt
module1.4gl
module2.4gl
module3.4gl
$ fglcomp -M -W all @mysources.txt
```

Note that the arguments file can contain expansion patterns as in the fglcomp
command line:

```
$ cat mysources.txt
module[123]*.4gl
$ fglcomp -M -W all @mysources.txt
```

Multiple `@argfile` files can be provided, and mixed with
filenames and patterns in the command line. The fglcomp compiler will build a
list of sources to be compiled.

## Handling fglcomp compiler errors

If an error occurs, the compiler writes an error file by default with the
.err extension.

```
$ cat prog.4gl
MAIN
  LET x = "hello"
END MAIN

$ fglcomp prog.4gl
Compilation was not successful. Errors found: 1.
 The file prog.4gl has been written.

$ cat prog.err
MAIN
  LET x = "hello"
| The symbol 'x' does not represent a defined variable.
| See error number -4369. 
END MAIN
```

With the `-M` option, you can force the compiler to display an error message
instead of generating an .err error file:

```
$ fglcomp prog.4gl
xx.4gl:2:8 error:(-4369) The symbol 'x' does not represent a defined variable.
```

## Producing compiler warnings with -W

To improve code quality, enable compiler warnings with the `-W` option:

```
$ cat prog.4gl
MAIN
  DATABASE test1
  SELECT COUNT(*) FROM x, OUTER(y) WHERE x.k = y.k
END MAIN

$ fglcomp -W stdsql prog.4gl
xx.4gl:3: warning: SQL statement or language instruction with specific SQL syntax.
```

By default warnings go to the stderr stream. When creating a
.err file, warnings can be redirected to the .err file
with the `-W to-err-file option`.

When a warning is raised, you can use the `-W error` option to force the compiler
to stop as if an error was found.

Some warnings are generated by default (without using the `-W` option), when the
source code uses a feature that is considered as "fragile" yet to be supported for backward
compatibility.

The `-W unused` warning option is a good starting point to cleanup your code: When
a source code is modified several times, it can happen that a variable definition is left, even if
the variable is no longer used. The same can happen for function parameters, as well as unused [`DECLARE`](../10_sql-support/1150-declare-result-set-cursor.md "Associates a database cursor with an SQL statement producing a result set.") and [`PREPARE`](../10_sql-support/1142-prepare-sql-statement.md "Prepares an SQL statement for execution.") statements and unused [imported modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.").

To detect such mistakes and cleanup the source, fglcomp provides the
`-W unused` option:

```
$ cat p.4gl
MAIN
    DEFINE v INT
END MAIN
FUNCTION func1(p1 INT, p2 INT) RETURNS()
    DISPLAY p1
END FUNCTION

$ fglcomp -M -Wunused p.4gl
p.4gl:2:12:2:12:warning:(-6615) The symbol 'v' is unused.
p.4gl:4:24:4:25:warning:(-6615) The parameter 'p2' is unused.
```

The `-W unused-parameter` option is dedicated to the detection of unused function
parameters. It can be used individually, or it can be combined to `-W unused` in its
negative form `-W no-unused-parameter`, to get all unused symbols warnings,
except for function parameters.

For the complete list of warning options, see [`-W`
option](2516-fglcomp.md) in fglcomp command reference.

## Compiling in make mode

The purpose of the `--make` option is to optimize the compilation when a large set
of modules is provided to the compiler.

With the `--make` option, fglcomp compiles a
.4gl source, if the .42m file is older than the
.4gl file. If the .42m is up-to-date with the
.4gl, it is not recompiled.

As with a regular compilation (not using the `--make` option), for a provided
.4gl module, imported modules will be [automatically compiled](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") when needed.

To test the compilation process and dependencies without actually compiling the modules, add the
`--simulate` option, to see what modules would be compiled. The
`--simulate` option can also be used without the `--make` option:
fglcomp will then print the modules that would be compiled, even if the
.42m file is up to date.

The fglcomp compiler produces the .42m files in the
current working directory. In the next example, we use a bin directory, to hold
all .42m files, and run the compiler in this directory by providing relative
paths of the .4gl sources. It is also possible to use the
`--output-dir` option:

```
$ tree .
.
|-- bin
|-- dir1
|   |-- module1.4gl
|   |-- module2.4gl
|   `-- module3.4gl
|-- prog1.4gl
`-- prog2.4gl

2 directories, 5 files

$ head *.4gl */*.4gl
==> prog1.4gl <==
IMPORT FGL module1
MAIN
    CALL module1.function1()
END MAIN

==> prog2.4gl <==
IMPORT FGL module2
MAIN
    CALL module2.function2()
END MAIN

==> dir1/module1.4gl <==
FUNCTION function1()
END FUNCTION

==> dir1/module2.4gl <==
IMPORT FGL module3
FUNCTION function2()
    CALL module3.function3()
END FUNCTION

==> dir1/module3.4gl <==
FUNCTION function3()
END FUNCTION
```

Go to the bin directory:

```
$ cd bin
```

Assuming there are no .42m in this directory tree, first check what would be
compiled if we pass all .4gl sources to
fglcomp:

```
$ fglcomp --make --simulate ../*.4gl ../*/*.4gl
[fglcomp ../dir1/module1.4gl]
[fglcomp ../prog1.4gl]
[fglcomp ../dir1/module3.4gl]
[fglcomp ../dir1/module2.4gl]
[fglcomp ../prog2.4gl]
```

Now let's do the real compilation by replacing the `--simulate` option by
`--verbose`, to see how modules are compiled (fglcomp produces the
.42m files for all the sources, since no .42m
exists):

```
$ fglcomp --make --verbose ../*.4gl ../*/*.4gl
[fglcomp ../dir1/module1.4gl]
[parsing ../dir1/module1.4gl]
[building module1]
[writing module1.42m]
[fglcomp ../prog1.4gl]
[parsing ../prog1.4gl]
[building prog1]
[writing prog1.42m]
[fglcomp ../dir1/module3.4gl]
[parsing ../dir1/module3.4gl]
[building module3]
[writing module3.42m]
[fglcomp ../dir1/module2.4gl]
[parsing ../dir1/module2.4gl]
[building module2]
[writing module2.42m]
[fglcomp ../prog2.4gl]
[parsing ../prog2.4gl]
[building prog2]
[writing prog2.42m]

$ ls 
module1.42m  module2.42m  module3.42m  prog1.42m  prog2.42m
```

Update module3.4gl and see how fglcomp recompiles the
modules, based on the `IMPORT FGL` dependence tree. Since
prog2.4gl imports module2.4gl, and
module2.4gl imports module3.4gl, fglcomp recompiles first
module3.4gl, then module2.4gl. The module
prog2.4gl is not recompiled, because module2.4gl was not
changed:

```
$ sleep 1 && touch ../dir1/module3.4gl
$ fglcomp --make --verbose ../*.4gl ../*/*.4gl
[fglcomp ../dir1/module3.4gl]
[parsing ../dir1/module3.4gl]
[building module3]
[writing module3.42m]
[fglcomp ../dir1/module2.4gl]
[parsing ../dir1/module2.4gl]
[building module2]
[writing module2.42m]
```

Touch module2.4gl, only this module and prog2.4gl are
compiled:

```
$ sleep 1 && touch ../dir1/module2.4gl
$ fglcomp --make --verbose ../*.4gl ../*/*.4gl
[fglcomp ../dir1/module2.4gl]
[parsing ../dir1/module2.4gl]
[loading module3.42m]
[building module2]
[writing module2.42m]
[fglcomp ../prog2.4gl]
[parsing ../prog2.4gl]
[building prog2]
[writing prog2.42m]
```

Execute again the command without touching sources files, nothing is
compiled:

```
$ fglcomp --make --verbose ../*.4gl ../*/*.4gl
[fglcomp: nothing to do]
```

## Producing make-style dependency rules

The fglcomp --dependencies option can be used to produce makefile-style
dependency rules for imported modules.

The dependency graph is constructed from the .4gl sources passed to the
fglcomp command. Package modules and modules located in the current working
directory that are imported by one of the .4gl modules specified to
fglcomp, are automatically included in the dependency rules. However, when
importing a module that is not part of a package and not located in the current working directory,
it needs to be specified in the .4gl list, in order to be included in the
dependency rules.

The FGLCOMP environment variable must be defined with the compiler command, to be used by the
make utility. This solution provides maximum flexibility to set the
fglcomp options adapted to your
needs:

```
$ export FGLCOMP="fglcomp -M -Wall"
```

In the next example, the main.4gl source imports `module1`
and thus the main.42m compilation depends from
module1.42m:

```
$ head *.4gl
==> main.4gl <==
IMPORT FGL module1

MAIN
    CALL module1.function1()
END MAIN

==> module1.4gl <==
FUNCTION function1()
END FUNCTION

$ fglcomp --dependencies main.4gl
main.42m: main.4gl module1.42m
	$(FGLCOMP) $<
module1.42m: module1.4gl
	$(FGLCOMP) $<
```

When a module is importing a non-existing module by mistake, it will be ignored by the
dependencies generation process.

Circular imports are handled with proper dependency rules that can even be used for parallel
builds (with make -j).

> **Tip:**
>
> Try to perform the command fglcomp --dependencies \*.4gl on all your sources
> and see the output.

The following generic makefile (requires GNU make) can be used to automatically produce a
`"depend"` file with .42m/`.4gl` module
dependencies, and compile your sources in an efficient way:

```
MODS=$(patsubst %.4gl,%.42m,$(wildcard *.4gl))
all:: depend $(MODS)
depend: 
        fglcomp --dependencies *.4gl > depend
clean:
        rm -f *.42? depend
FGLCOMP = fglcomp -M -Wall
-include depend
```

> **Note:**
>
> If you copy/paste the above code, make sure to start command lines in rules by a TAB
> character!

## Related links

**Related concepts**  

[Importing modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
