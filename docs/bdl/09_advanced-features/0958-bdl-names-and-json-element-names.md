---
title: "BDL names and JSON element names"
source: "fgl-topics/c_fgl_json_utils_names.html"
breadcrumb: "Advanced features > JSON support > BDL names and JSON element names"
type: "concept"
---

# BDL names and JSON element names

> To identify elements, JSON standards use different format as Genero BDL variable names.

As in many other programming languages, Genero BDL variables are named with simple identifiers.
These identifiers are case-insensitive, must start with a letter or underscore, and cannot contain
special characters such as spaces.

By default, JSON/BDL element name matching is case-insensitive. For example, if the Genero BDL
record member is defined as `CustNo`, and the JSON data string contains the
`"custno":999` name/value pair, the value will be assigned. However, since Genero BDL
record member names are used as-is to write JSON data, it is strongly recommended to define the
Genero BDL records with the exact names used in JSON data string. Since JSON is case-sensitive, make
sure the names of the Genero BDL record members match exactly the names expected in the resulting
JSON data string: `CustNo` will be different from `custNo`.

> **Important:**
>
> JSON specifications allow you to define element
> names with characters that cannot be used in Genero BDL identifiers. For example, a JSON element
> name can be `"customer.name"` or `"customer:name"`. To work around
> this issue, use the `json_name` attribute when defining the BDL variable.

If the JSON element name cannot be defined as a Genero BDL variable identifier, it is possible to
define the BDL variable with the `json_name` attribute, to specify the exact name of
the corresponding JSON element.

In the next example, the BDL variable `cust_name` will be mapped to the JSON
element `"Customer
Name"`:

```
DEFINE cust_name INTEGER ATTRIBUTES(json_name="Customer Name")
```

When converting JSON to BDL structures, elements in the JSON string that do not match an Genero
BDL record member are ignored; no error is thrown if there is no corresponding Genero BDL
member.

## Related links

**Related concepts**  

[Attributes on variable definitions](../08_language-basics/0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information.")
