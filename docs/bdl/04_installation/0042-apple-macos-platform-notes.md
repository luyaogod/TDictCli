---
title: "Apple macOS platform notes"
source: "fgl-topics/c_fgl_installation_macosx.html"
breadcrumb: "Installation > Platform specific notes > Apple® macOS® platform notes"
type: "concept"
description: "DYLD_LIBRARY_PATH denied in macOS® 10.11 Starting with macOS® 10.11 (El Capitan), if the System Integrity Protection (SIP) is enabled, the DYLD_LIBRARY_PATH environment variable is no longer exported ..."
---

# Apple macOS platform notes

## DYLD\_LIBRARY\_PATH denied in macOS® 10.11

Starting with macOS® 10.11 (El Capitan), if the [System Integrity Protection (SIP)](https://developer.apple.com/library/archive/documentation/Security/Conceptual/System_Integrity_Protection_Guide/RuntimeProtections/RuntimeProtections.html) is enabled, the
DYLD\_LIBRARY\_PATH environment variable is no longer exported in sub processes. This variable defined
the shared library search path for software components used by the Genero runtime system. This was
required especially for database client libraries installed in directories other than
/usr/lib and /usr/local/lib (the default location for
shared libraries).

As DYLD\_LIBRARY\_PATH cannot be used, a good practice is to create symbolic links to the required
shared libraries in $FGLDIR/lib.

For example, with a PostgreSQL client
(18):

```
$ export PGDIR="/opt/homebrew/Cellar/postgresql@18/18.4"
$ ln -s $PGDIR/lib/postgresql/libpq.5.dylib $FGLDIR/lib/libpq.5.dylib
```

> **Important:**
>
> If Genero BDL is installed in the /Applications directory, symbolic link
> creation needs "Full Disk Access", otherwise you will get an "Operation not permitted" error. For
> more details, check “Settings/Privacy & Security/Full Disk Access” on your mac.

To make sure that all required libraries can be found, check the dependencies on a Genero binary
(typically, the ODI driver) with the otool -L and otool -l
commands:

```
$ otool -L $FGLDIR/dbdrivers/dbmpgs_9.dylib
... check dependencies and paths

$ otool -l $FGLDIR/dbdrivers/dbmpgs_9.dylib
... check load sequences
```

## Java Interface

Check [maxOS platform notes](../14_extending-the-language/2660-platform-specific-notes-for-the-jvm.md) when using the [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") on this operating system.
