---
title: "Desupported environment variables"
source: "fgl-topics/c_fgl_Mig0000_008.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > Installation and setup topics > Desupported environment variables"
type: "concept"
---

# Desupported environment variables

> The Four Js Business Development Suite (BDS) environment variables no longer supported (or replaced) in Genero Business Development Language.

| Entry | Description of the BDS environment variable | Genero equivalent |
| --- | --- | --- |
| `FGLDBS` | FGLDBS defines the type and version of the database driver, used when linking fglrun with fglmkrun. | Database drivers are loaded dynamically by fglrun. |
| `FGLCC` | FGLCC defines the name of the C compiler. | The fglrun tool does not need to be created, it's fully dynamic. |
| `FGLLIBSQL` | FGLLIBSQL defines the list of database client software libraries to be used to link fglrun with fglmkrun. | Database drivers are loaded dynamically by fglrun. |
| `FGLLIBSYS` | FGLLIBSYS defines the list of system libraries to be used to link fglrun with fglmkrun. | The fglrun tool does not need to be created, it's fully dynamic. |
| `FGLSHELL` | FGLSHELL defined the name of the fglrun program, for example when using tools like fglschema. | The name of the runtime system tool is fglrun and does not need to be changed. |
