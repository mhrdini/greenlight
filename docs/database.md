# Database

## Setup

**NOTE:** The variables in the enclosed brackets should have the same values as the environment variable
equivalents stored in the `.env` file.

### Creating the database (`DB_NAME`)

```bash
> sudo -u postgres psql
# Within psql
> create database <DB_NAME>;  # create the DB under the user 'postgres'
> \c <DB_NAME>                # connect to the DB
```

### Creating the user (`DB_USER`) and password (`DB_PASSWORD`)

Still within `psql`:

```bash
> create role <DB_USER> with login password '<DB_PASSWORD>';  # note the enclosing single quotes
> create extension if not exists citext;                      # for case-insensitive fields, i.e. email
> alter database <DB_NAME> owner to <DB_USER>;
```

To ensure you've create the password-authenticated user:

```bash
> export GREENLIGHT_DB_DSN=postgres://<DB_USER>:<DB_PASSWORD>@localhost/<DB_NAME>?sslmode=disable
> psql $GREENLIGHT_DB_DSN
# or
> psql --host=localhost --dbname=<DB_NAME> --username=<DB_USER>
Password:
```

And you should be within the shell after entering the password.

## Migrations

Install the following dependency:

```bash
> go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

This tool will help to keep track which migrations have already been applied.

### Creating migration files

The following command will create an empty up and down SQL files prefixed with a sequential, incrementing
number (e.g. 000001) at the `/migrations` subdirectory:

```bash
> migrate create -seq -ext=.sql -dir=./migrations <MIGRATION_NAME>
```

"Up" migration files contain the SQL statements necessary to implement changes.
"Down" migration files contain the SQL statements to reverse or _roll-back_ the changes.

### Running migrations

Ensuring you have the DSN string exported under `GREENLIGHT_DB_DSN`, you can perform up and down
migrations:

```bash
> migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
> migrate -path=./migrations -database=$GREENLIGHT_DB_DSN down
```

## Troubleshooting

If the following shows up when connecting to `psql`:

```bash
> psql ...
psql: error: connection to server on socket "/tmp/.s.PGSQL.5432" failed: No such file or directory
        Is the server running locally and accepting connections on that socket?
```

Then write the following command:

```bash
> pg_ctl start
pg_ctl: another server might be running; trying to start server anyway
...
# now you can connect to psql normally
> psql ...
```

Connect to the database once inside `psql`. For example if the `DB_NAME` is "greenlight" and
`DB_USER` is "postgres":

```bash
postgres=# \c greenlight
You are now connected to database "greenlight" as user "postgres".
```

## PostgreSQL Commands

### Full-Text Search

On a `title` field, using a `WHERE` clause:

```sql
...
WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
...
```

#### Explanation

- `to_tsvector('simple', title)` function takes a movie title and splits it into lexemes
- `simple` configuration means that lexemes are just lowercase version of the words in the title

**Example:** "The Breakfast Club" would be split into the lexemes `'breakfast'` `'club'` `'the'`.

- `plainto_tsquery('simple', $1)` function takes a search value and turns it into a formatted query term that PostgreSQL full-text search can understand
- using the `simple` configuration, it normalises the search value, strips any special charactes,
  and inserts the `and` operator `&` between the words

**Example:** "The Club" would results in the query term `'the' & 'club'`.

- `@@` operator is the matches operator, used to check whether the generated query term matches the lexemes

### Pagination

- `LIMIT` to set maximum number of records that an SQL query returns
- `OFFSET` to 'skip' a specific number of rows before starting to return records from the query

Both written after `ORDER BY` clause:

```sql
ORDER BY ...
LIMIT 5 OFFSET 10
```
