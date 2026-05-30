# Gator

Gator is a small command-line RSS feed aggregator written in Go. It stores users,
feeds, follows, and fetched posts in Postgres.

## Requirements

Before running Gator, install:

- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

You will also need a running Postgres server and a database for Gator.

## Install

Install the `gator` CLI with `go install`:

```bash
go install github.com/mayaanabbasi/gator@latest
```

Make sure your Go binary directory is on your `PATH`. It is usually:

```bash
export PATH="$PATH:$HOME/go/bin"
```

## Database setup

Create a Postgres database:

```bash
createdb gator
```

Run the schema migrations from the repo. The migration files use
[goose](https://github.com/pressly/goose):

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
```

Update the connection string to match your local Postgres username, password,
host, port, and database name.

## Config

Gator reads its config from a JSON file in your home directory:

```bash
~/.gatorconfig.json
```

Create that file with your Postgres connection string:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

The `current_user_name` field is updated automatically when you register or log
in.

## Usage

Run commands with:

```bash
gator <command> [args...]
```

Common commands:

```bash
gator register alice
gator login alice
gator users
gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"
gator feeds
gator follow "https://blog.boot.dev/index.xml"
gator following
gator agg 1m
gator browse 5
gator unfollow "https://blog.boot.dev/index.xml"
```

Command notes:

- `register <name>` creates a user and sets them as the current user.
- `login <name>` switches the current user.
- `addfeed <name> <url>` adds a feed and follows it for the current user.
- `follow <url>` follows an existing feed.
- `agg <time_between_reqs>` fetches RSS posts on a loop, for example `agg 1m`.
- `browse [limit]` prints recent posts for the current user's followed feeds.
- `reset` deletes all users, which also removes related feeds and follows.
