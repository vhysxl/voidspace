from fastapi import Depends
from app.repository.likes_repository import LikeRepository
from app.services.posts_service import PostService
from app.common.decorators.handle_db_exc import handle_db_errors
from app.error.error_posts import NotFound
from app.schemas.posts_schemas import Posts


class LikeService:
    def __init__(
        self, repo: LikeRepository = Depends(), post_service: PostService = Depends()
    ):
        self.repo = repo
        self.post_service = post_service

    @handle_db_errors("Liking post")
    def like_post(self, *, post_id: str, username: str) -> bool:
        post = self.post_service.get_post(post_id=post_id)
        result = self.repo.like_post(post_id=post.id, username=username)
        return result

    @handle_db_errors("Removing like")
    def unlike_post(self, *, post_id: str, username: str) -> bool:
        liked_post = self.repo.get_liked_post(post_id=post_id, username=username)
        if not liked_post:
            raise NotFound
        result = self.repo.unlike_post(post=liked_post)
        return result

    @handle_db_errors("Getting liked posts")
    def get_liked_posts(self, *, username: str) -> list[Posts]:
        liked_posts = self.repo.get_liked_posts(username=username)
        return [Posts.model_validate(post) for post in liked_posts]
