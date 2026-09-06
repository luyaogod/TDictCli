---
title: "Getopt.initialize()"
source: "fgl-topics/c_fgl_utility_functions_getopt_initialize.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.initialize()"
type: "concept"
---

# Getopt.initialize()

> Initializes a variable defined with the Getopt type for command line argument processing.

## Syntax

```
FUNCTION (r Getopt) initialize(
   prog_name STRING,
   argv DYNAMIC ARRAY OF STRING,
   options GetoptOptions
   )
```

1. prog\_name is the name of the program (typically, [`arg_val(0)`](2726-arg-val.md "Returns a command line argument by position.")).
2. argv is a dynamic array of strings containing all command line arguments to
   be processed.
3. options is a [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.") array that holds the definition of the command line
   options.

## Usage

This is a method for the [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") type, used to initialize the variable of this type.

> **Tip:**
>
> Consider using the [`initDefault()`](2903-getopt-initdefault.md "Initializes a variable defined with the Getopt type.") method instead of `initialize()`, if the
> program name must be `arg_val(0)` and the start index to scan command line arguments
> is 1.

A variable of the type `Getopt` must be defined, as well as a [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.")
dynamic array containing the definitions of the command line options for the current program.

The first parameter defines the name of the program. This is typically
`arg_val(0)`, but it can be customized.

The second argument specifies the list of command line arguments to be processed. This is
typically [`copyArguments(1)`](2898-getopt-copyarguments.md "Returns a dynamic array of string with all command line arguments starting from the provided index."). For example, start at the second command line argument
with `copyArguments(2)`, to implement a command syntax with a verb as first
argument:

```
$ fglrun myprog capture --verbose --filename=file1
$ fglrun myprog duplicate --source=file1 --destination=file2
```

The third argument contains the definition of the command options that are available with this
program, defined in a [`GetoptOptions`](2895-getopt-getoptoptions-type.md "The GetoptOptions structured array type that holds the definition of command line options.") array.

After initializing the `Getopt` record, use the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method to parse
and validate command line options.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    DEFINE _options getopt.GetoptOptions = [ ... ]

    CALL g.inititalize("myprog", getopt.copyArguments(2), _options)
    WHILE g.getopt() == getopt.SUCCESS
        ...
    END WHILE

END MAIN
```

## Related links

**Related concepts**  

[Getopt.initDefault()](2903-getopt-initdefault.md "Initializes a variable defined with the Getopt type.")
