---
title: "Built-in front-end icons desupport"
source: "fgl-topics/c_fgl_Migrate_to_300_builtin_icons.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Built-in front-end icons desupport"
type: "concept"
---

# Built-in front-end icons desupport

> Image resources included in front-ends are desupported with Genero 3.00.

Starting with Genero 3.00, the icon files distributed in front-end packages are no
longer provided (as before in the GDC-installation-dir/pics for example)

Common icons for buttons, toolbars, topmenus, and other items using icons can be
centralized on the application side where the program executes. It is recommended that this feature
be used to provide the same icons on different types of front-ends, or use specific icons, but from
the same central icon directory. For more details, see [Providing the image resource](../11_user-interface/1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.").

Note that mobile front-ends will display default icons, for default action
views, if no `IMAGE` attribute is specified for the action. See [Action views on mobile devices](../11_user-interface/2297-action-views-on-mobile-devices.md "Action views are rendered following mobile specific standards.") for more details.

## Related links

**Related concepts**  

[FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.")
