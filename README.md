# Go bucketlist API

This is a simple bucketlist API written in [Go](https://golang.org/).

### Dependencies
- [GORM](http://gorm.io/docs/) - An opensource ORM for golang
- [Mux](https://github.com/gorilla/mux) - a URL router and dispatcher for golang
- [Godotenv](https://github.com/joho/godotenv) - loads environment variables

### Endpoints

| Endpoint | Function | Request payload |
| ------ | ------ | ------ |
| POST /bucketlists | Creates a bucketlist |```{"name": "bucketlist name", "description": "bucketlist description"}```
| GET /bucketlists | Fetch all bucketlists |
| GET /bucketlists/{id} | Fetch a single bucketlist |
| PUT /bucketlists/{id} | Update bucketlist | ```{"name": "bucketlist name", "description": "bucketlist description"}```
| DELETE /bucketlists/{id} | Delete bucketlist |
| POST /bucketlists/{id}/items | Creates a bucketlist items |```{"description": "bucketlist description"}```
| GET /bucketlists/{id}/items | Fetch all bucketlist items |
| GET /bucketlists/{id}/items/{itemId} | Fetch a single item |
| PUT /bucketlists/{id}/items/{itemId} | Update an item | ```{"name": "bucketlist name"}, "description": "bucketlist description"```
| DELETE /bucketlists/{id}/items/{itemId} | Delete an item |
