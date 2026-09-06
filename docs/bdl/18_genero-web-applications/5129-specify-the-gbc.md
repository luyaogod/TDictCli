---
title: "Specify the GBC"
source: "fgl-topics/c_fgl_gwa_specify_gbc.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Specify the GBC"
type: "concept"
---

# Specify the GBC

> A GBC must be bundled with the GWA.

The [gwabuildtool](5150-gwabuildtool.md "The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser.") selects the GBC at build time in the
following order:

1. The first option, is to use the GBC specified by the `gwabuildtool --gbc
   gbc-directory` build option.
2. If the `--gbc` option is not specified, gwabuildtool will look
   for the [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode.") environment variable.
3. Finally, it defaults to the FGLDIR/web\_utilities/gbc/gbc directory if none
   of the previous options is set.

GBC is loaded in an iframe inside the GWA application, which means that the
index.html of GWA is different from the index.html of
GBC.

Any customization of the GBC is expected to work. GWA requires that
GBCDIR/index.html exists and GBCDIR/VERSION exists.

## Using GBC query strings

You can specify GBC query string options –
"`debugmode`","`contextmenu`","`mobileui`",
"`theme`" and so on – as you would with standard server applications.

For example, to load the GWA application in debug mode, start your application with [gwarun](5152-helper-tools.md "The gwarun tool runs a GWA program in the browser on your desktop."). The URL may be something like
this:

```
http://localhost:9103/d/r/index.html?fglapp=main.42m&uuid=gwarun2024_12_11_18_00_31_769&viaMiniWS=1
```

Add
`&debugmode=1` at the end of the URL and reload the application. On reload, the
GBC debug icon should be visible in the chromebar at the top.

For more information about GBC query string options, refer to Query string parameters
page in the Genero Browser Client User Guide.

## Related links

**Related concepts**  

[Package web components](5130-package-web-components.md "Web components must be packaged with the GWA application because they are preloaded and must also be available offline.")

[Specify the start-up module](5131-specify-the-start-up-module.md "Specify the module you want to use as the main entry point.")

[Set environment with fglprofile](5128-set-environment-with-fglprofile.md "Configure environment settings with FGLPROFILE entries.")

[Customizing GWA apps](5132-customizing-gwa-apps.md "This section describes ways you may customize your GWA apps, such as adding a favicon, customizing the index page, and internationalizing your app.")
