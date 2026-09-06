---
title: "Organizing string resources"
source: "fgl-topics/c_fgl_localized_strings_organizing.html"
breadcrumb: "Advanced features > Localization > Localized strings > Organizing string resources"
type: "concept"
---

# Organizing string resources

> Good practice in use of localized strings.

## What application strings to localize?

When modifying sources to add `%` prefixes to the strings that have to be
localized, you need to consider which string is the subject of internationalization, and leave other
strings without a `%` prefix.

For example, strings used to build an SQL statement at runtime obviously must not be
localized:

```
LET sql = "SELECT * FROM customers ", where_part
```

Typical strings to be localized are application
messages:

```
MESSAGE %"The customer name is mandatory!"
```

Furthermore, you may also need to localize application data. In this case, you can store
localized string identifiers in the database, and use the [`LSTR()`](../08_language-basics/0638-lstr-function.md "The LSTR() operator returns a localized string.") function at runtime to get the localized
string from your string resource files:

```
SELECT order_warning INTO rec.order_warning
    FROM orders WHERE ...
LET msg = LSTR(rec.order_warning)
DISPLAY BY NAME msg
```

## Messages with parameters

Applications often display messages with variable parts. The message text is usually built at
runtime with comma concatenation expression, where the message is split into different string
literals:

```
LET msg = "There are ", ord_count USING "<<<&", " orders not yet validated for ", rec.cust_name, "."
```

To simplify translation, consider reviewing the message construction by using the [`SFMT()`](../08_language-basics/0639-sfmt-function.md "The SFMT() operator replaces place holders in a string with values.") operator, to set the variable
parameters in your
messages:

```
LET msg = SFMT(%"orders.message.valid_count", ord_count, rec.cust_name)
```

You can then easily define the corresponding localized strings with `%n`
placeholders:

```
-- English string file:
orders.message.valid_count = "There are %1 orders not yet validated for %2."

-- French string file:
orders.message.valid_count = "%1 commandes ne sont pas encore validées pour %2."
```

Note that in `SFMT()` calls, the `%n` placeholders can be
specified at different positions, depending on the language needs.

## Development and runtime locale

The character set encoding (LANG/LC\_ALL locale) used in sources and at runtime must match. For
more details, see [Application locale](0864-application-locale.md "The application locale defines the language and codeset for your application.").

A good practice is to have sources in ASCII, and have string resources in the locale of your
choice: The runtime locale can be a specific ISO8859-? encoding for each language, or UTF-8, to have
a common encoding for all languages to be supported. However, you can also use UTF-8 in sources
and at runtime, if you want to use original texts as string identifiers in your sources.

The locale to be used at runtime will depend on the database locale used. You may need to support
a set of string files using ISO8859-? and a set of files using UTF-8, if you need to deploy your
application with ISO8859-? databases and UTF-8 databases.

## String identifiers

A localized string must be identified with a unique name. By default, if you add a
`%` prefix before existing strings and you extract the strings with fglcomp
-m or fglform -m, you will get string identifiers with the original
text:

```
"OK" = "OK"
"Cancel = "Cancel"
"Close" = "Close"
"There are %1 orders not yet validated for %2." = "There are %1 orders not yet validated for %2."
```

At this point, you can keep the original text for string identifiers, or re-define more abstract
identifiers (without quotes, such as `common.button.text.ok`).

Using the original text as string identifier has the advantage of been fast. It also simplifies
translation because the original text is directly visible for the translator. However, the character
set encoding should be UTF-8.

If you want to leave the original text for string identifiers, you must make sure that the locale
used at compile time matches the runtime locale. If the languages to be supported do not fit in a
single encoding like ISO8859-15, you will have to convert your sources to UTF-8 and use UTF-8 at
runtime.

Using abstract identifiers allows to maintain the sources in pure ASCII. Additionally, you can
give an indication of the usage context by using a clear identifier. it is recommended that you also
group common messages in a single string resource file. Using abstract identifiers will simplify
uniqueness checking.

common.str:

```
common.button.text.accept = "OK"
common.button.text.cancel = "Cancel"
common.button.text.close = "Close"
common.topmenu.text.accept = "Validate"
...
```

orders001.str:

```
orders.messages.valid_count = "There are %1 orders not yet validated for %2."
...
```

## Create directories for each language

At runtime, the .42s string resource files to be loaded must be declared
with the `fglrun.localization.*` FGLPROFILE entries.

In order to provide a set of string files for each language you want to support, organize the
string files in directories dedicated to a given
language:

```
/opt/app/resource/strings/en_US.iso8859-15  -- English strings in iso8859-15 code-set
/opt/app/resource/strings/fr_FR.iso8859-15  -- French strings in iso8859-15 code-set
/opt/app/resource/strings/jp_JP.utf8        -- Japanese strings in utf-8 code-set
```

## Related links

**Related concepts**  

[Loading localized strings at runtime](0909-loading-localized-strings-at-runtime.md "Understand the rules for using localized strings at runtime.")
