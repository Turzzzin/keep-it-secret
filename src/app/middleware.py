from fastapi import Request, Response
from starlette.middleware.base import BaseHTTPMiddleware

class CustomMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        client_host = request.client.host
        if client_host in ("127.0.0.1", "::1", "localhost"):
            response: Response = await call_next(request)
            return response
        response: Response = await call_next(request)
        return response
