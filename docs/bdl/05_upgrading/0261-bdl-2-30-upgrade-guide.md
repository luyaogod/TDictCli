---
title: "BDL 2.30 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_230.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide"
type: "concept"
---

# BDL 2.30 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 2.30.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 2.30 new features](0070-bdl-2-30-new-features.md "Features added in 2.30 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 2.21 upgrade guide](0272-bdl-2-21-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.21.").

## Child topics

- [GUI server auto start](0262-gui-server-auto-start.md): FGLSERVER defaults the server defined by wsmap settings, when starting GUI server
- [Form compiler is more strict](0263-form-compiler-is-more-strict.md): The .per grammar parser has been reviewed to deny invalid code.
- [ORACLE and INTERVAL columns](0264-oracle-and-interval-columns.md): INTERVAL storage bug fix needs a review of existing databases in production.
- [DIALOG.setCurrentRow() changes row selection flags](0265-dialog-setcurrentrow-changes-row-selection-flags.md): Row selection flags are reset by a call to setCurrentRow().
- [Schema extractor needs table owner](0266-schema-extractor-needs-table-owner.md): The fgldbsch schema extractor requires a -ow option to distinguish different database users/shemas.
- [Windows installation for all users only](0267-windows-installation-for-all-users-only.md): Installation on Windows platforms is for all users.
- [MenuAction close no longer created by default](0268-menuaction-close-no-longer-created-by-default.md): The close action is no longer created by default in MENU dialog.
- [Emulated scrollable cursor temp files in DBTEMP](0269-emulated-scrollable-cursor-temp-files-in-dbtemp.md): Directory of scrollable cursor data storage can be defined with DBTEMP.
- [Modifying tree view data during dialog execution](0270-modifying-tree-view-data-during-dialog-execution.md): Use ui.Dialog methods to insert/append/delete treeview nodes.
- [Presentation styles changes](0271-presentation-styles-changes.md): Modifications to consider when using presentation styles.
