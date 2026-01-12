import http.server
import socketserver
import json

# --- Configuration ---
HOST = "0.0.0.0"  # Listen on all available interfaces
PORT = 8989       # Use a standard, non-privileged port

# --- JSON Response Data ---
RESPONSE_DATA = {"message": "present day, present time"}
JSON_RESPONSE = json.dumps(RESPONSE_DATA).encode('utf-8')
CONTENT_LENGTH = str(len(JSON_RESPONSE))

# --- Custom Request Handler ---
class SimpleRequestHandler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        """Handle GET requests."""
        if self.path == '/hello-world' or self.path == '/':
            # 1. Send the response header
            self.send_response(200)
            self.send_header("Content-type", "application/json")
            self.send_header("Content-Length", CONTENT_LENGTH)
            self.end_headers()

            # 2. Send the JSON body
            self.wfile.write(JSON_RESPONSE)
        else:
            # Handle all other paths (404 Not Found)
            self.send_response(404)
            self.send_header("Content-type", "text/plain")
            self.end_headers()
            self.wfile.write("404 Not Found".encode('utf-8'))

# --- Server Execution ---
try:
    # Set up the server
    with socketserver.TCPServer((HOST, PORT), SimpleRequestHandler) as httpd:
        print(f"--- Server started on http://{HOST}:{PORT} ---")
        print("Test route: http://localhost:8000/hello-world")
        print("Press Ctrl+C to stop.")
        # 
        
        # Start serving requests indefinitely
        httpd.serve_forever()

except KeyboardInterrupt:
    print("\nServer stopped by user.")
except PermissionError:
    print(f"\nError: Could not bind to port {PORT}. You might need administrator privileges or choose a port > 1024.")
except Exception as e:
    print(f"\nAn error occurred: {e}")
