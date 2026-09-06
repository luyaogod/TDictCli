---
title: "What is my current character set?"
source: "fgl-topics/c_fgl_localization_014.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > What is my current character set?"
type: "concept"
description: "UNIX™ platforms UNIX-like platforms provide the LANG / LC_ALL environment variables to define the locale. Each process / terminal can set its own locale. On recent UNIX-like systems, the locale is ..."
---

# What is my current character set?

## UNIX™ platforms

UNIX-like platforms provide the LANG / LC\_ALL environment
variables to define the locale.

Each process / terminal can set its own locale. On recent UNIX-like systems, the locale is
en\_US.utf8 by default.

Check for available locales with the locale -a command. Some systems come with
only a few locales installed, you must then install an additional package to get more languages.

You must also define the correct character set in the terminal (xterm or gnome-term), otherwise
non-ASCII characters will not display properly.

## Microsoft™ Windows® platforms

On Windows platforms, for non-UNICODE (that means
non-UTF-16/UCS-2) applications, you have ACP and OEMCP code pages.

ACP stands for ANSI Code Page and were designed by Microsoft for GUI applications specifically, while OEMCP defines old code pages for MS/DOS
console applications.

If needed, select the default ACP/OEMCP code pages for non-UNICODE application in the language
and regional settings panel of Windows (make sure you
define the settings for non-UNICODE applications.

Code page can be changed in each console window with the chcp command.

With Genero Business Development Language, use the LANG environment variable on Windows to define the character set for BDL. However, it is strongly
recommended to use the default Windows system locale and
avoid setting LANG on Windows.

## Related links

**Related concepts**  

[Checking the locale configuration on UNIX platforms](0897-checking-the-locale-configuration-on-unix-platforms.md "Checking the locale configuration on UNIX platforms")
