from fastapi import Depends, APIRouter
from app.common.response import SuccessResponse
from app.services.likes_service import LikeService
from app.schemas.posts_schemas import Posts


router = APIRouter()


@router.post("/like", response_model=SuccessResponse)
def like_post(post_id: str, username: str, service: LikeService = Depends()):
    service.like_post(username=username, post_id=post_id)
    return SuccessResponse(success=True, message="Successfully liked post")


@router.delete("/like", response_model=SuccessResponse)
def unlike_post(post_id: str, username: str, service: LikeService = Depends()):
    service.unlike_post(post_id=post_id, username=username)
    return SuccessResponse(success=True, message="Successfully liked post")
