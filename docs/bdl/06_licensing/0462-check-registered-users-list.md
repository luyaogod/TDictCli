---
title: "Force a check of registered users list"
source: "genero-install-topics/t_license_controller_force_check.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Managing the active users > Check registered users list"
type: "task"
---

# Force a check of registered users list

> After a system crash or an abnormal termination of an application, use this procedure to force a check of the registered users list on the current machine.

This license controller feature forces the synchronization of the real users list and the
registered users list for the current machine.

To check the registered users list, execute the command:

```
fglWrt -u
```

## Related links

**Related tasks**  

[Clean registered users list](0463-clean-registered-users-list.md "Generally, you never need to clean the registered users list on a host, but if, for example, a DVM crashes and the associated unreleased licenses would not be recovered otherwise, use this procedure to clear the list.")

[Display the process Id List](0464-display-the-process-id-list.md "Use this procedure to display the current list of processes on your machine.")

[Drop session](0465-drop-session.md "Use this procedure to drop a session referenced by a specified process id on a host.")

**Related reference**  

[License details reference](0451-license-details-reference.md "A reference to license details.")
