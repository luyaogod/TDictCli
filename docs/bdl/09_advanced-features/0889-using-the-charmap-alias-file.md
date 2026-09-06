---
title: "Using the charmap.alias file"
source: "fgl-topics/c_fgl_localization_027.html"
breadcrumb: "Advanced features > Localization > Application locale > Using the charmap.alias file"
type: "concept"
---

# Using the charmap.alias file

> The charmap.alias file can be used to map a system specific locale to a standard IANA locale.

The name of the character set defined within the LANG/LC\_ALL environment variables can wary
from system to system. For example, on a given platform, the ISO-8859-1 character set may be
named "iso88591", while others platform will use "8859-1".

On UNIX™-like systems, installed locales can be listed with
the locale -a command.

Example of locale configuration on Linux/Debian:

```
$ unset LANG
$ export LC_ALL=en_US.iso885915
$ locale
LANG=
LANGUAGE=
LC_CTYPE="en_US.iso885915"
LC_NUMERIC="en_US.iso885915"
LC_TIME="en_US.iso885915"
LC_COLLATE="en_US.iso885915"
LC_MONETARY="en_US.iso885915"
LC_MESSAGES="en_US.iso885915"
LC_PAPER="en_US.iso885915"
LC_NAME="en_US.iso885915"
LC_ADDRESS="en_US.iso885915"
LC_TELEPHONE="en_US.iso885915"
LC_MEASUREMENT="en_US.iso885915"
LC_IDENTIFICATION="en_US.iso885915"
LC_ALL=en_US.iso885915

$ locale charmap
ISO-8859-15
```

To communicate with other components like front-ends, or to identify the encoding of XML files,
Genero programs must use a normalized name for character sets. This normalized name must follow the
IANA specifications [[RFC2978]](http://www.ietf.org/rfc/rfc2978.txt).

In order to convert the operating system specific locale codeset name to an IANA name, the
runtime system uses the charmap.alias mapping file, located in
$FGLDIR/etc. Add your operating system specific locale, if not listed in the
`charmap.alias` file.

The charmap.alias file has the following
syntax:

```
system-codeset-name  IANA-codeset-name
[...]
```

Comment lines can be added by using a `#` hash character at the beginning of the
line.

Example of charmap.alias
file:

```
# Linux
ISO_646.IRV:1983 ASCII
ANSI_X3.4-1968 ASCII
# macOS
US-ASCII ASCII
# IBM AIX
ISO8859-1 ISO-8859-1
ISO8859-2 ISO-8859-2
ISO8859-5 ISO-8859-5
ISO8859-6 ISO-8859-6
ISO8859-7 ISO-8859-7
ISO8859-8 ISO-8859-8
ISO8859-9 ISO-8859-9
ISO8859-15 ISO-8859-15
...
```
