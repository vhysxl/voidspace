from fastapi import Depends, APIRouter
from app.common.response import SuccessResponse
from app.middleware.middleware import get_user_from_gateway
from app.schemas.posts_schemas import Posts, PostCreate
from app.services.posts_service import PostService
from typing import List
from app.api.v1.endpoint import likes

router = APIRouter()

# Pydantic model => Body parameter
# primitif(str, int, bool, float) => Query parameter
# Path variables with {}  => Path parameter


@router.get("", response_model=SuccessResponse[List[Posts]])
def get_posts(service: PostService = Depends()):
    posts = service.get_posts()
    return SuccessResponse(
        success=True, message="Posts retrieved successfully", data=posts
    )


@router.post("/", response_model=SuccessResponse[Posts])
def create_post(
    post_data: PostCreate,
    service: PostService = Depends(),
    # current_user: dict = Depends(get_user_from_gateway),
):
    new_posts = service.create_posts(post_data=post_data)
    return SuccessResponse(
        success=True, message="Posts Created Successfully", data=new_posts
    )


@router.get("/{post_id}", response_model=SuccessResponse[Posts])
def get_post(post_id: str, service: PostService = Depends()):
    post = service.get_post(post_id=post_id)
    return SuccessResponse(
        success=True, message="Post retrieved successfully", data=post
    )


@router.patch("/{post_id}", response_model=SuccessResponse[Posts])
def edit_post(post_id: str, post_data: PostCreate, service: PostService = Depends()):
    updated_post = service.edit_posts(post_data=post_data, post_id=post_id)
    return SuccessResponse(
        success=True, message="Post edited successfully", data=updated_post
    )


@router.delete("/{post_id}", response_model=SuccessResponse)
def delete_post(post_id: str, service: PostService = Depends()):
    service.delete_post(post_id=post_id)
    return SuccessResponse(success=True, message="Post Deleted successfully")


router.include_router(
    likes.router, prefix="/{post_id}"
)  # subrouter ke like jadi: /posts/{post_id}/like
