---
title: "Front calls changes"
source: "fgl-topics/c_fgl_Migrate_to_300_front_calls.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Front calls changes"
type: "concept"
---

# Front calls changes

> Describes changes applied to front calls.

Front call modifications in BDL version 3.00:

- Before version 3.00, the `connectivity` front call accepted a host name as
  parameter. Starting with version 3.00, this front call no longer uses a host name. It will only
  check the available network type. For more details, see [the mobile.connectivity front call](../15_library-reference/3447-mobile-connectivity.md "Returns the type of network available for the mobile device.").

## Related links

**Related concepts**  

[Front calls](../09_advanced-features/0962-front-calls.md "Front call functions execute on the platform where the front-end is installed.")
