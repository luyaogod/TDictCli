---
title: "Identifying elements in ui.Form methods"
source: "fgl-topics/c_fgl_ClassForm_identify_elements.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > Usage > Identifying elements in ui.Form methods"
type: "concept"
description: "In ui.Form methods such as setElementHidden() and setFieldHidden() , the first parameter identifies an AUI node by its name attribute. With ui.Form.setElement*() methods, the name specified as first ..."
---

# Identifying elements in ui.Form methods

In `ui.Form` methods such as [`setElementHidden()`](3156-ui-form-setelementhidden.md "Show or hide form elements.") and [`setFieldHidden()`](3161-ui-form-setfieldhidden.md "Show or hide a form field."), the first parameter identifies an AUI node by its
`name` attribute.

With `ui.Form.setElement*()` methods, the name specified as first parameter
defines the AUI node(s) to be modified.

When several AUI nodes share the same name, the `ui.Form.setElement*()` methods
will change all matching nodes.

With `ui.Form.setField*()` methods, the name specified as first parameter is the
name of a form field node, the parent node of the view node to be modified.

Methods like `ui.Form.setFieldHidden()` change rendering attributes like
`hidden` and `style`, while methods such as [`ui.Dialog.setFieldActive()`](3226-ui-dialog-setfieldactive.md "Enable and disable form fields.") are
functional and related to dialog management.

AUI nodes are identified in XML/DOM trees by the `name` attribute and are case
sensitive. When defining form elements in the form specification file, element names will be
converted to lowercase letters in the .42f file. The name of the element passed
as parameter to the `ui.Form` methods can use the same letter case as the in the form
definition: The lookup is case-insensitive.

In the form file:

```
LAYOUT
...
ATTRIBUTES
EDIT f1 = Customer.CustAddr;
...
END
```

In the .42f
file:

```
<FormField name="customer.custaddr" ... >
   <Edit ... />
</FormField>
```

In the program
code:

```
INPUT BY NAME r_cust.*
   ...
   ON ACTION hide_field
      CALL DIALOG.getForm().setFieldHidden("Customer.CustAddr", 1)
   ...
```

## Related links

**Related concepts**  

[Identifying fields in ui.Dialog methods](3239-identifying-fields-in-ui-dialog-methods.md "Identifying fields in ui.Dialog methods")
