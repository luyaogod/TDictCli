---
title: "FGLPROFILE"
source: "fgl-topics/c_fgl_EnvVariables_FGLPROFILE.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLPROFILE"
type: "concept"
---

# FGLPROFILE

> Defines the configuration files to be used by the runtime system.

## Usage

The FGLPROFILE environment variable defines a list of configuration files
to be used by the runtime system.

The runtime system always reads entries from the default configuration file located in [$FGLDIR/etc/fglprofile](0523-fgldir.md "Defines the installation directory of Genero Business Development Language."), then the
file(s) defined in the FGLPROFILE environment variable.

FGLPROFILE can define one unique configuration file, or a list of files to
be loaded sequentially.

FGLPROFILE must contain a list of file paths, separated by the operating system specific path
separator. The path separator is **":"** on UNIX™ platforms
and **";"** on Windows® platforms.

On mobile devices,
you must deploy a file with the name "`fglprofile`" in the appdir
directory. See [FGLPROFILE for mobile apps](0490-fglprofile-for-mobile-apps.md "The name of the FGLPROFILE file matters for mobile applications.") for more details.

For more details, see also [The FGLPROFILE file(s)](0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files").

## Example

On UNIX platforms:

```
$ FGLPROFILE="/opt/myapp/gui_settings.prf:/opt/myapp/sqldb_settings.prf"
$ export FGLPROFILE
```

On Windows platforms:

```
C:\> set FGLPROFILE=C:\myapp\gui_settings.prf;C:\myapp\sqldb_settings.prf
```
