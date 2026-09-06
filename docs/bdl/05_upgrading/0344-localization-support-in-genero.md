---
title: "Localization support in Genero"
source: "fgl-topics/c_fgl_MigI4GL_007.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > Installation and setup topics > Localization support in Genero"
type: "concept"
---

# Localization support in Genero

> I4GL and Genero BDL use different libraries and environment variables in support of localization.

To support language-specific and country-specific locales, as well as multibyte character sets
like BIG5, IBM® Informix® 4GL uses the Informix GLS library.

For locale support, Genero Business Development Language (BDL) does not use the Informix GLS
library, to be independent from GLS libraries. Genero uses the standard C library functions for
character data handling, based on the POSIX `setlocale()` function.

I4GL uses the CLIENT\_LOCALE environment variable to define the locale for the application.
With Genero BDL, you must use the LANG/LC\_ALL environment variables to specify the locale of
the application. However, CLIENT\_LOCALE is still needed to define the locale for the
IBM Informix database client.

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")
