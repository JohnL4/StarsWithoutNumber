The SQLite file stays on the machine that runs the server. A database created
on Windows opens on Linux. The published database is created on the server, in
a directory the service user can write, such as `/var/lib/swn/swn.db`.

Git ignores `/data/*.db`, `/data/*.db-wal`, and `/data/*.db-shm`.
