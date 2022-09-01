Reference repository for a typical modern cloud application. 

Integrates with CI/CD practices

TODO: Add a reference architecture diagram

Design decisions
This is a log of the design decions for the template repository in order to force me to justify why a decision was made
- For develeopment fully orchestrated with docker compose to make life easier for devs
- all in one repository (monorep microservice application) so that engineers can see the whole codebase and make more informed design decisions
- custom reverse-proxy/api gateway to not be tied to a cloud provider
- Infra as code so as to version control infrastrucrure and create more robust understandable systems


<!-- TODO: Enable protected routes with the reverse proxy and auth service -->
<!-- TODO: Finish auth service -->

https://kompose.io/architecture/


Template project to setup GCP project microservice architectyres 

In the isAuthenticated function use token as signature i.e. is the token present and is the token matching to the user trying to access his individual user information - get the user not by getting the user value from the frontend but by getting the user value associated with the token - this means it has to be a token for the user
https://cdn.shopify.com/shopifycloud/shopify_dev/assets/partners/jwt-request-flow-8377bd9698797d2d23713676585a01f9da42c80596ebdc673b971a1e577c65d4.png


test curl commands for the api

curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signup


curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signin

```bash
curl -v --cookie "session_token=[TOKEN]" http://localhost:8000/auth/test
```
