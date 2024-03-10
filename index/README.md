# Sign Indexer

Updates sign search index in MeiliSearch with data from the `vwindexsign` in PostgreSQL database.

## Usage
Configure environment variables:
- RSP_POSTGRES_HOST - Host name of PostgreSQL server
- RSP_POSTGRES_USER - User with permissions to read from `vwindexsign`
- RSP_POSTGRES_PASSWORD - Password for the user
- RSP_POSTGRES_PORT - Port to connect to the PostgreSQL server
- RSP_POSTGRES_DB - Database name hosting `vwindexsign`
- RSP_MEILISEARCH_HOST - Host name of MeiliSearch server to update index on
- RSP_MEILISEARCH_KEY - API key for MeiliSearch server
- RSP_IMAGE_HOSTING_URL - Base url for the path to the image being indexed, used to preview pictures in the MeiliSearch dashboard

```bash
cargo run
```
