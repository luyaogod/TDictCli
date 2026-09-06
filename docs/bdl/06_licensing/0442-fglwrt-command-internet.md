---
title: "License BDL from the command line (via internet)"
source: "genero-install-topics/t_bdl_license_fglwrt_with_internet.html"
breadcrumb: "Licensing > Steps to BDL license installation > fglWrt command (internet)"
type: "task"
---

# License BDL from the command line (via internet)

> You can license a Genero Business Development Language (BDL) product using the command line tool, fglWrt.

**Internet is required**

**Before you begin:** Internet
access is required to activate the license; a temporary installation does not require
internet.

> **Important:**
>
> If your installation directory is in the C:\Program Files path, you must run
> as administrator when you license the product. This avoids any permission issues.

> **Tip:**
>
> Run the `envcomp` script before starting. This ensures that the installation
> environment for Genero BDL is set.

**Which procedure should I follow?**

To install your
license, you must have either a license file, a license string, or your license number, license key,
and customer login.

If you are not
provided with a license string, you can generate a license string from your license number, license
key, maintenance/subscription key, and login.

You can
install the license using one of these methods:

- Using a license file
- Using a license string
- Using a license number and key

## Install using a license file

For this task you need a license file.

1. Start the command line interface.

   Open a command prompt.

   - On Linux®/UNIX®/macOS™, open a command prompt. "sudo" may be required.
   - On Windows®, open the Command Prompt
     from the
     Start menu .
2. Install license for your
   Genero product
   . 

   Type the command `fglWrt` `-f license-file`

   The license is installed in temporary mode.
   > **Important:**
   >
   > The license number and an installation number is displayed in the output. The installation number
   > is needed for activating the license.
3. Go to the License your products page on the [Four Js web site](https://4js.com/support/registration/)
   and follow the procedure described in [Register your license](0466-register-license.md "To validate your license, you must register it on the Four Js website. You can access the website from any machine or smart phone.").

   With your license registered, you have the required keys.
4. Activate the license.
   1. Run the following command:

      ```
      fglWrt -k lnum
      ```

      Replace lnum with your license
      number.
   2. At the prompt, enter the installation key.

      ```
      Enter the installation KEY (call your vendor to obtain it)
      ```

   A message is displayed in the output to say the license installation was successful.

   License installation is now completed.

## Install using a license string

For this task you need a license string. This command activates the license with Four Js over the
internet.

1. Start the command line interface.

   Open a command prompt.

   - On Linux/UNIX/macOS, open a command prompt. "sudo" may be required.
   - On Windows, open the Command Prompt
     from the
     Start menu .
2. Install license for your
   Genero product
   . 

   Type the command `fglWrt` `--install-license-string license-string
   auto`

   The installation of the license is complete.

## Install using a license number and key

This process validates your license with Four Js over the internet using your license number,
license key, and customer login.

1. Start the command line interface.

   Open a command prompt.

   - On Linux/UNIX/macOS, open a command prompt. "sudo" may be required.
   - On Windows, open the Command Prompt
     from the
     Start menu .
2. Type fglWrt -l and when prompted, enter the license number and the license
   key.
3. At the prompt: 

   ```
   Do you want to continue using HTTP ? (y/n)
   ```

   Enter:
   `y`
4. When prompted, enter your customer code:

   ```
   Enter your customer code >
   ```
5. At the following prompt: 

   ```
   Do you need to configure an HTTP proxy ? (y/n)
   ```

   Enter
   `y` if your access to the internet is through a
   proxy. Provide the required information for HTTP proxy, port, and
   authentication when prompted. Otherwise, enter `n`.
6. At the command line enter the command to install the maintenance/subscription key:

   ```
   fglWrt -m mkey
   ```

   Where mkey is the maintenance/subscription key to enter.

   A message is displayed in the output to say the license installation was successful.

   License installation is now completed.
