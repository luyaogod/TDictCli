---
title: "Syntax of FGLPROFILE entries"
source: "fgl-topics/c_fgl_fglprofile_003.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > Syntax of FGLPROFILE entries"
type: "concept"
---

# Syntax of FGLPROFILE entries

> The syntax of FGLPROFILE entries is in the key = value form.

## Syntax

```
  # comment
| entry-definition
```

where entry-definition is:

```
entry = value
```

where entry is:

```
ident [.ident [.ident] [...] ] ]
```

and value is:

```
  [-]{ digit [...] }[ . digit [...] ]
| " alphanum [...] "
| {true|false}
```

1. comment is a line of text that is started by a `#` hash.
2. entry identifies the name of the entry. This can be a dot-separated list of
   identifiers.
3. value is a numeric value, a string literal, or a boolean value
   (true/false).

## Usage

An FGLPROFILE entry is a line in the configuration file associating a parameter name
to a value that can be specified as a numeric, string or boolean.

> **Important:**
>
> The encoding of FGLPROFILE files must match the application locale of the
> program. For more details, see [FGLPROFILE encoding](0488-fglprofile-encoding.md "The encoding of the FGLPROFILE files must match the runtime system locale.").

The entries are defined by a name composed of a list of identifiers separated by a dot
character.

FGLPROFILE entry names are case insensitive. In order to avoid any confusion, it is recommended
to write FGLPROFILE entry names in lower case.

If an entry is defined several times in the same file, the last entry found in the
file is used. No error is raised.

The value can be a numeric literal, a string literal,
or a boolean (true/false).

Numeric values are composed by an
optional sign, followed by digits, followed by an optional decimal
point and digits:

```
my.numeric.entry = -1566.57
```

String values must be delimited by single or double quotes. The escape character is backslash,
`\t`
`\n`
`\r`
`\f` are interpreted as TAB, NL, CR, FF. Double the backslash to
write a backslash character
(`\\`):

```
my.string.entry = "C:\\data\\test1.dbf"
```

Boolean values must be either the `true` or `false`
keyword:

```
my.boolean.entry = true
```

## Example

```
# Last modification: 2013-03-12/mike
report.aggregatezero = true 
gui.connection.timeout = 100
dbi.database.stores.source = "C:\\data\\test1.dbf"
dbi.database.stores.prefetch.rows = 200
```
