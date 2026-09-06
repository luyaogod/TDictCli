---
title: "FGLWRTUMASK removal"
source: "fgl-topics/c_fgl_Migrate_to_400_FGLWRTUMASK.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > FGLWRTUMASK removal"
type: "concept"
---

# FGLWRTUMASK removal

> The FGLWRTUMASK environment variable is no longer supported.

Before BDL 4.00 (using FLM 5), the FGLWRTUMASK environment variable could be defined to specify a
Unix file permission mask for the FGLDIR/lock directory creation.

Starting with BDL 4.00 (now using FLM 6), the FGLWRTUMASK environment variable is no longer
supported and the FGLDIR/lock directory will be created with
`rwxrwxrwx` Unix file permissions.

For more details, see Four Js License Manager User Guide.
