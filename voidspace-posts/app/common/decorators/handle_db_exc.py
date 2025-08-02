from functools import wraps
from app.error.error_posts import DatabaseExcetion, NotFound
import logging


logger = logging.getLogger(__name__)


def handle_db_errors(operation_name: str):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            try:
                return func(*args, **kwargs)
            except SQLAlchemyError as e:
                logger.error(
                    f"Database error during {operation_name}: {str(e)}", exc_info=True
                )
                raise DatabaseExcetion

        return wrapper

    return decorator
