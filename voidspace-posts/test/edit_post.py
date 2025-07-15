from fastapi.testclient import TestClient
import pytest
from app.error.error_posts import NotFound
from app.schemas.posts_schemas import PostCreate, Posts as PostResult
from datetime import timezone, datetime
from app.main import app
from app.services.posts_service import PostService

class DummyPostService:
    def edit_posts(self, post_id: str, post_data: PostCreate):
        if post_id == "1":
            return PostResult(
                id=post_id,
                post_content=post_data.post_content,
                post_author_id=post_data.post_author_id,
                post_created_at=datetime.now(timezone.utc),
                post_updated_at=datetime.now(timezone.utc),
                post_image=post_data.post_image,
                likes=0
            )
        else:
            raise NotFound(f"Post with id {post_id} not found")
        
@pytest.fixture
def client():
    app.dependency_overrides[PostService] = lambda: DummyPostService()
    yield TestClient(app)
    app.dependency_overrides.clear()

def test_edit_post(client):
    post_id = "1"
    payload ={
        "post_content": "Updated Content",
        "post_author_id": "author123",
        "post_image": ["image1.jpg", "image2.jpg"]
    }

    response = client.put(f"/api/v1/posts/edit/{post_id}", json=payload)
    assert response.status_code == 200
    assert response.json()["success"] is True
    assert response.json()["message"] == "Post updated successfully"
    assert response.json()["data"]["post_title"] == "Updated Post"