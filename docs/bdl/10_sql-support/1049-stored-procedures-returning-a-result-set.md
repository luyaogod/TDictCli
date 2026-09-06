---
title: "Stored procedures returning a result set"
source: "fgl-topics/c_fgl_sql_programming_008.html"
breadcrumb: "SQL support > SQL programming > SQL portability > Stored procedures > Stored procedures returning a result set"
type: "concept"
description: "With some database servers it is possible to execute stored procedures that produce a result set, and fetch the rows as normal SELECT statements, by using DECLARE , OPEN , FETCH . Some databases can ..."
---

# Stored procedures returning a result set

With some database servers it is possible to execute stored procedures
that produce a result set, and fetch the rows as normal `SELECT` statements,
by using `DECLARE`, `OPEN`, `FETCH`.
Some databases can return multiple result sets and cursor handles
declared in a stored procedure as output parameters, but Genero supports
only unique and anonymous result sets. See the examples.
