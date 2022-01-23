# README
Based on [this medium article](https://levelup.gitconnected.com/lets-go-and-build-graphql-api-with-gqlgen-bfea2f346ea1). The setup should be based upon [gqlgen.com](https://gqlgen.com/) though.

## Example
After adding the pokemon schemas to `schema.graphqls`, implement the "create pokemon mutation" and the "list pokemon" query resolver.

Run `go run ./server.go`.

Then head to the browser on `http://localhost:8080/` and execute:

```graphql
mutation {
  CreatePokemon(input:{
    name:"Pikachu",
    Power:"Fire",
    Description:"Fluffy",
    date:"2022-01-32"
}) {
  id
  name
  description
}
}
```

And with the following you can query the pokemons:
```graphql
{
  ListPokemon {
    name
    description
    power
  }
}
```