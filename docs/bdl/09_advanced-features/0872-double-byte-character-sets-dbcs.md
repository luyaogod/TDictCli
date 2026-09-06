---
title: "Double-byte character sets (DBCS)"
source: "fgl-topics/c_fgl_localization_009.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > Double-byte character sets (DBCS)"
type: "concept"
description: "A double-byte character set defines the encoding for characters on two bytes. The size of a character is always two bytes. Example of double-byte character sets include UCS-2, used by SQL Server in ..."
---

# Double-byte character sets (DBCS)

A double-byte character set defines the encoding for characters
on two bytes. The size of a character is always two bytes.

Example of double-byte character sets include UCS-2, used by SQL
Server in NCHAR and NVARCHAR columns. Note that UTF-16 is not a (fixed)
double-byte character set: You can have characters encoded on 2 or
4 bytes. UCS-2 is actually a subset of UTF-16.

Note that Genero Business Development Language does not support double-byte character
sets.

## Related links

**Related concepts**  

[Length semantics settings](0881-length-semantics-settings.md "Length semantics settings")
