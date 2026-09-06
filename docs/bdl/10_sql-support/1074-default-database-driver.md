---
title: "Default database driver"
source: "fgl-topics/c_fgl_Connections_005.html"
breadcrumb: "SQL support > Database connections > Connection parameters > Default database driver"
type: "concept"
description: "The dbi.default.driver FGLPROFILE entry defines a default database driver to be loaded, if the driver is not specified by the connection parameters. dbi.default.driver = \" driver-name \" The driver ..."
---

# Default database driver

The `dbi.default.driver` [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
entry defines a default database driver to be loaded, if the driver is not specified by the
connection parameters.

```
dbi.default.driver = "driver-name"
```

The driver name must be specified without the .so or .DLL extension.

If this configuration entry is not defined, the driver name defaults to
`dbmdefault`.

## Related links

**Related concepts**  

[Database driver specification (driver)](1073-database-driver-specification-driver.md "Database driver specification (driver)")
