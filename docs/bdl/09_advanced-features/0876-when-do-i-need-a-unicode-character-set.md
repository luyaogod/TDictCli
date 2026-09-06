---
title: "When do I need a UNICODE character set?"
source: "fgl-topics/c_fgl_localization_012.html"
breadcrumb: "Advanced features > Localization > Application locale > Locale and character set basics > When do I need a UNICODE character set?"
type: "concept"
description: "With internationalization, people want to use different languages within the same application; for example, to have Chinese, Japanese, English, French and German addresses of customers in their ..."
---

# When do I need a UNICODE character set?

With internationalization, people want to use different languages within the same application;
for example, to have Chinese, Japanese, English, French and German addresses of customers in their
database.

UNICODE is a character encoding specification that defines characters for all
languages. More and more databases use an UNICODE character set on the database server, because it
"standardizes" all data from different client applications.

If needed, the client application can then use a different character set like ISO-8859-1 or BIG5:
The database software takes care of character set conversions. However, if the end user needs to
deal with different languages, all components of the system (from database backend to GUI front-end)
must work in UNICODE.

The UNICODE character set supported by Genero Business Development Language is UTF-8.
Double-byte based UNICODE character sets such as UCS-2 or UTF-16 are not supported. The
database server can however store character data in another UNICODE character set, as long as
the database client is able to handle to conversion to/from UTF-8 for the Genero runtime
system.

## Related links

**Related concepts**  

[The UNICODE Standard](0875-the-unicode-standard.md "The UNICODE Standard")
