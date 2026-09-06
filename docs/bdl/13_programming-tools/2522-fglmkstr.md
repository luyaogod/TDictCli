---
title: "fglmkstr"
source: "fgl-topics/c_fgl_tools_fglmkstr.html"
breadcrumb: "Programming tools > Command reference > fglmkstr"
type: "concept"
---

# fglmkstr

> The fglmkstr tool compiles .str localized string resource files.

## Syntax

```
fglmkstr [options] source[.str]
```

1. options are described in Table 1.
2. source.str is the string source file. The file extension
   is optional.

## Options

| Option | Description |
| --- | --- |
| `-V` | Displays version information. |
| `-h` | Displays options for the tool. |

## Usage

The fglmkstr command line tool is used to compile .str
localized string files into .42s files.

If case of error, the fglmkstr command execution status is different from
zero. Consequently, it is possible to detect compilation errors in scripts and makefiles.

## Related links

**Related concepts**  

[Compiling message files (.msg)](../11_user-interface/1596-compiling-message-files-msg.md "The .msg message files must be compiled to .iem binary files, in order to be loaded by the runtime system.")

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
