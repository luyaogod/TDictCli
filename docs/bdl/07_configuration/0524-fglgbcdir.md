---
title: "FGLGBCDIR"
source: "fgl-topics/c_fgl_EnvVariables_FGLGBCDIR.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLGBCDIR"
type: "concept"
---

# FGLGBCDIR

> Defines the GBC component to be used in GUI direct mode.

The FGLGBCDIR environment variable can be used to define the directory where the GBC component
must be loaded from.

The FGLGBCDIR environment variable is provided to select a given GBC when using the GUI direct
mode. FGLGBCDIR is ignored when using the GAS: As a general pattern, if a specific GBC is to be
coupled with a particular application in production, consider shipping GBC in
appdir/gbc. Otherwise, use the GBC available by default.

The GBC component will be searched in the following directories:

1. The appdir/gbc directory, where
   appdir is the directory where the [program
   file](../09_advanced-features/0829-executing-programs.md "There are different ways to execute compiled programs, depending on the configuration and the development or production context.") is located,
2. The directory defined in the [FGLGBCDIR](0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode.")
   environment variable,
3. The [$FGLDIR/web\_utilities/gbc/gbc](0523-fgldir.md "Defines the installation directory of Genero Business Development Language.") directory.

If defined, make sure that FGLGBCDIR is set to a directory containing the GBC component
files.

## Related links

**Related concepts**  

[Graphical mode rendering (GUI mode)](../11_user-interface/1518-graphical-mode-rendering-gui-mode.md "Graphical mode rendering (GUI mode)")

[Connecting with a front-end](../11_user-interface/1521-connecting-with-a-front-end.md "Connecting with a front-end")
