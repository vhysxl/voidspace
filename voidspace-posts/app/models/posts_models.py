from datetime import datetime
from typing import Optional
from sqlalchemy import DateTime
from sqlalchemy.orm import Mapped, mapped_column, relationship
from sqlalchemy import JSON, ForeignKey
from sqlalchemy.sql import func
from app.core.database import Base


class Posts(Base):
    __tablename__ = "posts"

    id: Mapped[str] = mapped_column(primary_key=True, nullable=False)
    post_content: Mapped[str] = mapped_column(nullable=False)
    post_author_id: Mapped[str] = mapped_column(nullable=False)
    post_created_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True), nullable=False, default=func.now()
    )
    post_updated_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True), nullable=False, default=func.now(), onupdate=func.now()
    )
    post_image: Mapped[Optional[list[str]]] = mapped_column(JSON, nullable=True)
    likes: Mapped[int] = mapped_column(default=0, nullable=False)
    posts_likes: Mapped[list["PostsLike"]] = relationship(
        "PostsLike", back_populates="post", cascade="all, delete"
    )


class PostsLike(Base):
    __tablename__ = "posts_likes"

    user_id: Mapped[str] = mapped_column(primary_key=True)
    post_id: Mapped[str] = mapped_column(ForeignKey("posts.id"), primary_key=True)
    like_created_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True), nullable=False, default=func.now()
    )

    # relationship
    post: Mapped["Posts"] = relationship("Posts", back_populates="posts_likes")
