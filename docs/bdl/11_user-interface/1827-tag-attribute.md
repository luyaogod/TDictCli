---
title: "TAG attribute"
source: "fgl-topics/c_fgl_FSFAttributes_TAG.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > TAG attribute"
type: "concept"
---

# TAG attribute

> The TAG attribute can be used to identify the form item with a specific string.

## Syntax

```
TAG = "tag-string"
```

1. tag-string is user-defined string.

## Usage

This attribute is used to identify form items with a specific string. It can be queried in the program to perform
specific processing.

You are free to use this attribute as you need. For example, you can define a numeric identifier for each field
in the form in order to show context help, or group fields for specific input verification.

If you need to handle multiple data, you can format the text, for example, by using a pipe separator, or even
the JSON notation.

## Example

```
EDIT f001 = customer.fname, TAG = "name";
EDIT f002 = customer.lname, TAG = "name|optional";
```

## Related links

**Related concepts**  

[ui.ComboBox.setDefaultInitializerFunction](../15_library-reference/3250-ui-combobox-setdefaultinitializerfunction.md "Define the default initializer for combobox form items.")

[ui.Form.setDefaultInitializerFunction](../15_library-reference/3146-ui-form-setdefaultinitializerfunction.md "Define the default initializer for all forms.")
