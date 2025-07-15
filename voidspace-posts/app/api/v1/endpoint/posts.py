from fastapi import Depends, APIRouter, HTTPException
from app.api.v1 import router
from app.common.response import SuccessResponse
from app.error.error_posts import DatabaseExcetion
from app.middleware.middleware import get_user_from_gateway
from app.schemas.posts_schemas import Posts, PostCreate
from app.services.posts_service import PostService
from typing import List

router = APIRouter()  # initializer endpoint ini


@router.get("/")
async def get_posts(service: PostService = Depends()) -> SuccessResponse[List[Posts]]:
    try:
        posts = service.get_posts()
        return SuccessResponse[List[Posts]](
            success=True, message="Posts retrieved successfully", data=posts
        )
    except DatabaseExcetion as e:
        raise HTTPException(
            status_code=500, detail="Internal Server Error, error " + str(e)
        )


@router.post("/create")
async def create_post(
    post_data: PostCreate,
    service: PostService = Depends(),
    current_user: dict = Depends((get_user_from_gateway)),
) -> SuccessResponse[Posts]:
    try:
        new_posts = service.create_posts(post_data)
        return SuccessResponse[Posts](
            success=True, message="Posts Created Successfully", data=new_posts
        )
    except DatabaseExcetion as e:
        raise HTTPException(status_code=500, detail="Internal Server Error " + str(e))


# @router.patch("/edit/{post_id}")
# async def edit_post(post_id: str, post_data: PostCreate, service: PostService = Depends()):
#     try:
#         updated_post = service.

# @router.get("/posts/{post_id}")
# async def get_post(post_id: str, db: Session = Depends(get_db)):

# @router.delete("/posts/{post_id}")
# async def delete_post(

# @router.get("/users/{user_id}/posts")
# async def get_user_posts(
