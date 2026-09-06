---
title: "Commenting a method"
source: "fgl-topics/c_fgl_AutoDoc_type_methods.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting a method"
type: "concept"
description: "To document a method for type, add some lines starting with #+ , before the FUNCTION declaration. The comment body is composed of paragraphs separated by blank lines. The first paragraph of the ..."
---

# Commenting a method

To document a method for type, add some lines starting with `#+`, before the
`FUNCTION` declaration.

The comment body is composed of paragraphs separated by blank lines. The first paragraph of the
comment is a short description of the method. This description will be placed in the method
summary table of the type. The next paragraph is long text describing the function in detail.
Other paragraphs must start with a tag to identify the type of paragraph; a tag starts with the
@ "at" sign.

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

| Tag | Description |
| --- | --- |
| `@code` | Indicates that the next lines show a code example using the method. |
| `@param name description` | Defines a method parameter identified by *name*, explained by a *description*.*name* must match the parameter name in the method declaration. |
| `@return description` | Describes the values returned by the method.Several @return comment lines can be written. |

```
#+ Rectangle class
PUBLIC TYPE Rectangle RECORD
   height FLOAT,
   width FLOAT
END RECORD

#+ Computes the area of the rectangle
#+
#+ @param ratio Applies a ratio to the area
#+
#+ @return The computed area
#+
PUBLIC FUNCTION (r Rectangle) area(ratio FLOAT) RETURNS FLOAT
   RETURN ratio * r.height * r.width
END FUNCTION
```
