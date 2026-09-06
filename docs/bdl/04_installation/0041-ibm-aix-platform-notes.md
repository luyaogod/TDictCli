---
title: "IBM AIX platform notes"
source: "fgl-topics/c_fgl_installation_017.html"
breadcrumb: "Installation > Platform specific notes > IBM® AIX® platform notes"
type: "concept"
description: "XL C++ Runtime for AIX version Depending on the AIX version used, you must install the the corresponding \"IBM XL C++ Runtime for AIX\" . To check the installed C++ runtime version, execute the ..."
---

# IBM AIX platform notes

## XL C++ Runtime for AIX version

Depending on the AIX version used, you must install the the corresponding "IBM XL C++
Runtime for AIX".

To check the installed C++ runtime version, execute the following
commands:

```
$ lslpp -l | grep "IBM XL C++ Runtime for AIX"
  xlC.aix61.rte             13.1.3.0  COMMITTED  IBM XL C++ Runtime for AIX 6.1
  xlC.rte                   13.1.3.0  COMMITTED  IBM XL C++ Runtime for AIX
```

If the C++ runtime is not installed, Web Services library usage will produce the following error
message:

```
Could not load C extension library 'com'.
Reason: A file or directory in the path name does not exist.
```

For more details about the C++ runtime package installation on AIX, see <https://www.ibm.com/support/pages/fileset-information-xlcrte>.

## LIBPATH environment variable

The LIBPATH environment variable defines the search path for shared libraries. Make sure
LIBPATH contains all required library directories, including the system library path
/lib and /usr/lib.

## The dump command

On IBM® AIX®,
you can check the library dependencies with the dump command:

```
$ dump -Hv -X64 libstckp.so
```

## Unloading shared libraries from memory

In production environments, AIX loads shared libraries into
the system shared library segment in order to improve program load time. Once a shared library is
loaded, other programs using the same library are attached to that memory segment.

Once a shared library is loaded by the system, you cannot copy the executable file unless you
unload the library from the system memory. This problem will occur when installing a new version of
the software, even if it is installed in a different directory. Since shared libraries have the same
name, AIX will not allow multiple versions of the same library
to load. Therefore, before installing a new version, make sure all shared libraries are unloaded
from memory.

The genkld command prints the list of shared libraries currently loaded into
memory. The slibclean command unloads a shared library from the system shared
library segment.

## POSIX Threads and shared libraries

When using a thread-enabled shared library like Oracle's `libclntsh`, the program
using the shared object must be linked with thread support, otherwise you can experience problems
(like segmentation fault when the runner program ends). IBM
recommends using the xlc\_r compiler to link a program with pthread support.

By default, the runtime system provided for AIX platforms is linked with pthread support.

## Java Interface

Check [IBM AIX platform notes](../14_extending-the-language/2660-platform-specific-notes-for-the-jvm.md) when using the [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") on this operating system.
