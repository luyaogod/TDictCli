---
title: "Commenting a module constant"
source: "fgl-topics/c_fgl_AutoDoc_module_constants.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting a module constant"
type: "concept"
description: "To document a module constant definition, add #+ lines just before the CONSTANT declaration. The comment body is composed of paragraphs separated by blank lines. The first paragraph of the comment is ..."
---

# Commenting a module constant

To document a module constant definition, add `#+` lines just before the
CONSTANT declaration.

The comment body is composed of paragraphs separated by blank lines. The first paragraph of the
comment is a short description of the constant. This description will be placed in the constant
summary table. The next paragraph is long text describing the constant in detail. Other paragraphs
must start with a tag to identify the type of paragraph; a tag starts with `@` (the
"at" sign).

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

| Tag | Description |
| --- | --- |
| `@code` | Indicates that the next lines show a code example using the constant. |

## Example

```
#+ This is the constant Pi
#+
#+ To be used in trigo computing
#+
#+ @code
#+ DISPLAY util.Math.cos( Pi / 2 )
#+
PUBLIC CONSTANT Pi = 3.14159
```
