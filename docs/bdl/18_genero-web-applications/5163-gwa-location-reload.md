---
title: "gwa.location.reload"
source: "fgl-topics/c_gwa_api_location_reload.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.location module > gwa.location module > gwa.location.reload"
type: "concept"
---

# gwa.location.reload

> Reloads the application page.

## Syntax

```
FUNCTION reload()
```

Reloads the application page either immediately or with a 50ms delay.

## Usage

The `gwa.location.reload()` method is a wrapper for the JavaScript API
`window.location.reload` object. For more information about the JavaScript API, refer
to [Location: reload() method](https://developer.mozilla.org/en-US/docs/Web/API/Location/reload) (external link). This method reloads the page you
are running on, in other words it restarts your GWA application. Its use is only recommended for the
purpose of updating a GWA application immediately.

For an example, run the GWA "update" demo in your GWA installation directory. If installed in FGLDIR, you will find the demo in
$FGLDIR/demo/update or if installed in a separate directory, you will find
it in gwa-install-dir/demo/update.

## Related links

**Related concepts**  

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")

[gwa.location.get](5162-gwa-location-get.md "Returns the location description of the application window.")
