---
title: "Best practices for localized strings"
source: "fgl-topics/c_fgl_localized_strings_best_practices.html"
breadcrumb: "Advanced features > Localization > Localized strings > Best practices for localized strings"
type: "concept"
---

# Best practices for localized strings

> This section describes good practices to localize your application messages.

## Program files and runtime language charset

Localization implies choosing a locale character set to use when executing a program.

The character set used at runtime must be the same as (or compatible with) the development
character set: The compiled program files (.42m, .42f) are
encoded in the character set used during compilation.

If the character set used in existing source files is different from the character set used at
runtime for a different target language, consider using ASCII only in your source code, and put
application messages in `.str` localized string files, using the appropriate
character set for each target language.

To support multiple languages at runtime, use the UTF-8 character set to encode your
`.str` (and `.42s`) localized string files. When using UTF-8 in your
source code and at runtime, the string keys can contain UTF-8 characters and you can use the
original text as string key.

For more details about character set usage, see [Application locale](0864-application-locale.md "The application locale defines the language and codeset for your application.").

## Defining the string key as identifier or as original text?

With localized strings, you have the choice of using the original text as string key, or of using
a more programmatic-type string identifier.

String key as original text (English):

```
"The transaction has been validated." = "La transaction a été validée."
```

String key as
identifier:

```
shipments.transaction.validated = "La transaction a été validée."
```

Pros and cons when using the original text as string key:

1. It is the fastest solution to localizing your application; you just need to add a percent sign
   (%) before texts in your source code, extract texts with fglcomp
   -m, and translate string files.
2. It simplifies the translation process, since the original text is always available as the string
   key.
3. If the charset used during development is different from the charset used at runtime, make sure
   to use ASCII only characters in the string key, otherwise your .str file will
   be dependent on the locale character set used in development.
4. Long texts can be used as string keys. For common messages that are often used in your sources,
   consider defining string constants in a dedicated module as described below.

Pros and cons when using identifiers as string key:

1. You can clearly distinguish messages depending on the context. For example, the text "Ok" may be
   the same in the English button label and transaction status, but may require different texts in
   other languages in a different context.
2. Distinct identifiers are easier to manage in the translation process, for example to store ids
   and texts in a database.
3. String identifiers can be used directly with the `%"ident"` notation in sources,
   and do not required you to define constants for common strings.
4. String identifiers require you to replace the original text in the source code by the string
   identifier preceded by the `%` sign. When using the original text as key, you just
   need to add the `%` sign in the source.

There is no constraint to exclusively using one of these patterns, you are free to use both
methods for the same application.

## Defining `CONSTANT` strings for common messages

Instead of repeating the same long string key in many places, group all common localized strings
in modules and define
constants:

```
-- mystrings.4gl module
PUBLIC CONSTANT STR_CONF_DEL_REC = %"Are you sure you want to delete this record?"
PUBLIC CONSTANT STR_TX_COMMITTED = %"The transaction has been committed."
...
```

When using the [`IMPORT FGL`](0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.")
instruction, constant definitions are available with [code
completion](../13_programming-tools/2544-configure-vim-for-genero-bdl.md).

## Using parameterized strings

If a message contains a variable part, or must display a value that is only known at runtime,
consider using the `SFMT()` operator to replace
`%n` placeholders in your
strings.

```
orders.item.validated = "The item %1 has been validated."
```

Since `%n` placeholders are replaced by position, it is easy to
put the placeholder at the position required by the language grammar.

English strings
file:

```
stock.items.count = "Stock %1 contains now %2 aditional items."
```

French strings
file:

```
stock.items.count = "%2 elements ajoutés dans le stock %1."
```

When using the original text as string key, it is good practice to identify parameter
placeholders with a different notation to the `%n` notation of
`SFMT()`. For example, you can use `P1`, `P2`,
etc:

```
"The item P1 has been validated." = "The item %1 has been validated."
```

Composed messages can be defined with `%n` and
`SFMT()`:

```
"P1 must be entered." = "%1 must be entered."
"Product code" = "Product code"
"Customer code" = "Customer code"
```

Then in the program
code:

```
LET str = SFMT(%"P1 must be entered.", %"Customer code")
```

To get the string `"Customer Code must be entered."`
