---
title: "A form is displayed with invalid characters"
source: "fgl-topics/c_fgl_localization_031.html"
breadcrumb: "Advanced features > Localization > Application locale > Troubleshooting locale issues > A form is displayed with invalid characters"
type: "concept"
description: "You may have different codesets on the development machine and the production server. The typical mistake that can happen is the following: You have edited and compiled a form-file on your Windows® ..."
---

# A form is displayed with invalid characters

You may have different codesets on the development machine and the production server.

The typical mistake that can happen is the following: You have edited and compiled a form-file on
your Windows® development workstation with the CP1253
encoding; When using the .42f form-file on a UNIX-server with encoding
ISO-8859-7, invalid characters will appear.

Keep in mind that all source files must be created/edited in the encoding of the server (where
fglcomp and fglrun will be executed). Consider writing your
sources in pure ASCII, and put language/codeset specific messages in localized strings.

## Related links

**Related concepts**  

[Understanding locale settings](0865-understanding-locale-settings.md "This is an introduction to application locale definition.")

[Language and character set settings](0880-language-and-character-set-settings.md "Language and character set settings")

[Compiling source files](../13_programming-tools/2530-compiling-source-files.md "Describes how to build the runtime files from source files.")
