---
title: "Debian / Ubuntu Linux platform notes"
source: "fgl-topics/c_fgl_installation_018.html"
breadcrumb: "Installation > Platform specific notes > Debian / Ubuntu Linux platform notes"
type: "concept"
description: "Installing NCurses on Debian or Ubuntu Important: Depending on the OS compatibility code of the Genero BDL package for Linux, fglrun requires the version 5 or version 6 of the wide-char NCurses shared ..."
---

# Debian / Ubuntu Linux platform notes

## Installing NCurses on Debian or Ubuntu

> **Important:**
>
> Depending on the OS compatibility code of the Genero BDL package for Linux,
> fglrun requires the version 5 or version 6 of the wide-char NCurses shared
> library to be installed on the system.
>
> - Package with OS code `l64xl217` needs NCurses version 5
> - Package with OS code `l64xl228` needs NCurses version 6
>
> After installing the Genero package, execute the ldd -r command on
> $FGLDIR/bin/fglrun, to identify what version of the NCurses shared library is
> required, and install the corresponding NCurses package if needed.

**Installing NCurses V6 on Debian or Unbuntu**

On Debian Linux 10+ or Ubuntu 22+, the NCurses library version 6 must be installed as
follows:

```
# apt-get install libncurses6
# apt-get install libncursesw6
```

The libncursesw.so.6 library for UTF-8 support in text mode is part of the
`libncursesw6` package: The `libncurses6` package does not include the
libncursesw.so.6 library.

**Installing NCurses V5 on Debian 9**

On Debian Linux 9+ the NCurses library version 5 must be installed as
follows:

```
# apt-get install libncurses5
# apt-get install libncursesw5
```

The libncursesw.so.5 library for UTF-8 support in text mode is part of the
`libncursesw5` package: The `libncurses5` package does not include the
libncursesw.so.5 library.

**Installing NCurses V5 on Debian 10 and Ubuntu 22**

The version 5 of the NCurses library is not available on Debian 10 or Ubuntu 22.

## Related links

**Related concepts**  

[TERMINFO terminal capabilities](../11_user-interface/1532-terminfo-terminal-capabilities.md "TERMINFO terminal capabilities")
