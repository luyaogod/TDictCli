---
title: "License BDL from the command line (no internet)"
source: "genero-install-topics/t_bdl_license_fglwrt_without_internet.html"
breadcrumb: "Licensing > Steps to BDL license installation > fglWrt command (no internet)"
type: "task"
---

# License BDL from the command line (no internet)

> Without internet access, you register the license with Four Js from another machine via the internet, or by phone. Then use the fglWrt command line tool to license your Genero Business Development Language (BDL) product.

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
- [Using a license number and key](0443-fglwrt-command-no-internet.md)

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
   > is needed for registering the license.
3. From a machine that has internet access or a smart phone, go to the
   License your products page on the [Four Js web site](https://4js.com/support/registration/) and follow the procedure described in [Register your license](0466-register-license.md "To validate your license, you must register it on the Four Js website. You can access the website from any machine or smart phone.").

   With your license registered, you have the required keys.
4. Activate the license: 
   1. Type `fglWrt -k lnum`

      Where lnum is the license number to enter.
   2. At the prompt: 

      ```
      Enter the installation KEY (call your vendor to obtain it)
      ```

      Enter the
      installation key.

   License installation is now completed.

## Install using a license string

For this task you need a license string.

1. Start the command line interface.

   Open a command prompt.

   - On Linux/UNIX/macOS, open a command prompt. "sudo" may be required.
   - On Windows, open the Command Prompt
     from the
     Start menu .
2. Install license using license string. 

   Type the command `fglWrt --install-license-string license-string`
   > **Important:**
   >
   > If you don't have a license string, you can generate one from your license number, license key,
   > maintenance/subscription key, and login.

   The license is installed in temporary mode.
   > **Important:**
   >
   > The license number and an installation number is displayed in the output. The installation number
   > is needed for registering the license.
3. From a machine that has internet access or a smart phone, go to the
   License your products page on the [Four Js web site](https://4js.com/support/registration/) and follow the procedure described in [Register your license](0466-register-license.md "To validate your license, you must register it on the Four Js website. You can access the website from any machine or smart phone.").

   With your license registered, you have the required keys.
4. Activate the license: 
   1. Type `fglWrt -k lnum`

      Where lnum is the license number to enter.
   2. At the prompt: 

      ```
      Enter the installation KEY (call your vendor to obtain it)
      ```

      Enter
      the installation key.

   License installation is now completed.

## Install using a license number and key

For this task you need your license number, license key, and customer login.

1. Start the command line interface.

   Open a command prompt.

   - On Linux/UNIX/macOS, open a command prompt. "sudo" may be required.
   - On Windows, open the Command Prompt
     from the
     Start menu .
2. Type fglWrt -l and when prompted, enter the license number and the license
   key.

   An **installation number** is generated and is displayed in the output.
3. At the prompt: 

   ```
   Do you want to continue using HTTP ? (y/n)
   ```

   Enter:
   `n`
4. At the prompt:

   ```
   Enter your maintenance KEY (Empty to continue) >
   ```

   Enter the
   maintenance/subscription key.

**Register the license**

5. From a machine that has internet access or a smart phone, go to the
   License your products page on the [Four Js web site](https://4js.com/support/registration/) and follow the procedure described in [Register your license](0466-register-license.md "To validate your license, you must register it on the Four Js website. You can access the website from any machine or smart phone.").

   With your license registered, you have the required keys.
6. Activate the license.
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

   License installation is now completed.
