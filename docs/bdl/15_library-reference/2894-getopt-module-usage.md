---
title: "Getopt module usage"
source: "fgl-topics/c_fgl_utility_functions_getopt_lib_usage.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt module usage"
type: "concept"
---

# Getopt module usage

> The getopt.4gl module provides command line argument processing.

## Features of getopt.4gl

The getopt.4gl module implements types, functions and methods to process
command line arguments in a standard way.

Use this library to implement arguments processing, and provide a common command line option
syntax for all your programs.

The options can be defined with a short name (like `-c`) and a long name (like
`--compare`).

Partial option specification is supported. For example if `--compare` is defined,
a command line argument `--comp` will match `--compare`. If the
partial option matches several long option names, the error "ambiguous match" is displayed.

Options can have a mandatory or optional value to be provided with the
`--option-name=option-value`
form:

```
$ fglrun myprog --verbose --level=5
```

Option values can also be concatenated to short option
name:

```
$ fglrun myprog --verbose -l5
```

If an argument on the command line is in the form `@filename`,
options will be read from the file:

```
$ fglrun myprog @myoptions
```

Options can start at a given index in the command line arguments, to support for example commands
with verbs followed by
options:

```
$ fglrun myprog capture --verbose --filename=file1
$ fglrun myprog duplicate --source=file1 --destination=file2
```

Command line arguments can end with a set of free arguments, to be processed after the
options:

```
$ fglrun myprog capture --verbose file1 file2 file3
```

## Steps to implement arguments processing

To implement command line arguments processing with `getopt`, do the following steps:

1. Import the `getopt` module:

   ```
   IMPORT FGL getopt
   ```
2. Define a variable with the [`Getopt` type](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options."):

   ```
       DEFINE g getopt.Getopt
   ```
3. Define a variable with the [`GetoptOptions` type](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.") that will contain the definition of the options (fill
   this array of options with a [variable
   initializer](../08_language-basics/0691-variable-initializers.md "Variables can be initialized in their definition.")):

   ```
       DEFINE _options getopt.GetoptOptions
           = [(name: "version",
               description: "Version information",
               opt_char: 'v',
               arg_type: getopt.NONE),
              ...
   ```
4. Initialize the `Getopt` variable with the [`initialize()`](2904-getopt-initialize.md "Initializes a variable defined with the Getopt type for command line argument processing.") or [`initDefault()`](2903-getopt-initdefault.md "Initializes a variable defined with the Getopt type.") method,
   passing the GetoptOptions array as
   parameter:

   ```
       CALL g.initialize("myprog", getopt.copyArguments(2), _options)
   ```

   If
   needed, use the [`copyArguments(index)`](2898-getopt-copyarguments.md "Returns a dynamic array of string with all command line arguments starting from the provided index.") function, to provide the second parameter of the
   `initialize()` method, to start command line argument parsing at a given index.
5. Use the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.")
   method in a [`WHILE`](../08_language-basics/0685-while.md "The WHILE statement executes a block of statements until the specified condition becomes false.") loop, to process
   all command line arguments that correspond to an option
   definition:

   ```
       WHILE g.getOpt() == getopt.SUCCESS
           CASE g.opt_char
               WHEN 'v'
                   DISPLAY "Version 1.50"
                   EXIT PROGRAM 0
               WHEN 'h'
                   CALL g.displayUsage("file ...")
                   EXIT PROGRAM 0
               WHEN 'o'
                   LET outfile = g.opt_arg
           END CASE
       END WHILE
   ```

   Additionally, the
   [`isEof()`](2906-getopt-iseof.md "Checks if there are more command line options to be read.") method can
   be used to check if all possible options are processed.
6. After the `WHILE` loop, check for the processing status with the methods [`invalidOptionSeen()`](2905-getopt-invalidoptionseen.md "Checks if the command line options are misused.") or
   [`isSuccess()`](2907-getopt-issuccess.md "Checks if a command line option parsing succeeded."):

   ```
       IF g.invalidOptionSeen() THEN
           CALL g.displayUsage("file ...")
           EXIT PROGRAM 1
       END IF
   ```
7. If additional arguments are possible, use the [`getMoreArgumentCount()`](2900-getopt-getmoreargumentcount.md "Returns the number of command line arguments left to be processed after the known options.") and [`getMoreArgument()`](2901-getopt-getmoreargument.md "Returns the additional argument at the specified index.")
   methods, to process these non-option
   arguments:

   ```
           LET cnt = g.getMoreArgumentCount()
           IF cnt == 0 THEN
               DISPLAY "ERROR: No files were provided..."
               EXIT PROGRAM 1
           ELSE
               FOR ind = 1 TO cnt
                   DISPLAY SFMT("File to process: %1", g.getMoreArgument(ind))
               END FOR
           END IF
   ```

For a complete example, see [Getopt.getopt()](2902-getopt-getopt.md "Process the next command line option.").
