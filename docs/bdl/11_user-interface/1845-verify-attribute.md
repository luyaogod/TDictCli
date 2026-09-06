---
title: "VERIFY attribute"
source: "fgl-topics/c_fgl_FSFAttributes_VERIFY.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > VERIFY attribute"
type: "concept"
---

# VERIFY attribute

> The VERIFY attribute requires users to enter data in the field twice to reduce the probability of erroneous data entry.

## Syntax

```
VERIFY
```

## Usage

This attribute supplies an additional step in data entry to ensure the integrity of your data.
After the user enters a value into a `VERIFY` field and presses the Return or Tab
key, the runtime system erases the field and requests reentry of the value. The user must enter
exactly the same data each time, character for character: 15000 is not exactly the same as
15000.00.

The `VERIFY` attribute takes effect in `INPUT` or `INPUT
ARRAY` instructions only, it has no effect on `CONSTRUCT`
statements.

## Related links

**Related concepts**  

[REQUIRED attribute](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.")

[NOT NULL attribute](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.")
