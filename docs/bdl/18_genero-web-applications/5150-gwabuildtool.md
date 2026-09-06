---
title: "gwabuildtool"
source: "fgl-topics/c_fgl_tools_gwa_gwabuildtool.html"
breadcrumb: "Genero Web applications > GWA Reference > Tools and commands > gwabuildtool"
type: "concept"
---

# gwabuildtool

> The gwabuildtool is a utility to build a GWA application with all the necessary files to run on a browser.

## Syntax

```
gwabuildtool [ options ]
```

1. option can be an option described in Options.

## Options

| Option | Short option | Description |
| --- | --- | --- |
| `--program-dir` directory | `-p` | Directory of program assets such as compiled modules (.42m), compiled form files (.42f) icons, images, web components, and so on. |
| `--output-dir` directory | `-o` | Directory where the application is created.Default is "gwa\_dist" |
| `--main-module` module\_name | No short option | Name of the main module.The default is "main" |
| `--extra-asset` directory | `-x` | Add asset outside of program-dir |
| `--gbc` directory | No short option | Directory of the Genero Browser Client. |
| `--version` | `-V` | Displays version information. |
| `--help` | `-h` | Display options for the tool. |
| `--app-version` version\_number | `-a` | Application version number in the format: `?prefix?major.minor.build.suffix?` |
| `--webcomponent` directory | `-w` | Adds web component from outside the program-dir/webcomponents directory. |
| `--title` name | `-t` | Title for the web app |
| `--include` files `[,...]` | `-i` | Include custom extensions/files in the bundle. Multiple files may be specified separated by commas. For example, `gwabuildtool -i '*.per, Makefile'` |
| `--exclude` files `[,...]` | `-e` | Exclude custom extensions/files in the bundle. Multiple files may be specified separated by commas. For example, `gwabuildtool -e '*.txt,bigfile'` |
| `--verbose` | `-v` | Display verbose information. |
| `--W-excluded` | No short option | Warns about files in the program-dir not copied into the application bundle. |
| `--app-dir-is-pwd` | No short option | Sets the initial working directory of your GWA application to app in memory instead of the default (home/build\_directory\_basename). For a usage example and to see how it affects persistence, go to Set initial working directory to app (--app-dir-is-pwd) |
| `--qa-version` | No short option | Prints a simple version number. |

## Usage

To run gwabuildtool, you need to have it in your PATH (see [Install Genero Web Application](../04_installation/0053-install-genero-web-application.md "To build and package Genero Web applications, you must first install Genero Web Application (GWA).")).

The gwabuildtool does not compile .4gl and
.per assets, it is assumed that the compilation process has been done already.

For instance, if you have compiled your Genero application in a directory called
`myapp` with a module `main.42m` containing `MAIN`, run
the following command to create a GWA application:

```
gwabuildtool -p /usr/myapp
```

A directory called `gwa_dist` with all the necessary files is created in your
current path. This `gwa_dist` can be packaged using [fglgar gwa](../13_programming-tools/2526-fglgar.md) (see
[Packaging gwa files](../13_programming-tools/2650-packaging-gwa-files.md "Using the fglgar tool to build a Genero Web Application (gwa) file allows you to deploy applications that are ready to run in a browser.")) and uploaded to a web space and a web server can serve
the application at the URL
`https:servername/gwa_dist/index.html`.

You can change the output directory with option
`-o`.

```
gwabuildtool -p /usr/myapp -o mydist
```

This mydist directory can be uploaded to a web space and a web server can
serve the application at the URL
`https:servername/mydist/index.html`.

## Set initial working directory to app (`--app-dir-is-pwd`)

By default, your application runs with its initial working directory set to
home/build\_directory\_basename. Bundled files always reside
in `/app` in the browser, so relative paths like
`ch.openFile("foo.txt","r")` only work if the working directory is set to
`/app`. Use `--app-dir-is-pwd` when your program assumes assets are in
the current working directory — for example, the FGL demo launcher, which runs several other
programs and expects each one to access resources relative to the launcher’s working directory.

```
gwabuildtool --program-dir . --app-dir-is-pwd
```

This command changes the initial working directory to app, which is
temporary and stored in memory only.
> **Tip:**
>
> If you use the `--app-dir-is-pwd` option, any files you save in
> app will be lost when the application or browser closes. To keep data
> persistent, copy files into your home directory
> home/build\_directory\_basename using operations such as [os.Path.copy](../15_library-reference/3705-os-path-copy.md "Creates a new file by copying an existing file."). For details, go to [Manage persistence in the file system](5125-file-system.md "As directories in a GWA application are in memory, this means the lifetime of those directories is the same as the lifetime of the application. Understanding the file system will help you develop a strategy for making data persistent.").

## Related links

**Related concepts**  

[gwarun](5152-helper-tools.md "The gwarun tool runs a GWA program in the browser on your desktop.")

[gwasrv](5152-helper-tools.md "The gwasrv is a mini web server written in Genero that allows you to run applications generated by gwabuildtool in the browser on your desktop.")

[Creating GWA apps with Genero](5126-creating-gwa-apps-with-genero.md "Prepare the environment to build GWA applications.")

[gwa.webmanifest file](5153-gwa-webmanifest-file.md "Example of a customized gwa.webmanifest file, providing information about a Genero Web Application (GWA).")

[Packaging gwa files](../13_programming-tools/2650-packaging-gwa-files.md "Using the fglgar tool to build a Genero Web Application (gwa) file allows you to deploy applications that are ready to run in a browser.")
