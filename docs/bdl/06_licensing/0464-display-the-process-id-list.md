---
title: "Display the process Id List"
source: "genero-install-topics/t_license_controller_display_process_id_list.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Managing the active users > Display the process Id List"
type: "task"
---

# Display the process Id List

> Use this procedure to display the current list of processes on your machine.

To display processes, execute the command:

```
fglWrt -a ps
```

A list of process Ids are displayed as shown.

![Image shows list of process ids currently running on a machine where the fglWrt -a ps command is executed](../_images/fglwrt_a_ps.jpg)

*Display the ids of Processes Currently Running*

## Related links

**Related tasks**  

[Force a check of registered users list](0462-check-registered-users-list.md "After a system crash or an abnormal termination of an application, use this procedure to force a check of the registered users list on the current machine.")

[Clean registered users list](0463-clean-registered-users-list.md "Generally, you never need to clean the registered users list on a host, but if, for example, a DVM crashes and the associated unreleased licenses would not be recovered otherwise, use this procedure to clear the list.")

[Drop session](0465-drop-session.md "Use this procedure to drop a session referenced by a specified process id on a host.")

**Related reference**  

[License details reference](0451-license-details-reference.md "A reference to license details.")
