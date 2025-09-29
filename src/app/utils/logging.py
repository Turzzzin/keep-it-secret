"""
Logging configuration for the application.
"""

import logging
import sys
import structlog
from structlog.stdlib import LoggerFactory

def configure_logging() -> None:
    """"
    Configure structured logging for the application.
    """
    logging.basicConfig(
        format="%(message)s",
        stream=sys.stdout,
        level=logging.INFO,
    )

    structlog.configure(
        processors=[
            structlog.stdlib.add_logger_name,
            structlog.stdlib.add_log_level,
            structlog.processors.TimeStamper(fmt="iso"),
            structlog.processors.StackInfoRenderer(),
            structlog.processors.format_exc_info,
            structlog.processors.UnicodeDecoder(),
            structlog.processors.JSONRenderer(),
        ],
        context_class=dict,
        logger_factory=LoggerFactory(),
        wrapper_class=structlog.stdlib.BoundLogger,
        cache_logger_on_first_use=True,
    )

def get_logger(name: str) -> structlog.stdlib.BoundLogger:
    """
    Get a structured logger by name.

    Args:
        name (str): The name of the logger.

    Returns:
        structlog.stdlib.BoundLogger: The configured structured logger.
    """
    return structlog.get_logger(name)