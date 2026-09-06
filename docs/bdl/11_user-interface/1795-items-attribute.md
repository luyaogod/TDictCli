---
title: "ITEMS attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ITEMS.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ITEMS attribute"
type: "concept"
---

# ITEMS attribute

> The ITEMS attribute defines a list of possible values that can be used by the form item.

## Syntax 1: Single-valued item list

```
ITEMS = ( value [,...] )
```

1. The single-value list is a comma-separated list of single values.
2. value is a numeric or string literal, or one of the following keywords:
   `NULL`, `TRUE`, `FALSE`.

## Syntax 2: Key-label item list

```
ITEMS = ( ( key, label ) [,...] )
```

1. The key-label list is a comma-separated list of (key,label) pairs within parentheses.
2. key is a numeric or string literal, or one of the following keywords:
   `NULL`, `TRUE`, `FALSE`.
3. label is a numeric literal, a string literal, or a localized string.

## Usage

The list must be delimited by parentheses, and each element of the list can be a simple literal
value or a pair of literal values delimited by parentheses.

This attribute is not used by
the runtime system to validate the field, you must use the
`INCLUDE` attribute to force the possible
values.

This example defines a list of simple values:

```
ITEMS = ("Paris", "London", "New York")
```

This example defines a list of pairs:

```
ITEMS = ((1,"Paris"),(2,"London"),(3,"New York"))
```

This attribute can be used, for example, to define the list of a `COMBOBOX` form
item:

```
COMBOBOX cb01 = FORMONLY.combobox01,
        ITEMS = ((1,"Paris"), (2,"London"),(3,"New York"));
```

In this example, the first value of a pair (1,2,3)
defines the data values of the form field and the second value of
a pair ("Paris", "London", "New York") defines the value to be displayed
in the selection list.

When used in a `RADIOGROUP` form item, this attribute defines the list of radio
buttons:

```
RADIOGROUP rg01 = FORMONLY.radiogroup01,
        ITEMS = ((1,"Paris"), (2,"London"),(3,"New York"));
```

In this case, the first value of a pair (1,2,3)
defines the data values of the form field and the second value of
a pair ("Paris", "London", "New York") defines the value to be displayed
as the radio button label.

You can specify item labels with localized strings, but this is only possible when you specify a
key and a label:

```
ITEMS = ((1,%"item1"),(2,%"item2"),(3,%"item3"))
```

It is allowed to define a `NULL` value for an item (An empty string is equivalent
to `NULL`):

```
ITEMS = ((NULL,"Enter bug status"),(1,"Open"),(2,"Resolved"))
```

In this case, the behavior of the field depends on the item type used.

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")

[COMBOBOX item type](1687-combobox-item-type.md "Defines a line-edit with a drop-down list of values.")

[RADIOGROUP item type](1699-radiogroup-item-type.md "Defines a mutual exclusive set of options field.")
