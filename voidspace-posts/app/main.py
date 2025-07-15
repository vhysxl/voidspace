from fastapi import FastAPI, Depends
from sqlalchemy.orm import Session
from sqlalchemy import text  # Add this import
from app.core.database import get_db
from app.api.v1.router import api_router

app = FastAPI(title="Posts Service", version="1.0.0")
app.include_router(api_router, prefix="/api/v1")

@app.get("/")
async def root():
    return {"message": "Posts API is running"}

@app.get("/health")
async def health_check(db: Session = Depends(get_db)):
    try:
        db.execute(text("SELECT 1"))
        return {"status": "healthy", "message": "Database connected"}
    except Exception as e:
        return {"status": "unhealthy", "error": str(e)}