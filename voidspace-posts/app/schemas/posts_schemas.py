from datetime import datetime
from pydantic import BaseModel, ConfigDict

class PostCreate(BaseModel):
    post_content: str
    post_author_id: str
    post_image: list[str] | None = None

class Posts(BaseModel):
    id: str
    post_content: str
    post_author_id: str
    post_created_at: datetime
    post_updated_at: datetime
    post_image: list[str] | None = None
    likes: int = 0

    model_config = ConfigDict(from_attributes=True) # ini agar orm bisa di convert ke pydantic
