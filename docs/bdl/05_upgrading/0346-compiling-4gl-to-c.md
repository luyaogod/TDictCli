---
title: "Compiling 4GL to C"
source: "fgl-topics/c_fgl_MigI4GL_009.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > Installation and setup topics > Compiling 4GL to C"
type: "concept"
---

# Compiling 4GL to C

> Genero's use of a p-code architecture removes restrictions on which platforms you use to develop your application.

The IBM® Informix® 4GL compilers include a p-code
based runtime system called RDS as well as a C-compiled solution,
the c4gl compiler. The RDS solution is typically used in a development
environment, supporting a debugger, while the Informix 4GL C compiler is traditionally
used to maximize performance on production sites. However, the
C compiled binaries need to be built on the same target platform
as the production system.

Genero Business Development Language supports a p-code architecture,
which is as fast as the C-compiled version of IBM Informix 4GL.
Since p-code files are portable, you can develop your application
on a platform that is different from the production platform,
saving porting procedures and simplifying deployment tasks.
