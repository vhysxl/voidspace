from fastapi import Depends
from app.error.error_posts import DatabaseExcetion, NotFound
from app.repository.posts_repository import PostRepository
from sqlalchemy.exc import SQLAlchemyError
from app.schemas.posts_schemas import PostCreate, Posts

class PostService:
    def __init__(self, repo: PostRepository = Depends()):
        self.repo = repo

    def get_posts(self) -> list[Posts]:
        try:
            posts = self.repo.get_posts()
            return [Posts.model_validate(post) for post in posts]
        except SQLAlchemyError as e:
            raise DatabaseExcetion(f"Failed to retrieve posts: {str(e)}")
        
    def create_posts(self, post_data: PostCreate) -> Posts:
        try:
            new_post = self.repo.create_posts(post_data)
            if not new_post:
             raise DatabaseExcetion("Failed to create post")
            
            return Posts.model_validate(new_post)

        except SQLAlchemyError as e:
            raise DatabaseExcetion(f"Failed to create post: {str(e)}")
        
    def edit_posts(self, post_data: PostCreate) -> Posts:
        try:
            new_post = self.repo.edit_posts(post_data)
            if not new_post:
             raise DatabaseExcetion("Failed to create post")
            
            return Posts.model_validate(new_post)

        except SQLAlchemyError as e:
            raise DatabaseExcetion(f"Failed to create post: {str(e)}")
    