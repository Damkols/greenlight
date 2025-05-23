## This file will detail the way this API code is being written.

- Imported some packages such as flag, fmt, log/slog, net/http, os, time that will be needed as i proceed with the build

- Created a config struct to hold port values and current environment value (development| staging| production).

- Created an application struct that will hold all app dependencies, this will be useful for dependency injection.

- Initialize an instance of config struct

- Using the flag import to set-up port and environment values, if no commandline values are provided for port and env, the values default to "4000" and "development" respectively. cmd flags(--port, --env)

- Setup a logger that writes to os.Stdout

- Create an instance of application struct with logger and config values

- mux(http.NewServeMux) returns a new ServeMux struct instance that maps route to a specific handler, HandleFunc registers the route.