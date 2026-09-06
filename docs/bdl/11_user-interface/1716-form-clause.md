---
title: "FORM clause"
source: "fgl-topics/c_fgl_FormSpecFiles_FORM_clause.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > LAYOUT section > FORM clause"
type: "concept"
---

# FORM clause

> Reuse the definition of a form in the current form.

## Syntax

```
FORM "form-file"
```

1. form-file is the form to be included (without .per
   extension).

## Usage

The `FORM` clause includes an external form at the current layout position,
enforcing form re-usability, or to solve form complexity when using a `DIALOG`
instruction; for example to define a common form header for several application forms.

Wherever
a layout container can be specified, the layout of an external
form can be merged into the layout of the current form, with the
`FORM` clause. See [External form inclusion](1681-external-form-inclusion.md "Form inclusion allows to reuse the same form part in different forms.").

The .per source of the included form must be readable. If the compiled
version (.42f) does not exist, or is older than the .per
source, fglform will automatically compile the included form. The included forms
can be located in a different directory as the main form.

The form compiler searches for the external form relative to the path of the current
compiled form. For example, with `fglform dir1/dir2/main.per`, when the main form
includes an external form with `FORM "../otherdir/subform"`,
fglform will include the form file located in
`dir1/otherdir/subform.per`.

The form compiler performs an up-to-date test of the compiled form. Error [-6842](../15_library-reference/4483-genero-bdl-errors.md) is thrown if the up-to-date
test fails.

If the external form contains a `TOOLBAR` or a `TOPMENU`
section, error [-6841](../15_library-reference/4483-genero-bdl-errors.md) is
thrown.

The external form must not define a `SCREEN RECORD` or use a
`TABLE` already defined in the current form, otherwise error [-2024](../15_library-reference/4483-genero-bdl-errors.md) is thrown. Consider using the
[table alias syntax](1726-tables-section.md "Defines the list of database tables referenced by form field definitions.") to avoid duplicate
table names in merged forms.

The external form can define its own `ACTION DEFAULTS` section.
The action defaults of the external file will be merged into the
action defaults of the current form.

The `TABINDEX` attributes
of the elements of the result form will be adjusted. As the result
tabbing (`OPTIONS FIELD ORDER FORM` in programs) keeps
the visual order of the layout.

## Example

```
LAYOUT
  FOLDER
  PAGE page1 (TEXT = "Customer")
    FORM "customer"
  END
  PAGE page2 (TEXT = "Orders")
    FORM "orders"
  END
  END
END
```

## Related links

**Related concepts**  

[Screen records / arrays](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition.")

[TOOLBAR section](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.")

[TOPMENU section](1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions.")

[The SUBDIALOG clause](2087-the-subdialog-clause.md "The SUBDIALOG clause")
