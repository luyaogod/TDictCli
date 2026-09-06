---
title: "Extracting strings from sources"
source: "fgl-topics/c_fgl_localized_strings_008.html"
breadcrumb: "Advanced features > Localization > Localized strings > Extracting strings from sources"
type: "concept"
---

# Extracting strings from sources

> Localized strings can be easily extracted from .4gl and .per source files.

Use the [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") and [fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.") compilers with the
-m option to extract localized
strings.

```
$ fglcomp -m mymodule.4gl
```

The compilers dumps all localized string to stdout. This output can be redirected to a
file to generate the default source string file with all the localized strings used in the
source file. Source string files can then be re-organized, to centralize common messages
in a unique .str file, and can then be compiled by fglmkstr into
.42s files to be used by the runtime system.

## Related links

**Related concepts**  

[Creating source string files](0903-creating-source-string-files.md "A source string file contains localized string definitions for a given language (or localization context).")
