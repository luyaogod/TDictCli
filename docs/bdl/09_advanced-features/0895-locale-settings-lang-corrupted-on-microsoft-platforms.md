---
title: "Locale settings (LANG) corrupted on Microsoft platforms"
source: "fgl-topics/c_fgl_localization_030.html"
breadcrumb: "Advanced features > Localization > Application locale > Troubleshooting locale issues > Locale settings (LANG) corrupted on Microsoft™ platforms"
type: "concept"
description: "On Microsoft™ Windows® XP / 2000 platforms, some system updates (Services Pack 2) or Office versions do set the LANG environment variable with a value for Microsoft applications (for example 1033). ..."
---

# Locale settings (LANG) corrupted on Microsoft platforms

On Microsoft™ Windows® XP / 2000 platforms, some system updates (Services Pack 2) or Office versions do set
the LANG environment variable with a value for Microsoft
applications (for example 1033).

Such value is not recognized by Genero as a valid locale specification.

Make sure that the LANG environment variable is properly set in the context of Genero
applications.

## Related links

**Related concepts**  

[What is my current character set?](0878-what-is-my-current-character-set.md "What is my current character set?")

[Language and character set settings](0880-language-and-character-set-settings.md "Language and character set settings")
