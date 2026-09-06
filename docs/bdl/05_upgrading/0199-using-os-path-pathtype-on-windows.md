---
title: "Using os.Path.pathType() on Windows"
source: "fgl-topics/c_fgl_Migrate_to_310_path_type.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Using os.Path.pathType() on Windows"
type: "concept"
description: "How to identify absolute file paths on Windows, whether using a drive letter or not."
---

# Using os.Path.pathType() on Windows

> How to identify absolute file paths on Windows®, whether using a drive letter or not.

Before version 3.10, on Windows
platforms, a filename starting with a directory-separator (slash or backslash) was not interpreted
by `os.Path.pathType()` as an absolute name. As result, such files where considered
as relative path and searched by concatenating the path to the elements defined in FGLLDPATH,
FGLRESOURCEPATH, FGLIMAGEPATH.

Starting with version 3.10, a filename starting with a directory-separator or starting with a
drive letter, is considered an absolute file path.

This change has also an impact on the `os.Path.join`
method, which returns the second parameter only if it is identified as an absolute path.

## Related links

**Related concepts**  

[FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.")

[FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.")

[FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.")

[os.Path.pathType](../15_library-reference/3734-os-path-pathtype.md "Checks if a path is a relative path or an absolute path.")
