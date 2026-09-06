---
title: "RHEL / CentOS / Rocky Linux / Oracle Linux platform notes"
source: "fgl-topics/c_fgl_installation_019.html"
breadcrumb: "Installation > Platform specific notes > RHEL® / CentOS® / Rocky Linux® / Oracle Linux® platform notes"
type: "concept"
description: "Installing NCurses on RHEL, CentOS, Rocky Linux and Oracle Linux Important: Depending on the OS compatibility code of the Genero BDL package for Linux, fglrun requires the version 5 or version 6 of ..."
---

# RHEL / CentOS / Rocky Linux / Oracle Linux platform notes

## Installing NCurses on RHEL, CentOS, Rocky Linux and Oracle Linux

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

**Installing NCurses V6 on Red Hat Enterprise Linux, CentOS, Rocky Linux or Oracle Linux**

On Red Hat® Enterprise Linux® 8+, CentOS® 7+, Rocky Linux® 8+ or Oracle Linux® 8+, the
NCurses library version 6 must be installed as
follows:

```
# dnf install ncurses-libs
```

The `ncurses-libs` package includes the libncursesw.so.6
library for UTF-8 support in text mode.

**Installing NCurses V5 on Red Hat Enterprise Linux 7 / CentOS 7**

On Red Hat® Enterprise Linux® 7 and CentOS® 7, the NCurses library version 5 must be installed as
follows:

```
# yum install ncurses-libs
```

The `ncurses-libs` package includes the libncursesw.so.5
library for UTF-8 support in text mode.

## Related links

**Related concepts**  

[TERMINFO terminal capabilities](../11_user-interface/1532-terminfo-terminal-capabilities.md "TERMINFO terminal capabilities")
