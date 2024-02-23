NGINX API GATEWAY
===

## Installation

- Add ```127.0.0.1    api-gateway.local``` inside ``/etc/hosts ``
- Clone the project
- Run ``` cd nginx-api-gateway ```
- Run ``` docker-compose up --buiild -d ```
- curl api-gateway.local

Your service is now accessible by this urls 
- api-gateway.local/service1
- api-gateway.local/service2

### Token Validation

```bash
curl --request GET \
  --url http://api-gateway.local/service1/ \
  --header 'Authorization: Bearer token'
```
If you don't pass the token or wrong spelling of token It will return 401