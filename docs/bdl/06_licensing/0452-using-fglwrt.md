---
title: "Get license information using fglWrt"
source: "genero-install-topics/t_bdl_license_get_info_using_fglwrt.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Get license information > Using fglWrt"
type: "task"
---

# Get license information using fglWrt

> Get details of the status of the license of your Genero product.

**About this task:**

Use this procedure, for example, to find out the
date your maintenance contract or your subscription license expires.
> **Tip:**
>
> You can also use the procedure described in this topic to retrieve the **installation number**
> when you need to register the license and finalize the installation.

From a command prompt, enter the command:

```
fglWrt -a info license
```

The output from the command displays details about the installed
license. This is a temporary runtime license for the Four Js Dynamic Virtual Machine. The license is
date-limited and will expire on the date shown. After this date, the license can no longer be used
as a temporary installation—you must register and activate it to continue using it. Currently, this
temporary installation will expire in 85 days, so activation should be completed before that
deadline. The installation number shown will be required during the registration
process.

```
License      : GSA-XXXXXXXX
License key  : HBZ2LV7Z5YFEG24A45KT6VK24FLE6XE76T9RLA5BZJJQM4RSVSHUI8L0ZSEQOBQC
Product      : Four Js Dynamic Virtual Machine
Type         : Runtime version
Users        : 5
Extensions   :
        - Open Database Interface
        - Enhanced license key format
        - Strict licensing option
This license is date limited and expires the 2026/01/31 (Year/Month/Day).
Warning! This is a temporary license, installation number is 
'PJZHAM36OE74IANZ7CGRGYZTP6XYOGCJGEA0MYI24FGMVYFD4P2YBP25DGEPRQAWRLFJGQ5A4HE'.
This temporary installation will expire in 85 day(s).
End of maintenance date: 2026/01/30 (Year/Month/Day).
```

## Related links

**Related reference**  

[License details reference](0451-license-details-reference.md "A reference to license details.")
