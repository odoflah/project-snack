# Reference template for microservice architecture apps

## Abstract and architecture

This is a reference repository for a typical modern cloud application. The project is broken down into several backend services and the frontend service. The backend services are services that provide application logic and frontend is a siplay service that allows users to interface with all the various backend services throiugh one coherent GUI. Backends referes to all backend services (including authentication). Each of the backend services can be interfaced with via a single URL which is achieved by an API gateway (a reverse-proxy server) which sits in front of all the backends and handles routing to the correct service and implements route protection. The frontend is a display layer which provides a single interface with which users may interact with the application through a GUI.

The application architecture is depicted bellow:
<!-- TODO: Add a reference architecture diagram -->
```mermaid
flowchart LR
    A(Frontend) -->|public| B(API gateway)
    B <--> C(Authentication)
    B --> D(Backend service)
    B --> E(Backend service)
```

Application architecture within the context of Kubernetes

Application architecture within the context of GCP

## Why a microservice architecture
- allows for non-uniform scalling

## Design decisions
In order to promote best design practices we have a log which forces is us to record each design desision and justify in order to follow best practices and provide a generalised solution...
This is a log of the design decions for the template repository in order to force me to justify why a decision was made
- For develeopment fully orchestrated with docker compose to make life easier for devs
- all in one repository (monorep microservice application) so that engineers can see the whole codebase and make more informed design decisions
- custom reverse-proxy/api gateway to not be tied to a cloud provider
- Infra as code so as to version control infrastrucrure and create more robust understandable systems
- route protection by default - all services are protected other than the aut service which by definition cxan't be protected
- document each service locally within the service directory - means that the documentation is naturally organised by the repository structure and broken down into manageable chunks per service + quick to access when working on the service

## Setting up the repository for development

### Orchestration

One of the diffciulties when working with microserviuces is that there are many moving components that are required to make the app run - this can get in the way of developmnent as some services may have dependencies on other services in order for them to be worked on. You could run each service locally on the machine in different shell sessions but this can quickly become tedious specially if the architetcure has many services. This is where an automated orchestration tool can be of use. An orchestration tool 'orchestrates' all the different services. It records their dependencies and allows them to communicate with each other. It is used to build all the necessery copnents of the application and bring them up together so that they can communicate with one another. In addition, all teh application logs are aggregated into one shell which makes debugging easier. An orchestration tool such a docker-compose, which is geared towards local orchestration for developes can be an easy way to orchestrate your mciroservices on development machines.

Build each of the services with:
```bash
docker compose -f compose-dev.yml build
```

Then, bring the whole application up with:
```bash
docker compose -f compose-dev.yml up
```

Exit the session with `ctrl+c`

And bring all the services down with
```bash
docker compose -f compose-dev.yml down
```


### Testing the API

There are many way to develop and test an API. You can use excelent tools such as postman with GUIs but the simplest way is to just use the `cURL` programme pre-installed on most unix machines. 

test curl commands for the api

curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signup


curl -X POST -v -H "Content-Type: application/json" -d '{"username": "[USERNAME]", "password": "[PASSWORD]"}' http://localhost:8000/auth/signin

```bash
curl -v --cookie "session_token=[TOKEN]" http://localhost:8000/auth/test
```

## Going into production

While microservice architectures have many benefits they are tricker to work with both in development and production because of the need for orchestration and managing dependencies between each of the services.

While docker-compose works well in the context of development it is limited to running each service on the same machine. This is where Kubernetes comes in - its function is the same as docker compose: orchestration but is able to spread computation across a cluster of machines.