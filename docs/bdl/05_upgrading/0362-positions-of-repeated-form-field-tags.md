---
title: "Positions of repeated form field tags"
source: "fgl-topics/c_fgl_MigI4GL_038.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Positions of repeated form field tags"
type: "concept"
---

# Positions of repeated form field tags

> The Genero Abstract User Interface definition (AUI tree) does not support misaligned repeating form field tags.

In a `SCREEN` section of IBM® Informix® 4GL (I4GL) form specification files, it is possible
to repeat [form item tags](../11_user-interface/1679-item-tags.md "Item tags define the position and size in a grid-based container.") at various
positions, with variable spaces between items.

In this example, the `f01` form item tags are aligned horizontally, but use
different spacing:

```
SCREEN
{
[f01]        [f01] [f01]  [f01]
}
END
```

In the next example, the `f01` form item tags are skewed vertically at different
column positions:

```
SCREEN
{
[f01]
  [f01]
    [f01]
      [f01]
}
END
```

With Genero form specification files, repeating form item tags must align vertically or
horizontally, with a uniform spacing.

Form item tags organized horizontally with 2 spaces between each item:

```
SCREEN
{
[f01]  [f01]  [f01]  [f01]
}
END
```

Form item tags organized vertically, with one line between each item:

```
SCREEN
{
  [f01]

  [f01]

  [f01]

  [f01]
}
END
```

If the form item tags are misaligned, the fglform compiler will report the
error -6832:

```
SCREEN
{ 
[f01]  [f01]  [f01]   [f01]
# Repeated screen tags 'f01' are misaligned, must align on X or Y.
# See error number -6832.
}
END
ATTRIBUTES
f01 = FORMONLY.field1;
END
```

You need to review the forms using misaligned item tags, and in some case you will have to rename
and define new form fields, when the data displayed in these fields is different from the fields
that are used for another type of data.

For example, a typical use case is to define a bunch of form item tags with the same field name,
to create a matrix with a single SCREEN RECORD, that can be used to display a
specific value with DISPLAY value TO sr[line].col1:

```
SCREEN
{
 Col1   Col2   Col3    Total
----------------------------
[f01]  [f02]  [f03]    [f04]
[f01]  [f02]  [f03]    [f04]
[f01]  [f02]  [f03]    [f04]
----------------------------
[f01]  [f02]  [f03]    [f04]
General total          [f04]
}
END
ATTRIBUTES
f01 = FORMONLY.col01;
f02 = FORMONLY.col02;
f03 = FORMONLY.col03;
f04 = FORMONLY.col04;
END
INSTRUCTIONS
SCREEN RECORD sr(FORMONLY.*);
END
```

In such case you will have to define individual form item tags and corresponding form
fields:

```
SCREEN
{ 
 Col1   Col2   Col3    Total
----------------------------
[f01]  [f02]  [f03]    [f04]
[f01]  [f02]  [f03]    [f04]
[f01]  [f02]  [f03]    [f04]
----------------------------
[f11]  [f12]  [f13]    [f14]
General total          [f99]
}
END
ATTRIBUTES
f01 = FORMONLY.col01;
f02 = FORMONLY.col02;
f03 = FORMONLY.col03;
f04 = FORMONLY.col04;
f11 = FORMONLY.col11;
f12 = FORMONLY.col12;
f13 = FORMONLY.col13;
f14 = FORMONLY.col14;
f99 = FORMONLY.col99;
END
INSTRUCTIONS
SCREEN RECORD sr1(FORMONLY.col01 THRU FORMONLY.col04);
END
```

## Related links

**Related concepts**  

[Form tags](../11_user-interface/1677-form-tags.md "Form tags define layout elements inside a grid-based container.")
