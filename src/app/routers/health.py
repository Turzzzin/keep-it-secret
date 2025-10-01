from fastapi import APIRouter, Depends
from fastapi.responses import JSONResponse

router = APIRouter()

@router.get("/health")
async def health_check():
    """"
    Health check endpoint to verify if the API is running.
    """
    return JSONResponse(
        content={
            "status":"healthy",
            "message": "Keep it Secret API is running. Let's keep it safe!"
        }
    )
