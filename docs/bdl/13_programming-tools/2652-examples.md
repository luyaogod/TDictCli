---
title: "Packaging examples"
source: "fgl-topics/c_fgl_packaging_applications_examples.html"
breadcrumb: "Programming tools > Packaging web applications > Examples"
type: "concept"
---

# Packaging examples

> With the fglgar tool you can immediately see the results of a Genero Web Application or Web Service during your development stage without having to install a Web server or a GAS package.

## Packaging and running applications with fglgar

If you have a directory named MyFirstApp with a Genero
4GL application and/or a service that you wish to test, then with just three
fglgar commands applications are packaged and can be viewed in a browser.

```
$ cd MyFirstApp
$ fglcomp -M HelloWorld.4gl
$ fglcomp -M MyWebService.4gl
$ fglgar gar --application HelloWorld.42m --service MySebService.42m
$ fglgar war --input-gar MyFirstApp.gar
```

Where:

1. The application and service source files (4gl) are compiled.
2. fglgar gar  command is run to create an archive file with the compiled
   applications and services.
3. fglgar war command is run to create a war file that
   embeds the gar file.
4. For information on running the applications, go to [Running applications from a war file](2651-running-applications-from-a-war-file.md "Applications packaged in a Java Web Archive (war) file and deployed in any existing Java Enterprise Edition container such as Apache Tomcat, Jetty, or Glassfish can be run in a browser.").

## Adding a customized GBC to the package

If you have customized Genero Browser Client (GBC) in a directory named
gbc\_customized, then you can include it in the package with the
`--gbc` option. The `--gbc` option is optional but it allows you
specify a GBC that you have customized specifically for your applications. Otherwise, the default
gbc installed with the FGLGWS package is embedded in the war
file.

```
$ cd MyFirstApp
$ fglcomp -M HelloWorld.4gl
$ fglcomp -M MyWebService.4gl
$ fglgar gar --application HelloWorld.42m --service MyWebService.42m
$ fglgar war --input-gar MyFirstApp.gar --gbc c:\dev\gbc_customized
```

Where:

1. The application and service source files (4gl) are compiled.
2. fglgar gar  command is run to create an archive file with the compiled
   applications and services.
3. fglgar war command is run to create a war file that
   embeds the gar file.
4. For information on running the applications, go to [Running applications from a war file](2651-running-applications-from-a-war-file.md "Applications packaged in a Java Web Archive (war) file and deployed in any existing Java Enterprise Edition container such as Apache Tomcat, Jetty, or Glassfish can be run in a browser.").

## Packaging multiple applications

If you have multiple applications to deploy in your MyApps directory that
you wish to test, you can add them to the fglgar gar with `--application` options, as shown. The applications are bundled in the gar and run
from the war file.

```
$ cd MyApps
$ fglcomp -M HelloWorld.4gl
$ fglcomp -M app1.4gl
$ fglgar gar --application HelloWorld.42m --application app1.42m 
$ fglgar war --input-gar MyApps.gar
```

Where:

1. The application and service source files (4gl) are compiled.
2. fglgar gar  command is run to create an archive file with the compiled
   applications and services.
3. fglgar war command is run to create a war file that
   embeds the gar file.
4. For information on running the applications, go to [Running applications from a war file](2651-running-applications-from-a-war-file.md "Applications packaged in a Java Web Archive (war) file and deployed in any existing Java Enterprise Edition container such as Apache Tomcat, Jetty, or Glassfish can be run in a browser.").

## Related links

**Related concepts**  

[Introducing the GAS and JGAS](2646-introducing-the-gas-and-jgas.md "The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.")
