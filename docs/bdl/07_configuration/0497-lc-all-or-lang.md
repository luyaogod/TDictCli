---
title: "LC_ALL (or LANG)"
source: "fgl-topics/c_fgl_EnvVariables_LC_ALL.html"
breadcrumb: "Configuration > Environment variables > Operating system environment variables > LC_ALL (or LANG)"
type: "concept"
description: "Defines the current application locale on UNIX platforms."
---

# LC_ALL (or LANG)

> Defines the current application locale on UNIX™ platforms.

The LC\_ALL (or LANG) environment variable defines language, territory and codeset for programs
running on UNIX platforms.

The codeset defined in LC\_ALL is used by the runtime system to handle character strings.

It is important to set this variable properly to the character set used by your
application.

If LC\_ALL is not defined, LANG is used instead.

Read the UNIX man page of the
`setlocale()` C runtime function for more details about this variable.

On Microsoft™ Windows™ platforms, the application locale is defined by the regional settings. However, when
the the LANG environment variable is defined, Genero compilers and runtime system use this variable.
Refer to [Language and character set settings](../09_advanced-features/0880-language-and-character-set-settings.md) for more details.

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")

[TZ](0504-tz.md "Defines the timezone for date/time values handling.")
