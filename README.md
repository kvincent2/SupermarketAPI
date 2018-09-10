# Supermarket API

The purpose of this API is to manage produce inventory for a grocery store.

## Getting Started:
Use Dep to gather required dependencies
 ```
 dep init
 dep ensure - update
 ```
 
### Endpoints:
/GetProduce  --  Gets all produce in inventory in JSON format.

/GetProduceByID -- Takes a ProduceCode URL parameter and returns the record for that item in JSON format.

Example: 
```
localhost:8080/GetProduceByID?ProduceCode=A12T-4GH7-QPL9-3N4M
```

/PostProduce -- Pass a JSON object as the body of your request to add a new record to the database.
Example:
```
{  
   "Name":"Pineapple",
   "ProduceCode":"A23K-4GH7-QPL9-1B2U",
   "UnitPrice":4.26
}
```

/DeleteProduce -- Takes a ProduceCode URL parameter and deletes that item from the database.
Example:
```localhost:8080/DeleteProduce?ProduceCode=A12T-4GH7-QPL9-3N4M```

##### Author
**Kristina Vincent** 