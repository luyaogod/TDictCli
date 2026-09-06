---
title: "License an upgraded installation"
source: "genero-install-topics/c_bdl_license_upgrade_install.html"
breadcrumb: "Licensing > Manage Genero BDL local license > License an upgraded installation"
type: "concept"
---

# License an upgraded installation

> There are some considerations for licensing when upgrading an existing Genero installation to a newer version.

## Notes about upgrading

It is not necessary to re-enter the license of the product if the new version:

- Is installed into an existing installation directory
- Is not a major version number change

Alternatively, if you install the new product version in a directory separate to your existing
product installation, you can import the license from your existing product installation using your
preferred licensing method's `import-license` command: license controller
(fglWrt or greWrt) or License Manager
(flmprg). Importing a license is an alternative to uninstalling and reinstalling
a license. An additional advantage is that you do not need to reactivate the license by registering
it with Four Js.

## Maintenance contract and subscription license renewal

When you purchase a new maintenance contract or renew a
subscription license, a new maintenance/subscription key must be installed. To install your new key,
you do not need to reinstall your license. Updating the maintenance/subscription key is handled
separately. There is also no need to uninstall the old maintenance/subscription key.

## Related links

**Related tasks**  

[Uninstall (BDL license)](0454-uninstall-license.md "You may need to uninstall a license if you are making some hardware or software changes.")

**Related reference**  

[Steps to BDL license installation](0439-steps-to-bdl-license-installation.md "Preparing to license involves evaluating your options based on what licensing tool (GUI or command line) is available to you, and whether you can validate the installed license with Four Js over the internet. Choose the option that is right for you.")
