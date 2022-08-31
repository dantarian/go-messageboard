# go-messageboard - Message Board API

go-messageboard is an API for a simple message board system written in Go. It's
not intented as a production system - it's an experiment in writing an API in Go
using Clean Architecture principles.

## Getting Started

This project is based on a VS Code devcontainer using Docker Compose. Start by
copying `example.env` to `.env` and fill in the variables in the file. Then open
the project in VS Code, and re-open in container when prompted.

Alternatively, if you've got Go installed and a PostrgreSQL server handy, just
configure your `.env` and `config.yaml` files appropriately and run locally.

Either way, start the server by running:

```bash
go run main.go
```

## Development

As mentioned above, this project uses Clean Architecture. The principle is that
the code is arranged in layers, and inner layers are not allowed to refer to
things in outer layers; this means that business-driven code doesn't concern
itself with purely mechanical matters. For example, if an operation requires
that a message be persisted, it's of no interest to the operation exactly _how_
that happens - whether it's to a PostgreSQL database, or MongoDB, or even
printed onto paper tape.

With that in mind, the various packages in the application are arranged as
follows:

### Entities Layer

This is the innermost layer. The code here is independent of the rest of the
application, and is entirely focused on the domain objects that we're operating
on and their fundamental properties (e.g. validation). The code for this layer
is stored in the `entities` folder.

### Operations Layer

This is the second layer (counting out from the core). The code here is
concerned with the business processes (or use cases) that operate on the
entities defined in the innermost layer (found in the `operations` folder); it
also includes definitions of the repository interfaces that the operations
require for data persistence (found in the `repositories` folder).

### Controllers Layer

At the third layer, we have the controller code that marshalls inputs, invokes
the use cases in the Operations layer, and shapes the responses to the user.
This code can be found in the `controllers` folder.

### Frameworks Layers

This is the outermost layer, and handles:

- The mechanics of data persistence (see the `data` folder, where repository
  interfaces have their concrete implmeentations).
- Configuration (see the `config` folder).
- Routing (see the `api` folder).
- Running the server (see the`server` folder).

### Other

The one exception to this structure is logging. Logging isn't really a domain-
level concern - it's a cross-cutting concern that we may need to access at all
levels of our application. As such it lives in the `util` folder and is broadly
available.
