### This template has integrated the following modules:
* Consul as the registration center
* Jaeger for distributed tracing
* JWT authentication
* Service rate limiting
## Usage
```bash
cookiecutter https://github.com/Fl0rencess720/kratos-grpc-service-template.git
``` 
After configuring the template, run the following commands:
```bash
make config
wire cmd/server
```
