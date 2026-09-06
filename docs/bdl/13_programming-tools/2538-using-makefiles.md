---
title: "Using makefiles"
source: "fgl-topics/c_fgl_CompilingPrograms_006.html"
breadcrumb: "Programming tools > Compiling source files > Using makefiles"
type: "concept"
---

# Using makefiles

> Describes how to define program construction rules in makefiles.

Most UNIX™ platforms provide the make
utility program to compile projects. The make program is an interpreter of
makefiles. These files contain directives to compile and link programs
and/or generate other kind of files.

When developing on Microsoft™
Windows® platforms, you may use the NMAKE utility provided
with Visual C++. However, this tool does not have the same behavior as the UNIX make program. To have a compatible make on Windows, you can install a GNU make or third party UNIX tools such as Cygwin.

For more details about the make utility, see the platform-specific
documentation.

File makeincl with generic makefile rules to be included in other
Makefiles:

```
.SUFFIXES: .42s .42f .42m .42r .str .per .4gl .msg .hlp

FGLFORM=fglform -M
FGLCOMP=fglcomp -M
FGLMKMSG=fglmkmsg
FGLMKSTR=fglmkstr

all::
.msg.hlp:
        $(FGLMKMSG) $*.msg $*.hlp
.str.42s:
        $(FGLMKSTR) $*.str $*.42s
.per.42f:
        $(FGLFORM) $*.per
.4gl.42m:
        $(FGLCOMP) $*.4gl

clean::
        rm -f *.hlp *.42? *.out
```

Makefile to compile customer management application:

```
include makeincl

FORMS=\
 customers.42f\
 orderlist.42f\
 itemlist.42f

MODULES=\
 customer_mgmt.42m\
 zoom_orders.42m\
 zoom_items.42m

all:: $(FORMS) $(MODULES)
```

In this example, the `customer_mgmt` module imports `zoom_orders`
and `zoom_items` modules.

## Related links

**Related concepts**  

[Importing modules](2535-importing-modules.md "Describes how to define module interdependence with IMPORT FGL.")
