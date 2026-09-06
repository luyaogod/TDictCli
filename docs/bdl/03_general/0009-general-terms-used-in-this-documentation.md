---
title: "General terms used in this documentation"
source: "fgl-topics/c_fgl_terms_001.html"
breadcrumb: "General > General terms used in this documentation"
type: "concept"
---

# General terms used in this documentation

> This documentation uses general terms that must be clarified for a good understanding.

**Application / Program**

The terms application and program are often synonyms, but can also be
used to reference different concepts. The term application can define all software
components that compose the information system managing a given domain. A program is an
individual component that implements business logic and runs as an application process. Programs are
executed by the runtime system. Program components are typically p-code modules, forms
and additional resource files.

**Application data**

Application data defines the data manipulated by the application. It is typically
managed by one or more database systems. The application data has a volatile state when loaded in
the runtime system, and it has a static state when stored in the database system.

**Database**

The database is a logical entity regrouping the application data. It is managed by
the database system.

**Database system**

The database system is the software that manages data storage and searching; it is
usually installed on the database server machine and is supported by a tier software vendor.

**Developer**

The developer is the person in charge of the conception and implementation of the
application components.

**Deprecated feature**

A deprecated feature is a feature, design, or practice whose use is discouraged
although not prohibited. Typically, a deprecated feature has been superseded or is no longer
considered safe, but it is not yet removed from the system. Four Js provides support for deprecated
features. Bugs will be fixed but enhancements will not be made.

**Desupported feature**

A desupported feature is a feature, design, or practice that is no longer supported.
A desupported feature may still exist, but bugs will not be fixed. The code supporting the feature
may be removed without notice.

**End user**

The end user is the person that uses the application;
that person works on hardware called the workstation.

**Experimental feature**

An experimental feature is a new feature, design, or practice that is provided in a
production software package, but that should only be used for testing in the development
environment, as it can be subject of changes in a next version.

**Front-end**

The front-end is the software that manages the display and input on the end-user
device. It is the software handling the presentation. Available front-ends are GDC, GAS, GMI and
GMA.

**Open Database Interface (ODI)**

The acronym ODI refers to the components used by the runtime system to connect
to an SQL database server, using the SQL client software provided by the database vendor.

**Runtime system (DVM)**

The runtime system is the software that manages the execution of the programs,
where the business logic is processed. The runtime system is also known as the Dynamic
Virtual Machine (DVM - fglrun).

**User interface**

The user interface defines the parts of the programs that interact with the end
user, including interactive elements like windows, screens, input fields, buttons and menus. It
is managed by the front-end.

**Workstation**

The workstation identifies the hardware used by the end user to interact with the
front-end. It can be an dumb terminal, a computer, or mobile device, as long as a front-end is
available on the hardware.
