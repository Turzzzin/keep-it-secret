from fastapi import APIRouter, Depends
from fastapi.responses import JSONResponse
from src.app.core.logging import get_logger

logger = get_logger(__name__)

router = APIRouter()

@router.get("/health")
async def health_check():
    """"
    Health check endpoint to verify if the API is running.
    """
    logger.info("Health check endpoint called")
    return JSONResponse(
        content={
            "status":"healthy",
            "message": "Keep it Secret API is running. Let's keep it safe!"
        }
    )
