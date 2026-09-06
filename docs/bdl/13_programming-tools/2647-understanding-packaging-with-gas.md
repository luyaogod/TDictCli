---
title: "Understanding packaging with GAS"
source: "fgl-topics/c_fgl_packaging_applications_003.html"
breadcrumb: "Programming tools > Packaging web applications > Understanding packaging with GAS"
type: "concept"
---

# Understanding packaging with GAS

> Before deploying applications on a GAS installation, you need to package the required compiled files and resources in a Genero archive file. The fglgar command provides this function and you can find out here what is involved in this process.

Multiple applications may be bundled in a gar file. As a prerequisite, all
of the required files and resources must be in a root directory. When you run the
fglgar tool with the `gar` command a Genero archive is
created.

```
fglgar gar --application helloWorld.42r --service welcomeService.42m
```

## Example gar file

The gar file contains a MANIFEST file, your application
modules, form files, configuration files, for example:

- MANIFEST
- helloWorld.42m
- helloWorld.42r
- helloWorld.42f
- welcomeService.42m
- helloWorld.xcf
- welcomeService.xcf

## The MANIFEST file

The MANIFEST is an XML file that essentially provides a list of the
applications and services in the gar to make available. It can be created
automatically by fglgar at the command line for applications specified with the
`--application` and/or `--service` option.

Or alternatively, if you have many applications to package, you may find it easier to first
create the MANIFEST by hand. Running fglgar (without the
`--application` and/or `--service` option) it checks if a
MANIFEST file is present in the directory, and uses it to create the
gar. For more information see the MANIFEST file topic in the Genero Application Server User Guide.

## Application configuration files

You can provide the application configuration (xcf) files if you wish but
they are created automatically for you if you provide the `--application` and/or
`--service` options with the name of the executable (42r or
42m) files instead of xcf files.

The xcf file is created based on default configurations in the GAS
configuration file, as.xcf:

- With option `--application`, fglgar creates a default
  xcf based on the "`defaultwa`" configuration.
- With option `--service`, fglgar creates a default
  xcf based on "`ws.default`" configuration.

In both cases, the generated xcf file is given the name of the
42r or 42m provided at the command line.

## Deploy the gar file and enable its applications

You deploy the archive on your GAS installation locally using the gasadmin
tool, or the FGL Web services PublishGAR tool that is provided in the Genero BDL
installation.

If the GAS is on a remote Web server, you can only deploy applications with
PublishGAR. You can also use other tools such as curl, etc.

These tools unpack the contents of the archive in the GAS
`$(res.appdata.path)`/deployment path. The applications deployed
are identified by the name of the xcf. To make the applications available to
end users, enable the archive using either the gasadmin or the
PublishGAR tool. For examples, see [Deploying applications on GAS](2648-deploying-applications-on-gas.md "Follow these procedures to build archives for deploying applications and services with fglgar.").

## Related links

**Related concepts**  

[Introducing the GAS and JGAS](2646-introducing-the-gas-and-jgas.md "The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.")
