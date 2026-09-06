---
title: "fglgar"
source: "fgl-topics/c_fgl_tools_fglgar.html"
breadcrumb: "Programming tools > Command reference > fglgar"
type: "concept"
---

# fglgar

> The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).

## Syntax

```
fglgar command [options] [ argument [...] ]
```

1. command is an fglgar command, it can be one of the following:
   - `gar`: To create a Genero Archive file (.gar).
   - `gwa`: To create a Genero Web Application (.gwa) file.
   - `war`: To create a Java Web Archive file (.war).
2. options are specific to each command as described
   below.
3. argument is a parameter for command.

The fglgar tool supports the following commands:

- ```
  fglgar gar [options]
  ```

  The
  gar command creates a .gar Genero Archive file. Options are
  described in Table 2.
- ```
  fglgar gwa [options]
  ```

  The
  gwa command creates a .gwa file. Options are described in
  Table 3.
- ```
  fglgar war [options]
  ```

  The
  war command creates a .war Java Web Archive file embedding a
  Genero Archive. It also provides the option to add a customized Genero Browser Client (GBC). Options
  are described in Table 4.

## Options

| Option | Description |
| --- | --- |
| `-V` or `--version` | Displays version information. |
| `-h` or `--help` | Displays options for the tool. |

| Option | Description |
| --- | --- |
| `-h` or `--help` | Displays help for the gar command. |
| `-v` or `--verbose` | Displays the verbose output of additional information. |
| `-q` or `--quiet` | Operates in silent mode |
| `-o` or `--output output_file_path` | The relative or absolute path to the archive file to create. If not specified, the archive defaults to the name of current directory where the command is run. |
| `-s` or `--input-source directory` | Directory to archive. |
| `--resource directory` | Specifies the Genero Archive resource directory where the application's public images are found. (These are common or can be shared by all your applications.) |
| `--trigger-component component_name` | Specifies Genero Archive application trigger execution component. (Optional). |
| `--deploy-trigger command` | (Optional) Specifies a command to execute after deploying the archive. |
| `--undeploy-trigger command` | (Optional) Specifies a command to execute after undeploying the archive. |
| `--application application_file` | Specifies the application configuration or executable file. If you specify executable (42r or 42m) files instead of xcf file, xcf files are created automatically based on default configuration defined in the GAS as.xcf file. Multiple applications may be specified. Generates a MANIFEST file, if none exists. |
| `--service application_file` | Specifies the service configuration file or executable.If you specify executable (42r or 42m) files instead of xcf file, xcf files are created automatically based on default configuration defined in the GAS as.xcf file. Multiple applications may be specified. Generates a MANIFEST file, if none exists. |

| Option | Description |
| --- | --- |
| `-h` or `--help` | Displays help for the gwa command. |
| `-v` or `--verbose` | Displays the verbose output of additional information. |
| `-q` or `--quiet` | Operates in silent mode |
| `-o` or `--output output_file_path` | The relative or absolute path to the archive file to create. |
| `--gwa` `gwa_directory` | Genero Web Application directory to archive. For usage example, go to [Packaging gwa files](2650-packaging-gwa-files.md "Using the fglgar tool to build a Genero Web Application (gwa) file allows you to deploy applications that are ready to run in a browser.") |

| Option | Description |
| --- | --- |
| `-h` or `--help` | Displays help for the war command. |
| `-v` or `--verbose` | Displays the verbose output of additional information. |
| `-q` or `--quiet` | Operates in silent mode |
| `-o` or `--output output_file_path` | Specifies the relative or absolute path to the war file to create. If not specified, the war defaults to the name of the current directory where the command is run. |
| `-g` or `--input-gar gar-file` | Specifies the Genero Archive (gar) file you want to use to create the war. This option is mandatory. |
| `-w` or `--web-content directory` | Specifies Java Web content directories you want to use. This option must specify the path to the Java Web Content directory of your JGAS installation, for example MyJgas/WebContent, when creating the war. If a second directory is specified, for example to package some additional Java applications or files, its contents overwrite the first directory:`fglgar war --web-content MyJgas/WebContent --web-content myWebContent/WebContent`Using the `--web-content` option more than twice will raise an error. |
| `-c` or `--gbc directory` | Specifies your customized Genero Browser Client (GBC). By default, the gbc installed with the FGLGWS package is embedded in the war file. |

## Usage

The fglgar command line tool is used to create Genero archive
(gar), Genero Web Application (gwa), or Java Web archive
(war) files. You can use it for the following methods of application
deployment:

- You can package web applications or services in a gar file and deploy on
  any web server where the GAS is installed.
- You can package Genero Web Applications in a gwa file and deploy on any web
  server where the GAS is installed and then with the browser you can access your applications.
- You can package Genero applications and services in a war file, deploy on a
  Java EE server like Apache Tomcat®, or Glassfish (via war) and
  then with the browser you can access your applications and services.

> **Important:**
>
> - FGLGWS needs to be installed on the server to interpret the Genero 4GL
>   applications and services.
> - You need a Java Runtime Environment (JRE) version that is at least version 1.8 or greater.

For usage options, run the command fglgar without any other parameters or with
the help ( -h ) option. For help with using a specific command, for example
gar, running fglgar gar, is the same as running
fglgar gar --help.

If an error occurs, the execution status of the command will be non-zero. This allows for the
detection of errors in scripts and makefiles when the command is run from those contexts.

## Related links

**Related concepts**  

[Packaging web applications](2645-packaging-web-applications.md "Describes methods of packaging the runtime files and resources of your web applications and services using the fglgar tool.")
