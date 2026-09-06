---
title: "Application updates"
source: "fgl-topics/c_fgl_gwa_application_updates_reload.html"
breadcrumb: "Genero Web applications > Application updates"
type: "concept"
---

# Application updates

> The GWA API queries the status of the application and can update it automatically even when not active.

In a GWA application, when the application ends, an "Application Ended" page is displayed. By
default, a Reload page button is included on the application ended page,
allowing users to easily restart the application.

If the application has been updated server side in the meantime, clicking the Reload
page button will start the updated version of the application.

## Automatic application updates

As the GWA uses JavaScript service workers, the application can even be updated when the program
is not active; the browser automatically checks the server side for a new version and downloads the
new version if there is one.

By default the application is also updated automatically when it is reloaded. This can happen
when the user clicks the Reload page button on the application ended page, by
using the browser reload menu action (F5 usually), or by reopening the browser with the same
application URL.

To keep track of application updates, the [gwabuildtool](5150-gwabuildtool.md "The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser.") has an
`--app-version` switch to set the application version. This version can be queried by
using methods in the `gwa.app` package module.

## User-driven update

The [`gwa.app`](5165-gwa-app-module.md "Functions and types of the gwa.app module for querying version numbers and for working with application updates.") package module has
a function to query the up-to-date status of the application and perform a user-driven update
immediately if a version difference is detected.

If a network connection is not available, the current installed version in the browser is used.
Internally, the GWA tries first to contact the server for an update, if this fails, after 500ms the
cached application version is used.

It is recommended that you code to check for updates after your application starts by calling the
`gwa.app.getServerVersionAndBuild()` method.

For an example performing a user-driven update, explore the GWA "update" demo in your GWA
installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/gwa/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/gwa/update.

## Related links

**Related concepts**  

[GWA API](5157-genero-web-application-api.md "The Genero Web Application API is a package included with GWA installations, offering modules like gwa.location and gwa.app to assist with application development.")

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")

[Customize the index page](5134-customize-the-index-page.md "The GWA allows customization of index.html using custom CSS and HTML files.")
