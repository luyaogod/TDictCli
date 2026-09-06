---
title: "Reading FGLPROFILE entries"
source: "fgl-topics/c_fgl_fglprofile_009.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > Reading FGLPROFILE entries"
type: "concept"
---

# Reading FGLPROFILE entries

> Entries of FGLPROFILE files can be read by program.

Use the [`base.Application.getResourceEntry()`](../15_library-reference/2975-base-application-getresourceentry.md "Returns the value of a FGLPROFILE entry.") method to read an entry of FGLPROFILE
files.

Consider the following facts when reading FGLPROFILE entries:

1. FGLPROFILE entries are case insensitive. However, you should use the exact name to read the
   entry. Consider using lowercase identifiers like `myapp.myentry`.
2. Multiple files can be defined in the FGLPROFILE environment variable, and multiple entries with
   the same name can be defined. In such case, the last found wins.
3. The encoding of the FGLPROFILE file must match the encoding of the runtime system.
4. Use a `STRING` variable to hold the value of the FGLPROFILE entry to be
   read.

Example:

```
MAIN
    DISPLAY base.Application.getResourceEntry("myapp.myentry")
END MAIN
```

With the p1.prf file
defining:

```
myapp.myentry = "aaa"
```

And the p2.prf file
defining:

```
myapp.myentry = "bbb"
myapp.myentry = "ccc"
```

Setting FGLPROFILE environment variable to this
value:

```
export FGLPROFILE="p1.prf:p2.prf"
```

The last found entry is
read:

```
$ fglcomp main.4gl && fglrun main.42m
ccc
```
