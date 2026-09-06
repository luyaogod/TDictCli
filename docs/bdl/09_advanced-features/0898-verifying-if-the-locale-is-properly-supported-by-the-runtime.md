---
title: "Verifying if the locale is properly supported by the runtime system"
source: "fgl-topics/c_fgl_localization_033.html"
breadcrumb: "Advanced features > Localization > Application locale > Troubleshooting locale issues > Verifying if the locale is properly supported by the runtime system"
type: "concept"
description: "Check the current LANG/LC_ALL locale settings by using the -i option of fglrun : $ fglrun -i Charmap : UTF-8 Multibyte : yes Stateless : yes Length Semantics : CHAR The lines printed with this option ..."
---

# Verifying if the locale is properly supported by the runtime system

Check the current LANG/LC\_ALL locale settings by using the `-i` option of
fglrun:

```
$ fglrun -i
Charmap          : UTF-8
Multibyte        : yes 
Stateless        : yes
Length Semantics : CHAR
```

The lines printed with this option indicate if the locale can be supported by the operating
system libraries.

If the locale settings are wrong or unsupported, the command will display an
error:

```
$ fglrun -i
Error: locale not supported by C library, check LANG/LC_ALL.
```

| Namea | Description |
| --- | --- |
| `Charmap` | This is the normalized IANA name of the [character set](0880-language-and-character-set-settings.md) used by the runtime system to communicate with external components (front-end, I/O of XML files). The mapping from the system locale name to a normalized name is defined in [$FGLDIR/etc/charmap.alias](0889-using-the-charmap-alias-file.md "The charmap.alias file can be used to map a system specific locale to a standard IANA locale."). |
| `Multibyte` | This line indicates if the character set is a MBCS ([Multi-Byte Character Set](0873-multibyte-character-sets-mbcs.md)) and an SBCS [Single Byte Character Set](0871-single-byte-character-sets-sbcs.md) (). |
| `Stateless` | A few character sets use an internal state that can change during the character flow. Only stateless character sets can be supported (the value must be 'yes'). |
| `Length Semantics` | `BYTE` indicates Byte Length Semantics, while `CHAR` indicates character length semantics. See [Length semantics settings](0881-length-semantics-settings.md) for more details. |
