---
title: "Application locale"
source: "fgl-topics/c_fgl_localization_038.html"
breadcrumb: "Advanced features > Localization > Application locale"
type: "concept"
---

# Application locale

> The application locale defines the language and codeset for your application.

The application locale defines:

- The language (for messages),
- The country or territory (for currency symbols and date formats),
- The code set (for character set encoding).

A program needs to be able to determine its locale and act accordingly, to support different
languages and character sets.

> **Important:**
>
> The same code point can represent distinct glyphs in different characters sets. Even if the
> glyphs/characters seem to display properly on the screen, an invalid locale configuration in one of
> the software components will result in invalid characters in your database system.
>
> Take for example a client application configured to display glyphs (font) for CP437. If the
> application gets a 0xA2 (decimal 162) code point, it displays an o-acute character. Now imagine that
> the DB client is configured with character set CP1252. In this character set, the code point 0xA2 is
> actually the cent currency sign. As a result, if the user enters the o-acute char (0xA2 in CP437) in
> the database, it will actually be interpreted as cent sign (0xA2 in CP1252) by the database server.
> When fetching that character back to the client, the database server returns the 0xA2 code point,
> which displays correctly as o-acute on the CP437 configured client, and the end user sees what was
> entered before. But with a different client application configured properly with CP1252 as DB client
> codeset, the end user will see the cent currency sign instead of the o-acute character.

## Child topics

- [Understanding locale settings](0865-understanding-locale-settings.md): This is an introduction to application locale definition.
- [Quickstart guide for locale settings](0866-quickstart-guide-for-locale-settings.md): This is a quick step-by-step guide to properly configure locale settings for your Genero application.
- [Locale and character set basics](0867-locale-and-character-set-basics.md): This section is an introduction to locale and character set basics.
- [Defining the application locale](0879-defining-the-application-locale.md): This section describes the settings defining the application locale, changing the behavior of the compilers and runtime system.
- [Database client settings](0885-database-client-settings.md): This section describes the settings defining the locale for the database client.
- [Front-end locale configuration](0886-front-end-locale-configuration.md): The host operating system on the front-end workstation must be able to handle the character set and fonts.
- [Locale matters when writing programs](0887-locale-matters-when-writing-programs.md): The language locale used when writing source code defines the runtime locale, except when developing in ASCII.
- [Runtime system messages](0888-runtime-system-messages.md): This section describes how to translate default English runtime system message files in a different language.
- [Using the charmap.alias file](0889-using-the-charmap-alias-file.md): The charmap.alias file can be used to map a system specific locale to a standard IANA locale.
- [Date, numeric and monetary formats](0890-date-numeric-and-monetary-formats.md): This section describes how Genero BDL handles date, time, numeric and monetary formats.
- [Using the Ming Guo date format](0891-using-the-ming-guo-date-format.md): Genero BDL can be configured to use the The Ming Guo calendar.
- [User's preferred language](0892-user-s-preferred-language.md): An application can get the user's preferred language and territory as configured on the front-end platform.
- [Right-to-left languages support](0893-right-to-left-languages-support.md): Genero supports right-to-left languages, such as Arabic and Hebrew.
- [Troubleshooting locale issues](0894-troubleshooting-locale-issues.md): This section describes common issues related to language locale definition.
