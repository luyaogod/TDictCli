---
title: "Compiling string resource files (.str)"
source: "fgl-topics/c_fgl_localized_strings_009_2.html"
breadcrumb: "Programming tools > Compiling source files > Compiling string resource files (.str)"
type: "concept"
---

# Compiling string resource files (.str)

> The .str source string files must be compiled to .42s binary files, in order to be loaded by the runtime system.

To compile a source string file, use the [fglmkstr](2522-fglmkstr.md "The fglmkstr tool compiles .str localized string resource files.")
compiler.

```
$ fglmkstr filename.str
```

The fglmkstr tool generates a .42s file with the
filename prefix.

> **Important:**
>
> When compiling a .str source string file, you
> must set the locale (character set) corresponding to the encoding used in the
> .str file.

## Related links

**Related concepts**  

[Extracting strings from sources](../09_advanced-features/0906-extracting-strings-from-sources.md "Localized strings can be easily extracted from .4gl and .per source files.")
