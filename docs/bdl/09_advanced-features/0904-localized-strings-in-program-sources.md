---
title: "Localized strings in program sources"
source: "fgl-topics/c_fgl_localized_strings_003.html"
breadcrumb: "Advanced features > Localization > Localized strings > Localized strings in program sources"
type: "concept"
---

# Localized strings in program sources

> How to specify a localized string in .4gl and .per sources?

## Defining localed strings in programs

A localized string is specified in the source code of program modules or form
specification files with the `%"string"` notation, to identify the string that must
be replaced at runtime by the corresponding text found in compiled string files.

In programs, localized strings can also be loaded dynamically with the `LSTR()`
operator.

## Syntax 1: Static localized string

```
%"sid"
```

1. sid is a character string literal that defines both the string identifier and
   the default text.

## Syntax 2: Dynamic localized string

```
LSTR(eid)
```

1. eid is a character string expression used at runtime as the string identifier
   to load the text.

## Static localized strings

A static localized string specification begins with a percent sign (`%`), followed
by the identifier of the string which will be used to find the text to be loaded. Since the identifier
is a string, you can use any type of characters, but it is recommended that you use a naming convention.
For example, you can specify a path by using several names separated by a
dot:

```
MAIN
  DISPLAY %"common.message.welcome"
END MAIN
```

The string after the percent sign defines both the localized string identifier and the default text
to be used for extraction, or the default text when no string resource files are provided at runtime.

You can use this notation in form specification files any place where a string literal can be
used.

```
LAYOUT
  VBOX
    GROUP g1 (TEXT=%"group01")
...
```

It is not possible to specify a static localized string directly in the area of containers like
`GRID`, `TABLE`, `TREE` or `SCROLLGRID`.
Use static [`LABEL`](../11_user-interface/1696-label-item-type.md "Defines a simple text area to display a read-only value.") form items to
define localized strings in layout labels:

```
LAYOUT
  GRID
  {
    [lab01  |f001              ]
  }
  END
END
ATTRIBUTES
LABEL lab01: TEXT=%"myform.label01";
EDIT f001 = FORMONLY.field01;
END
```

## Dynamic localized strings

The language provides a special operator to load a localized string dynamically, using an
expression as string identifier.

The name of this operator is [`LSTR()`](../08_language-basics/0638-lstr-function.md "The LSTR() operator returns a localized string.").

The following code example builds a localized string identifier with an integer and loads the
corresponding string with the `LSTR()`
operator:

```
MAIN
  DEFINE n INTEGER
  LET n = 234
  DISPLAY LSTR("str"||n)  -- loads string 'str234'
END MAIN
```

## Related links

**Related concepts**  

[Text literals](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.")
