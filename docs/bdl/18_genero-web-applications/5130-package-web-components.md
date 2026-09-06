---
title: "Package web components"
source: "fgl-topics/c_fgl_gwa_web_component.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Package web components"
type: "concept"
---

# Package web components

> Web components must be packaged with the GWA application because they are preloaded and must also be available offline.

The directory programdir/webcomponents is the default
location for web components in your program directory. All components located there will be
automatically packaged.

Furthermore, the [gwabuildtool](5150-gwabuildtool.md "The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser.") has a
`--webcomponent` switch to bundle a web component from another location in the file
system other than programdir/webcomponents. The switch can
be used multiple times in the build command.

## Including built-in web components

The gwabuildtool also scans all form files (.42f) in your
application and checks them for the `COMPONENTTYPE` attribute. This affects the build in the following way:

1. It adds the specified [Built-in web components](../11_user-interface/2415-built-in-web-components.md "Genero provides a set of ready-to-use web components.") from
   $FGLDIR/webcomponents to the build. For example, if a
   `COMPONENTTYPE` matches "fglsvgcanvas", the fglsvgcanvas web component is
   automatically included in the GWA.
2. It raises a warning if, for a particular `COMPONENTTYPE`, no web component has
   been packaged.

A call to [`ui.Interface.filenameToURI`](../15_library-reference/3098-ui-interface-filenametouri.md "Converts a filename to a URI to be used as a web component image resource.") (asset\_in\_filesystem) can be
used to produce URLs to make a web component visible.

> **Tip:**
>
> **Identifying web component URL paths**
>
> Notice that URL paths starting with
> "webcomponent" refer to files bundled in the
> programdir/webcomponents subdirectory and URLs starting with
> HTTP(S) can refer to assets in the internet.

The `gwa_webcos` demo is a good source to check the various asset paths.

## Related links

**Related concepts**  

[Specify the start-up module](5131-specify-the-start-up-module.md "Specify the module you want to use as the main entry point.")

[Customizing GWA apps](5132-customizing-gwa-apps.md "This section describes ways you may customize your GWA apps, such as adding a favicon, customizing the index page, and internationalizing your app.")

[Specify the GBC](5129-specify-the-gbc.md "A GBC must be bundled with the GWA.")

[Set environment with fglprofile](5128-set-environment-with-fglprofile.md "Configure environment settings with FGLPROFILE entries.")

[GWA demos and examples](5156-gwa-demos-and-examples.md "Demos and examples are provided for Genero Web Application.")
