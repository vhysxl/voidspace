import pytest
from fastapi.testclient import TestClient
from app.main import app
from app.schemas.posts_schemas import Posts as PostResult, PostCreate
from app.services.posts_service import PostService
import uuid
from datetime import datetime, timezone

class DummyPostService:
    def create_posts(self, post_data: PostCreate):
        return PostResult(
            id=str(uuid.uuid4()),
            post_content=post_data.post_content,
            post_author_id=post_data.post_author_id,
            post_created_at=datetime.now(timezone.utc), 
            post_updated_at=datetime.now(timezone.utc), 
            post_image=post_data.post_image,
            likes=0 
        )
    
@pytest.fixture
def client():
    app.dependency_overrides[PostService] = lambda: DummyPostService()
    yield TestClient(app)
    app.dependency_overrides.clear()

def test_create_post(client):
    payload = {
        "post_content": "new Content", 
        "post_author_id": "author123",
        "post_image": ["image1.jpg", "image2.jpg"]
    }

    response = client.post("/api/v1/posts/create", json=payload)
    assert response.status_code == 200
    assert response.json()["success"] is True
    assert response.json()["message"] == "Posts Created Successfully"
    assert response.json()["data"]["post_content"] == "new Content"
    assert response.json()["data"]["id"] is not None




