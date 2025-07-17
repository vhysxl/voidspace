from fastapi import APIRouter
from app.api.v1.endpoint import posts, user_posts

api_router = APIRouter()
api_router.include_router(posts.router, prefix="/posts", tags=["posts"])
api_router.include_router(user_posts.router, prefix="/users", tags=["users-stuff"])
