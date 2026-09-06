---
title: "Runner creation is no longer needed"
source: "fgl-topics/c_fgl_Migrate_to_200_dynamic_runner.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Runner creation is no longer needed"
type: "concept"
---

# Runner creation is no longer needed

> Starting with version 2.00, you no longer need to recompile/build a runner.

The runtime system architecture is now based on shared libraries (or DLLs on Windows®), and the database drivers are automatically loaded
according to [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") configuration parameters.

If you have C extensions, you must rebuild them as shared libraries.

> **Important:**
>
> Database vendor client libraries must be provided as shared objects
> (or DLL on Windows).

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
