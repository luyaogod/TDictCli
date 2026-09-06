---
title: "License BDL in a container"
source: "genero-install-topics/t_bdl_license_bdl_in_container.html"
breadcrumb: "Licensing > License BDL in a container"
type: "task"
---

# License BDL in a container

> License a Genero Business Development Language (BDL) product in a container using the fglWrt command-line tool. Container licenses do not require activation.

You need a container license (a special license for containerized environments) provided as one
of the following: a license file, a license string, or a combination of license number, license key,
and maintenance key.

Container licenses are for containers only and never require activation. If you try to install a
container license outside of a container environment, you’ll see the following error:

```
(FLM-153): Container licenses can only be used in containers.
License cannot be installed.
Contact your support center.
```

Manage container licenses directly using command-line tools; the Four Js License Manager does not
support them. At the start of the container, install the license using the fglWrt
command.

1. Install the license using one of the following methods:

   - **Using a license file:**

     ```
     fglWrt -f license-file
     ```
   - **Using a license
     string:**

     ```
     fglWrt --install-license-string license-string
     ```
   - **Using a license number and key:**

     ```
     fglWrt -L license-number license-key maintenance-key
     ```

   The license is installed in the container.
2. (Recommended) When the container stops, uninstall the license.

   ```
   fglWrt -d
   ```
3. To check the license expiry date from within the container, use:

   ```
   fglWrt -a info
   ```

   The license information is displayed, including the expiry date. The expiry date matches
   the contract signature date.
