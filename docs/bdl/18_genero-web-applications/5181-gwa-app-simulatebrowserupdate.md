---
title: "gwa.app.simulateBrowserUpdate"
source: "fgl-topics/c_gwa_app_api_simulateBrowserUpdate.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.simulateBrowserUpdate"
type: "concept"
---

# gwa.app.simulateBrowserUpdate

> Simulate a browser update (service worker update).

## Syntax

```
FUNCTION simulateBrowserUpdate()
  RETURNS STRING
```

Returns `NULL` if the update was successful; otherwise, the string contains an
error message

## Usage

Use the `gwa.app.simulateBrowserUpdate()` method to simulate a browser update
(service worker update). You can reload the page using `gwa.app.set_href_to_new_index` if no error occurs.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
