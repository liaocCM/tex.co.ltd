import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: {
    host: "0.0.0.0", 
    port: 5173,
    allowedHosts: ["42dc55684645.ngrok-free.app"],
    proxy: {
      "/api": "http://127.0.0.1:8080"
    }
  }
});
