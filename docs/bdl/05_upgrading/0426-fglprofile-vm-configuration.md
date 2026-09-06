---
title: "FGLPROFILE: VM configuration"
source: "fgl-topics/c_fgl_Mig0000_025.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > 4GL Programming topics > FGLPROFILE: VM configuration"
type: "concept"
---

# FGLPROFILE: VM configuration

> Identify the Four Js BDS FGLPROFILE virtual machine configuration entries that are no longer supported, and the Genero BDL equivalent (where relevant).

Genero Business Development Language (BDL) comes with redesigned
software components and features. Some Four Js Business Development
Suite (BDS) specific FGLPROFILE entries have been desupported.
This section describes what configurations settings are no longer
supported, and point to Genero equivalent features if they exist.

This table shows BDS FGLPROFILE entries related to runtime system
configuration which are desupported in Genero. See the FGLPROFILE
description page for supported entries:

| Entry | Description of the BDS feature | Genero BDL equivalent |
| --- | --- | --- |
| `fglrun.checkDecimalPrecision` | Controls decimal variable assignment when overflow occurs. For example, a value of 1000.0 does not fit in a DECIMAL(2,0).Is false by default = no overflow error, value assigned. | *There is no equivalent in Genero.**By default Genero assigns NULL to a decimal when overflow occurs. Can be trapped by WHENEVER ANY ERROR.* |
| `fglrun.ix6` | Controls Informix® version 6.x compatibility.By default BDS is compatible with I4GL 4.x | *There is no equivalent in Genero.**By default Genero is compatible to Informix 4gl 7.32.* |
| `fglrun.cmd.winnt`, `fglrun.cmd.win95` | Defines the command line to be executed for a RUN WITHOUT WAITING on Windows® platforms. | With Genero the command program can be defined with the COMSPEC environment variable. |
| `fglrun.database.listvar`, `fglrun.remote.envvar` | These entries were used by Informix driver to set environment variables with the ifx\_putenv() function on Windows platforms. | *There is no equivalent in Genero.* |
| `fglrun.setenv.*`, `fglrun.defaultenv.*` | These entries define environment variables for all programs. | *There is no equivalent in Genero.* |
| `fgllic.*` | License controller related entries | With Genero you configure license settings with the `flm.*` entries.See license manager documentation for more details. |
| `fglrun.server.*` | These entries define X11 front-end automatic startup. | GUI server autostart is desupported. |

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
