curl localhost:8080/prov/put \
-H "Content-Type: application/json" \
--data-binary @- << EOF 
{
 "entities": [
  {
   "id":"entity1"
  },
  {
  "id":"entity2", 
  "wasDerivedFrom":["entity1"], 
  "wasGeneratedBy":["activity"], 
  "wasAttributedTo":["agent"]
  }
 ],
 "activities":[
  {
   "id":"activity",
   "used":["entity1"],
   "wasAssociatedWith":["agent"]
  }
 ],
 "agents": [
  {
   "id": "agent",
   "actedOnBehalfOf":["organization"]
  }
 ]
}
EOF
