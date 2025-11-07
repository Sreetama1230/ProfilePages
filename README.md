Understood One to Many Relationship between the entities
Create user with profile
POST http://localhost:8080/users
```
{
  "name": "Alice Johnson",
  "email": "alice.johnson@example.com",
  "profile": {
    "bio": "Backend developer and Go enthusiast",
    "twitter": "@alice_codes"
  }
}

```
Get by Id
http://localhost:8080/users/<id>
