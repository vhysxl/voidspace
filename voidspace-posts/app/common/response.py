from typing import Any, Dict, Optional, TypeVar, Generic
from fastapi import HTTPException
from pydantic import BaseModel

T = TypeVar("T")  # generic type


class SuccessResponse(BaseModel, Generic[T]):
    success: bool = True
    message: str
    data: Optional[T] = None
