---
title: "Structured ARRAYs in list dialogs"
source: "fgl-topics/c_fgl_Migrate_to_300_struct_arr.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Structured ARRAYs in list dialogs"
type: "concept"
---

# Structured ARRAYs in list dialogs

> ARRAYs with sub-records can be used in list dialogs, to simplify array definition based on database tables, requiring additional information at runtime.

Starting with Genero version 3.00, `ARRAY` variables defined with sub-records can
be bound to [DISPLAY ARRAY](../11_user-interface/1964-variable-binding-in-display-array.md) and [INPUT ARRAY](../11_user-interface/2010-variable-binding-in-input-array.md) screen records.

This is especially useful when you need to define arrays from database tables, and handle
additional row information at runtime, for example, to hold an image resource for each row, to be
displayed with the `IMAGECOLUMN` attribute.

An array is usually defined with a flat list of members:

```
SCHEMA shop
DEFINE a_items DYNAMIC ARRAY OF RECORD LIKE items.*
...
```

With version 3.00, arrays structured with sub-records can now be used within a
`DISPLAY ARRAY` or `INPUT ARRAY` dialog. The
array members and the form fields used by the screen array are bound by
position:

```
SCHEMA shop
DEFINE a_items DYNAMIC ARRAY OF RECORD
                   item_data RECORD LIKE items.*,
                   it_image STRING,
                   it_count INTEGER
               END RECORD
...
DISPLAY ARRAY a_items TO sr.*
   ...
```

For more details about program variable to form field binding in dialogs, see
[Binding variables to form fields](../11_user-interface/2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values."), [Example 4: DISPLAY ARRAY with structured array](../11_user-interface/2004-example-4-display-array-with-structured-array.md).
