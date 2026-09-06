---
title: "ui.Interface.frontCall"
source: "fgl-topics/c_fgl_ClassInterface_frontCall_2.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.frontCall"
type: "concept"
---

# ui.Interface.frontCall

> ui.Interface.frontCall performs a function call to the current front-end.

## Syntax

> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

```
ui.Interface.frontCall(
   moduleName STRING, functionName STRING,
   [ parameters ], [ results ] )
```

where parameters is a comma-separated list of
expressions:

```
value [,...]
```

and where results is a comma-separated list of
variables:

```
variable [,...]
```

1. moduleName defines the shared library or classpath where the function is
   implemented.
2. functionName defines the name of the function to be called.
3. The input value(s) and the receiving variable(s) must be specified between `[ ]`
   square brackets. These are provided as [a variable
   parameter list](../08_language-basics/0643-variable-parameter-list.md "Variable parameter list delimiters.").
4. value is an expression of a [primitive
   type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), that will be used as input value for the front call. A comma-separated list of simple
   values can be provided. The [`record.*`](../08_language-basics/0720-accessing-record-members.md "Record members are accessed with the dot notation.") notation can be used, to expand all record
   fields as a list of simple input values. [Dynamic arrays](../08_language-basics/0734-dynamic-arrays.md)
   can also be used as input value list.
5. variable is a program variable of a [primitive type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), that will receive the return value of the front call. A comma-separated list
   of simple variables can be provided. The variables for output parameters are passed by reference to
   the `frontCall()` method. The [`record.*`](../08_language-basics/0720-accessing-record-members.md "Record members are accessed with the dot notation.") notation can be used, to get output values into
   record members. [Dynamic arrays](../08_language-basics/0734-dynamic-arrays.md) can also be used as output
   receiving list.
6. primitive-type is a [primitive data
   type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") such as `INTEGER`, `VARCHAR(50)`, etc.

## Usage

The `ui.Interface.frontCall()` class method can be used to execute a procedure on
the front-end workstation through the front-end software component. You can for example launch a
front-end specific application like a browser or a text editor, or manage the clipboard content.

> **Important:**
>
> When calling the `ui.Interface.frontCall()` method, the connection to the
> front-end is initiated, if it is not yet established. Consequently, front calls cannot be used in
> batch programs or when using the [text mode](../11_user-interface/1517-text-mode-rendering-tui-mode.md). Each front
> call will [sync the AUI tree](../11_user-interface/1509-the-dynamic-user-interface.md "The dynamic user interface is the base concept of the Genero user interaction components.") with the front end.

The method takes four parameters:

1. The module, identifying the shared library (.so or
   .DLL) or the Java class (GMA) implementing the front call function.
2. The function of the module to be executed.
3. The list of input parameters, using the square brackets notation.
4. The list of output parameters, using the square brackets notation.

Input and output parameters are provided as a [variable list of parameters](../08_language-basics/0643-variable-parameter-list.md "Variable parameter list delimiters."), by using the square brackets notation
(`[param1,param2,...]` ):

- Input and output parameters can be of any simple type like `INTEGER`, a
  `RECORD` or a `DYNAMIC ARRAY`.
- An empty list of input or output parameters is specified with `[]` .
- Input parameters can be an expression such as `(10 * var)`.
- Output parameters must be variables only, to receive the returning values.
- Output parameters are optional. If the front call returns values, these values will be ignored
  by the runtime system, if no output parameters are provided to receive these values.

Simple front call example:

```
FUNCTION get_front_end_name() RETURNS STRING
  DEFINE info STRING
  CALL ui.Interface.frontCall( "standard", "feInfo", ["feName"], [info] )
  RETURN info
END FUNCTION
```

Some front calls need a file path as parameter. File paths must follow the syntax of the front
end workstation file system. The following example shows how to pass a file path to a front-end
running on a Microsoft™ Windows®
workstation:

```
FUNCTION choose_directory() RETURNS STRING
   DEFINE path, dir STRING
   LET path = "C:\\work dir\\subdir"
   CALL ui.Interface.frontCall( "standard", "openDir",
           [path,"Choose a directory"], [dir] )
   RETURN dir
END FUNCTION
```

When using `RECORD` and `DYNAMIC ARRAY` as front call input or
output parameters, the runtime system will use JSON serialization, to pass and
return such structured data to/from the front-end. This is important to know when
implementing your own [custom front
calls](../14_extending-the-language/2718-user-defined-front-calls.md "Front-ends can be extended with custom functions to access specific features."). Note that one can use the [`json_null` and
`json_name` variable definition attributes](../08_language-basics/0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information.") to control
JSON
serialization:

```
DEFINE optrec RECORD
              mode INTEGER ATTRIBUTES(json_null="null"),
              filter STRING ATTRIBUTES(json_name="Data Filter")
          END RECORD
DEFINE flags DYNAMIC ARRAY OF INTEGER ATTRIBUTES(json_null="undefined")
DEFINE result_list DYNAMIC ARRAY OF STRING
LET optrec.mode = 999
LET optrec.filter = "*A*"
LET flags[1] = 111
LET flags[3] = 333
CALL ui.Interface.frontCall( "m1", "fc1", [optrec, flags], [result_list] )
```

## Front call cost

A front call is a remote procedure call requiring a full network round trip between the server
app and the front-end.

Depending on the current network speed, this may result in delays in the millisecond to sub
second range.

In mobile application [development](../17_mobile-applications/5083-mobile-development-mode.md "Set up a development environment to display app forms on a mobile front-end.") or [`runOnServer`](3458-mobile-runonserver.md "Run an application from the Genero Application Server using the specified URL.") mode, the
execution time of a front can be much slower when running the app on the server, compared to
embedded apps.

## Front call error handling

Exception handling instructions can be used to check the execution status of a front call. Both
[`WHENEVER ERROR`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") directives or [`TRY/CATCH`](../09_advanced-features/0853-try-catch-block.md "Use TRY / CATCH blocks to trap runtime exceptions in a delimited code block.") blocks can surround the front
call to avoid program stopping in case of error, and to check the error number returned in the [`status`](../09_advanced-features/0939-status.md "status is a predefined variable that contains the execution status of the last instruction.") variable.

There is no need to surround front calls with exception handlers such as
`TRY/CATCH`, if the front call is always supposed to execute without error. For
example, the [feInfo](3399-standard-feinfo.md "Queries general front-end properties.") front call will never
produce an exception.

Example of front call error handling with a `TRY/CATCH` block:

```
FUNCTION takePhoto()
   DEFINE path STRING
   TRY -- This front call may fail if the front-end is not a mobile device:
       CALL ui.Interface.frontCall( "mobile", "takePhoto", [], [path] )
   CATCH
       MESSAGE "Cannot take photo: ", status, " ", err_get(status)
       LET path = NULL
   END TRY
   RETURN path
END FUNCTION
```

If the front call module name or the function name is invalid, the errors [-6331](4483-genero-bdl-errors.md) or [-6332](4483-genero-bdl-errors.md) will be raised,
respectively.

If the front call execution fails for some reason, the error [-6333](4483-genero-bdl-errors.md) will be raised. The
description of the problem can be found in the second part of the error message, returned by a call
to the [ERR\_GET()](2732-err-get.md "Returns the text corresponding to an error number.") function.

The error [-6334](4483-genero-bdl-errors.md) can be
raised in case of input or output parameter mismatch. The control of the number of input and output
parameters is in the hands of the front-end. Most of the standard front calls have optional
returning parameters and will not raise error -6334, if the output parameter list is left empty.
However, front-end specific extensions or user-defined front-end functions may return an invalid
execution status in case of input or output parameter mismatch, raising error -6334. If the
front-end sends a call execution status of zero (OK), and the number of returned values does not
match the number of program variables, the runtime system will set unmatched program variables to
`NULL`. As a general rule, it is recommended that the program provides the expected
input and output parameters as specified in the documentation.

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
