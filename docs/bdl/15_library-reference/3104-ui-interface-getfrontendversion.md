---
title: "ui.Interface.getFrontEndVersion"
source: "fgl-topics/c_fgl_ClassInterface_getFrontEndVersion.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.getFrontEndVersion"
type: "concept"
---

# ui.Interface.getFrontEndVersion

> Returns the version of the front-end currently in use.

## Syntax

```
ui.Interface.getFrontEndVersion()
  RETURNS STRING
```

## Usage

The `ui.Interface.getFrontEndVersion()` class method returns the version number
of the front-end used by the program.

This method is primarily used for debugging purposes.

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

[ui.Interface.getFrontEndName](3103-ui-interface-getfrontendname.md "Returns the type of the front-end currently in use.")
