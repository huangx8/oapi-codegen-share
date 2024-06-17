curl http://localhost:8888/cards

for _ in {1..5}; do
  curl -X POST http://localhost:8888/cards -d '{"owner": "new"}'
done

curl http://localhost:8888/cards | jq .

curl http://localhost:8888/cards/0cefa3a4-e022-453d-84f7-e2e1c3bafef1
