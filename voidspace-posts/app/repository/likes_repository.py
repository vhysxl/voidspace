from datetime import datetime, timezone
from app.core.database import SessionDb
from app.models.posts_models import Posts, PostsLike


class LikeRepository:
    def __init__(self, db: SessionDb):
        self.db = db

    def like_post(self, *, post_id: str, username: str):
        liked_post = PostsLike(
            post_id=post_id,
            username=username,
            like_created_at=datetime.now(timezone.utc),
        )

        self.db.add(liked_post)
        self.db.commit()
        return True

    def get_liked_post(self, *, post_id: str, username: str):
        return (
            self.db.query(PostsLike)
            .filter(PostsLike.post_id == post_id, PostsLike.username == username)
            .first()
        )

    def unlike_post(self, *, post: Posts):
        self.db.delete(post)
        self.db.commit()
        return True

    def get_liked_posts(self, *, username: str):
        return (
            self.db.query(Posts)
            .join(PostsLike, Posts.id == PostsLike.post_id)
            .filter(PostsLike.username == username)
        ).all()
