---
title: "fglrun"
source: "fgl-topics/c_fgl_tools_fglrun.html"
breadcrumb: "Programming tools > Command reference > fglrun"
type: "concept"
---

# fglrun

> The fglrun tool is the runtime system program that executes p-code programs.

## Syntax 1: Executing programs

```
fglrun [exec-options] program [ argument [...] ]
```

1. In this form, fglrun [executes the
   program](../09_advanced-features/0828-program-execution.md "This section describes program execution and language instructions related to program execution.").
2. exec-options can be execution
   options as well as trace options.
3. program is a .42r program, or a
   .42m p-code module containing the `MAIN` definition. This can be
   a simple base name, or a filename including the path to the program file.
4. argument is an argument passed to the program.

## Syntax 2: Starting the debugger

```
fglrun -d [exec-options] program
```

1. In this form, fglrun starts the program [in debug mode](2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs.").
2. exec-options can be execution options (trace options are not supported).
3. program is a .42r program, or a
   .42m p-code module containing the `MAIN` definition. This can be
   a simple base name, or a filename including the path to the program file.

## Syntax 3: Starting the debug-server

```
fglrun --da-listen [host:]port
       --da-pid pid
```

Or:

```
fglrun --da-listen [host:]port
       --da-run program [ argument [...] ]
```

1. When using the `--da-listen` option, fglrun [starts in debug-server mode](2585-using-the-debug-server.md "The Genero BDL debug-server is a proxy for fglrun processes which can not be accessed directly by the debugger."), to let the real debugger attach
   to a running process that cannot be accessed directly.
   - With the `--da-pid pid` option, the debug-server binds to an
     already running fglrun process, identified by its process id.
   - With the `--da-run program` option, the debug-server starts a
     new fglrun process with the specified program, then bind to this process.
2. host is the name of a host or IPv4 address to be used to establish the TCP
   connection. If host is not specified, the debug-server will listen on localhost.
   If you want to listen to any interfaces, use `"0.0.0.0"`.
3. port is the TCP port number to be used to listen to debugger connections. The
   TCP port is mandatory. When specifying `0` (zero), the runtime will automatically
   select a free port in the range of ephemeral TCP ports.
4. pid as a process id.
5. program is a .42r program, or a
   .42m p-code module containing the `MAIN` definition. This can be
   a simple base name, or a filename including the path to the program file.
6. argument is an argument passed to the program.

## Syntax 4: Linking programs or libraries

```
fglrun -l -o outfile{.42r|.42x} [link-options] file-list
```

1. In this form, fglrun [links a
   program or library](2530-compiling-source-files.md "Describes how to build the runtime files from source files.").
2. link-options are described in Table 4.
3. outfile is the .42r program (or a
   .42x library) to produce from the link.

where file-list
is:

```
{ { module.42m | library.42x }
| pattern
| @argfile
} [...]
```

1. module.42m is a p-code module
   compiled with fglcomp.
2. library.42x is a name of a library to be used for
   linking.
3. pattern is a `MATCHES`-style pattern to
   find files, like `'[a-z]*.42m'`.
4. argfile defines a file that contains a list of
   .42m or .42x files. Each line must specify a filename or a
   pattern.

## Syntax 5: Using diagnostic options

```
fglrun diagnostic-option file-list
```

1. In this form, fglrun produces diagnostic data.
2. diagnostic-option is one of the diagnostic options.

where file-list is:

```
{ module.42m
| pattern
| @argfile
} [...]
```

1. module.42m is a p-code module
   compiled with fglcomp.
2. pattern is a `MATCHES`-style pattern to
   find files, like `'[a-z]*.42m'`.
3. argfile defines a file that contains a list of .42m
   files. Each line must specify a filename or a pattern.

## Syntax 6: Information options

```
fglrun info-option
```

1. info-option can be any of the informational options.

## Options

| Option | Description |
| --- | --- |
| `-e extfile[,...]` | Specify a [C extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") module to be loaded. This option can take a comma-separated list of extensions. |
| `--java-option=option` | Passes Java runtime options when initializing the JNI interface.See [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") for more details. |

| Option | Description |
| --- | --- |
| `-p` | Generate profiling information to stderr. See [Program profiler](2622-program-profiler.md "Find out what function is causing the bottleneck in your program."). |
| `--start-guilog=logfile` | Log all GUI protocol exchange in a file. The GUI log file can then be replayed with the `--run-guilog` option. If the log file contains a `%p` placeholder, it is replaced by the current process id. |
| `--run-guilog=logfile` | Replays a GUI log created with the `--start-guilog` option. |
| `--gui-listen=port` | Instructs the runtime system to listen to a TCP port for incoming GUI connections. For more details see [Connecting with a front-end](../11_user-interface/1521-connecting-with-a-front-end.md). |
| `--trace` | Starts the program by printing function call stack trace with parameter and return values. For more details, see [Execution trace](2632-execution-trace.md "Print a function call stack of your program."). |

| Option | Description |
| --- | --- |
| `-b` | Displays compiler version information of the module, see [Compiling source files](2530-compiling-source-files.md "Describes how to build the runtime files from source files."). |
| `--print-imports` | Loads the specified modules and prints all `IMPORT FGL` instructions required in each module. See [Identifying modules to be imported](../09_advanced-features/0821-identifying-modules-to-be-imported.md "Use the --print-missing-imports and --print-imports options to identify missing IMPORT FGL instructions."). |
| `--print-missing-imports` | Loads the specified modules and prints all missing `IMPORT FGL` instructions for each module. See [Identifying modules to be imported](../09_advanced-features/0821-identifying-modules-to-be-imported.md "Use the --print-missing-imports and --print-imports options to identify missing IMPORT FGL instructions."). |
| `--module-size` | Displays the pcode size for a single module.The argument of the `--module-size` option must be a .42m module. See [Computing the p-code size of a module or program](2539-42m-module-information.md). |
| `--program-size` | Displays the pcode size for an entire program.The argument of the `--program-size` option can be a .42r program file (for linked programs), or the .42m module containing `MAIN` (for programs using `IMPORT FGL`). See [Computing the p-code size of a module or program](2539-42m-module-information.md). |
| `--merge-cov name.4gl [name.42m.cov]` | Merges FGLCOV coverage data files with source files to produce a readable file name.4gl.cov module. If the .42m.cov file is not located beside the source file, you can specify the full path with the second optional file path. See [Source code coverage](2628-source-code-coverage.md "Collect information about used source lines"). |

| Option | Description |
| --- | --- |
| `-e extfile[,...]` | Specify a [C extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") module to be loaded. This option can take a comma-separated list of extensions. |
| `--print-imports` | Loads the specified modules and prints all `IMPORT FGL` instructions required in each module. See [Identifying modules to be imported](../09_advanced-features/0821-identifying-modules-to-be-imported.md "Use the --print-missing-imports and --print-imports options to identify missing IMPORT FGL instructions."). |
| `--print-missing-imports` | Loads the specified modules and prints all missing `IMPORT FGL` instructions for each module. See [Identifying modules to be imported](../09_advanced-features/0821-identifying-modules-to-be-imported.md "Use the --print-missing-imports and --print-imports options to identify missing IMPORT FGL instructions."). |

| Option | Description |
| --- | --- |
| `-V` or `--version` | Displays version information. |
| `-h` or `--help` | Displays options for the tool. |
| `-i [ mbcs ]` | Displays information about the current locale / character set settings. See [Application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application."). |

## Usage

The fglrun command line tool executes p-code programs, for
example:

```
fglrun myprogram.42r -x 123
```

The program file must contain the [`MAIN`](../09_advanced-features/0789-the-main-block-function.md "The MAIN block is the starting point of the program.")
routine.

The arguments passed to the program can be queried with the [arg\_val()](../15_library-reference/2726-arg-val.md "Returns a command line argument by position.") built-in function.

The .42r or .42m extension is optional:

```
fglrun myprogram -x 123
```

The program name passed to fglrun can be a simple base name of the program
file, or an absolute or relative filename including the path to the program file.

To find the program file, fglrun does the following:

1. First fglrun tries to find the program file with the filename as provided in
   the command line. The filename can have a different extension than .42r or
   .42m.
2. If the file is not found, a new search is done by adding the .42r
   extension.
3. If the file is still not found, fglrun tries with the
   .42m extension.
4. If the file is still not found, fglrun produces the error -4448.

## Related links

**Related concepts**  

[Executing programs](../09_advanced-features/0829-executing-programs.md "There are different ways to execute compiled programs, depending on the configuration and the development or production context.")
