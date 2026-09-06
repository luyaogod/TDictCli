---
title: "Install or update BDL maintenance/subscription key from the command line (via internet)"
source: "genero-install-topics/t_bdl_install_maintenance_key_fglwrt_with_internet.html"
breadcrumb: "Licensing > Steps to apply maintenance/subscription key (BDL) > Apply maintenance/subscription key with fglWrt (internet)"
type: "task"
---

# Install or update BDL maintenance/subscription key from the command line (via internet)

> You can install or update the maintenance/subscription key of your BDL product using the command fglWrt -m auto.

**Internet is required**

**Before you begin:** Internet
access is required to activate the license; a temporary installation does not require
internet.

> **Important:**
>
> If your installation directory is in the C:\Program Files path, you must run
> as administrator when you license the product. This avoids any permission issues.

1. Start the command line interface.

   Open a command prompt.

   - On Linux®/UNIX®/macOS™, open a command prompt. "sudo" may be required.
   - On Windows®, open the Command Prompt
     from the
     Start menu .
2. At the command line enter the command: 

   ```
   fglWrt -m auto
   ```
3. When prompted, enter your customer code:

   ```
   Enter your customer code >
   ```
4. At the following prompt: 

   ```
   Do you need to configure an HTTP proxy ? (y/n)
   ```

   Enter
   `y` if your access to the internet is through a
   proxy. Provide the required information for HTTP proxy, port, and
   authentication when prompted. Otherwise, enter `n`.

   The maintenance/subscription key is then retrieved from the Four Js Activation
   server and a message is displayed:

   ```
   Maintenance key installation successful.
   ```
