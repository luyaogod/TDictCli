---
title: "Emulated scrollable cursor temp files in DBTEMP"
source: "fgl-topics/c_fgl_Migrate_to_230_scroll_cursor_tmpfile.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > Emulated scrollable cursor temp files in DBTEMP"
type: "concept"
---

# Emulated scrollable cursor temp files in DBTEMP

> Directory of scrollable cursor data storage can be defined with DBTEMP.

On UNIX™ platforms, starting with 2.30, the temporary files
for emulated scrollable cursors will be created in the directory defined by the [DBTEMP](../07_configuration/0517-dbtemp.md "Defines the directory for temporary files.")
environment variable when defined, otherwise TMPDIR, TEMP or TMP will be used. Using DBTEMP for
database files conforms to DBTEMP usage for temporary files of TEXT and BYTE data storage.

## Related links

**Related concepts**  

[TMPDIR, TMP, TEMP](../07_configuration/0503-tmpdir-tmp-temp.md "Defines the directory for temporary files.")

[TEXT](../08_language-basics/0569-text.md "The TEXT data type stores large text data.")

[BYTE](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.")
