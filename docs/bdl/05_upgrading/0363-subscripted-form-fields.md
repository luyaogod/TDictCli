---
title: "Subscripted form fields"
source: "fgl-topics/c_fgl_MigI4GL_021.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Subscripted form fields"
type: "concept"
---

# Subscripted form fields

> Subscripted form fields must be located and redefined when moving to Genero BDL.

IBM® Informix® 4GL forms can define subscripted fields with multiple
field definition entries in the `ATTRIBUTES` section, each defining a
piece of the data displayed by the field, as in this
example:

```
DATABASE stores 
SCREEN
{
   1234567890
  [f01       ]
  [f02       ]
}
END
ATTRIBUTES
f01 = customer.cust_name[1,10];
f02 = customer.cust_name[11,20];
END
```

In the `ATTRIBUTES` section, the name of the field is immediately followed
by a subscript specification defining the piece of sub-data the screen tag displays and
allows to input.

This feature is not supported by Genero BDL, all fields must be defined as a whole.

## Related links

**Related concepts**  

[LAYOUT section](../11_user-interface/1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.")
