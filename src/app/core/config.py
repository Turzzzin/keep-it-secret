"""Application configuration Settings"""

from pathlib import Path
from typing import Dict, List, Optional

from dotenv import load_dotenv
from pydantic import AnyHttpUrl, computed_field, Field
from pydantic_settings import BaseSettings, SettingsConfigDict

env_file = Path(__file__).parent.parent.parent / ".env"
if env_file.exists():
    load_dotenv(env_file) 

class Settings(BaseSettings):
    PROJECT_NAME: str = Field(default="Keep it Secret API", env="PROJECT_NAME")
    DESCRIPTION: str = Field(default="An API for managing secrets securely.", env="DESCRIPTION")
    VERSION: str = Field(default="1.0.0", env="VERSION")
    API_V1_STR: str = Field(default="/api/v1", env="API_V1_STR")


settings = Settings()