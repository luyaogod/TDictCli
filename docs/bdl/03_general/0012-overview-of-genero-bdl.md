---
title: "Overview of Genero BDL"
source: "fgl-topics/c_fgl_intro_BDL_002.html"
breadcrumb: "General > Introduction to Genero BDL programming > Overview of Genero BDL"
type: "concept"
---

# Overview of Genero BDL

> Genero Business Development Language (BDL) is a program language designed to write an interactive database application.

A Genero BDL application is a set of programs that handle the interaction between a user
and a database. Programs communicate with the database server with Structured Query
Language (SQL), and execute interactive instruction controlling application forms, to
manage user input.

![Interactive database application diagram](../_images/TUT104.jpg)

*Interactive database applications with Genero*

An important feature of the language is the ease with which you can design applications that
allow the user to access and modify data in a database. The language syntax includes a set
of SQL statements to manipulate the database, powerful interactive instructions that
provide simple record input, read-only and read-write record list handling, as well as
database query to search the database, by using forms supporting a large variety of
graphical widgets.

The program sources are compiled to p-code modules, which can be interpreted on
different platforms by the Dynamic Virtual Machine (the Runtime system).

## Child topics

- [Separation of business logic and user interface](0013-separation-of-business-logic-and-user-interface.md): Genero BDL separates business logic and the user interface to provide maximum flexibility.
- [Portability - write once, deploy anywhere](0014-portability-write-once-deploy-anywhere.md): Genero application can be deployed for different kinds of display devices, operating systems and database servers, by using the same source code.
