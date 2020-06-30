# Golang API Fiber
> REST API using golang fiber with React on frontend

## Demo
https://tuxrace.github.io/golang-api-fiber/.

![screen](./screen.gif "Screen")

## Install
```bash
   npm install
   npm run start:dev
```
This will run both `front end` and `back end`

## API
| Method   |      API               |
|----------|------------------------|
<<<<<<< HEAD
| GET      |  /api/appliances`      |
| GET      |  /api/appliances/1   |
=======
| GET      | /api/appliances        |
| GET      | /api/appliances/1      |
>>>>>>> 2693a9c95f6534db370d3c02dec8df3b1f05f7fa
| POST     | /api/appliances        |
| DELETE   | /api/appliances/1      |
| GET      | /api/appliances-search?category=model&search=Tub |

## API Usage

### Get all records
```bash
CURL /api/v1/appliances
```

### Add a record
```bash
CURL -X POST -H /api/appliances --data "{\"serial_number\": \"1111\", \"brand\": \"Mayer\",
\"model\": \"Gas Range\"}" /api/v1/appliances
```

### Get a record
```bash
CURL /api/v1/appliances/1
```

### Delete a record
```bash
CURL -X DELETE /api/v1/appliances/1
```
