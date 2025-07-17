from fastapi import Depends
from app.error.error_posts import NotFound
from app.repository.posts_repository import PostRepository
from app.schemas.posts_schemas import PostCreate, Posts
from app.common.decorators.handle_db_exc import handle_db_errors


class PostService:
    def __init__(self, repo: PostRepository = Depends()):
        self.repo = repo

    @handle_db_errors("Retrieving posts")
    def get_posts(self) -> list[Posts]:
        posts = self.repo.get_posts()

        return [Posts.model_validate(post) for post in posts]

    @handle_db_errors("Creating post")
    def create_posts(self, *, post_data: PostCreate) -> Posts:
        new_post = self.repo.create_posts(post_data=post_data)

        return Posts.model_validate(new_post)

    @handle_db_errors("Editing post")
    def edit_posts(self, *, post_data: PostCreate, post_id: str) -> Posts:
        post = self.repo.get_post(post_id=post_id)
        if not post:
            raise NotFound
        edited_post = self.repo.edit_posts(post_data=post_data, post=post)
        return Posts.model_validate(edited_post)

    @handle_db_errors("Retrieving post")
    def get_post(self, *, post_id: str) -> Posts:
        post = self.repo.get_post(post_id=post_id)
        if not post:
            raise NotFound
        return Posts.model_validate(post)

    @handle_db_errors("Deleting post")
    def delete_post(self, *, post_id: str) -> bool:
        post = self.repo.get_post(post_id=post_id)
        if not post:
            raise NotFound
        result = self.repo.delete_post(post=post)
        return result

    @handle_db_errors("Retrieving posts")
    def get_user_posts(self, *, username: str) -> list[Posts]:
        posts = self.repo.get_user_posts(author=username)
        return [Posts.model_validate(post) for post in posts]
