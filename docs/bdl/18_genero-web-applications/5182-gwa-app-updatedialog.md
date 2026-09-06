---
title: "gwa.app.updateDialog"
source: "fgl-topics/c_gwa_app_api_updateDialog.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.updateDialog"
type: "concept"
---

# gwa.app.updateDialog

> Checks the version of the loaded application against the server version and prompts the user to update if required.

## Syntax

```
FUNCTION updateDialog(reportWhenUpToDate BOOLEAN)
   RETURNS gwa.app.TUpdateDialogResult
```

1. If reportWhenUpToDate returns `TRUE`, a message box displays
   that the application is up to date.

It returns a `gwa.app.TUpdateDialogResult` record about the update success.

## Usage

Use the `gwa.app.updateDialog()` method when you want to implement a user-driven
update. The method checks if the application is up to date by calling the [gwa.app.checkUpToDate](5172-gwa-app-checkuptodate.md "Checks if the application is up to date by comparing versions and date time stamps.") method internally. When the loaded version differs from
the server version, a dialog is displayed that prompts the user for action to update or cancel.
Otherwise, a message box displays that the application is up to date.

For a complete example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
