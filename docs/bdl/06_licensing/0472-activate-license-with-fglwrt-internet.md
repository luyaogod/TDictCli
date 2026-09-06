---
title: "Activate your BDL license from the command line (via internet)"
source: "genero-install-topics/t_bdl_finalize_temporary_license_fglwrt_with_internet.html"
breadcrumb: "Licensing > Steps to activate temporary BDL license > Activate license with fglWrt (internet)"
type: "task"
---

# Activate your BDL license from the command line (via internet)

> If a Genero Business Development Language (BDL) license has been installed, you can activate it using the command fglWrt -k auto. The license is registered with Four Js via the internet and the installation and maintenance/subscription key is installed.

**Internet is required**

**Before you begin:** Internet
access is required to activate the license; a temporary installation does not require
internet.

The license must be installed but not yet
activated.

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
   fglWrt -k auto
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

   The installation key and the maintenance/subscription key are
   installed automatically and the message is displayed.

   ```
   Installing maintenance key 'MMMMMMMMMMMM'
   License installation successful.
   ```

   License
   installation is now completed.

> **Note:**
>
> There is no need to run `fglWrt -m auto` to install the maintenance or
> subscription key.
