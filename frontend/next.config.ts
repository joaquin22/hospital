import type { NextConfig } from "next";
import path from 'path';
const isProd = process.env.NODE_ENV === "production";
const nextConfig: NextConfig = {
  /* config options here */
  // output: "export",  // Uncomment the following line only for building purposes. By default, this line should remain commented out.
  // reactStrictMode: true,
  trailingSlash: true,
  basePath: isProd ? "/tailwind/page/xintra-ts/preview" : undefined,
	assetPrefix : isProd ? "/tailwind/page/xintra-ts/preview" : undefined,
  images: {
    loader: "imgix",
    path: "/",
  },
  typescript: {
    ignoreBuildErrors: true,
  },
  sassOptions: {
    includePaths: [path.join(__dirname, 'public/assets/scss')],
    silenceDeprecations: ['legacy-js-api'],
    quietDeps: true, 
  },
  reactStrictMode: false, // Disable Strict Mode if necessary
  eslint: {
    ignoreDuringBuilds: true,
  },
};

export default nextConfig;
