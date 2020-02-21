# protoc-gen-graphql

MVP:
* [ ] Given a set of Protobuf `message`s and `service`s, generate a:
  * [ ] GraphQL Schema
    * [ ] Allow introspection into the generated schema
  * [ ] GraphQL Queries that correspond with the `rpc`s in each of the `service`s
  * [ ] Implement resolvers that communicate with the `service`s.
* [ ] Accept a single Github or Gitlab organization as input, parse all projects, look for protobuf files, and use those as inputs for generation.
* [ ] Authentication
  * [ ] Bearer token
  * [ ] Scopes
