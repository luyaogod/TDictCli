---
title: "Migrating from IBM Informix 4GL to Genero BDL"
source: "fgl-topics/c_fgl_MigI4GL_001.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL"
type: "concept"
description: "Product differences you must be aware of (and plan for) when migrating from IBM Informix 4GL to Genero Business Development Language."
---

# Migrating from IBM Informix 4GL to Genero BDL

> Product differences you must be aware of (and plan for) when migrating from IBM® Informix® 4GL to Genero Business Development Language.

## IBM Informix 4GL and Genero BDL products

IBM
Informix 4GL (I4GL) and Genero Business Development
Language (BDL) are distinct development tools. The goal of Genero BDL is to be as compatible as
possible with I4GL, and it is very close. The success of Genero BDL depends on the ability to
compile and run legacy code with minimum code changes. For text-mode applications, the migration
steps are often reduced to recompile-and-run.

Genero BDL extends the I4GL language with advanced features, such as a Graphical User Interface
and SQL access to non-Informix databases.
This leads to some differences that you have to handle, but these incompatibilities are minor
compared to the added value.

In some rare cases, the Genero BDL team decided to take a different path to implement an I4GL
feature, because we considered that the IBM
Informix 4GL solution was not adaptable. For example,
[dynamic arrays in I4GL and Genero BDL](0375-dynamic-arrays.md "Support for dynamic arrays differs between I4GL and Genero BDL.") have different
semantics.

This guide will help you identify the differences and find solutions to make the migration from
IBM
Informix 4GL easier.

## IBM Informix 4GL reference version

Several versions of the IBM
Informix 4GL language have been released. It started in
the mid-80s with I4GL version 4.x; then came version 6.x in 1996. I4GL version 7.2 was released in
1998; then versions 7.31, 7.32, and finally the version: 7.50 came out.

There have been several bug fixes and enhancements over the life of I4GL, resulting in releases
that slightly differ. Supporting strict compatibility with all versions of I4GL is not possible for
Genero BDL.

The Genero BDL compatibility level with IBM
Informix 4GL is achieved by comparing with the latest
version of I4GL, which is version 7.50 at the time of this writing.

## Child topics

- [Installation and setup topics](0342-installation-and-setup-topics.md): When migrating from I4GL to Genero BDL, review the differences between the installation and setup of the two products. Reviewing the differences allows you to plan and prepare for a smooth migration.
- [User interface topics](0347-user-interface-topics.md): When migrating from I4GL to Genero BDL, review the differences between how windows and form content is rendered between the two products. Reviewing the differences allows you to plan and prepare for a smooth migration.
- [4GL programming topics](0373-4gl-programming-topics.md): When migrating from I4GL to Genero BDL, review the programming differences between the two products. Reviewing the differences allows you to plan and prepare for a smooth migration.
