---
title: "The FIELD form item type and .val schema file"
source: "fgl-topics/c_fgl_Migrate_to_251_field_val.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.51 upgrade guide > The FIELD form item type and .val schema file"
type: "concept"
---

# The FIELD form item type and .val schema file

> Form files using the FIELD item type and/or .val attribute definitions must be reviewed.

Starting with version 2.51, the `FIELD` item type defining abstract fields in
forms, based on `.val` schema file attributes is deprecated.

Furthermore, any non-I4GL attribute defined in the `.val` schema file must be
avoided. Reading attributes in the `.val` is now only supported for compatibility
with I4GL projects.

With Genero, it is recommended to define all form item attributes in the form definition
file.

## Related links

**Related concepts**  

[Form specification files](../11_user-interface/1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.")
