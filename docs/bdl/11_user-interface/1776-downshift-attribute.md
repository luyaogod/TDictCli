---
title: "DOWNSHIFT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_DOWNSHIFT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > DOWNSHIFT attribute"
type: "concept"
---

# DOWNSHIFT attribute

> The DOWNSHIFT attribute forces character input to lowercase letters.

## Syntax

```
DOWNSHIFT
```

## Usage

When defining the `DOWNSHIFT` attribute for a character field, the characters
typed by the user are automatically converted to lowercase.

> **Important:**
>
> For maximum security, when the data is stored in a database, define a `CHECK`
> constraint on the SQL column corresponding to the field, in order to allow only downshifted
> characters, when an `INSERT` or `UPDATE` statement is executed.

In text mode, the results of conversions between uppercase and lowercase letters are based on the
[locale settings](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.").

Because using uppercase and lowercase letters results in different values in the database,
storing character strings in one or the other format can simplify sorting and querying.

Values send by the runtime system are displayed as is by the front-ends: No conversion will occur
if you display `"AbCdE"` to a field using `UPSHIFT` or
`DOWNSHIFT` attributes.

## Example

```
EDIT f001 = FORMONLY.name, DOWNSHIFT;
```

## Related links

**Related concepts**  

[UPSHIFT attribute](1838-upshift-attribute.md "The UPSHIFT attribute forces character input to uppercase letters.")
