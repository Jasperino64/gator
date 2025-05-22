# gator

You will need Postgres and Go installed on your system to run the program.

## Install

go install # to install gator onto the system

## Configuration

Create `.gatorconfig.json` file in your home directory with:
```
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

## Commands

- `login`      - Log in as a user
- `register`   - Register a new user
- `reset`      - Reset the database
- `users`      - List all users
- `agg`        - Aggregate feeds
- `addfeed`    - Add a new feed (requires login)
- `feeds`      - List all feeds
- `follow`     - Follow a feed (requires login)
- `following`  - List feeds you are following (requires login)
- `unfollow`   - Unfollow a feed (requires login)
- `browse`     - Browse feeds (requires login)

