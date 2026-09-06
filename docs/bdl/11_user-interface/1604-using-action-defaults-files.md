---
title: "Using action defaults files"
source: "fgl-topics/c_fgl_action_defaults_files_usage.html"
breadcrumb: "User interface > Form definitions > Action defaults files > Using action defaults files"
type: "concept"
---

# Using action defaults files

> To use action default files, you must understand how they work and how to structure the code.

Global action defaults are defined in an XML file with the
.4ad extension. By default, the runtime system searches for a file named
default.4ad in several directories as described in [the FGLRESOURCEPATH reference topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files."). If no
file was found, standard action default settings are loaded from the
$FGLDIR/lib/default.4ad file.
> **Important:**
>
> Global action defaults must be defined in a unique file; you cannot combine
> several `4ad` files.

Action defaults files usage is related to action configuration concepts. For more details, see
[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.").

## Related links

**Related concepts**  

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")
