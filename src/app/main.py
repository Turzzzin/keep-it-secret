from fastapi import FastAPI, Security
from fastapi.middleware
from fastapi.responses import JSONResponse
from contextlib import asynccontextmanager
from typing import AsyncGenerator

from src.app.core.config import settings

from src.app.routers.health import router as health_router

def create_app() -> FastAPI:
    application = FastAPI(
        title=settings.PROJECT_NAME,
        description=settings.DESCRIPTION,
        version=settings.VERSION,
    )

