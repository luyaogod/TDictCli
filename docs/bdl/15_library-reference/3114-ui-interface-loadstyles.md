---
title: "ui.Interface.loadStyles"
source: "fgl-topics/c_fgl_ClassInterface_loadStyles.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.loadStyles"
type: "concept"
---

# ui.Interface.loadStyles

> Load the presentation styles file.

## Syntax

```
ui.Interface.loadStyles(
   path STRING )
```

1. path is the name of presentation styles file, without the extension.

## Usage

Use the `ui.Interface.loadStyles()` class method to load a
.4st file defining presentation styles for all program forms.

If the interface already contains a set of presentation styles, it will be replaced by the new
presentation styles loaded by this method.

Specify the presentation styles filename without the ".4st"
extension.

The resource file is searched in several directories in a given order, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

## Example

```
CALL ui.Interface.loadStyles("mystyles")
```

For a complete example, see [Example 3: Loading custom resources](3127-example-3-loading-custom-resources.md).

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[Presentation styles](../11_user-interface/1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")
