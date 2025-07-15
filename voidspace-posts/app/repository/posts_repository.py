from datetime import datetime, timezone
from fastapi import Depends
from sqlalchemy import desc, table, update
from sqlalchemy.orm import Session
import uuid

from app.core.database import get_db
from app.error.error_posts import Forbidden, NotFound
from app.models.posts_models import Posts as PostsModel
from app.schemas.posts_schemas import PostCreate


class PostRepository:
    def __init__(self, db : Session = Depends(get_db)):
        self.db = db

    def get_posts(self):
        posts_data = self.db.query(PostsModel).order_by(desc(PostsModel.post_created_at)).all()
        return posts_data
    
    def create_posts(self, post_data: PostCreate):
        new_post = PostsModel(
            id=str(uuid.uuid4()),
            post_content=post_data.post_content,
            post_author_id=post_data.post_author_id,
            post_created_at=datetime.now(timezone.utc),
            post_updated_at=datetime.now(timezone.utc),
            post_image=post_data.post_image,
            likes=0
        )

        self.db.add(new_post)
        self.db.commit()
        self.db.refresh(new_post)
    
        return new_post
    
    def edit_posts(self, post_id: str, post_data: PostCreate, current_user_id: str):
        # gaperlu update, karena session di tracking
        post = self.db.query(PostsModel).filter(PostsModel.id == post_id).first()
        if not post:
            raise NotFound(f"Post with id {post_id} not found")
        
        if post.post_author_id != current_user_id:
            raise Forbidden("You can only edit your own posts")

        post.post_content = post_data.post_content
        post.post_image = post_data.post_image
        post.post_updated_at = datetime.now(timezone.utc)

        self.db.commit()
        self.db.refresh(post)
        return post