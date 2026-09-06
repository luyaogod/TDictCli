---
title: "Dameng"
source: "fgl-topics/c_fgl_Connections_037.html"
breadcrumb: "SQL support > Database connections > Database client environment > Dameng"
type: "concept"
description: "Make sure the PATH environment variable is set the path to find Dameng® client program tools. On UNIX™, LD_LIBRARY_PATH (or equivalent) must hold the path to the client libraries libdmdpi.so . Check ..."
---

# Dameng

1. Make sure the PATH environment variable is set the path to find Dameng® client program tools.
2. On UNIX™, LD\_LIBRARY\_PATH (or equivalent) must hold the
   path to the client libraries `libdmdpi.so`.
3. Check the Dameng client configuration parameters in the dm\_svc.conf
   configuration file. Verify the value of the DM\_SVC\_PATH environment variable.
4. The Dameng database client character set is defined by the `PAGE_MODE` parameter
   in the dm\_svc.conf configuration file.
5. You can make a connection test with the Dameng disql command line
   tool:

   ```
   disql sysdba/@[::]:5237
   password:
   ```

## Related links

**Related concepts**  

[Database client settings](../09_advanced-features/0885-database-client-settings.md "This section describes the settings defining the locale for the database client.")

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1218-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
