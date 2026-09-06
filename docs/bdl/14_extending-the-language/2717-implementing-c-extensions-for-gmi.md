---
title: "Implementing C-Extensions for GMI"
source: "fgl-topics/c_fgl_CExtensions_ios.html"
breadcrumb: "Extending the language > C-Extensions > Implementing C-Extensions for GMI"
type: "concept"
---

# Implementing C-Extensions for GMI

> This section describes how to program C-Extensions for the GMI VM.

## C-Extensions for GMI

With C-Extensions for GMI, you can address specific needs on iOS platforms, that are not
available by default in the Genero language. For example, implement functions to
interface with mobile specific hardware like sensors, card readers, scanners,
bluetooth, etc.

The runtime system virtual machine build in the GMI for iOS platforms can be extended
with the C-Extension technology. The basics to implement C-Extensions are the same
for iOS as for UNIX®/Windows® platforms, but
there are some differences, explained in this section.

The main difference is that user libraries cannot be loaded dynamically on iOS and
thus require a re-link of the GMI binary with the user-defined C-Extension
library.

## Writing C-Extension sources for GMI

C-Extension source files can be organized in several .c or
.m files, but the final library name must be
`userextension`.

For a first test, we recommend that you group all your C-Extension functions in a
single source file called userextension.m.

In the Objective C source file, add the following lines, to include typical iOS
header
files:

```
#include <Foundation/Foundation.h>
#include <UIKit/UIKit.h>
```

The Genero runtime system header file must be included as
well:

```
#include "f2c/fglExt.h"
```

The C-Extension functions must be registered as usual, in a [`UsrFunction`](2708-the-c-interface-file.md "To make your C functions visible to the runtime system, you must define all the functions in the C interface file.") array,
defining the number of input and output
parameters:

```
UsrFunction usrFunctions[]={
  {"get_user_info",get_user_info,1,1},
  ...
  {NULL,NULL,0,0}
};
```

## Using iOS C-Extensions in your program

The application code needs to be compiled on the development platform before it is
deployed on the iOS device or simulator, by using the C-Extension library build for
the development platform.

In your Genero program, import the C-Extension module with [`IMPORT` userextension](2710-loading-c-extensions-at-runtime.md "The runtime system can load several C-Extensions libraries, allowing you to properly split your libraries by defining each group of functions in separate C interface files."). You can also omit
this `IMPORT` instruction, because the runtime system tries to find and load the
`userextension` library by default. Note also that C-Extension functions have a
global scope, so you can omit the prefix of the function name with the lib/module
name:

```
IMPORT userextension
MAIN
   DEFINE info STRING
   LET info = get_user_info()
   ...
END MAIN
```

Compiler behavior regarding `IMPORT userextension` usage:

- With `IMPORT userextension`: The compiler can check references to functions
  defined in the extension. The programmer can qualify a function-name as
  `userextension.function-name`. But in
  this case, the userextension.so shared library must exist
  on the development platform.
- Without `IMPORT userextension`: The compiler can not check references
  to those functions. The compiler does not load the
  `userextension` module implicitly. C-Extension function names
  can not be qualified. In this case, the userextension.so
  library is not required for compilation, but it will be needed if the final
  program is linked, or if you want to execute/test the application in
  client/server development mode.

## Compiling and linking with C-Extensions on the development platform

On the development machine (Unix, Mac or Windows), if you [link 42r programs](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file."), or if you want the compiler to
check for missing symbols (with the `-r` option), the `userextension`
library must exist in the development environment. The extension library will be loaded at first
extension function call.

When the application is deployed on the iOS device, the extension library will be part of the
GMI/VM binary (because it is statically linked).

To create the `userextension` library for the development environment,
you must build an Objective-C shared library.

If the C-Extension contains iOS API calls, it will not be possible to compile the
extension library as is on the development machine: Write conditional pre-processor
macros to hide the iOS specific code, and simulate the function behavior for the
development platform:

```
#ifndef EMULATE_IOS
#include <Foundation/Foundation.h>
#include <UIKit/UIKit.h>
#endif
...
int get_user_info(int pc)
{
    char prop[101];
    char value[101];
    int z = (int) sizeof(prop);
    assert(pc==1);
    popvchar(prop, z);
#ifndef EMULATE_IOS
    ... here goes the iOS specific code ...
#else
    value[0] = '\0';
#endif
    pushvchar(value, (int) strlen(value));
    return 1;
}
```

Command line example to create a shared library with the Xcode® environment (note that we define the
NOT\_IOS\_IMPL constant to compile the code without iOS specific API
calls):

```
$ cc -shared -o userextension.dylib userextension.c \
-D EMULATE_IOS -I $FGLDIR/include -L $FGLDIR/lib -lfgl
```

## Building the iOS app with C extensions

Genero iOS apps are created with the gmibuildtool command-line
tool.

In order to build your iOS app with C extensions, you need to create the static
library with the `staticlib` target of $GMIDIR/lib/Makefile-gmi
and then pass the static library to the linker with the `--extension-libs`
option of gmibuildtool.

For more details, see [Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.")
