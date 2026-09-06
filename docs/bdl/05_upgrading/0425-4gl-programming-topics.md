---
title: "4GL Programming topics"
source: "fgl-topics/c_fgl_Mig0000_024.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > 4GL Programming topics"
type: "concept"
---

# 4GL Programming topics

> When migrating from Four Js BDS to Genero BDL, review the programming differences between the two products. Reviewing the differences allows you to plan and prepare for a smooth migration.


## Child topics

- [FGLPROFILE: VM configuration](0426-fglprofile-vm-configuration.md): Identify the Four Js BDS FGLPROFILE virtual machine configuration entries that are no longer supported, and the Genero BDL equivalent (where relevant).
- [The fgl_init4gl() function](0427-the-fgl-init4gl-function.md): The fgl_init4js() function has no effect in Genero BDL.
- [Static versus Dynamic Arrays](0428-static-versus-dynamic-arrays.md): Support for dynamic arrays in Genero BDL may result in a need for some redesigns in your application.
- [Debugger syntax changed](0429-debugger-syntax-changed.md): While Four Js BDS and Genero BDL both provide a program debugger, the commands used and how it is used can differ.
- [fgl_system() function](0430-fgl-system-function.md): With Genero BDL, the fgl_system() function no longer raises a terminal window by default, but some front-ends offer a workaround.
- [The Channel:: methods](0431-the-channel-methods.md): Review your code and replace Channel:: calls with the new base.Channel API.
- [The Dialog:: methods](0432-the-dialog-methods.md): Review your code and replace Dialog:: calls with the new ui.Dialog API.
- [STRING versus CHAR/VARCHAR](0433-string-versus-char-varchar.md): Genero BDL supports the STRING data type in addition to CHAR and VARCHAR. While the STRING data type is useful in certain situations, there are times when you should use CHAR or VARCHAR instead.
- [Review user-made C routines](0434-review-user-made-c-routines.md): Genero BDL provides libraries which may replace some of the C routines required by I4GL applications.
- [Variable identification in SQL statements](0435-variable-identification-in-sql-statements.md): Program variable identification in static SQL statements is more strict in version 2.20 than older versions.
- [Default action of WHENEVER ANY ERROR](0436-default-action-of-whenever-any-error.md): By default, the WHENEVER ANY ERROR action is to CONTINUE the program flow.
- [Database driver features](0437-database-driver-features.md): Some ODI features are no longer supported in Genero BDL.
