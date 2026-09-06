---
title: "Compile and run the service"
source: "fgl-topics/t_gws_restful_high_level_quick_start_start_server.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 1: RESTful server application > Compile and run the service"
type: "task"
---

# Compile and run the service

> Compile and execute a REST Web Services server in direct mode.

When the server starts, it listens on the port specified by FGLAPPSERVER.

1. Compile the server application.

   fglcomp wsserver.4gl

   fglcomp also compiles the imported module and creates
   wsserver.42m and myservice.42m.
2. Start the Genero RESTful Web Service server.

   Use the fglrun command to execute the web service server application.

   ```
   fglrun wsserver.42m
   ```

   The server starts. Output is written to stdout (and
   stderr when FGLWSDEBUG is set). As the server processes requests, the GWS
   outputs the WS debug information:

   ```
   WS-DEBUG (Info)
   Default zlib loaded
   WS-DEBUG END

   WS-DEBUG (Security Info)
   OpenSSL 3.1.1 30 May 2023
   WS-DEBUG END

   WS-DEBUG (Security Warning)
   Crypto library wasn't compiled with support of ZLIB compression.
   WS-DEBUG END

   WS-DEBUG (Security Warning)
   SSL library wasn't compiled with support of RLE compression.
   SSL library wasn't able to initiate the ZLIB compression library.
   WS-DEBUG END

   ...
   Server started
   ```
