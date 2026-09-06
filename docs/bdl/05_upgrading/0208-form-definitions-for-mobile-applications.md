---
title: "Form definitions for mobile applications"
source: "fgl-topics/c_fgl_Migrate_to_300_mobile_forms.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Form definitions for mobile applications"
type: "concept"
---

# Form definitions for mobile applications

> Genero version 3 supports grid and stack-based layout for mobile applications.

## Support for grid-based and stack-based layout

Before Genero version 3.00 (i.e., with Genero Mobile version 1.1), the GMI front-end
could only support a stack-based layout and it was required to create different
forms for iOS apps and other front-ends supporting grid-based layout. In fact, to
get a stack-based layout, grid-based .per forms were
automatically transformed on the fly when displayed on the GMI front-end.

Starting with Genero 3.00, all mobile front-ends support now grid-based layout and stack-based
layout, and a new `STACK` layout container was introduced to define stack-based layout
forms explicitly. Therefore, you can now use the same form definition for all mobile front-ends, by
implementing the layout type of your choice. It is even possible to mix grid-based or stack-based
forms in the same app.

> **Note:**
>
> The `STACK` layout feature is deprecated, see [Universal Rendering as standard](0134-universal-rendering-as-standard.md "All front-ends (GDC, GAS, GMA, GMI) use now the GBC as rendering engine.").

## Loading different forms based on front-end type

If you want to use a grid-based or stack-based form for the front-end, you can load the form with
`OPEN FORM` (or `OPEN WINDOW`) and implement the layout type based on
the front-end name returned by the [`ui.Interface.getFrontEndName()`](../15_library-reference/3103-ui-interface-getfrontendname.md "Returns the type of the front-end currently in use.") method:

```
MAIN
    ...
    OPEN FORM f1 FROM IIF( ui.Interface.getFrontEndName()=="GMI",
                           "myform_stack", "myform_grid" )
    DISPLAY FORM f1
    ...
END MAIN
```

## Related links

**Related concepts**  

[Form rendering](../11_user-interface/1538-form-rendering.md "The section explains the layout rules to render forms on graphical front-ends.")
