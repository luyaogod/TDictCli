---
title: "Quickstart guide for locale settings"
source: "fgl-topics/c_fgl_localization_003.html"
breadcrumb: "Advanced features > Localization > Application locale > Quickstart guide for locale settings"
type: "concept"
---

# Quickstart guide for locale settings

> This is a quick step-by-step guide to properly configure locale settings for your Genero application.

Setting the locale involves different components, which must all be configured properly.

> **Tip:**
>
> This is a quickstart guide for locale settings. It is highly recommended that
> you read the complete set of articles regarding localization.

1. The application locale is defined by the character set used in your source files
   (.4gl, .per, .str). The same
   character set will be used in the compiled files (.42m,
   .42f, .42s). The current application locale can be checked
   with fglrun -i command, which displays information such as the character set name
   and length semantics.
2. Set the [operating system locale](0880-language-and-character-set-settings.md)
   corresponding to the application locale.
   - On UNIX™ based systems (including Mac® OS-X®), define the LANG, LC\_ALL or LC\_CTYPE environment variable. Use
     `locale -a` command to check if the locale exists on the machine. If not, it must be
     installed. If not set, LANG defaults to POSIX (ASCII).
   - On Windows® platforms, check if the regional
     settings for non-UNICODE applications match the application locale. If the regional settings do no
     match, you can define the LANG environment variable with a locale name supported by Microsoft® C Runtime Library, such as `French_France.1252`,
     or set `LANG=.utf8` for the UTF-8 character set.
   - On Android™ and iOS mobile devices, the
     application locale is always UTF-8 and the length semantics is CHAR. This cannot be changed.
3. Check that the operating system character set name is listed in the [charmap.alias file](0889-using-the-charmap-alias-file.md "The charmap.alias file can be used to map a system specific locale to a standard IANA locale."), to allow the
   runtime system to map it to a standard IANA character set name.
4. When using UTF-8 as character encoding, define the [length semantics](0881-length-semantics-settings.md) with the
   FGL\_LENGTH\_SEMANTICS={BYTE|CHAR} environment variable. On server platforms, Genero uses Byte Length
   Semantics by default for compatibility reasons. It's highly recommended to set
   FGL\_LENGTH\_SEMANTICS=CHAR to use Character Length Semantics. On mobile platforms, character length
   semantics is the default (this means that FGL\_LENGTH\_SEMANTICS does not need to be defined when
   running on a mobile device, it defaults to CHAR, and cannot be set to BYTE).
5. Set the [database client locale](0885-database-client-settings.md "This section describes the settings defining the locale for the database client.") with
   a character set corresponding to the application locale. For example, with Informix®, this is defined with the CLIENT\_LOCALE environment variable. The
   actual name of the database client locale may be different from the application locale name. But
   remember that the application and database client character sets must match.
6. Check the length semantics used by the database. For example, with Oracle, it is recommended that you set the
   database option `NLS_LENGTH_SEMANTICS='CHAR'`, if the application uses CLS (typically
   with UTF-8).
7. With UTF-8, use the proper SQL character data type to store UTF-8 data: This data
   type can be different depending on the type of database server. For more details, see [SQL character type for Unicode/UTF-8](../10_sql-support/1014-sql-character-type-for-unicode-utf-8.md "This section explains database server specifics regarding Unicode / UTF-8 support with character string SQL types.").
8. Set the [front-end locale](0886-front-end-locale-configuration.md "The host operating system on the front-end workstation must be able to handle the character set and fonts.") (and
   font). By front-end, we mean the program the end user interacts with. This can be a Genero front-end
   or a terminal emulator like Gnome-term, Putty, or a Windows Console. When using a Genero front-end, the front-end character set is defined by the
   type of the front-end, and the conversion from/to application character set is automatic, but you
   may need to select a font that is different from the system default. If you want to execute a TUI
   application in a terminal emulator, you must be sure that the terminal is configured to display the
   correct character set. This is for example defined with the *chcp* command on Windows, or in the "*Set Character Encoding*" menu option of a
   Gnome-term.
9. Define the [date, numeric and monetary
   formats](0890-date-numeric-and-monetary-formats.md "This section describes how Genero BDL handles date, time, numeric and monetary formats.") with the DBDATE, DBMONEY, DBFORMAT environment variables. On server platforms such as
   UNIX™ and Windows, these default to US
   formats (month/day/year for dates, the dot as decimal separator and $ as currency symbol). On mobile
   platforms, these default to the regional settings defined on the device.
