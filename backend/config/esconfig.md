Meme Index Creation

PUT http://localhost:9200/memes

Request Body:

```json
{
	"mappings": {
		"properties": {
            "id": {
                "type": "text",
                "index": false
            },
			"text": {
				"type": "text"
			},
			"description": {
				"type": "text"
			},
			"name": {
				"type": "text"
			},
			"tags": {
				"type": "text"
			},
            "origin": {
				"type": "text"
			},
			"url": {
				"type": "text",
				"index": false
			},
            "created_at": {
                "type": "date"
            },
            "updated_at": {
                "type": "date"
            }
		}
	}
}
```