from fastapi import FastAPI, Security
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
from contextlib import asynccontextmanager
from typing import AsyncGenerator

from src.app.core.config import settings
from src.app.middleware import CustomMiddleware

from src.app.routers.health import router as health_router

def create_app() -> FastAPI:
    application = FastAPI(
        title=settings.PROJECT_NAME,
        description=settings.DESCRIPTION,
        version=settings.VERSION,
        docs_url=f"{settings.API_V1_STR}/docs",
        redoc_url=f"{settings.API_V1_STR}/redoc",
        openapi_url=f"{settings.API_V1_STR}/openapi.json"
    )
    # Set up CORS
    application.add_middleware(
        CORSMiddleware,
        allow_origins=settings.BACKEND_CORS_ORIGINS,
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    application.add_middleware(CustomMiddleware)

    # Set up routers
    application.include_router(health_router)

    return application

app = create_app()
