---
title: "Drop session"
source: "genero-install-topics/t_license_controller_drop_session.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Managing the active users > Drop session"
type: "task"
---

# Drop session

> Use this procedure to drop a session referenced by a specified process id on a host.

This feature is available starting from version 5.19.01.

> **Important:**
>
> This feature only applies to User- type licenses, not to CPU licenses.

To drop a session, execute the command:

```
fglwrt -x  pid
```

Where pid is the process id.

## Related links

**Related tasks**  

[Force a check of registered users list](0462-check-registered-users-list.md "After a system crash or an abnormal termination of an application, use this procedure to force a check of the registered users list on the current machine.")

[Clean registered users list](0463-clean-registered-users-list.md "Generally, you never need to clean the registered users list on a host, but if, for example, a DVM crashes and the associated unreleased licenses would not be recovered otherwise, use this procedure to clear the list.")

[Display the process Id List](0464-display-the-process-id-list.md "Use this procedure to display the current list of processes on your machine.")

**Related reference**  

[License details reference](0451-license-details-reference.md "A reference to license details.")
