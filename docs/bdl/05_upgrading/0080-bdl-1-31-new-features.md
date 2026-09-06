---
title: "BDL 1.31 new features"
source: "fgl-topics/fgl_whatsnew_131.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 1.31 new features"
type: "topic"
---

# BDL 1.31 new features

> Features added in 1.31 release of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 1.31 upgrade guide](0330-bdl-1-31-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 1.31.").

Prior new features guide: [BDL 1.30 new features](0081-bdl-1-30-new-features.md "Features added in 1.30 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| C extensions can be loaded dynamically, no need to re-link runner. | See [C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code."). |
| The `FGL_WIDTH()` built-in function computes the number of print columns needed to represent a single or multibyte character. | See [fgl\_width()](../15_library-reference/2785-fgl-width.md "Returns the number of columns needed to represent the printed version of the expression."). |

| Overview | Reference |
| --- | --- |
| GUI protocol compression for slow networks. | This feature is desupported, see [Presentation styles changes in V4.00](0143-presentation-styles-changes.md "Modifications to consider when using presentation styles."). |
| Interruption handling with SSH port forwarding - only supported with GDC 1.31! | See [User interruption handling](../11_user-interface/2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure."). |
| New method `ui.Form.setFieldStyle()` to set a style for a field. | See [ui.Form.setFieldStyle](../15_library-reference/3162-ui-form-setfieldstyle.md "Change the style of a form field."). |
| Improved front-end identification when connecting to GUI client. | See [GUI front-end connection](../11_user-interface/1520-gui-front-end-connection.md "This section explains runtime to front-end connection in its simplest form."). |

| Overview | Reference |
| --- | --- |
| MySQL version 4.1.x is now supported (3.23 is desupported) | See [Oracle MySQL / MariaDB](../10_sql-support/1314-oracle-mysql-mariadb.md). |
| Oracle version 10g is now supported. | See [Oracle Database](../10_sql-support/1359-oracle-database.md). |
