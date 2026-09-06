---
title: "Install or update BDL maintenance/subscription key from the command line (no internet)"
source: "genero-install-topics/t_bdl_install_maintenance_key_fglwrt_without_internet.html"
breadcrumb: "Licensing > Steps to apply maintenance/subscription key (BDL) > Apply maintenance/subscription key with fglWrt (no internet)"
type: "task"
---

# Install or update BDL maintenance/subscription key from the command line (no internet)

> Without internet access, you register the license with Four Js from another machine via the internet, or by phone. Then use the fglWrt command line tool to install or update the maintenance/subscription key of your product.

> **Important:**
>
> If your installation directory is in the C:\Program Files path, you must run
> as administrator when you license the product. This avoids any permission issues.

1. From a machine with internet access, open a browser and navigate to the [Four Js web site](https://4js.com/support/registration/) to get the maintenance/subscription key. See the procedure described in [Get the maintenance/subscription key](0468-get-maintenance-subscription-key.md "You get the maintenance/subscription key from the Four Js website to ensure you have a valid key to install or update your product's license.").
2. Start the command line interface.

   Open a command prompt.

   - On Linux®/UNIX®/macOS™, open a command prompt. "sudo" may be required.
   - On Windows®, open the Command Prompt
     from the
     Start menu .
3. At the command line enter the command to install the maintenance/subscription key:

   ```
   fglWrt -m mkey
   ```

   Where mkey is the maintenance/subscription key to enter.

   A message is displayed in the output to say the license installation was successful.

   License installation is now completed.
