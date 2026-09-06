---
title: "License details reference"
source: "genero-install-topics/c_license_details_reference.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Get license information > License details reference"
type: "concept"
---

# License details reference

> A reference to license details.

The output from the license controller command, for example, fglWrt -a info license, and/or the Genero
Licenser application screen displays details about the installed license. Use Table 1 as a reference
to look up details of your license.
> **Note:**
>
> For specific license details, see your license agreement or contact your local Four Js sales
> office.

| Details displayed | Description |
| --- | --- |
| Server name | The name of the license manager server.**Note:**Indicates that licensing is managed by a License Manager. |
| TCP Port | The port where the license manager is listening.**Note:**Indicates that licensing is managed by a License Manager. |
| License | The value of the license number. |
| License Key | The value of the license key. |
| Product | The name of the product under license, for example:Four J's Universal Compiler (for Genero Enterprise and Genero Mobile products)Four Js Genero Report Writer (for Genero Report Engine) |
| Type | The type of license. This can be:`runtime``development` |
| Users | The maximum number of users allowed to use the license. |
| Extension(s) | Information on some specific conditions of the license. For example, "- Open Database Interface" means that your license allows you to use any database.The **installation number** is displayed if the license is a temporary license and the installation has not been finalized.If the license is **date-limited**, then the date when the license expires is shown. |
| End of subscription/maintenance date | This is the date your subscription license ends or the date your maintenance contract expires, depending on what type of license you have. The message will indicate the type of license you have, this can be:`End of subscription date:YYYY/MM/DD (Year/Month/Day)``End of maintenance date: YYYY/MM/DD (Year/Month/Day)` |

## Related links

**Related tasks**  

[Get license information using fglWrt](0452-using-fglwrt.md "Get details of the status of the license of your Genero product.")

[Get license information using BDL licenser](0453-using-bdl-licenser.md "Get details about your Genero Business Development Language (BDL) license using the Genero licensing graphical interface.")
