import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const isProduction = mode === 'production';

  return {
    base: '/Mangal/', // Base URL for GitHub Pages deployment
    plugins: [react()],
    // Use public as root for production builds (GitHub Pages)
    // Use standard setup for local development
    ...(isProduction && {
      root: 'public',
      build: {
        outDir: '../dist',
      },
    }),
    ...(!isProduction && {
      build: {
        outDir: 'dist',
      },
    }),
    test: {
      globals: true,
      environment: 'jsdom',
      setupFiles: './src/setupTests.ts',
    },
  };
});