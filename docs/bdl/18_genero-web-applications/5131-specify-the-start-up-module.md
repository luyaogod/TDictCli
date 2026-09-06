---
title: "Specify the start-up module"
source: "fgl-topics/c_fgl_gwa_main_module.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Specify the start-up module"
type: "concept"
---

# Specify the start-up module

> Specify the module you want to use as the main entry point.

By default, GWA will load the "main.42m" module as the main entry point, for
example, programdir/app/main.42m. You can change this
behavior in two ways:

1. In the URL query string, you can specify
   `fglapp=anothermodule`. GWA will then look in the embedded file
   system for a 42m of that name (for example,
   /app/anothermodule.42m)
2. Or with the [gwabuildtool](5150-gwabuildtool.md "The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser."), you can use the option
   `--main-module` to specify the initial 42m module to be
   loaded.

## Related links

**Related concepts**  

[Customizing GWA apps](5132-customizing-gwa-apps.md "This section describes ways you may customize your GWA apps, such as adding a favicon, customizing the index page, and internationalizing your app.")

[Specify the GBC](5129-specify-the-gbc.md "A GBC must be bundled with the GWA.")

[Set environment with fglprofile](5128-set-environment-with-fglprofile.md "Configure environment settings with FGLPROFILE entries.")

[Package web components](5130-package-web-components.md "Web components must be packaged with the GWA application because they are preloaded and must also be available offline.")
