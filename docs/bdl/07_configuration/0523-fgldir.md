---
title: "FGLDIR"
source: "fgl-topics/c_fgl_EnvVariables_FGLDIR.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLDIR"
type: "concept"
---

# FGLDIR

> Defines the installation directory of Genero Business Development Language.

The FGLDIR environment variable defines the [installation directory](../04_installation/0031-installation.md "This chapter contains installation and setup instructions.") of the runtime system and compilers of Genero.

When executing on a mobile device, the FGLDIR environment variable is automatically set
by the front-end component, before starting the runtime system component. As result, it is possible
to use the $FGLDIR keyword in [FGLPROFILE
environment variable settings](0495-setting-environment-variables-in-fglprofile-mobile.md) when executing on mobile devices.

Make sure that the FGLDIR environment variable is properly set, and that the PATH environment
variable contains FGLDIR/bin to find Genero BDL command line tools. The
behavior is undefined if the PATH is not in sync with FGLDIR.

## Related links

**Related concepts**  

[PATH](0499-path.md "Defines the list of paths to find executable files.")
