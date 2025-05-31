## This file will detail the way this API code is being written.

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