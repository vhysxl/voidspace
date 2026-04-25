im going to sleep, no need to ask approval you can direcly edit

1. implement delete user in settings
2. implement delete comment
3. implement search if possible
4. in components who to follow populate with my account (only 1)

{
    "success": true,
    "detail": "User retrieved successfully",
    "data": {
        "id": 41,
        "username": "vhysxl",
        "profile": {
            "display_name": "Lotan",
            "bio": "Developernya bang",
            "avatar_url": "https://storage.googleapis.com/assets_voidspace/avatars/1775756391697344648.jpg",
            "banner_url": "https://storage.googleapis.com/assets_voidspace/banners/1775749269524352400.png",
            "location": "",
            "followers": 1,
            "following": 1
        },
        "created_at": "2026-04-07T05:57:03.503497Z",
        "is_followed": true
    }
}


this user populate hard code in who to follows

		},
		{
			"name": "Search",
			"item": [
				{
					"name": "Search Users",
					"request": {
						"method": "GET",
						"header": [
							{ "key": "x-api-key", "value": "{{apiKey}}", "type": "text" }
						],
						"url": {
							"raw": "{{baseUrl}}/search?q=gemini&type=user",
							"host": ["{{baseUrl}}"],
							"path": ["search"],
							"query": [
								{ "key": "q", "value": "gemini" },
								{ "key": "type", "value": "user" }
							]
						}
					}
				},
				{
					"name": "Search Posts",
					"request": {
						"method": "GET",
						"header": [
							{ "key": "x-api-key", "value": "{{apiKey}}", "type": "text" }
						],
						"url": {
							"raw": "{{baseUrl}}/search?q=hello&type=post",
							"host": ["{{baseUrl}}"],
							"path": ["search"],
							"query": [
								{ "key": "q", "value": "hello" },
								{ "key": "type", "value": "post" }
							]
						}
					}
				},
				{
					"name": "Search Comments",
					"request": {
						"method": "GET",
						"header": [
							{ "key": "x-api-key", "value": "{{apiKey}}", "type": "text" }
						],
						"url": {
							"raw": "{{baseUrl}}/search?q=great&type=comment",
							"host": ["{{baseUrl}}"],
							"path": ["search"],
							"query": [
								{ "key": "q", "value": "great" },
								{ "key": "type", "value": "comment" }
							]
						}
					}
				}
			]
		}
	],