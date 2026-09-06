---
title: "gwa.app.checkUpToDate"
source: "fgl-topics/c_gwa_app_api_checkUpToDate.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.checkUpToDate"
type: "concept"
---

# gwa.app.checkUpToDate

> Checks if the application is up to date by comparing versions and date time stamps.

## Syntax

```
FUNCTION checkUpToDate() 
  RETURNS gwa.app.TUpdateDialogResult
```

Returns a `gwa.app.TUpdateDialogResult` record with up-to-date information,
version numbers, and error.

## Usage

Use the `gwa.app.checkUpToDate()` method to check if the application is up to
date. This method is called by default when [gwa.app.updateDialog](5182-gwa-app-updatedialog.md "Checks the version of the loaded application against the server version and prompts the user to update if required.") is
used.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
