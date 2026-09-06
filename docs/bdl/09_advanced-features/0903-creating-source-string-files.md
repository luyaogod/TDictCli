---
title: "Creating source string files"
source: "fgl-topics/c_fgl_localized_strings_007.html"
breadcrumb: "Advanced features > Localization > Localized strings > Creating source string files"
type: "concept"
---

# Creating source string files

> A source string file contains localized string definitions for a given language (or localization context).

## What is a source string file?

A source string file is basically a mapping table that defines an identifier for each string.

After compiling source string files, programs can load and use a given string by referencing its
identifier (or key).

By convention, the source files of localized strings have the .str extension.

## Syntax

Define a list of string identifiers, and the corresponding text, by using the following
syntax:

```
"string-identifier" = "string-text"
```

For example:

```
"Cancel" = "Annuler"
```

Localized string keys are case sensitive. Consider using lower case characters only to avoid
mistakes.

As an alternative, you can define string identifiers as a dot-separated list of identifiers:

```
identifier. [...] = "string-text"
```

For example:

```
common.button.cancel = "Annuler"
```

If needed, you can add comment lines with the # or -- markers, like in other Genero source files:

```
# a comment
-- another comment
```

## Special characters

The [fglmkstr](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") compiler accepts the
backslash `\` as the escape character, to define non-printable characters such as
TAB, as well as the backslash and the double quote characters:

```
\l   \n   \r   \t   \\   \"
```

## Example

```
# A comment line
"Original text" = "Original text"
forms.customer.list = "Customer List"
special.characters.backslash = "\\"
special.characters.newline = "\n"
string.with.double_quote = "abc\"def"
```

## Related links

**Related concepts**  

[Extracting strings from sources](0906-extracting-strings-from-sources.md "Localized strings can be easily extracted from .4gl and .per source files.")
