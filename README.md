# Blogo

Sandbox project to learn new stack
 
## Roadmap
- Threaded comment with
  - SQL
    - Raw
    - Entity ORM (entgo)
  - Redis
  - MongoDB
  - Cassandra
  - Graph
- Searchable content (ES/Solr)
- DB Sharding & Replication 
- API
  - REST
  - GraphQL
- Oauth2
- Will add many more

## Development
### Requirement
- Go 1.16 above
- Docker
- docker-compose

### Swagger
1. Create new endpoint by adding via swagger specs in `/api/openapi/swagger.json`
2. Regenerate the server stub using `make gen-server` 

### Migration
1. To add new SQL migration `name=YOUR_MIGRATION_NAME make create-migration`
2. Then `make migrate-up`

### Misc
Other trivial things might not be explained here. Makefile is your best friend