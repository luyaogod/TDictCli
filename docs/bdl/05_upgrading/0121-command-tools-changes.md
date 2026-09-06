---
title: "Command tools changes"
source: "fgl-topics/c_fgl_Migrate_to_500_fgl_tools.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.00 upgrade guide > Command tools changes"
type: "concept"
---

# Command tools changes

> Modifications to consider regarding command line tools.

This topic describes changes that may need code review.

See also [new 5.00 features of commands](0058-bdl-5-00-new-features.md).

## Changes to fglgar

Starting with version 5.00, changes have been made to the fglgar command
features for creating Java Web Archive (war) files and running applications.

- When creating a war file, you must now use the
  `--web-content` option of the fglgar war  command to specify the
  path to the Java Web Content directory of the JGAS installation. As the JGAS is no longer embedded
  in the FGLGWS package, the `--web-content` option is now mandatory.
- The embedded Jetty web server and the standalone JGAS is removed from the JGAS installation.
  Where before you could execute applications deployed in a war  file with the
  standalone JGAS using the fglgar run command, this feature is no longer available
  and therefore the fglgar run  has been removed.

For details of other changes, such as creating war files for testing and deployment, review the
[Packaging web applications](../13_programming-tools/2645-packaging-web-applications.md "Describes methods of packaging the runtime files and resources of your web applications and services using the fglgar tool.") and [fglgar](../13_programming-tools/2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).") pages. For more information on JGAS, see the Genero Application Server for Java User Guide

## Changes to fgldbsch

Starting with version 5.01.00, the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool will by default deny SQL table or column names that are not
valid identifiers (like `"Customer Name"`), or when a table/column name is a
case-insensitive duplicate of other table/column names (like `"Stock"` vs
`"STOCK"`).

To bypass this control and allow invalid table or column names in the .sch
file, use the `-sl` (sloppy) option. This option has been implemented in case if an
SQL table mixes valid and invalid column names, and `DEFINE varname LIKE
tabname.colname` instructions are used with the
columns using regular identifier names.

To skip problematic tables and extract valid table defintions, use the already-existing
`-ie` (ignore errors) option.

Use the `-v` verbose option to get details of the database schema extraction
process.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases.

## Related links

**Related concepts**  

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
