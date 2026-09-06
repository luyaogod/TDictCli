---
title: "Unique session mode connection instructions"
source: "fgl-topics/c_fgl_connections_unique_session_mode.html"
breadcrumb: "SQL support > Database connections > Unique session mode connection instructions"
type: "concept"
---

# Unique session mode connection instructions

> Opening and closing a database for a unique session.

In unique-session mode, the `DATABASE` instruction initiates a connection the
database server and creates the current session. The database connection is terminated with the
`CLOSE DATABASE` instruction, or when another `DATABASE` instruction
is executed, or when the program ends.

## Child topics

- [DATABASE](1096-database.md): Opens a new database connection in unique-session mode.
- [CLOSE DATABASE](1097-close-database.md): Closes the current database connection created by a DATABASE instruction.
