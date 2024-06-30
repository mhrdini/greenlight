# Greenlight API

## Database setup

You must set up the database according to the instructions [here](https://github.com/mhrdini/greenlight/blob/main/docs/database.md).

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

Enter the following variables in an `.env` file:

- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
