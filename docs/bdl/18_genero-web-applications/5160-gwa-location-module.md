---
title: "gwa.location module"
source: "fgl-topics/r_gwa_location_api_functions.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.location module > gwa.location module"
type: "reference"
---

# gwa.location module

> Functions and types of the gwa.location module for interfacing GWA applications with HTML DOM API location object.

| Types | Description |
| --- | --- |
| TYPE TGWALocation RECORD href STRING protocol STRING host STRING hostname STRING port INTEGER pathname STRING search STRING hash STRING origin STRING approot STRING END RECORD | The TGWALocation type defines a record for a Genero wrapper around the JavaScript `window.location` object. |

| Function | Description |
| --- | --- |
| FUNCTION get() RETURNS gwa.location.TGWALocation | Returns the location description of the application window. |
| FUNCTION reload() | Reloads the application page. |

## Child topics

- [gwa.location.TGWALocation type](5161-gwa-location-tgwalocation-type.md): The TGWALocation type defines a record for a Genero wrapper around the JavaScript window.location object.
- [gwa.location.get](5162-gwa-location-get.md): Returns the location description of the application window.
- [gwa.location.reload](5163-gwa-location-reload.md): Reloads the application page.
