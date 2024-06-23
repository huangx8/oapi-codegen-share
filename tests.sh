echo "current cards"
curl http://localhost:8888/cards

echo "create 5 cards"
for _ in {1..5}; do
  curl -X POST http://localhost:8888/cards -d '{"owner": "new"}'
done

echo "current cards"
cards=$(curl http://localhost:8888/cards | jq .)
echo "$cards"

id=$(echo "$cards" | jq -r '.[0] | .id')
echo "fetch card $id"
curl http://localhost:8888/cards/"$id"
