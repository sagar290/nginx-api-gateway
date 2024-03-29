NGINX API GATEWAY
===

### Installation

- Add ```127.0.0.1    api-gateway.local``` inside ``/etc/hosts ``
- Clone the project
- Run ``` cd nginx-api-gateway ```
- Run ``` docker-compose up --buiild -d ```
- curl api-gateway.local

Your services are now accessible through these URLs 
- api-gateway.local/service1
- api-gateway.local/service2

### Token Validation

```bash
curl --request GET \
  --url http://api-gateway.local/service1/ \
  --header 'Authorization: Bearer token'
```
If you don't pass the token or the wrong spelling of the token It will return 401

[Click Here for Full Details](https://sagardash.me/make-your-nginx-api-gateway-with-auth-validation-7efd122a18d3)
