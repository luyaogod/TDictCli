---
title: "Locale matters when writing programs"
source: "fgl-topics/c_fgl_localization_020.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale matters when writing programs"
type: "concept"
---

# Locale matters when writing programs

> The language locale used when writing source code defines the runtime locale, except when developing in ASCII.

## Development and runtime character set must match

When writing a form or program source file, you use a specific character set. This character set
depends upon the text editor or operating system settings you are using on the development platform.
For example, when writing a string constant in a .4gl module, containing Arabic
characters, you probably use the ISO-8859-6 character set. The character set used at runtime (during
program execution) must match the character set used to write programs.

At runtime, a Genero program can only work in a specific character set. However, by using
localized strings, you can start multiple instances of the same compiled program using
different locales. For a given program instance the character set used by the strings
resource files must correspond to the locale. Make sure the string identifiers use ASCII
only.

## Byte length semantics and substring expressions

When using Byte Length Semantics (BLS), all character positions in strings are actually byte
positions. In a multibyte environment, if you don't pay attention to this, you can end
up with invalid characters in strings. For example, an expression using a subscript
operator `[x,y]` might refer to a byte position which is in fact in the
middle of a multibyte character. If possible, use Character Length Semantics (CLS) with
a multibyte locale to avoid such problems, or use only `STRING` methods
to parse character strings.

## Related links

**Related concepts**  

[Defining the application locale](0879-defining-the-application-locale.md "This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.")

[Length semantics settings](0881-length-semantics-settings.md "Length semantics settings")

[Language and character set settings](0880-language-and-character-set-settings.md "Language and character set settings")

[Localized strings](0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")
