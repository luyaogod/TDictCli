---
title: "File system"
source: "fgl-topics/c_fgl_gwa_filesystem.html"
breadcrumb: "Genero Web applications > File system"
type: "concept"
---

# File system

> When the GWA application is launched in the browser, a virtual UNIX-like file system emulation of your GWA application is created in memory.

By default, the following directories are created when a GWA starts:

- /app directory contains the .42m modules and program
  resources (with possible subdirectories). The best way to access a resource in your application,
  regardless of mobile or desktop, is using `base.Application.getProgramDir()`.
  For
  example:

  ```
  VAR file_in_app_dir=os.Path.join(base.Application.getProgramDir(),"main.42")
  DISPLAY os.Path.exists(file_in_app_dir)
  ```
- /fgl contains the Genero runtime assets.
- /tmp is for temp files.
- /home is the default working directory (pwd) where applications start. Each
  application has its own unique subdirectory in /home,
  /home/build\_directory\_basename – the
  build\_directory\_basename corresponds to the name of the directory you specified
  when building the application.

  The
  /home/build\_directory\_basename is also the "persistent"
  directory of the GWA because it stores data even if the application is restarted (or the browser is
  restarted). For more information about managing persistence in the GWA, go to [Manage persistence in the file system](5125-file-system.md "As directories in a GWA application are in memory, this means the lifetime of those directories is the same as the lifetime of the application. Understanding the file system will help you develop a strategy for making data persistent.").

These are the default directories; other directories can be created with the
`os.Path.mkdir` function as needed.

## Related links

**Related concepts**  

[File transfer support in GWA deployment](5147-file-transfer-support-in-gwa-deployment.md "File transfer operations are supported in Genero Web Application (GWA) deployments using fgl_putfile and fgl_getfile.")

## Manage persistence in the file system

As directories in a GWA application are in memory, this means the lifetime of those
directories is the same as the lifetime of the application. Understanding the file system will help
you develop a strategy for making data persistent.

### Bundled and runtime directories

At startup, GWA pre-populates the browser file system with the files you bundle at build time. The following table summarizes the directories and their persistence characteristics.

| Directory | Contents | Persistent |
| --- | --- | --- |
| app | Stores all files bundled by `gwabuildtool`, such as compiled modules (`.42m`), form files (`.42f`), icons, images, web components, and other build-time assets. | No |
| home/build\_directory\_basename | Holds runtime data and user-generated files for the application. See the [Home directory](5125-file-system.md) section for details. | Yes |

### Home directory

Each application has its own home directory in the browser file system at home/build\_directory\_basename. The value of build\_directory\_basename is the directory name you specified when building the application.

For example, if you run `gwabuildtool --program-dir=/myGwas/test1`, the build
directory name is test1, so the application's home directory becomes
home/test1. Because the browser stores this directory in [IndexedDB](https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API) (external link), files placed in this directory persist across sessions and
survive browser restarts.

The home directory provides isolation between applications:

- Each GWA receives its own persistent directory.
- Program A cannot access Program B’s data.
- This avoids conflicts and maintains data integrity.

By default, the initial working directory for the application is home/build\_directory\_basename. If you use the [gwabuildtool](5150-gwabuildtool.md) `--app-dir-is-pwd` option, the initial working directory becomes app, which is temporary and stored in memory only.

> **Note:**
>
> If you use the `--app-dir-is-pwd` option, any files saved in app are lost when the application or browser closes. To make data persistent, copy files into your home directory home/build\_directory\_basename using operations such as [os.Path.copy](../15_library-reference/3705-os-path-copy.md "Creates a new file by copying an existing file."). See the [Database persistence](5125-file-system.md) section for an example.

### Database persistence

Database persistence is handled through an SQLite database file stored in home/build\_directory\_basename. Files in this directory persist across sessions.

You may need a database in one of the following cases:

- You create a new SQLite database at application startup.
- Your application includes a predefined database that must be copied into the persistent directory the first time the application runs.

To create a new SQLite database at startup, write `4GL` code to generate the file before calling the `CONNECT` instruction.

If your application ships with an initial database file, bundle it into the
app directory at build time. At runtime, check whether the database already
exists in home/build\_directory\_basename. If it does not
exist, copy the bundled version from app into the persistent directory, as
shown below:

```
# check for database
VAR home_db="/home/mygwa/mydatabase.db"
#...
IF NOT os.Path.exists(home_db) THEN --Copy to home directory first time
   --the initial mydatabase.db is bundled by gwabuildtool into /app
  IF NOT os.Path.copy("/app/mydatabase.db", home_db) THEN
    DISPLAY "Can't copy initial db"
    EXIT PROGRAM 1
  END IF
END IF
# connect to database
```

### Accidental clearing of the cache

Explicitly cleaning the browser cache for the site hosting a GWA application will remove all
stored data of a GWA application.

As there is always the danger that an end user may clear browser cache and thus delete the
contents of a GWA application home/build\_directory\_basename
directory unintentionally, your GWA application should be prepared to send the data to a REST server
via REST API, or synchronise data with the server by other combinations of GWS calls.
