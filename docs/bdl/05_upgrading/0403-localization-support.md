---
title: "Localization support"
source: "fgl-topics/c_fgl_Mig0000_005.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > Installation and setup topics > Localization support"
type: "concept"
---

# Localization support

> Four Js BDS and Genero BDL use different libraries and environment variables in support of localization.

IBM® Informix® 4GL (I4GL) and Four Js Business
Development Suite (BDS) use the Informix GLS
library for localization support (i.e. to support non-ASCII character
sets such as BIG5). This implies a strong dependency to the proprietary
GLS library.

Genero Business Development Language (BDL) does not use the GLS
library; Genero BDL uses the standard C library functions for character
set handling, based on the setlocale() POSIX conformant function.

While I4GL/BDS need the CLIENT\_LOCALE environment variable to define
the locale for the application, you must now use the LANG/LC\_ALL environment
variables to specify the locale of the Genero application. Note, however,
that CLIENT\_LOCALE is still needed when connecting to an IBM
Informix database.

In Four Js BDS, you could select the locale library with the fglmode tool, to
select either GLS or ASCII mode. This tool is no longer needed in Genero.

## Related links

**Related concepts**  

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")
