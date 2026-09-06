---
title: "Front calls changes"
source: "fgl-topics/c_fgl_Migrate_to_400_frontcalls.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Front calls changes"
type: "concept"
---

# Front calls changes

> Modifications to consider when using front calls.

## Deprecated front calls

The following front calls are deprecated:

- Standard front calls
  - [`standard.shellExec`](../15_library-reference/3417-standard-shellexec.md "Opens a file on the front-end platform with the program associated to the file extension."):
    Has security issues, use [`standard.launchURL`](../15_library-reference/3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.") instead.

## Desupported front calls

The following front calls are desupported:

- Standard front calls
  - `getWindowId`
  - `feInfo` options:
    - `dictionariesDirectory`
    - `isActiveX`
- GDC WINCOM, WINDDE, WinMail front calls: Use [Apache
  POI](../14_extending-the-language/2701-example-2-using-the-apache-poi-framework.md) instead.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases. For more details, see [Front
call changes in BDL 3.20](0162-front-calls-changes.md "Modifications to consider when using front calls.").

Notable changes introduced in maintenance releases:

- [monitor.update parameter for elevation prompt (GDC only)](0162-front-calls-changes.md), also available in GDC 4.00.01.

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
