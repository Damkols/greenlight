## This file details the structure and approach to writing the API.

- Imported some packages such as flag, fmt, log/slog, net/http, os, time that will be needed as i proceed with the build

- Created a config struct to hold port values and current environment value (development| staging| production).

- Created an application struct that will hold all app dependencies, this will be useful for dependency injection.

- Initialize an instance of config struct

- Using the flag import to set-up port and environment values, if no commandline values are provided for port and env, the values default to "4000" and "development" respectively. cmd flags(--port, --env)

- Setup a logger that writes to os.Stdout

- Create an instance of application struct with logger and config values

- mux(http.NewServeMux) returns a new ServeMux struct instance that maps route to a specific handler, HandleFunc registers the route.

- Create an http Server that listens on the provided port uses the provided handler for routing, add sensible timeouts and writes any log messages to the structured logger at Error level.

- Log starting server, port and env values to logger Info level

- Start http server with ListenAndServe and log errors to logger Error level

- Created Handlers file for all app handlers, a few package imports, send status, environment and version to the user.

- Added httprouter third party package for proper route handling

- Added a routes file, initialized a new route instance and registered /v1/healthcheck route, using httprouter package

- Created and registered createMovie route handler
- Created and registered showMovie route handler
- Created helpers go file

- Required imports to create a readIDParam helper

- Created an helper function readIDParam that reads the ID of a wildcard segment

- Used the helper function readIDParam to get ID of specific movies to show in showMovieHandler

- Created a writeJson helper function, to send response to users in json format.

- Used writeJson helper to send healthcheck response to user

- Created internal package to hold all custom data and logic for interacting with database

- Created a data type struct Movie

- Used writeJson helper to send a specific Movie json response to user

- Using - (Hyphen) and omitzero directives, we use (-)directive to hide a field thats not relevant to the user e.g (CreatedAt) and omitzero to hide a field with a zero value e.g(Year)

-Using omitempty and string directives

- Used json.MarshalIndent to format json responses

- Envelope Json responses using map

- Created a logError helper method
- Created errorResponse helper method
- Created serverErrorResponse helper method
- Created notFoundResponse helper method
- Created methodNotAllowedResponse helper method

- Use error helper method in healthcheckHandler
- Use error helper method in showMovieHandler
- Use error helpers in routes

-Setup a recoverPanic middleware
- Warp router with panic recover middleware

- Added a createMovie Handler function and route
- Declare an anonymous struct to hold info from the user
- Decode request body content into input struct
- Send contents of input struct as HTTP response
- readJson helper func
- If there is an error during decoding, start the triage
- Use the errors.As() funtion to check whether error has type of *json.SyntaxError
- Check for more errors with errors.Is
- Catch any *json.UnmarshalTypeError errors
- Catch more errors and return error messages
- Use readJson helper
- Created badRequestResponse errors func
- Use badRequestResponse in createMovieHandler
- Use http.MaxBytesReader to limit size of request body to 1mb
- Call DisallowUnknownFields on json decoder
- Create a maxBytesError Variable
- Added more error switch cases to readJson helper
- Call Decode on an empty anonymous struct as the destination and check for more errors
- Create and use Runtime type

- Regexp and slices import in validators package.
- Validators struct type with Errors field map
- func New returns a pointer to Validators struct instance with empty Errors map
- func Valid returns true if errors map contains no entry
- func AddError adds error message to the map
- func Check adds an error message to the map only if validation check is not ok
- func returns true if a specific value is in a list of permitted values
- func Matches returns true if a string value matches a specific regexp pattern
- func returns if all values in a slice are unique
- Declare regexp for sanity checking
- func failedValidationResponse error helper

- Initialize a new Validator instance
- Use Check() method to execute validation checks on user title input
- Use Check() method to execute validation checks on user year input
- Use Check() method to execute validation checks on user runtime input
- Use Check() method to execute validation checks on user genres input
- Use Valid() method to see if any of the checks failed
- Create ValidateMovie function and add all validation checks in the function
- Copy values from input struct to a new movie struct
- Call ValidateMovie func
- Database imports
- Add a Db struct to config struct

- Read DSN value from the db-dsn command-line flag
- Call OpenDB helper function
- Check for errors and log errors
- Defer call to db.Close
- Log a meesage to say connection pool has been succesfully established

- Create openDB func which returns sql.DB connection pool
- Use sql.Open to create an empty connection pool
- Create a context with 5 second timeout
- Use PingContext() to establish a new connection to the database, passing in the context we created above.
- Use value of env
- maxOpenConns, maxIdleConns, maxIdleTime fields for connection pool config
- Read connection pool settings from cmd flags into config struct
- Set maximum number of open connections in the pool
- Set maximum number of idle connections in the pool
- Set maximum idle timeout for connections in the pool

- Generated a pair of migrations file
- Add movies PostgreSQL field and data types to up migrations table
- Add DROP TABLE to down migrations (to reverse up migrations)

- Create second pair of migrations file
- Add CHECK constraints to up migrations
- Drop Check constraints to Down migrations

- Defined a MovieModel struct which wraps sql.DB connection pool
- Placeholder method of type MovieModel for inserting a new record in the movies table
- Placeholder method of type MovieModel for fetching a specific record from the movies table
- Placeholder method of type MovieModel for updating a specific record from the movies table
- Placeholder method of type MovieModel for deleting a specific record from the movies table

- Create Models data file and declare a Models struct to hold MovieModel
- NewModels func returns Models struct with initialized MovieModel struct
- Import data package and add Models struct to application dependencies
- Use NewModels func to initialize a Models struct and pass in db connection pool
- Define SQL query for inserting a new record in the movies table
- Args slice containing values for placeholder params for the Movie struct
- Use QueryRow method to execute sql query on our conection pool.
- Call Insert method on movies model
- Set Location header for newly created resource URL
- Write Json response with a 201 created status code\
- Check if movie Id is greter than 1
- Define sql query for retrieving the movie data
- Declare movie struct to hold data returned from query
- Execute query using QueryRow and scan response data into fields of the movie struct
- Define sql query for updating the movie data
- Args slice containing values for placeholder params for updating the Movie struct
- Use QueryRow method to execute sql query and scan the new version value into movie struct
- Extract movie Id from URL with readIDParam helper
- Fetch existing movie record from database and send a 404 not found if we cant find a matching record
- Declare input struct to hold expected data from client
- Read JSON request body into input struct
- Copy values of request body into appropriate field of movie record
- Validate updated movie record
- Pass updated movie record to Update method
- Write updated movie record in a JSON response
- Delete movie handler function
- Return record not found error if Movie ID is less than 1
- Define sql query for deleting a movie data
- Execute sql Query using Exec() method
- Use RowsAffected methid to get numbers of rows affected by the query
- Check if rowsaffected is equal to zero
- Extract movie id from URL
- Delete movie from database, send a 404 error Not found response to client if no matching record
- Return a 200 ok status code along with success message