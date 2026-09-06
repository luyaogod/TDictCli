---
title: "Commenting a module type"
source: "fgl-topics/c_fgl_AutoDoc_module_types.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting a module type"
type: "concept"
description: "To document a module type definition, add #+ lines just before the TYPE declaration. The comment body is composed of paragraphs separated by blank lines. The first paragraph of the comment is a short ..."
---

# Commenting a module type

To document a module type definition, add `#+` lines just before the
TYPE declaration.

The comment body is composed of paragraphs separated by blank lines. The first paragraph of the
comment is a short description of the type. This description will be placed in the type summary
table. The next paragraph is long text describing the type in detail. Other paragraphs must start
with a tag to identify the type of paragraph; a tag starts with the @ "at" sign.

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

| Tag | Description |
| --- | --- |
| `@code` | Indicates that the next lines show a code example using the type. |

## Example

```
#+ This is the customer type
#+
#+ Define variables with this type to hold customer records.
#+
#+ @code
#+ DEFINE myvar t_cust
#+
PUBLIC TYPE t_cust RECORD
              cust_num INTEGER,
              cust_name VARCHAR(50),
              cust_addr VARCHAR(100)
            END RECORD
```
