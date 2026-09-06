---
title: "ui.Interface.getUniversalClientVersion"
source: "fgl-topics/c_fgl_ClassInterface_getUniversalClientVersion.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.getUniversalClientVersion"
type: "concept"
---

# ui.Interface.getUniversalClientVersion

> Returns the version of the front-end used for Universal Rendering.

## Syntax

```
ui.Interface.getUniversalClientVersion()
  RETURNS STRING
```

## Usage

The `ui.Interface.getUniversalClientVersion()` class method returns the version
number of the GBC front-end used to render application forms.

This method is primarily used for debugging purposes.

When customizing your own GBC, the value returned by this API will contain the base version,
followed by the value set for the `customizationSuffix` parameter of the
.json GBC customization configuration file.

> **Important:**
>
> When calling this method, the user interface module of the runtime system will initialize. As a
> result, in text mode, the terminal will be initialized and get some escape sequences. This may
> corrupt standard output when executing batch programs. Consider testing the FGLGUI environment
> variable, if you want to check that the batch program is executing in text mode.
