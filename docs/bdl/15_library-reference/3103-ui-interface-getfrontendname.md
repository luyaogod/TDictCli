---
title: "ui.Interface.getFrontEndName"
source: "fgl-topics/c_fgl_ClassInterface_getFrontEndName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.getFrontEndName"
type: "concept"
---

# ui.Interface.getFrontEndName

> Returns the type of the front-end currently in use.

## Syntax

```
ui.Interface.getFrontEndName()
  RETURNS STRING
```

## Usage

The `ui.Interface.getFrontEndName()` class method returns the type of the
front-end used by the program.

| Front-end name | Description |
| --- | --- |
| `GDC` | Genero Desktop Client front-end |
| `GMA` | Genero Mobile front-end for Android™ |
| `GMI` | Genero Mobile front-end for iOS |
| `GBC` | Genero Browser Client front-end |
| `GWA` | Genero Web Application front-end |
| `Console` | Text front-end (dumb terminal) |

> **Important:**
>
> When calling this method, the user interface module of the
> runtime system will initialize. As result, in text mode, the terminal will be initialized and get
> some escape sequences. This may corrupt standard output when executing batch programs. Consider
> testing the FGLGUI environment variable if you want to check that the batch program is executing in
> text mode, instead of using the `getFrontEndName()` method.

## Related links

**Related concepts**  

[Example 1: Get the type and version of the front-end](3125-example-1-get-the-type-and-version-of-the-front-end.md "Example 1: Get the type and version of the front-end")

[ui.Interface.getFrontEndVersion](3104-ui-interface-getfrontendversion.md "Returns the version of the front-end currently in use.")

[ui.Interface.getUniversalClientName](3110-ui-interface-getuniversalclientname.md "Returns the name of the front-end used for Universal Rendering.")
