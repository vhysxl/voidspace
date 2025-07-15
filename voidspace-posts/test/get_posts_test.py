import pytest
from fastapi.testclient import TestClient
from app.main import app
from app.models.posts_models import Posts
from app.services.posts_service import PostService

class DummyPostService:
    def get_posts(self):
        return [
            Posts(
                id="1",
                post_content="Test content",
                post_author_id="123",
                post_created_at="2023-01-01",
                post_updated_at="2023-01-01",
                post_image=None,
                likes=0
            )
        ]

@pytest.fixture
def client():
    # override dependency
    app.dependency_overrides[PostService] = lambda: DummyPostService()
    yield TestClient(app)
    app.dependency_overrides.clear()

def test_get_posts(client):
    response = client.get("/api/v1/posts/")
    assert response.status_code == 200
    assert response.json()["success"] is True
    assert response.json()["message"] == "Posts retrieved successfully"
    assert isinstance(response.json()["data"], list)
    assert response.json()["data"][0]["post_content"] == "Test content"