import { defineConfig, mergeConfig } from 'vite'
import react from '@vitejs/plugin-react'
import type { UserConfig as VitestUserConfig } from 'vitest/config'; // Alias UserConfig from vitest/config

// https://vitejs.dev/config/
const viteConfig = defineConfig({
  base: '/Mangal/', // Base URL for GitHub Pages deployment
  plugins: [react()],
  root: 'public', // Set the root to the public directory
  build: {
    outDir: '../dist', // Output to a 'dist' folder outside of 'public'
  },
});

const vitestConfig: VitestUserConfig = {
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: './src/setupTests.ts',
  },
};

export default mergeConfig(viteConfig, vitestConfig);