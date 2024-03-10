# Random Sign Generator
Cloudflare Worker that generates a random sign based on data from the Workers KV store.  Uses Bun and the Hono framework.

### Api
- `GET /` - Returns a random sign as a 302 redirect back to the referring url.  In production this is https://roadsign.pictures
- `GET /state/:state?idonly=true` - Returns a random sign filtered by state slug.  If the `idonly=true` query parameter is present, the response will be a json object with the id of the sign.  Otherwise, the response will be a 302 redirect back to the referring url.
- `GET /state/:state?idonly=true` - Returns a random sign filtered by state slug.  If the `idonly=true` query parameter is present, the response will be a json object with the id of the sign.  Otherwise, the response will be a 302 redirect back to the referring url.
- `GET /statesubdivision/:county?idonly=true` - Returns a random sign filtered by state/county slug.  If the `idonly=true` query parameter is present, the response will be a json object with the id of the sign.  Otherwise, the response will be a 302 redirect back to the referring url.
- `GET /place/:place?idonly=true` - Returns a random sign filtered by state/place slug.  If the `idonly=true` query parameter is present, the response will be a json object with the id of the sign.  Otherwise, the response will be a 302 redirect back to the referring url.

- To Run Locally
```bash
bun install
bun run dev
```

To deploy
```bash
bun run deploy
```
