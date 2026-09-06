---
title: "Searching for image files on the application server"
source: "fgl-topics/c_fgl_Migrate_to_220_image_files.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Searching for image files on the application server"
type: "concept"
---

# Searching for image files on the application server

> For security reasons, the image file transfer mechanism has been slightly modified in version 2.20.

(This modification has also been back-ported in 2.11.14):

If FGLIMAGEPATH is set, the current working directory is no longer
searched as in previous versions. You must explicitly add "." to the
list of directories. By default, if FGLIMAGEPATH is not defined, the
runtime system still searches the current directory.

If FGLIMAGEPATH is defined, the image files used in IMAGE form fields or in the IMAGE attribute
must be located below one of the directories listed in the environment variable. This
constraint does not exist if FGLIMAGEPATH is not set and has been relaxed in 2.21.00 for
image fields displayed by program.

Starting with 2.21.00, images displayed by program to IMAGE fields are considered as valid files
to be transferred to the clients without risk and do not follow the FGLIMAGEPATH security
restrictions. Images are however searched for in the path list defined in FGLIMAGEPATH.

## Related links

**Related concepts**  

[FGLIMAGEPATH](../07_configuration/0527-fglimagepath.md "Defines a list of paths and filenames for image resources.")

[IMAGE attribute](../11_user-interface/1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.")

[IMAGE item type](../11_user-interface/1695-image-item-type.md "Defines an area that can display an image resource.")
