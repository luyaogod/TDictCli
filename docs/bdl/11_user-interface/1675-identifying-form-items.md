---
title: "Identifying form items"
source: "fgl-topics/c_fgl_FormSpecFiles_ident_items.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Identifying form items"
type: "concept"
---

# Identifying form items

> Elements defined in a form file can be identified with a name, to be used in programs.

Form fields are implicitly identified by the `tabname.colname` specification after
the equal sign, while other (non-field) form items such as static labels and group boxes can get an
optional item name.

The form item name defined in the form file will be copied to the `name` attribute
of the corresponding node in the .42f file. It can then be used by programs to
select a form element at runtime, to introspect or modify its attributes.

For example, specify the name for a `GROUP` container by writing an identifier
after the layout container type:

```
GROUP group1 (TEXT="Customer")
```

Here the group name is '`group1`', and it can be used in a program to identify the
group
element:

```
DEFINE w ui.Window 
DEFINE g om.DomNode 
LET w = ui.Window.getCurrent()
LET g = w.findNode("Group","group1")
CALL g.setAttribute("text","Another text")
```

Helper methods are provided for common tasks on form elements. For example, to hide a group with
the identifier `group1`, you can use the `setElementHidden()` method
on a [`ui.Form`
object](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction."):

```
DEFINE f ui.Form
...
   LET f = DIALOG.getForm()
   ...
   CALL f.setElementHidden("group1", TRUE)
```

Consider defining unique names to identify form elements, and to simplify the search at runtime.
A good practice is the use of a prefix based on the type of form element (`g_` for
groups, `l_` for labels for example).

Static items (labels) in a `GRID` container cannot get a name, because these are
self-defined with the layout part of the item:

```
GRID
{
Name: [f1               ]
...
}
END
```

In the above example, the label `"Name:"` cannot be identified. In order to give a
name to such a label, use an item tag and add a `LABEL` item definition in the
`ATTRIBUTES` section, to specify the name of the label after the
colon:

```
GRID
{
[l1   ][f1               ]
...
}
END
...
ATTRIBUTES
LABEL l1: l_name, TEXT="Name:";
...
```

## Related links

**Related concepts**  

[ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
