---
title: "Clean registered users list"
source: "genero-install-topics/t_license_controller_clean_reg_user_list.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Managing the active users > Clean registered users list"
type: "task"
---

# Clean registered users list

> Generally, you never need to clean the registered users list on a host, but if, for example, a DVM crashes and the associated unreleased licenses would not be recovered otherwise, use this procedure to clear the list.

> **Important:**
>
> The fglWrt
> -i  should be used with caution because it invalidates implicitly **all**
> existing sessions and eventually all running applications will stop with a license error. In other
> cases, the -u option to force a check of the registered users should be preferred
> to the -i option.

To clean the registered users list, execute the license controller command:

fglWrt -i

All users registered are disconnected when you use this option.

![Image shows the warning about deleting all existing sessions that is shown when the fglWrt -i command is executed](../_images/fglwrt_i.png)

*Clean the Registered Users with fglWrt -i*

## Related links

**Related tasks**  

[Force a check of registered users list](0462-check-registered-users-list.md "After a system crash or an abnormal termination of an application, use this procedure to force a check of the registered users list on the current machine.")

[Display the process Id List](0464-display-the-process-id-list.md "Use this procedure to display the current list of processes on your machine.")

[Drop session](0465-drop-session.md "Use this procedure to drop a session referenced by a specified process id on a host.")

**Related reference**  

[License details reference](0451-license-details-reference.md "A reference to license details.")
