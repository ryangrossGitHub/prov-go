# PROV-GO
A service for inserting and retrieving provenance data over REST to a graph data store.

### Example Usage
curl localhost:8080/prov/put -H "Content-Type: application/json" -d '{"id":"entity2", "wasDerivedFrom":"entity1", "wasGeneratedBy":"activity", "wasAttributedTo":"agent"}'
