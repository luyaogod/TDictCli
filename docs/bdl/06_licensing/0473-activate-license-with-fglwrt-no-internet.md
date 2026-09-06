---
title: "Activate your BDL license from the command line (no internet)"
source: "genero-install-topics/t_bdl_finalize_temporary_license_fglwrt_without_internet.html"
breadcrumb: "Licensing > Steps to activate temporary BDL license > Activate license with fglWrt (no internet)"
type: "task"
---

# Activate your BDL license from the command line (no internet)

> Without internet access, you register the license with Four Js from another machine via the internet, or by phone. Then activate your BDL license using the fglWrt command line tool.

> **Important:**
>
> If your installation directory is in the C:\Program Files path, you must run
> as administrator when you license the product. This avoids any permission issues.

1. From a machine that has internet access or a smart phone, go to the
   License your products page on the [Four Js web site](https://4js.com/support/registration/) and follow the procedure described in [Register your license](0466-register-license.md "To validate your license, you must register it on the Four Js website. You can access the website from any machine or smart phone.").

   With your license registered, you have the required keys.
2. Start the command line interface.

   Open a command prompt.

   - On Linux®/UNIX®/macOS™, open a command prompt. "sudo" may be required.
   - On Windows®, open the Command Prompt
     from the
     Start menu .
3. At the command line enter the command to install the installation key:

   ```
   fglWrt -k installation-key
   ```

   A message is displayed in the output to say the installation key was installed
   successfully.
4. At the command line enter the command to install the maintenance/subscription key:

   ```
   fglWrt -m mkey
   ```

   Where mkey is the maintenance/subscription key to enter.

   A message is displayed in the output to say the license installation was successful.

   License installation is now completed.
