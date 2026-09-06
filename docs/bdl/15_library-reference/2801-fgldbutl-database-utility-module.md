---
title: "fgldbutl: Database utility module"
source: "fgl-topics/r_fgl_utility_functions_fgldbutl.html"
breadcrumb: "Library reference > Utility modules > fgldbutl: Database utility module"
type: "reference"
description: "Table 1. Database utility functions (fgldbutl.4gl) Function Description Important: This feature is deprecated, its use is discouraged although not prohibited. FUNCTION db_get_database_type () RETURNS ..."
---

# fgldbutl: Database utility module

| Function | Description |
| --- | --- |
| **Important:**This feature is deprecated, its use is discouraged although not prohibited.FUNCTION db_get_database_type() RETURNS STRING | Returns the database type for the current connection. |
| FUNCTION db_get_last_serial( emultype STRING, tabname STRING ) RETURNS BIGINT | Retrieves the last generated serial for a given serial emulation and database table. |
| FUNCTION db_get_sequence( id STRING ) RETURNS BIGINT | Generates a new sequence for a given identifier. |
| FUNCTION db_start_transaction() RETURNS INTEGER | Starts a nested transaction call. |
| FUNCTION db_finish_transaction( commit INTEGER ) RETURNS INTEGER | Terminates a nested transaction call. |
| FUNCTION db_is_transaction_started() RETURNS INTEGER | Indicates whether a nested transaction call is started. |
