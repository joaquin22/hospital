import { NextResponse, type NextRequest } from "next/server";

const TOKEN_KEY = "hospital_token";

export function middleware(request: NextRequest) {
  const isLoginPage = request.nextUrl.pathname === "/";
  const token = request.cookies.get(TOKEN_KEY);

  if (isLoginPage && token?.value) {
    return NextResponse.redirect(new URL("/components/doctores/", request.url));
  }

  if (!isLoginPage && !token?.value) {
    return NextResponse.redirect(new URL("/", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico|assets).*)"],
};
