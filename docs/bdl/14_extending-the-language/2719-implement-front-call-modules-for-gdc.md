---
title: "Implement front call modules for GDC"
source: "fgl-topics/c_fgl_frontcalls_user_ext_gdc.html"
breadcrumb: "Extending the language > User-defined front calls > Implement front call modules for GDC"
type: "concept"
---

# Implement front call modules for GDC

> Custom front call modules for the desktop front-end are implemented by using the API for GDC front calls in C language.

## GDC custom front call basics

In order to extend the GDC with your own front calls, you must be familiar with C++
programming, and have a C++ compiler installed on your development platform.

GDC front call modules must be implemented as a Dynamic Linked Library
(.DLL) on Windows® platforms, as a
shared library (.so) on Linux®, or as a
Dynamic Library (.dyLib) under Mac® Os X. This shared library must be deployed on each platform where the GDC front-end
executes.

The GDC is able to automatically load the front call module and find the function,
based on the module name and function name used in the Genero BDL front call
(`ui.Interface.frontCall`).

The API for GDC front calls is based on the `frontEndInterface`
front call interface structure, that is used to interface with the GDC core, in
order to pass/return values to/from a front call.

Follow these steps to implement a custom front call module for the GDC:

1. Create a C source to implement your front call functions.
2. In the front call functions body:
   1. Check the number of parameters passed with the `getParamCount()` function.
   2. Pop parameter values with one of the `pop*()` functions.
   3. Perform the function task.
   4. Push the result values with one of the `push*()` functions.
   5. Return 0 on success, -1 otherwise.
3. Compile and link the shared library.
4. Deploy the shared library to the platform where GDC executes.

## The front call interface structure

Information required to execute the front call is transmitted to the extension module
through the front call interface structure. This structure contains a list of function pointers
to:

- manage the stack (push or pop for each handled data type)
- get information about the function (number of in and out parameters)
- get information about the front-end (front call environment variables)

The following defines the front call interface structure:

```
struct frontEndInterface
  {
      short (* getParamCount) ();
      short (* getReturnCount) ();
      void (* popInteger) (long &, short &);
      void (* pushInteger) (const long, short);
      void (* popString) (char *, short &, short &);
      void (* pushString) (const char *, short, short);
      void (* getFrontEndEnv) (const char *, char *, short &);
      void (* popWString) (wchar_t *, short &, short &);
      void (* pushWString) (const wchar_t*, short, short);
  };
```

> **Important:**
>
> The front call interface structure is defined for the C++ language.

## Prototype of a front call function implementation

The prototype of each front call function must be:

```
int function_name ( const struct frontEndInterface &fci );
```

1. function\_name is the name of your function.
2. `fci` is the front call interface structure.

The `fci` structure will be filled in by the GDC and passed to the
custom function. You can then use this structure to pop/push values from/to the stack, and get
environment information from the core GDC.

The function must return 0 on success, -1 otherwise.

## Front call environment variables

The front call function can query the GDC for front call environment variables, to
get information about the context.

The following front call environment variables are supported:

| Environment Variable | Description |
| --- | --- |
| `frontEndPath` | The path where the GDC front-end is installed. |

## Module initialization and finalization

The font-call module can define initialization and finalization functions. GDC will
automatically call these functions as follows:

- `void initialize();`

  This function is called when the front call module
  library is loaded. If needed, perform variable initialization and resource allocation in
  this function.
- `void finalize();`

  This function is called when the GDC front-end stops.
  If needed, perform resource release in this function.

## The API for custom front call implementation

| Function | Description |
| --- | --- |
| short getParamCount(); | This function returns the number of parameters given to the function called. |
| short getReturnCount(); | This function returns the number of returning values of the function called. |
| void ( * getFrontEndEnv ) (const char * name, char * value, short & length ); | This function is used to get context information from the front-end.`name` is the name of the front call environment variable.`value` is the char buffer to hold the value of the variable.`length` is the actual length of the value. |
| void popInteger( long & value, short & isNull ); | This function is used to get an integer from the stack.`value` is the reference to where the popped integer will be set.`isNull` indicates whether the parameter is null. |
| void pushInteger( const long value, short isNull ); | This function is used to push an integer on the stack.`value` is the value of the integer.`isNull` indicates whether the value is null. |
| void popString( char * value, short & length, short & isNull ); | This function is used to get a string from the stack.`value` is the pointer where the popped string will be set.`length` is the length of the string.`isNull` indicates whether the parameter is null. |
| void pushString( const char * value, short length, short isNull ); | This function is used to push a string on the stack.`value` is the value of the string.`length` the length of the string. A length of -1 indicates that the length is detected based on the content of the string.`isNull` indicates whether the parameter is null. |
| void ( * popWString ) (wchar_t *value, short & length, short & isNull); | This function is used to get a WideChar string from the stack.`value` is the pointer where the popped string will be set.`length` is the length of the string.`isNull` indicates whether the parameter is null. |
| void ( * pushWString ) (wchar_t *value, short length, short isNull); | This function is used to push a WideChar string on the stack.`value` is the value of the string.`length` the length of the string. A length of -1 indicates that the length is detected based on the content of the string.`isNull` indicates whether the parameter is null. |

## Calling the custom front call from BDL

In the Genero program, use the [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.") API to call the front-end function. This method
takes the front call module name as the first parameter and the front call function name as second
parameter. The front call module name is defined by the name of the dynamic library
(module\_name.DLL,
module\_name.so or
module\_name.dylib).

For example, if you implement a front call module with the name
"mymodule.so", the Genero program code must use the name
"`mymodule`" as front call module
name:

```
CALL ui.Interface.frontCall("mymodule", "myfunction", ["John DOE"], [msg])
```

## Compiling and linking the module

To create the front call module shared library on
Linux:

```
g++ -c -o mymodule.o mymodule.cpp
g++ -shared -o mymodule.so mymodule.o
```

To create the front call module DLL on Windows (with Visual C++ environment
set):

```
cl /LD mymodule.cpp
```

## Deploying the custom front call module

The shared library implementing the custom front call functions must be deployed on
the platform where the GDC executes. Copy your custom front call modules in the bin directory of the
GDC installation directory (%GDCDIR%\bin).

## Example

This example implements a simple front call function that computes
the sum of two integer numbers. It takes two parameters and returns two
values.

mymodule.h:

```
struct frontEndInterface
  {
      short (* getParamCount) ();
      short (* getReturnCount) ();
      void (* popInteger) (long &, short &);
      void (* pushInteger) (const long, short);
      void (* popString) (char *, short &, short &);
      void (* pushString) (const char *, short, short);
      void (* getFrontEndEnv) (const char *, char *, short &);
      void (* popWString) (wchar_t *, short &, short &);
      void (* pushWString) (const wchar_t*, short, short);
  };

#ifdef _WIN32
#define EXPORT extern "C" __declspec(dllexport)
#else
#define EXPORT extern "C"
#endif

EXPORT void initialize();
EXPORT void finalize();
EXPORT int mysum(const frontEndInterface &fx);
```

mymodule.cpp:

```
#include "mymodule.h"
#include <stdio.h>
#include <string.h>

void initialize() {
}

void finalize() {
}

int mysum(const struct frontEndInterface &fci) {
  long param1, param2;
  short isNull1, isNull2;
  long sum;
  char msg[255];

  if (fci.getParamCount() != 2 || fci.getReturnCount() != 2) {
      return -1;
  }

  fci.popInteger(param2, isNull2);
  fci.popInteger(param1, isNull1);

  sum = 0;

  if (!isNull1 && !isNull2) {
    sum = param1 + param2;
    sprintf(msg, "%ld + %ld = %ld", param1, param2, sum);
  } else {
    sum = 0;
    sprintf(msg, "Parameters are NULL");
  }

  fci.pushInteger(sum, 0);
  fci.pushString(msg, strlen(msg), 0);

  return 0;
}
```

To invoke the sum front-end function, use the `ui.Interface.frontCall()`
method in your Genero program:

```
MAIN
  DEFINE res INT, msg STRING
  MENU
    ON ACTION frontcall ATTRIBUTES(TEXT="Call custom front call")
      CALL ui.Interface.frontCall("mymodule", "mysum",
                                  [100,250], [res,msg])
      DISPLAY "Result: ", res, "\n", msg
    ON ACTION quit
      EXIT MENU
  END MENU
END MAIN
```
