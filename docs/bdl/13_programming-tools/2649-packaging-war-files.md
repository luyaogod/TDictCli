---
title: "Packaging war files"
source: "fgl-topics/c_fgl_packaging_applications_004.html"
breadcrumb: "Programming tools > Packaging web applications > Packaging war files"
type: "concept"
---

# Packaging war files

> Using the fglgar tool to build a Java Web Archive (war) file allows you to deploy applications that are ready to run.

The fglgar tool run with the `war` command creates a
war file that embeds the Genero Archive file (gar).

A Java Enterprise Edition (Java EE) server is required for the development and production
environment to access your applications and services via a browser.

A typical fglgar
`war` command is
shown:

```
fglgar war --web-content MyJgasDir/WebContent -g MyApp.gar --gbc c:/dev/gbc_custom -o ../MyGeneroJavaApps.war
```

The example uses four of the command's main options:

- The `--web-content` option is mandatory. It allows you to specify the path to the
  Java Web Content directory of your JGAS installation.
  > **Tip:**
  >
  > A second `--web-content` option may be used to specify a WebContent directory
  > where, for example, you have some additional Java applications or files that overwrite files copied
  > in the first option from the Web Content directory of your JGAS installation.
- The `--input-gar` (`-g` ) option is mandatory. It allows you to
  specify the Genero Archive you want to use to create the war file. This means
  that you need to have a gar file already created to include in the
  war file. See [Deploying applications on GAS](2648-deploying-applications-on-gas.md "Follow these procedures to build archives for deploying applications and services with fglgar.").
- The `--gbc` option is optional but it allows you specify a different Genero
  Browser Client (GBC), for example, one that you have customized specifically for your applications.
  Otherwise, the default gbc installed with the FGLGWS package is embedded in the
  war file.
- The `--output` or `-o` option specifies the relative or absolute
  path to the war file to generate.

## Related links

**Related concepts**  

[Introducing the GAS and JGAS](2646-introducing-the-gas-and-jgas.md "The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.")

[fglgar](2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).")
