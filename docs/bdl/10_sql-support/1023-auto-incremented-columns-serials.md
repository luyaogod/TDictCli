---
title: "Auto-incremented columns (serials)"
source: "fgl-topics/c_fgl_sql_programming_073.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Auto-incremented columns (serials)"
type: "concept"
---

# Auto-incremented columns (serials)

> How to implement automatic record keys.

IBM® Informix®
provides the SERIAL, BIGSERIAL or SERIAL8 data types which can be emulated with database drivers for
most non-Informix database engines by using
native sequence generators (when "`ifxemul.serial`" FGLPROFILE setting is
`true`).

But, this requires additional configuration and maintenance tasks. If you plan to review the
programming pattern of sequences, it is recommended that you use a portable implementation instead
of the serial emulation provided by the database drivers.

This section describes different solutions to implement auto-incremented fields. The
preferred implementation is the solution using SEQUENCES.

## Child topics

- [Solution 1: Use database specific serial generators](1024-solution-1-use-database-specific-serial-generators.md)
- [Solution 2: Generate serial numbers from your own sequence table](1025-solution-2-generate-serial-numbers-from-your-own-sequence-ta.md)
- [Solution 3: Use native SEQUENCE database objects](1026-solution-3-use-native-sequence-database-objects.md)
