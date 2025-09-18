# Keep It Secret API
```JSON
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "FastAPI (Uvicorn)",
      "type": "debugpy",
      "request": "launch",
      "module": "uvicorn",
      "args": [
        "app.main:app",
        "--reload"
      ],
      "jinja": true
    }
  ]
}
```