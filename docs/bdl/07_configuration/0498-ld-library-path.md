---
title: "LD_LIBRARY_PATH"
source: "fgl-topics/c_fgl_EnvVariables_LD_LIBRARY_PATH.html"
breadcrumb: "Configuration > Environment variables > Operating system environment variables > LD_LIBRARY_PATH"
type: "concept"
description: "Defines a list of paths to find shared libraries on UNIX platforms."
---

# LD_LIBRARY_PATH

> Defines a list of paths to find shared libraries on UNIX™ platforms.

The LD\_LIBRARY\_PATH environment variable defines the list of search paths for shared libraries
loaded by the *dynamic linker* on UNIX platforms.

On some operating systems, the environment variable defining the shared library search path may
have a different name.

- On a system where a 32-bit and a 64-bit environment coexist, you may need to set
  LD\_LIBRARY\_PATH\_64 to execute the 64-bit programs.
- On macOS®, the usage of DYLD\_LIBRARY\_PATH
  is discouraged. Therefore, shared libraries that are not part of the Genero runtime system (such
  as database client libraries) must be found in the standard system directories
  (/usr/lib, /usr/local/lib)
- On AIX®, set `LIBPATH`.
