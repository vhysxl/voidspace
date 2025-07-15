from fastapi import HTTPException, Request


async def get_user_from_gateway(request: Request):
    user_id = request.headers.get("X-User-ID")
    user_role = request.headers.get("X-User-Role")

    if not user_id:
        raise HTTPException(status_code=401, detail="Missing user authentication")

    return {"user_id": user_id, "user_role": user_role}
