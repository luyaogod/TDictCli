---
title: "Desupported FGLPROFILE entries"
source: "fgl-topics/c_fgl_Mig0000_009.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > Installation and setup topics > Desupported FGLPROFILE entries"
type: "concept"
---

# Desupported FGLPROFILE entries

> Identify the Four Js BDS FGLPROFILE configuration entries that are no longer supported, and the Genero BDL equivalent (where relevant).

Genero Business Development Language comes with redesigned software components and features. Some
FGLPROFILE entries have been desupported. This table describes what configurations settings are no
longer supported, and point to Genero equivalent features if they exist.

| Entry | Description of the BDS FGLPROFILE entry | Genero equivalent |
| --- | --- | --- |
| `dbi.sqldirset.*` | SQL directive set definition. | No equivalent option. See [SQL directive sets](0320-sql-directive-set-removed.md "The SQL directive set specification has been removed."). |
| `dbi.sql.directives` | SQL directive set parser de-activation. | N/A. See [SQL directive sets](0320-sql-directive-set-removed.md "The SQL directive set specification has been removed."). |
| `dbi.sql.static.optimization.cache.size` | Static SQL cache size. | No equivalent option. See [Static SQL cache](0319-static-sql-cache-is-removed.md "The Static SQL Cache has been removed."). |
| `dbi.database.dbname.connector` | Database connectors. | [`dbi.database.dbname.driver`](../10_sql-support/1073-database-driver-specification-driver.md) |
| `dbi.connector.cname.*` | Database connector definition. | No equivalent option. |
| `fglrun.cursor.global` | Boolean to defines if SQL cursors have global scope or not. | No equivalent option. |
| `fglrun.ix6` | Boolean to control the Informix 4GL version 6 and higher behavior. Default is false (Informix 4GL 4.10 behavior) | No equivalent option: Genero implements Informix 4GL 4.10 behavior. |
| `fglrun.signalOOB` | Defines the Out Of Band signal. When pressing the INTERRUPT key in GUI mode, OOB data is sent to the application over the socket. Some UNIX systems use a specific identifier for OOB signals that you can define with this entry. Default is 0. | No equivalent option. |
| `fglrun.cmd.winnt` | On Windows NT and 2000, this entry defines the program to execute 4GL applications started by RUN WITHOUT WAITING. Default is "cmd /c". | N/A: Obsolete platforms. |
| `fglrun.cmd.win95` | On Windows 95 and XP, this entry defines the program to execute 4GL applications started by RUN WITHOUT WAITING. Default is "start /m". | N/A: Obsolete platforms. |
| `fglrun.interface` | Defines the TCL/TK configuration file to be used by the graphical server. The file is read from FGLDIR/etc directory. Default is "fgl2c.res". | N/A: No more TCL/TK client. |
| `fglrun.scriptname` | Defines the TCL/TK client script file to be used by the graphical server. The file is read from FGLDIR/etc directory. Default is "fgl2c.tcl". | N/A: No more TCL/TK client. |
| `fglrun.server.*` | Parameters to automatically start the graphical server. | [`gui.server.autostart`](../11_user-interface/1530-automatic-front-end-startup.md) |
| `Menu.style` | Defines if the options of MENU instructions must be displayed in a separate control frame or in the MENU LINE. | No equivalent option. |
| `gui.screen.*` | Define the layout of the application container window. | No equivalent option. |
| `gui.workspaceframe.*` | Parameters for the work frame (the area where 4GL windows and forms are displayed) | No equivalent option. |
| `gui.controlframe.*` | Parameters for the control frame (right menu) | No equivalent option. |
| `gui.display.source.*` | Display redirections for instructions like ERROR and MESSAGE. | [Presentation styles for MESSAGE and ERROR](../11_user-interface/1644-message-style-attributes.md "Message presentation style attributes apply to an ERROR or MESSAGE instruction."). |
| `gui.statusbar.*` | Status bar configuration and rendering. | No equivalent option. |
| `gui.toolbar.*` | Toolbar configuration and rendering. | [Toolbar definitions](../11_user-interface/1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms."). |

## Related links

**Related concepts**  

[FGLPROFILE: VM configuration](0426-fglprofile-vm-configuration.md "Identify the Four Js BDS FGLPROFILE virtual machine configuration entries that are no longer supported, and the Genero BDL equivalent (where relevant).")

[FGLPROFILE: GUI configuration](0412-fglprofile-gui-configuration.md "Identify the Four Js BDS FGLPROFILE GUI configuration entries that are no longer supported, and the Genero BDL equivalent (where relevant).")
