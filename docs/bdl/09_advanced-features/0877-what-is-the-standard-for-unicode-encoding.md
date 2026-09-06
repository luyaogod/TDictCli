---
title: "What is the standard for UNICODE encoding?"
source: "fgl-topics/c_fgl_localization_013.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > What is the standard for UNICODE encoding?"
type: "concept"
description: "UNICODE is the standard for internationalization, but not all platforms/systems use the same UNICODE character set / encoding. Recent UNIX™ distributions define UTF-8 as the default character set ..."
---

# What is the standard for UNICODE encoding?

UNICODE is the standard for internationalization, but not all platforms/systems use
the same UNICODE character set / encoding.

Recent UNIX™ distributions define UTF-8 as the default
character set locale, XML files are UTF-8 by default, while Microsoft™ Windows® standard character encoding is
UTF-16 (NTFS) / UCS-2 (SQL Server). Recent Microsoft Windows 10 updates support UTF-8 as system locale for
non-Unicode applications (in beta stage while writing these lines).

Genero BDL supports UTF-8 and this is the encoding that must be used with implement UNICODE
applications.

> **Important:**
>
> Files encoded in UTF-8 can start with the UTF-8 Byte Order Mark (BOM), a sequence of `0xEF
> 0xBB 0xBF` bytes, also known as UNICODE `U+FEFF`. When reading files, Genero
> BDL will ignore the UTF-8 BOM, if it is present at the beginning of the file. This applies to
> instructions such as `LOAD`, as well as I/O APIs such as
> `base.Channel.read()` and `readLine()`.

## Related links

**Related concepts**  

[The UNICODE Standard](0875-the-unicode-standard.md "The UNICODE Standard")
