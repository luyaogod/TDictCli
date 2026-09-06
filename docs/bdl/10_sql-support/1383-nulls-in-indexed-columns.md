---
title: "NULLs in indexed columns"
source: "fgl-topics/c_fgl_odiagora_049.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > NULLs in indexed columns"
type: "concept"
description: "Oracle btree indexes do not store null values, while Informix® btree indexes do. This means that if you index a single column and select all the rows where that column is null, Informix will do an ..."
---

# NULLs in indexed columns

Oracle btree indexes do not store null values, while Informix® btree indexes do. This means that
if you index a single column and select all the rows where that column
is null, Informix will
do an indexed read to fetch just those rows, but Oracle will do a
sequential scan of all rows to find them. Having an index unusable
for "is null" criteria can also completely change the behavior and
performance of more complicated selects without causing a sequential
scan.

## Solution

Declare the indexed columns as `NOT NULL` with a default value and change the
program logic. If you do not want to change the programs, partitioning the table so that the nulls
have a partition of their own will reduce the sequential scan to just the nulls (un-indexed)
partition, which is relatively fast.
