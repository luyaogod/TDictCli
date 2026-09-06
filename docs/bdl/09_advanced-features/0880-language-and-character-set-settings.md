---
title: "Language and character set settings"
source: "fgl-topics/c_fgl_localization_017.html"
breadcrumb: "Advanced features > Localization > Application locale > Defining the application locale > Language and character set settings"
type: "concept"
description: "Purpose of application locale definition The locale settings matters at compile time and at runtime. At runtime, the locale changes the behavior of the character handling functions, such as UPSHIFT ..."
---

# Language and character set settings

## Purpose of application locale definition

The locale settings matters at compile time and at runtime. At runtime, the locale
changes the behavior of the character handling functions, such as `UPSHIFT` and `DOWNSHIFT`. It also changes the handling of
the character strings, which can be single byte or multibyte encoded. Compilation errors will occur
if the source files contain characters that do not exist in the encoding defined by the current
locale.

Always check that the local environment variable matches the locale of your Genero
application, during development and at
runtime:

```
$ fglrun -i
Charmap      : UTF-8
Multibyte    : yes
Stateless    : yes
Length Semantics : CHAR
```

## Mobile platforms

On iOS and Android™ mobile platforms, the locale
is automatically defined to be UTF-8. This cannot be changed.

The language conventions and system messages are defined by the device settings.

## Microsoft™ Windows® platforms

On Windows platforms, when the LANG environment
variable is undefined, the language and character set defaults to the system locale which is defined
by the regional settings for "non-Unicode applications". For example, on a US-English
Windows, this defaults to the 1252 code page.

It is not recommended to set the LANG variable, unless your application uses a UTF-8 locale. For
Single Byte Character Set locales, use the Windows system
locale.

The values supported by Genero BDL for UTF-8 with the LANG environment variable are
`.fglutf8` (for backward compatibility) and
`.utf8`:

```
C:\> set LANG=.utf8
```

## UNIX™ platforms

On UNIX-based platforms, The LANG/LC\_ALL/LC\_CTYPE
environment variables define the local settings for the application.

The LC\_CTYPE environment variable can be set to overwrite the LANG/LC\_ALL setting. Make sure that
LC\_CTYPE does not define a different encoding than your application locale is requesting.

With the LANG/LC\_ALL/LC\_CTYPE environment variables, you define the
language, the territory (aka country) and the
codeset (aka character set or code page) to be used. The format of the value is
normalized as follows, but may be specific on some operating
systems:

```
language_territory.codeset
```

For example:

```
$ LC_ALL=en_US.iso88591; export LC_ALL
```

## What are possible locales on my platform?

Usually OS vendors define a specific set of values for the language,
territory and codeset. For example, on a UNIX platform, you typically have the value "en\_US.ISO8859-1" for a US English
locale. A list of available locales can be found by running the locale -a
command. You may also want to read the man pages of the *locale* command and the
setlocale function.

On Microsoft Windows, use the default system encoding. For more details about supported locales, search the
Microsoft MSDN documentation for "Language and
Country/Region Strings".

## UNICODE support (UTF-8)

To support multiple languages in your application, you must use UNICODE. The encoding
supported by Genero for UNICODE applications is UTF-8.

On UNIX platforms, UTF-8 locales are natively
supported with LANG/LC\_ALL.

On Windows platforms, defining the LANG
environment variable to code page 65001 will not work. According to Microsoft C++ [`setlocale()`](https://learn.microsoft.com/en-us/cpp/c-runtime-library/reference/setlocale-wsetlocale?view=msvc-170#utf-8-support) documentation, recent Windows versions support the UTF-8 encoding. This encoding can
be used with Genero, by setting the LANG environment variable to
`"language_territory.UTF8"`:

```
C:\> set LANG=en_us.UTF8
```

For system-wide locale settings, see also language options for "non-Unicode
applications", in the `Language & region` section of the Windows operating system settings.

If your Windows platform does not provide proper
support for UTF-8, Genero implements full UTF-8 support on Windows by setting the LANG environment variable to the value `.fglutf8` (for
backward compatibility) and `.utf8`:

```
C:\> set LANG=.utf8
```

## Related links

**Related concepts**  

[Checking the locale configuration on UNIX platforms](0897-checking-the-locale-configuration-on-unix-platforms.md "Checking the locale configuration on UNIX platforms")

[Locale settings (LANG) corrupted on Microsoft platforms](0895-locale-settings-lang-corrupted-on-microsoft-platforms.md "Locale settings (LANG) corrupted on Microsoft platforms")
