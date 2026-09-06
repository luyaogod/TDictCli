---
title: "Query operators in CONSTRUCT"
source: "fgl-topics/c_fgl_Construct_004.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > Query operators in CONSTRUCT"
type: "concept"
description: "The CONSTRUCT instruction supports a specific query syntax, using wildcard characters and comparison operators. Table 1. CONSTRUCT query operators Symbol Meaning value Search for this value. The ..."
---

# Query operators in CONSTRUCT

The `CONSTRUCT` instruction supports a specific query syntax, using wildcard
characters and comparison operators.

| Symbol | Meaning |
| --- | --- |
| `value` | Search for this value. The string may contain `*` `?` `\|` `:` wildcards as described below. |
| `=value` | Equals value as is (wilcards are part of the value) |
| `<> value` or `!= value` | Not equal to `value` (wilcards are part of the value) |
| `=` (without value) | Is NULL |
| `<>` or `!=` (without value) | Is not NULL |
| `> value` | Greater than `value` |
| `>= value` | Greater than or equal to `value` |
| `< value` | Less than `value` |
| `<= value` | Less than or equal to `value` |
| `value1:value2` or `value1..value2` | Range from `value1` to `value2` |
| `element1\|element2 [\|...]` | List of values or character string patterns.For character type fiels, elements can be simple values and matches search patterns using \* or ? wildcards. |

When preceding the value with the `=` equal sign, wildcard characters such as
`*` `?` `|` `:` will be ignored and be
part of the value.

| Symbol | Meaning |
| --- | --- |
| `*` | A sequence of zero to n of any characters |
| `?` | Any single-character at this position |
| `[c1-c2]` | A character in the specified range, at this position |
| `[^c1-c2]` | A character NOT in the specified range, at this position |
| `[c1c2 [...] ]` | A character in the specified set, at this position |
| `[^c1c2 [...] ]` | A character NOT in the specified set, at this position |

Queries based on character types are case sensitive, because SQL
is case sensitive, except if the database server is configured
to be case-insensitive.

The `*` (star) and `?` (question mark) wildcards are specific to
character string type queries, and will generate MATCHES expression or LIKE expressions, depending
on the type of database used.

When entering a `*` or `?`, the pattern can also contain a
character range specification with the square brackets notation `[a-z]` or
`[xyz]`. A caret ( `^`. ) as the first character within the square
brackets specifies the logical complement of the set, and matches any character that is not listed.
For example, the search value `[^AB]*` specifies all strings beginning with
characters other than A or B. If the criteria does not contain `*` or
`?`, character range in square brackets are considered par of the value
(`[a-z]` produces `col='[a-z]'`).

Some syntaxes can produce an "`Error in field`" dialog error if the feature is
supported by the pattern matching operator of the database server. For example, not all db servers
support the `[a-z]` character range specification in the `LIKE`
pattern.

To search for rows with values containing a `*` star, a `?`
question mark or a `\` backslash, escape the wildcard character with a backslash.
Another option is to put an equal sign at the beginning of the QBE string, to find values that match
exactly that string.

The end user can combine the `|` operator with `*` or
`?` wildcards, to specify a list of search patterns. For example, when entering the
string `A|B*|?X`, the generated SQL condition for Informix will be `(col='A'
or col matches 'B*' or col matches '?X')` .

| QBE input example | Matching values | Non matching values |
| --- | --- | --- |
| `100` | `100` | `99, 101, NULL` |
| `>=100` | `100, 101, 200` | `10, 99, NULL` |
| `!=100` | `98, 98, 101, 102` | `100, NULL` |
| `!=` | `98, 99, 100, 101` | `NULL` |
| `1:100` | `1, 2 ... 99, 100` | `0, 101, NULL` |
| `aaa:yyy` | `aaa, aab, ab, yy, yyy` | `zaa, NULL` |
| `abc` | `abc` | `bc, abcd, Abc, NULL` |
| `ABC` | `ABC` | `abc, aBC, NULL` |
| `abc*` | `abc, abcd, abcdef` | `bc, ABC, NULL` |
| `=abc*` | `abc*` | `abc, abcdef` |
| `*bc` | `abc, bc` | `acd, aBC, NULL` |
| `?bc` | `abc, xbc, zbc` | `aabc, aBC, NULL` |
| `*bc?` | `aaaabcd, abcd, bcd` | `abcdef, bcdef, NULL` |
| `a\|b*\|?c` | `axx, bxx, xxc` | `xyz, NULL` |
| `[a-z]bc` | `abc, ebc, zbc` | `2bc, +bc, Abc, NULL` |
| `[^abc]*` | `deee, feee, zyx, z` | `azzz, byy, d, NULL` |
| `a[bxy]c` | `abc, axc, ayc` | `a2c, azc, aBc, NULL` |
| `*[xyz]` | `abcx, eeeez` | `abcd, eeee, NULL` |
| `1\|2\|35` | `1, 2, 35` | `4, 5, 6, NULL` |
| `aa\|bb\|cc` | `aa, bb, cc` | `ab, dd, NULL` |
| `=aa\|bb\|cc` | `aa\|bb\|cc` | `aa, bb, cc` |
| `\\abc*` | `\abc, \abcdef` | `abc, NULL` |
| `\*bc` | `*bc` | `abc, bc, NULL` |
| `*\[?\]*` | `[a], a[b]c, xx[y]zz` | `a[bb]c, a[]c, NULL` |
