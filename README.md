# Greenlight API

## Running the server

Within the project directory:

```bash
> cd /cmd/api
> air # To close the server, press Ctrl+C
```

If the port is being used:

```bash
> kill $(lsof -t -i:4000) # replace 4000 with any other port number you want
```

## Environment variables

Enter the following

- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`

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
```

To ensure you've create the password-authenticated user:

```bash
> psql --host=localhost --dbname=<DB_NAME> --username=<DB_USER>
Password:
```

And you should be within the shell after entering the password.
