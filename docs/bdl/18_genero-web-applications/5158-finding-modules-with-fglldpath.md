---
title: "Finding modules with FGLLDPATH"
source: "fgl-topics/c_gwa_setting_fglldpath.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > Finding modules with FGLLDPATH"
type: "concept"
---

# Finding modules with FGLLDPATH

> For compilation, you may need to set the FGLLDPATH environment variable to find the modules in the GWA API package.

When you install GWA, you have the option to install it on top of FGLDIR or in a directory of its
own.

If GWA is installed on top of FGLDIR, there is no need to modify FGLLDPATH. The gwa package will
be found automatically by the compiler, as the package will sit in
$FGLDIR/lib/gwa.

If GWA is installed in its own directory, you must add $GWADIR/lib to FGLLDPATH:

- On UNIX™:

  ```
  export FGLLDPATH=$FGLLDPATH:gwa-install-dir/lib
  ```
- On Windows®:

  ```
  set FGLLDPATH=%FGLLDPATH%;gwa-install-dir\lib
  ```

> **Tip:**
>
> By default, FGLLDPATH is set to correctly locate gwa modules at runtime; you should not need to
> update FGLLDPATH unless you have custom modules that need locating.

## Related links

**Related concepts**  

[FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.")
