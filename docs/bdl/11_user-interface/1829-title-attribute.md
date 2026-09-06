---
title: "TITLE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_TITLE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > TITLE attribute"
type: "concept"
---

# TITLE attribute

> The TITLE attribute defines the title of a form item.

## Syntax

```
TITLE = [%]"string"
```

1. string defines the title to be associated with the form item, with
   the % prefix it is a localized string.

## Usage

The `TITLE` attribute defines the title of a form field.

For example, a form field `TITLE` can be used to define the header of a
`TABLE` or `TREE` column.

If you plan to internationalize your application, consider using localized strings with the
`%"string-id"` syntax.

The `TITLE` attribute supports accessibility when a screen reader is present. For
more details, go to [Screen readers](1591-screen-readers.md "How to integrate with platform screen readers?").

## Example

```
EDIT col4 = FORMONLY.ord_shipdate, TITLE="Ship date";
```

## Related links

**Related concepts**  

[Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.")

[TABLE container](1724-table-container.md "Defines a re-sizable table designed to display a list of records.")

[Form items](1665-form-items.md "The concept of form item includes all elements used in the definition of a form.")
