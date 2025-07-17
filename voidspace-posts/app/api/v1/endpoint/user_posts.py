from fastapi import APIRouter
from fastapi import Depends, APIRouter
from app.common.response import SuccessResponse
from app.services.likes_service import LikeService
from app.schemas.posts_schemas import Posts
from app.services.posts_service import PostService


router = APIRouter()


@router.get("/{username}/liked", response_model=SuccessResponse[list[Posts]])
async def get_liked_posts(username: str, service: LikeService = Depends()):
    posts = service.get_liked_posts(username=username)  
    message = "Successfully get liked posts" if posts else "You haven't liked any post"
    return SuccessResponse(success=True, message=message, data=posts)


@router.get("/{username}/posts", response_model=SuccessResponse[list[Posts]])
async def get_user_posts(username: str, service: PostService = Depends()):
    posts = service.get_user_posts(username=username)
    message = (
        "Posts retrieved successfully" if posts else "User hasn't posted anything yet"
    )
    return SuccessResponse(success=True, message=message, data=posts)
