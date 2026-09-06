---
title: "Checking the locale configuration on UNIX platforms"
source: "fgl-topics/c_fgl_localization_032.html"
breadcrumb: "Advanced features > Localization > Application locale > Troubleshooting locale issues > Checking the locale configuration on UNIX™ platforms"
type: "concept"
description: "On UNIX™ systems, the locale command without parameters outputs information about the current locale environment. The locale is typically set with the LANG or, preferably with the LC_ALL environment ..."
---

# Checking the locale configuration on UNIX platforms

On UNIX™ systems, the locale command
without parameters outputs information about the current locale environment.

The locale is typically set with the LANG or, preferably with the LC\_ALL environment variable.
When the LC\_ALL variable is defined, LANG is
ignored:

```
$ export LANG=en_US.ISO8859-1

$ unset LC_ALL

$ locale 
LANG=en_US.ISO8859-1
LANGUAGE=
LC_CTYPE="en_US.ISO8859-1"
LC_NUMERIC=en_US.UTF-8
LC_TIME=en_US.UTF-8
LC_COLLATE="en_US.ISO8859-1"
LC_MONETARY=en_US.UTF-8
LC_MESSAGES="en_US.ISO8859-1"
LC_PAPER=en_US.UTF-8
LC_NAME="en_US.ISO8859-1"
LC_ADDRESS="en_US.ISO8859-1"
LC_TELEPHONE="en_US.ISO8859-1"
LC_MEASUREMENT=en_US.UTF-8
LC_IDENTIFICATION="en_US.ISO8859-1"
LC_ALL=

$ export LC_ALL=POSIX

LANG=en_US.ISO8859-1
LANGUAGE=
LC_CTYPE="POSIX"
LC_NUMERIC="POSIX"
LC_TIME="POSIX"
LC_COLLATE="POSIX"
LC_MONETARY="POSIX"
LC_MESSAGES="POSIX"
LC_PAPER="POSIX"
LC_NAME="POSIX"
LC_ADDRESS="POSIX"
LC_TELEPHONE="POSIX"
LC_MEASUREMENT="POSIX"
LC_IDENTIFICATION="POSIX"
LC_ALL=POSIX
```

Use the fglrun -i command to check current locale settings for
FGL:

```
$ fglrun -i
Charmap      : ASCII
Multibyte    : no
Stateless    : yes
Length Semantics : BYTE
```

Here the charset used is the ASCII (POSIX) charset. Clearing the LC\_ALL environment variable and
setting the LC\_CTYPE to a different value as LANG gives the following:

```
$ unset LC_ALL

$ export LC_CTYPE=zh_TW.big5

$ locale
LANG=fr_FR.ISO-8859-1
LANGUAGE=
LC_CTYPE=zh_TW.big5
LC_NUMERIC=en_US.UTF-8
LC_TIME=en_US.UTF-8
LC_COLLATE="fr_FR.ISO-8859-1"
LC_MONETARY=en_US.UTF-8
LC_MESSAGES="fr_FR.ISO-8859-1"
LC_PAPER=en_US.UTF-8
LC_NAME="fr_FR.ISO-8859-1"
LC_ADDRESS="fr_FR.ISO-8859-1"
LC_TELEPHONE="fr_FR.ISO-8859-1"
LC_MEASUREMENT=en_US.UTF-8
LC_IDENTIFICATION="fr_FR.ISO-8859-1"
LC_ALL=

$ fglrun -i
Charmap      : BIG5
Multibyte    : yes
Stateless    : yes
Length Semantics : BYTE
```

If the current locale settings is not supported by FGL, you get the following output:

```
$ fglrun -i
Error: locale not supported by C library, check LANG/LC_ALL.
```

> **Tip:**
>
> If the locale environment is not correct, check the value of the following environment
> variables: LANG, LC\_ALL, LC\_CTYPE, LC\_NUMERIC, LC\_TIME, LC\_COLLATE, etc.

The next example shows how to define a UTF-8 character set with character length semantics
(FGL\_LENGTH\_SEMANTICS):

```
$ export LC_ALL=en_US.UTF-8

$ export FGL_LENGTH_SEMANTICS=CHAR

$ fglrun -i
Charmap      : UTF-8
Multibyte    : yes
Stateless    : yes
Length Semantics : CHAR
```

## Related links

**Related concepts**  

[Operating system environment variables](../07_configuration/0496-operating-system-environment-variables.md "Describes some well-known system environment variables that are used by Genero software components.")
