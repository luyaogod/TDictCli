---
title: "Steps to BDL license installation"
source: "genero-install-topics/c_bdl_steps_to_license_success.html"
breadcrumb: "Licensing > Steps to BDL license installation"
type: "concept"
---

# Steps to BDL license installation

> Preparing to license involves evaluating your options based on what licensing tool (GUI or command line) is available to you, and whether you can validate the installed license with Four Js over the internet. Choose the option that is right for you.

Before you begin Genero product licensing:

1. Locate your email message that contains the following:
   - license file (with the product license string)
   - product license string
   - product license key
   - product license number
   - product maintenance/subscription key
   - your customer codeIf you do not have this email, contact your local Four Js sales office.

   | Method | Best used when | Internet requirement | License registration |
| --- | --- | --- | --- |
| License file | Internet is unavailableManual registration is acceptableSingle-step installActivation is not forced | Not required | Manual |
| License string | Internet is availablePrefer automationSingle-step install and registration | Required for automatic registration | Manual or automatic (via `auto` option) |
| License key + number + login | Internet is availableMulti-step process | Required for automatic registration | Manual or automatic (via HTTP prompt) |
2. Decide which tool (GUI or command line) you are using to install the license:
   - Genero BDL Licenser (GUI application)
   - fglWrt (command line tool).
3. Determine if you have internet access from the machine you are licensing the product.
4. Use Table 2 to select which procedure to
   follow.

| You are using ... | Then ... |
| --- | --- |
| Genero BDL Licenser with internet | See [License BDL using GUI and internet](0440-bdl-licenser-internet.md "You can license a Genero Business Development Language (BDL) product using the Genero licensing graphical interface."). |
| Genero BDL Licenser without internet. | See [License BDL using GUI without internet](0441-bdl-licenser-no-internet.md "Without internet access, you register the license with Four Js from another machine via the internet, or by phone. Then use the user interface to license your Genero Business Development Language (BDL) product."). |
| fglWrt command line tool with internet. | See [License BDL from the command line (via internet)](0442-fglwrt-command-internet.md "You can license a Genero Business Development Language (BDL) product using the command line tool, fglWrt."). |
| fglWrt command line tool without internet. | See [License BDL from the command line (no internet)](0443-fglwrt-command-no-internet.md "Without internet access, you register the license with Four Js from another machine via the internet, or by phone. Then use the fglWrt command line tool to license your Genero Business Development Language (BDL) product."). |

## Related links

**Related tasks**  

[Uninstall (BDL license)](0454-uninstall-license.md "You may need to uninstall a license if you are making some hardware or software changes.")
