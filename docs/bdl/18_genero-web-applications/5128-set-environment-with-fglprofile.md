---
title: "Set environment with fglprofile"
source: "fgl-topics/c_fgl_gwa_envvariables_in_fglprofile.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Set environment with fglprofile"
type: "concept"
---

# Set environment with fglprofile

> Configure environment settings with FGLPROFILE entries.

The GWA behaves similarly to the GMA/GMI when it comes to setting the environment. For more
information, go to [Setting environment variables in FGLPROFILE (mobile)](../07_configuration/0495-setting-environment-variables-in-fglprofile-mobile.md).

Setting environment variables for your app must be done in an fglprofile
file, which should be located in the programdir directory,
alongside the main program module. The GWA app reads the default fglprofile
provided in $FGLDIR/etc/fglprofile. However, if `gwabuildtool`
detects a custom fglprofile in the
programdir, it ensures that the runtime locates and uses
this custom version.

The two main environment variables you may have to set for GWA are [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") and [FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources."):

1. If you have to set `FGLLDPATH` because some `42m` files are
   located on subdirectories, you must create a `fglprofile` with an entry called
   `mobile.environment.FGLLDPATH`
2. You can also set other environment variables such as `FGLIMAGEPATH`

   In this
   example, environments for FGLLDPATH and FGLIMAGEPATH are set. "`$FGLAPPDIR`" refers
   to the [virtual file system](5125-file-system.md "When the GWA application is launched in the browser, a virtual UNIX-like file system emulation of your GWA application is created in memory.")
   /app directory in memory.

   For a portable variant of your GWA
   application, set these environment variables as
   shown:

   ```
   #portable variant which works also for GMI/GMA
   mobile.environment.FGLLDPATH  = "$FGLAPPDIR/dbsync"
   mobile.environment.FGLIMAGEPATH="$FGLAPPDIR/image_dir:$FGLDIR/lib/image2font.txt"
   ```

   For
   a non-portable variant, or for GWA specific, set the environment variables as
   shown:

   ```
   # non portable: GWA specific
   mobile.environment.FGLLDPATH  = "/app/dbsync"
   mobile.environment.FGLIMAGEPATH  = "/app/image_dir:/fgl/lib/image2font.txt"
   ```

For more details about fglprofile settings, go to [Understanding FGLPROFILE](../07_configuration/0484-understanding-fglprofile.md "The runtime system uses one or more configuration files in which you can define options and parameters to change the behavior of the programs.").

## Related links

**Related concepts**  

[Specify the GBC](5129-specify-the-gbc.md "A GBC must be bundled with the GWA.")

[Package web components](5130-package-web-components.md "Web components must be packaged with the GWA application because they are preloaded and must also be available offline.")

[Specify the start-up module](5131-specify-the-start-up-module.md "Specify the module you want to use as the main entry point.")

[Customizing GWA apps](5132-customizing-gwa-apps.md "This section describes ways you may customize your GWA apps, such as adding a favicon, customizing the index page, and internationalizing your app.")
