---
title: "Optional SCREEN RECORD size for lists"
source: "fgl-topics/c_fgl_Migrate_to_310_screen_record_size.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Optional SCREEN RECORD size for lists"
type: "concept"
---

# Optional SCREEN RECORD size for lists

> A SCREEN RECORD definition can omit the number of rows of the corresponding list container.

Starting with Genero 3.10, a `SCREEN RECORD` definition can be defined without the size of the corresponding list container
(`TABLE`, `SCROLLGRID`, `TREE` or static field
list).

When specifying the size in `SCREEN RECORD`, it must match the exact number of
rows of the corresponding list container, otherwise fglform will throw error
[-2029](../15_library-reference/4483-genero-bdl-errors.md):

```
LAYOUT
GRID
{
<TABLE t1         >
[c1     |c2       ]
[c1     |c2       ]
<                 >
}
END
END
ATTRIBUTES
c1 = FORMONLY.cust_id;
c2 = FORMONLY.cust_name;
END
INSTRUCTIONS
SCREEN RECORD sr_cust[10](FORMONLY.*);
# Screen record array 'sr' has different component sizes.
# See error number -2029.
END
```

## Related links

**Related concepts**  

[Binding tables to arrays in dialogs](../11_user-interface/2318-binding-tables-to-arrays-in-dialogs.md "Program arrays act as data model that are bound to form tables, when implementing list dialogs.")

[Defining tables in the layout](../11_user-interface/2317-defining-tables-in-the-layout.md "Define table views in the LAYOUT section of the form definition file.")
