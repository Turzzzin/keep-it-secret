"""
Configuration settings for the application.
Loads environment variables from a .env file if it exists.
"""

from pathlib import Path
from typing import List
from dotenv import load_dotenv
from pydantic import AnyHttpUrl, Field
from pydantic_settings import BaseSettings, SettingsConfigDict

env_file = Path(__file__).parent.parent.parent / ".env"
if env_file.exists():
    load_dotenv(env_file)

class Settings(BaseSettings):
    model_config = SettingsConfigDict(env_file=env_file)

    PROJECT_NAME: str = Field(default="Keep it Secret API", alias="PROJECT_NAME")
    DESCRIPTION: str = Field(default="An API for managing secrets securely.", alias="DESCRIPTION")
    VERSION: str = Field(default="1.0.0", alias="VERSION")
    API_V1_STR: str = Field(default="/api/v1", alias="API_V1_STR")

    BACKEND_CORS_ORIGINS: List[str | AnyHttpUrl] = ["*"]

settings = Settings()
