---
title: "Packaging gwa files"
source: "fgl-topics/c_fgl_packaging_gwa_apps.html"
breadcrumb: "Programming tools > Packaging web applications > Packaging gwa files"
type: "concept"
---

# Packaging gwa files

> Using the fglgar tool to build a Genero Web Application (gwa) file allows you to deploy applications that are ready to run in a browser.

The fglgar tool run with the `gwa` command creates a
gwa file that embeds the Genero Browser Client.

Once you have built the app, (for details go to[Creating GWA apps with Genero](../18_genero-web-applications/5126-creating-gwa-apps-with-genero.md "Prepare the environment to build GWA applications.")) you can
package the `gwa_dist` directory using `fglgar gwa`, and deploy it on
the GAS. The application can then be accessed at a URL on the server where the GAS is running:
http://host:port/gas/gwa/myapp/index.html

A typical fglgar
`gwa` command is shown:

```
fglgar gwa --gwa gwa_dist -o myapp
```

The example uses two of the command's mandatory options:

- The `--gwa` option is mandatory. It allows you to specify the path to the Genero
  Web Application directory to create the gwa file. You must have a Genero Web
  Application already created to include in the gwa file. For details, go to
  [gwabuildtool](../18_genero-web-applications/5150-gwabuildtool.md).
- The `--output` or `-o` option is mandatory. It specifies the name
  and relative or absolute path to the gwa file to generate.

## Related links

**Related concepts**  

[Genero Web applications](../18_genero-web-applications/5123-genero-web-applications.md "These topics cover programming subjects about Genero Web applications")

[Introducing the GAS and JGAS](2646-introducing-the-gas-and-jgas.md "The Genero Application Server (GAS) is an engine that plugs in to a Web server for the purpose of delivering Genero Web applications and services. The Genero Application Server for Java (JGAS) is designed to run your applications on the Java EE servlet. A general knowledge of how they operate can be helpful in testing and deploying Web applications.")

[fglgar](2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).")
