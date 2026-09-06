---
title: "FGLPROFILE entries for core language"
source: "fgl-topics/c_fgl_fglprofile_004.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > FGLPROFILE entries for core language"
type: "concept"
---

# FGLPROFILE entries for core language

> This is a summary of FGLPROFILE entries supported by the core BDL.

Find more information for an FGLPROFILE entry by following the documentation link in the
description of the entry.

This topic describes FGLPROFILE entries for the BDL core language. Web services specific
FGLPROFILE entries description can be found in [FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.").

| Entry | Values | Default | Description |
| --- | --- | --- | --- |
| `Dialog.currentRowVisibleAfterSort` | boolean | false | Forces current row to be shown after a sort in a table.See [Dialog configuration with FGLPROFILE](../11_user-interface/2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior."). |
| `Dialog.fieldOrder` | boolean | false | Defines if the intermediate field triggers must be executed when a new field gets the focus with a mouse click.See [Dialog configuration with FGLPROFILE](../11_user-interface/2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior."). |
| `dbi.default.driver` | string | NULL | Defines the default database driver.See [Default database driver](../10_sql-support/1074-default-database-driver.md). |
| `dbi.database.dbname.driver` | string | NULL | Defines the database driver for a database name.See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md). |
| `dbi.database.dbname.source` | string | NULL | Defines the data source for a database name.See [Database source specification (source)](../10_sql-support/1072-database-source-specification-source.md). |
| `dbi.*` | N/A | N/A | Database interface configuration.See [Connections](../10_sql-support/1059-database-connections.md "Explains how to manage database connections in a program."). |
| `fglrun.arrayIgnoreRangeError` | boolean | false | Controls runtime behavior when array index is out of bounds.See [Arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") for more details. |
| `fglrun.decToCharScale2` | boolean | false | Formats `DECIMAL(P)` with 2 decimal digits globally. See [Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types."). |
| `fglrun.decToCharScale2.print` | boolean | false | Formats `DECIMAL(P)` with 2 decimal digits in `PRINT` statements. See [Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types."). |
| `fglrun.floatToCharScale2` | boolean | false | Formats `FLOAT/SMALLFLOAT` with 2 decimal digits globally. See [Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types."). |
| `fglrun.floatToCharScale2.print` | boolean | false | Formats `FLOAT/SMALLFLOAT` with 2 decimal digits in `PRINT` statements. See [Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types."). |
| `fglrun.defaults` | string | NULL | Defines the directory where program specific configuration files are located.See [Understanding FGLPROFILE](0484-understanding-fglprofile.md "The runtime system uses one or more configuration files in which you can define options and parameters to change the behavior of the programs."). |
| `fglrun.ignoreDebuggerEvent` | boolean | false | Defines whether the runtime system can swtich to debug mode.See [Integrated debugger](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs."). |
| `fglrun.ignoreLogoffEvent` | boolean | false | Defines whether the runtime system ignores a CTRL\_LOGOFF\_EVENT on Windows® platforms.See [Responding to CTRL\_LOGOFF\_EVENT](../09_advanced-features/0936-responding-to-ctrl-logoff-event.md "FGLPROFILE fglrun.ignoreLogoffEvent controls program behavior in case of logoff events on Windows platforms."). |
| `fglrun.localization.*` | N/A | N/A | Defines load parameters for localized string resource files.See [Localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site."). |
| `fglrun.mapAnyErrorToError` | boolean | false | Controls default action of `WHENEVER ANY ERROR`.See [Default exception handling](../09_advanced-features/0855-default-exception-handling.md "Default exception handling must be adapted to your programming pattern."). |
| `fglrun.mmapDisable`**Important:**This feature is deprecated, its use is discouraged although not prohibited. | boolean | false | Turns program files memory mapping off on Windows platforms.This entry is only provided to solve file overwrite issues when doing live program files updates on Windows platforms.See [Dynamic module loading](../09_advanced-features/0968-dynamic-module-loading.md). |
| `flm.*` | N/A | N/A | License management related entries.See licensing documentation. |
| `gui.connection.timeout` | integer | 30 | Defines the timeout delay (in seconds) the runtime system waits when it establishes a connection to the front-end. After this delay the program stops with an error.See [GUI connection timeout](../11_user-interface/1524-gui-connection-timeout.md). |
| `gui.key.add_function` | integer | none | If set, this entry defines the offset for function key mapping when using Shift-Fx and Control-Fx key modifiers.See [Graphical mode with Traditional Display](../11_user-interface/1519-graphical-mode-with-traditional-display.md). |
| `gui.programStoppedMessage` | string | none | Generic message to be displayed to the end user when a program stops because of a runtime error.When not specified, fglrun displays the detailed error message.See [Default exception handling](../09_advanced-features/0855-default-exception-handling.md "Default exception handling must be adapted to your programming pattern."). |
| `gui.protocol.pingTimeout` | integer | 600 | Defines the timeout delay (in seconds) the runtime system waits for a front-end ping when there is no user activity. After this delay the program stops with an error.See [Wait for front-end ping timeout](../11_user-interface/1525-wait-for-front-end-ping-timeout.md). |
| `gui.uiMode` | string | NULL | Defines the user interface mode, to render windows in traditional I4GL mode.Possible values are: `"default"` or `"traditional"`.Default is the new Genero GUI mode with real resizable windows.See [Graphical mode with Traditional Display](../11_user-interface/1519-graphical-mode-with-traditional-display.md). |
| `key.key-name.text`**Important:**This feature is deprecated, its use is discouraged although not prohibited. | string | N/A | Defines a label for an action defined with an ON KEY clause.This entry is provided for backward compatibility with BDS V3.See [Setting action key labels](../11_user-interface/2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers."). |
| `mobile.environment.name = "value"` | N/A | N/A | Define environment variable values in FGLPROFILE for mobile applications.See [Setting environment variables in FGLPROFILE (mobile)](0495-setting-environment-variables-in-fglprofile-mobile.md). |
| `Report.aggregateZero`**Important:**This feature is deprecated, its use is discouraged although not prohibited. | boolean | false | Defines if the report aggregate functions must return zero or NULL when all values are NULL.This entry is provided for backward compatibility with BDS V3.See [Report engine configuration](../12_reports/2508-report-engine-configuration.md "Report engine behavior can be controlled with FGLPROFILE settings."). |

## Related links

**Related concepts**  

[FGLPROFILE entries for web services](../16_web-services/4916-fglprofile-entries-for-web-services.md "The FGLPROFILE entries relating to Genero Web Services are divided between five categories: security, basic or digest HTTP authentication, proxy configuration, web server configuration, and XML cryptography.")
