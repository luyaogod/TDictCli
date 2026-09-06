---
title: "Specifying the URL source of a web component"
source: "fgl-topics/c_fgl_webcomponent_url_spec.html"
breadcrumb: "User interface > User interface programming > Web components > Using a URL-based web component > Specifying the URL source of a web component"
type: "concept"
---

# Specifying the URL source of a web component

> The content of URL-based web components is defined by the form field value. It can only be set by program.

## Setting the initial URL

When the current form defines a `WEBCOMPONENT` form item without the
`COMPONENTTYPE` attribute, it is a URL-based web component. The program can set the
URL dynamically in field value:

```
DISPLAY "wc-URL" TO wc-field
```

or with:

```
DEFINE wc_field STRING
LET wc_field = "wc-URL"
DISPLAY BY NAME wc_field
```

or by using the variable in an `INPUT` dialog with the
`UNBUFFERED` option:

```
DEFINE rec RECORD
   name STRING,
   mymap STRING
 END RECORD
...
LET rec.mymap = "http://www.openstreetmap.org"
INPUT BY NAME rec.* WITHOUT DEFAULTS
      ATTRIBUTES(UNBUFFERED)
   ...
```

Once the URL of the web component is defined, the initial URL content is shown by the
front-end, and the end user can interact with it.

## Changing the URL

During program execution, you can assigning another URL to the web component field value.
The content will be updated to show the new URL.

This example implements a `MENU` dialog with actions that set
different URLs to the web component field, changing the content based on the selected
action:

```
MENU "test"
   ON ACTION map_1
      DISPLAY "http://www.openstreetmap.org" TO wc_field
   ON ACTION map_2
      DISPLAY "http://www.wikimapia.org" TO wc_field
   ON ACTION map_3
      DISPLAY "http://maps.google.com" TO wc_field
END MENU
```

## Source of the URL

The URL resource set by the program code is typically an URL that can be accessed and downloaded
from a web server, on the internet or in a local network.

However, it is also possible to use a resource that is only available locally on the server where
the program executes. The URL-based web components can consume URIs computed in programs by [`ui.Interface.filenameToURI()`](../15_library-reference/3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.").
In this case, no web server is required: The URL-based `WEBCOMPONENT` field content
will be provided by the [intrinsic file
transfer](1586-providing-the-image-resource.md):

```
-- mypage.html is in current fglrun working directory:
DISPLAY ui.Interface.filenameToURI("mypage.html") TO wc_field
```
