from datetime import datetime, timezone
from sqlalchemy import desc
import uuid

from app.models.posts_models import Posts as PostsModel
from app.schemas.posts_schemas import PostCreate, Posts
from app.core.database import SessionDb


class PostRepository:
    def __init__(self, db: SessionDb):
        self.db = db

    def get_posts(self):
        return (
            self.db.query(PostsModel).order_by(desc(PostsModel.post_created_at)).all()
        )

    def create_posts(self, *, post_data: PostCreate):
        new_post = PostsModel(
            id=str(uuid.uuid4()),
            post_content=post_data.post_content,
            post_author=post_data.username,
            post_created_at=datetime.now(timezone.utc),
            post_updated_at=datetime.now(timezone.utc),
            post_image=post_data.post_image,
        )

        self.db.add(new_post)
        self.db.commit()
        self.db.refresh(new_post)
        return new_post

    def edit_posts(self, *, post: Posts, post_data: PostCreate):
        post.post_content = post_data.post_content
        post.post_image = post_data.post_image
        post.post_updated_at = datetime.now(timezone.utc)

        self.db.commit()
        self.db.refresh(post)
        return post

    def get_post(self, *, post_id: str):
        return self.db.query(PostsModel).filter(PostsModel.id == post_id).first()

    def delete_post(self, *, post: Posts):
        self.db.delete(post)
        self.db.commit()
        return True

    def get_user_posts(self, *, author: str):
        return (
            self.db.query(PostsModel)
            .filter(PostsModel.post_author == author)
            .order_by(PostsModel.post_created_at.desc())
            .all()
        )
