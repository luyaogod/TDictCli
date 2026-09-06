---
title: "Getopt.GetoptOptions type"
source: "fgl-topics/c_fgl_utility_functions_getopt_t_getoptoptions.html"
breadcrumb: "Library reference > Utility modules > getopt: Command line options module > Getopt.GetoptOptions type"
type: "concept"
---

# Getopt.GetoptOptions type

> The GetoptOptions structured array type that holds the definition of command line options.

## Syntax

```
TYPE GetoptOptions DYNAMIC ARRAY OF RECORD
    name STRING,
    description STRING,
    opt_char CHAR,
    arg_type INTEGER
END RECORD
```

1. `name` defines the long name of the command line option.
2. `description` is the text to explain the command line option.
3. `opt_char` is the single-char command line option name.
4. `arg_type` can be one of:
   - `getopt.NONE`: The option has no additional value argument.
   - `getopt.OPTIONAL`: The option can be used with an optional value argument
     (`--option[=value]`).
   - `getopt.REQUIRED`: The option needs a mandatory value argument
     (`--option=value`).

## Usage

This type defines a dynamic array of a record structure to hold command line options definitions
information.

Define a variable of the `getopt.GetoptOptions` type and fill it with an [initializer](../08_language-basics/0691-variable-initializers.md "Variables can be initialized in their definition.").

Once the array is initialized, it can be passed to the [`initDefault()`](2903-getopt-initdefault.md "Initializes a variable defined with the Getopt type.") or
[`initialize()`](2904-getopt-initialize.md "Initializes a variable defined with the Getopt type for command line argument processing.")
method, to setup a [`Getopt`](2896-getopt-getopt-type.md "The Getopt structured type is used to process command line options.") variable in order to process command line arguments with the [`getopt()`](2902-getopt-getopt.md "Process the next command line option.") method.

## Example

```
IMPORT FGL getopt

MAIN
    DEFINE g getopt.Getopt
    DEFINE _options getopt.GetoptOptions
        = [(name: "version",
            description: "Version information",
            opt_char: 'v',
            arg_type: getopt.NONE),
           (name: "help",
            description: "This help page",
            opt_char: 'h',
            arg_type: getopt.NONE),
           (name: "outfile",
            description: "Output filename",
            opt_char: 'o',
            arg_type: getopt.REQUIRED)]

    CALL g.initDefault(_options)
    ...
END MAIN
```
