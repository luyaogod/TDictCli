---
title: "The FGLDIR/lock directory"
source: "genero-install-topics/c_bdl_license_FGLDIR_lock.html"
breadcrumb: "Licensing > Manage Genero BDL local license > The FGLDIR/lock directory"
type: "concept"
---

# The FGLDIR/lock directory

> When using a local license, the license controller uses the FGLDIR/lock directory to store information (number of active users). A user running a BDL program must have access rights to this directory. Permissions can be restricted to specific users with FGLWRTUMASK.

This FGLDIR/lock directory must have access rights for any user running a
BDL program. If it does not exist, it is automatically created.

> **Note:**
>
> The FGLDIR/lock directory is **not** used when Genero is configured to
> use a [license manager](0480-license-genero-bdl-using-flm.md "License your Genero Business Development Language (BDL) installation using a license manager on the network.").

By default, the FGLDIR/lock directory is created with
`rwxrwxrwx` rights, to let any user access the directory and create files. To
restrict the access to a specific group or user, use the `FGLWRTUMASK` environment
variable to force fglWrt to use a specific mask when creating the lock
directory:

`$ FGLWRTUMASK="022"; export FGLWRTUMASK`

> **Important:**
>
> The `FGLWRTUMASK` environment variable must be set for all users executing BDL
> programs, because the FGLDIR/lock directory can be re-created by any user at
> first BDL program execution.
