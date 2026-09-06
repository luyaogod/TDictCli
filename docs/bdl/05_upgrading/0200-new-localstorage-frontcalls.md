---
title: "New localStorage frontcalls"
source: "fgl-topics/c_fgl_Migrate_to_310_localstorage.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > New localStorage frontcalls"
type: "concept"
---

# New localStorage frontcalls

> New localStorage frontcalls replace GAS specific session.setVar and session.getVar calls.

The `session` module front calls `setVar` and
`getVar` are deprecated. Starting with version 3.10, you can use the new
`localStorage` front calls, supported by all Genero front-ends.

For more details, see [Local storage front calls](../15_library-reference/3435-local-storage-front-calls.md "This section describes front calls to store data on the front-end platform.").

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
