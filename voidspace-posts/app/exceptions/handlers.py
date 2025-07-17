from fastapi import Request
from fastapi.responses import JSONResponse
from app.error.error_posts import DatabaseExcetion, Forbidden, NotFound


# todo: review
async def database_exception_handler(request: Request, exc: DatabaseExcetion):
    return JSONResponse(status_code=500, content={"detail": "Internal Server Error"})


async def not_found_exception_handler(request: Request, exc: NotFound):
    return JSONResponse(status_code=404, content={"detail": "User or Post not found"})


async def unauthorized_exception(request: Request, exc: Forbidden):
    return JSONResponse(status_code=403, content={"detail": "Unathorized"})
