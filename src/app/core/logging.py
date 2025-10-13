"""
Logging configuration for the application.
Provides structured logging via structlog, with JSON output suitable for
production and optional plain-text output for development.
"""

import logging
import sys
import structlog
from structlog.stdlib import LoggerFactory
from structlog.types import EventDict, Processor


def _normalize_log_level(_: logging.Logger, __: str, event_dict: EventDict) -> EventDict:
    """
    Normalize the log level to lowercase and strip whitespace.
    Prevents padded output like '[info     ]'.
    """
    level = event_dict.get("level")
    if isinstance(level, str):
        event_dict["level"] = level.strip().lower()
    return event_dict


def configure_logging(json_output: bool = True) -> None:
    """
    Configure structured logging for the application.

    Args:
        json_output (bool): If True, use JSON logs (production).
                            If False, use human-readable console logs (dev).
    """
    # --- Standard logging setup (for stdlib compatibility) ---
    logging.basicConfig(
        format="%(asctime)s [%(levelname)s] %(message)s",
        stream=sys.stdout,
        level=logging.INFO,
    )

    # --- Select renderer based on environment ---
    if json_output:
        renderer: Processor = structlog.processors.JSONRenderer()
    else:
        # Human-friendly, colorized logs with no padding
        renderer = structlog.dev.ConsoleRenderer(pad_event=None)

    # --- Configure structlog ---
    structlog.configure(
        processors=[
            structlog.stdlib.add_logger_name,
            structlog.stdlib.add_log_level,
            _normalize_log_level,  # ✅ Removes padding/whitespace
            structlog.processors.TimeStamper(fmt="iso"),
            structlog.processors.StackInfoRenderer(),
            structlog.processors.format_exc_info,
            structlog.processors.UnicodeDecoder(),
            renderer,
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


# Example usage (optional):
# if __name__ == "__main__":
#     configure_logging(json_output=False)
#     log = get_logger(__name__)
#     log.info("Health check endpoint called", status="ok")
