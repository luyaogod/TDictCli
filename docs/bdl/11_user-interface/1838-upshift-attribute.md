---
title: "UPSHIFT attribute"
source: "fgl-topics/c_fgl_FSFAttributes_UPSHIFT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > UPSHIFT attribute"
type: "concept"
---

# UPSHIFT attribute

> The UPSHIFT attribute forces character input to uppercase letters.

## Syntax

```
UPSHIFT
```

## Usage

When defining the `UPSHIFT` attribute for a character field, the characters typed
by the user are automatically converted to uppercase.

> **Important:**
>
> For maximum security, when the data is stored in a database, define a `CHECK`
> constraint on the SQL column corresponding to the field, in order to allow only upshifted
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
EDIT f001 = FORMONLY.name, UPSHIFT;
```

## Related links

**Related concepts**  

[DOWNSHIFT attribute](1776-downshift-attribute.md "The DOWNSHIFT attribute forces character input to lowercase letters.")
