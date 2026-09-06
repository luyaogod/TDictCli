---
title: "Multi-session mode connection instructions"
source: "fgl-topics/c_fgl_connections_multi_session_mode.html"
breadcrumb: "SQL support > Database connections > Multi-session mode connection instructions"
type: "concept"
---

# Multi-session mode connection instructions

> Opening and closing a database for a unique session.

In multi-session mode, open a database session with the `CONNECT TO` instruction.
Other connections can be created with subsequent `CONNECT TO` instructions. To switch
to a specific session, use the `SET CONNECTION` instruction; this suspends other
opened connections. Disconnect from a specific or from all sessions with the
`DISCONNECT` instruction. The end of the program disconnects all sessions
automatically.

## Child topics

- [CONNECT TO](1099-connect-to.md): Opens a new database session in multi-session mode.
- [SET CONNECTION](1100-set-connection.md): Selects the current session when in multi-session mode.
- [SET SESSION AUTHORIZATION](1101-set-session-authorization.md): Select the user under which database operations are performed in the current connection.
- [DISCONNECT](1102-disconnect.md): Terminates database sessions when in multi-session mode.
