from typing import TypeVar, Generic, Optional
from pydantic import BaseModel

T = TypeVar("T") # generic type

class SuccessResponse(BaseModel, Generic[T]):
    success: bool = True
    message: str
    data: T

