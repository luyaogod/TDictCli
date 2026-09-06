---
title: "Commenting a module variable"
source: "fgl-topics/c_fgl_AutoDoc_module_variables.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting a module variable"
type: "concept"
description: "To document a module variable declaration, add #+ lines just before the DEFINE declaration. The comment body is composed of paragraphs separated by blank lines. The first paragraph of the comment is a ..."
---

# Commenting a module variable

To document a module variable declaration, add `#+` lines just before the
DEFINE declaration.

The comment body is composed of paragraphs separated by blank lines. The first paragraph of the
comment is a short description of the variable. This description will be placed in the variable
summary table. The next paragraph is long text describing the variable in detail. Other paragraphs
must start with a tag to identify the type of paragraph; a tag starts with the @ "at" sign.

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

| Tag | Description |
| --- | --- |
| `@code` | Indicates that the next lines show a code example using the variable. |

## Example

```
#+ Customer array
#+
#+ Fill this array with customer records and use it
#+ in DISPLAY ARRAY to control a TABLE
#+
PUBLIC DEFINE cust_list DYNAMIC ARRAY OF RECORD
              cust_num INTEGER,
              cust_name VARCHAR(50),
              cust_addr VARCHAR(100)
            END RECORD
```
