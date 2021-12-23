# Testing using cURL and jq

## Testing GET
curl http://localhost:8080/cloud | jq

## Testing POST
curl -X POST -d '{"Id": "5", "Title": "Alibaba", "desc": "Alibaba Cloud", "content": "Alibaba Cloud"}' -H "Content-Type: application/json" http://localhost:8080/cloud | jq
curl http://localhost:8080/cloud | jq

## Testing PUT
curl -X PUT -d '{"Id": "5", "Title": "Alibaba v2", "desc": "Alibaba Cloud v2", "content": "Alibaba Cloud v2"}' -H "Content-Type: application/json" http://localhost:8080/cloud/5 | jq
curl http://localhost:8080/cloud | jq

## Testing DELETE
curl -X DELETE http://localhost:8080/cloud/5
curl http://localhost:8080/cloud | jq
