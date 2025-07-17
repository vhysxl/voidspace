from datetime import datetime
from typing import Optional
from sqlalchemy import DateTime
from sqlalchemy.orm import Mapped, mapped_column, relationship
from sqlalchemy import JSON, ForeignKey
from sqlalchemy.sql import func
from app.core.database import Base
from sqlalchemy.ext.hybrid import hybrid_property


class Posts(Base):
    __tablename__ = "posts"

    id: Mapped[str] = mapped_column(primary_key=True, nullable=False)
    post_content: Mapped[str] = mapped_column(nullable=False)
    post_author: Mapped[str] = mapped_column(nullable=False, index=True)
    post_created_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True), nullable=False, default=func.now(), index=True
    )
    post_updated_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True), nullable=False, default=func.now(), onupdate=func.now()
    )
    post_image: Mapped[Optional[list[str]]] = mapped_column(JSON, nullable=True)

    posts_likes: Mapped[list["PostsLike"]] = relationship(
        "PostsLike", back_populates="post", cascade="all, delete"
    )

    @hybrid_property  # count likes automatically
    def likes(self):
        return len(self.posts_likes)


class PostsLike(Base):
    __tablename__ = "posts_likes"

    username: Mapped[str] = mapped_column(primary_key=True)
    post_id: Mapped[str] = mapped_column(ForeignKey("posts.id"), primary_key=True) # composite to avoid duplikat like
    like_created_at: Mapped[datetime] = mapped_column(
        DateTime(timezone=True), nullable=False, default=func.now()
    )

    # relationship
    post: Mapped["Posts"] = relationship("Posts", back_populates="posts_likes")
