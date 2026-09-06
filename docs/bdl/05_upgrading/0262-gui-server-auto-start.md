---
title: "GUI server auto start"
source: "fgl-topics/c_fgl_Migrate_to_230_guiserver_autostart.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > GUI server auto start"
type: "concept"
---

# GUI server auto start

> FGLSERVER defaults the server defined by wsmap settings, when starting GUI server

> **Important:**
>
> This feature is desupported. It may still be
> available, but can be removed in a next version.

Before version 2.30, the runtime system was trying to connect to `localhost:0`
when FGLSERVER was not set, even if `gui.server.autostart` FGLPROFILE entries are
defined.

This behavior has been identified as a bug (FGL-1583) and has been fixed in 2.30, changing the
way fglrun proceeds with the GUI connection when autostart settings are defined; with 2.30, the
`wsmap` workstation mappings are now taken into account, so that FGLSERVER defaults
to `locahost:n`, where n is the GUI server
number found from the wsmap settings.

However, starting with version 6.00.01, the GUI server auto start feature has been removed.

## Related links

**Related concepts**  

[Automatic front-end startup](../11_user-interface/1530-automatic-front-end-startup.md "Automatic front-end startup")
