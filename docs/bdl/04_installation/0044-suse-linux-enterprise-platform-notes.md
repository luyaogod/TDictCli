---
title: "SUSE Linux Enterprise platform notes"
source: "fgl-topics/c_fgl_installation_016.html"
breadcrumb: "Installation > Platform specific notes > SUSE® Linux Enterprise platform notes"
type: "concept"
description: "Installing NCurses on SUSE Linux Enterprise Important: Depending on the OS compatibility code of the Genero BDL package for Linux, fglrun requires the version 5 or version 6 of the wide-char NCurses ..."
---

# SUSE Linux Enterprise platform notes

## Installing NCurses on SUSE Linux Enterprise

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

**Installing NCurses V6 on SUSE Linux Enterprise**

On SUSE® Linux Enterprise 12 and 15+, the NCurses library version 6 must be
installed as follows:

```
# zypper install libncurses6
```

The `libncurses6` package includes the libncursesw.so.6
library for UTF-8 support in text mode.

**Installing NCurses V5 on SUSE Linux Enterprise**

On SUSE® Linux Enterprise 12 and 15+, the NCurses library version 5 must be
installed as follows:

```
# SUSEConnect -p sle-module-legacy/15/x86_64
# zypper install libncurses5
```

The `libncurses5` package includes the
libncursesw.so.5 library for UTF-8 support in text mode.

## Related links

**Related concepts**  

[TERMINFO terminal capabilities](../11_user-interface/1532-terminfo-terminal-capabilities.md "TERMINFO terminal capabilities")
