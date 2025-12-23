package config

const template = `# --- Global Settings ---
debounce: 500ms
env:
    - APP_ENV: "development" # development | staging | production

# --- Ignore List ---
# Directory to ignore to save CPU
ignore:
    - ".git"
    - "node_modules"
    - "bin"
    - "vendor"

# --- Pipelines ---
pipelines:
    # Example: Go Application
    - name: "Run go app"
      extensions: [".go", ".mod"]
      command: "go run ."
      env:
          DB_USERNAME: "root"
          DB_PASSWORD: "root@123"
          DB_HOST: "localhost"
          DB_PORT: "5050"

    # Example: Python Application
    - name: "Run Python"
      extensions: [".py"]
      command: "python main.py"
`
