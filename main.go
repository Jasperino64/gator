package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type state struct {
	db *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading config:", err)
	}

	// Initialize the database connection
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	
	programState := &state{
		config: cfg,
		db: dbQueries,
	}
	commands := Commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerGetFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))

	
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}
	
	err = commands.run(programState, cmd)
	if err != nil {
		log.Fatalf("error running command: %s", err)
	}
}
