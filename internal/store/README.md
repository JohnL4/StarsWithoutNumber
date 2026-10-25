SQLite access will live in this package, behind a repository the domain can
call. Domain packages do not import `database/sql`.

Use the latest stable `modernc.org/sqlite` when this package is implemented.
It is a pure-Go SQLite driver registered as `sqlite` for `database/sql`, so
the server builds without gcc or MinGW. Open one connection with
`SetMaxOpenConns(1)`, and set `PRAGMA journal_mode=WAL` and
`PRAGMA busy_timeout=5000`.

The database file's location is in [data/README.md](../../data/README.md).
