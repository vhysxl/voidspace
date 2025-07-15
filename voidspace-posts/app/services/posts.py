from typing import List
from fastapi import Depends
from app.error.error_posts import DatabaseExcetion, NotFound
from app.repository.posts_repository import PostRepository
from sqlalchemy.exc import SQLAlchemyError
from app.schemas.posts_schemas import PostCreate, Posts


class PostService:
    def __init__(self, repo: PostRepository = Depends()):
        self.repo = repo

    def get_posts(self):
        try:
            post: List[Posts] = self.repo.get_posts()
            return post
        except SQLAlchemyError as e:
            raise DatabaseExcetion(f"Failed to retrieve posts: {str(e)}")
        
    def create_posts(self, post_data: PostCreate):
        try:
            new_post = self.repo.create_posts(post_data)
            if new_post:
                return Posts.model_validate(new_post)

        except SQLAlchemyError as e:
            raise DatabaseExcetion(f"Failed to create post: {str(e)}")
        