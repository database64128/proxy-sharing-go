# Configuring the database

proxy-sharing-go can be configured to use different database drivers.

| Database | `database.driver` | Driver | DSN |
| --- | --- | --- | --- |
| SQLite (CGO) | `sqlite3` | [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3) | `file:data.db?_fk=1&_journal_mode=WAL` |
| SQLite (no CGO) | `sqlite3` | [`modernc.org/sqlite`](https://gitlab.com/cznic/sqlite) | `file:data.db?_pragma=foreign_keys(1)&_pragma=journal_mode('WAL')` |
| MySQL | `mysql` | [`github.com/go-sql-driver/mysql`](https://github.com/go-sql-driver/mysql) | `user:password@tcp(localhost:3306)/proxy_sharing_go?parseTime=true&loc=Local` |
| PostgreSQL | `postgres` | [`github.com/jackc/pgx/v5`](https://github.com/jackc/pgx) | `postgres://user:password@localhost:5432/proxy_sharing_go` |

## SQLite

SQLite does not support concurrent writes. Therefore, both `maxOpenConns` and `maxIdleConns` must be set to `1` to prevent `database is locked` errors.

Binaries built with CGO use [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3) as the SQLite driver. Without CGO, the driver is [`modernc.org/sqlite`](https://gitlab.com/cznic/sqlite). These two drivers differ in their DSN format.
