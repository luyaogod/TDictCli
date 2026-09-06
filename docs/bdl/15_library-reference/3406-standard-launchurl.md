---
title: "standard.launchURL"
source: "fgl-topics/c_fgl_frontcall_standard_launchurl.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.launchURL"
type: "concept"
---

# standard.launchURL

> Opens a URL with the default URL handler of the front-end.

## Syntax

```
ui.Interface.frontCall("standard", "launchURL",
  [ url [, mode ] ], [] )
```

1. url - The URL / URI to invoke.
2. mode (optional) - When using GAS/GBC or GWA, by default, a new browser tab is
   opened for the specified URL. When mode is set to `"replace"`,
   the content of the specified URL will be displayed in the current browser tab where the application
   forms are displayed. With GWA, the current app will be killed when using the
   `"replace"` mode. The mode parameter is ignored by GMA/GMI/GDC.

## Usage

The "`launchURL`" front call opens a URL
with the default URL handler available on the front-end platform. This is typically the web browser
for "`HTTP:`" URLs, or the mailer for "`mailto:`" URLs, but the
corresponding application may also be dedicated to the type of object specified by the URL (for
example, a mapping service or to initiate a phone call).

> **Important:**
>
> Some types of URLs are not supported by all front-end platforms. Make sure
> that you test all target front-ends when using a `launchURL` front call. For example,
> when the GBC front-end is running from HTTP/HTTPS (through the GAS), the web browser will block
> `file://host/path` URIs from opening, as this
> would create a security hole.

This front call is a powerful feature: front-end applications can register themselves as URL
handlers, so you can start applications on the front-end through the `launchURL`
front call.

Supported schemes depend on your system configuration.

> **Tip:**
>
> It is possible to produce a URI with the [`ui.Interface.filenameToURI()`](3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.")
> method, from the file located on the application server. This URI can then be used with the
> `launchURL` front call for example, to show PDF files.

The mode parameter is optional and is interpreted
differently depending on the front-end type:

- When the mode parameter is set to "`replace`":
  - With the GAS/GBC, the new URL will be launched in the current browser tab used by application
    windows, instead of launching the specified URL a new browser tab.
  - With the GAS/GWA, the new URL will be launched in the current browser tab, and the application
    displayed in the current browser tab will be killed. Consequently, there is little use cases where
    this mode makes sense with GAS/GWA apps.
  - With GMA, GMI and GDC front-ends, the `"replace"` mode is ignored.

## Example

To invoke Google Play Store:

```
CALL ui.Interface.frontCall("standard", "launchURL",
 ["market://details?id=com.google.android.apps.currents"], [])
```

```
CALL ui.Interface.frontCall("standard", "launchURL",
 ["market://details?id=com.google.zxing.client.android"], [])
```

To open Google Maps:

```
CALL ui.Interface.frontCall("standard", "launchURL",
 ["geo:48.613363,7.711083?z=17"], [])
```

To open Google Street View:

```
CALL ui.Interface.frontCall("standard", "launchURL",
 ["google.streetview:cbll=48.613363,7.711083&cbp=1,0,,0,1.0&mz=17"], [])
```

To initiate a phone call:

```
CALL ui.Interface.frontCall("standard", "launchURL", ["tel:+336717623"], [])
```
