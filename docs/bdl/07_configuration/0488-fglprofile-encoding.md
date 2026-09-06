---
title: "FGLPROFILE encoding"
source: "fgl-topics/c_fgl_fglprofile_006.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > FGLPROFILE encoding"
type: "concept"
---

# FGLPROFILE encoding

> The encoding of the FGLPROFILE files must match the runtime system locale.

The [application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.") defines the character
encoding used by the runtime system.

Since no character set conversion is done when reading entries from FGLPROFILE files, it is
mandatory that the FGLPROFILE files use the same encoding as the runtime system.

If your application is designed to be executed with different encodings, consider defining
ASCII-only values in FGLPROFILE files.

If character strings of FGLPROFILE entries need to be in a specific language selected at runtime,
use string identifiers that point to [localized
strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site."). The language-specific messages can then be loaded dynamically with the [`LSTR(string-identifier)`](../08_language-basics/0638-lstr-function.md "The LSTR() operator returns a localized string.")
operator.
