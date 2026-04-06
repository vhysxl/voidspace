import { NextRequest, NextResponse } from "next/server";

const BACKEND_URL = process.env.INTERNAL_API_URL || "http://localhost:5000/api/v2";
const API_KEY = process.env.INTERNAL_API_KEY;

export async function GET(request: NextRequest, { params }: { params: Promise<{ path?: string[] }> }) {
  return proxyRequest(request, await params);
}

export async function POST(request: NextRequest, { params }: { params: Promise<{ path?: string[] }> }) {
  return proxyRequest(request, await params);
}

export async function PUT(request: NextRequest, { params }: { params: Promise<{ path?: string[] }> }) {
  return proxyRequest(request, await params);
}

export async function PATCH(request: NextRequest, { params }: { params: Promise<{ path?: string[] }> }) {
  return proxyRequest(request, await params);
}

export async function DELETE(request: NextRequest, { params }: { params: Promise<{ path?: string[] }> }) {
  return proxyRequest(request, await params);
}

async function proxyRequest(request: NextRequest, params: { path?: string[] }) {
  if (!API_KEY) {
    return NextResponse.json({ detail: "INTERNAL_API_KEY is not defined on server" }, { status: 500 });
  }

  const path = params.path?.join("/") || "";
  const searchParams = request.nextUrl.searchParams.toString();
  const backendUrl = `${BACKEND_URL}/${path}${searchParams ? `?${searchParams}` : ""}`;

  const headers = new Headers(request.headers);
  headers.set("x-api-key", API_KEY);
  // Remove host header to avoid issues with Cloud Run / proxying
  headers.delete("host");

  try {
    const body = ["GET", "HEAD"].includes(request.method) ? undefined : await request.arrayBuffer();

    const response = await fetch(backendUrl, {
      method: request.method,
      headers,
      body,
      cache: "no-store",
    });

    const data = await response.arrayBuffer();
    
    // Create a new response with the backend's data and status
    const proxyResponse = new NextResponse(data, {
      status: response.status,
      statusText: response.statusText,
      headers: response.headers,
    });

    // Forward cookies if any
    const setCookie = response.headers.get("set-cookie");
    if (setCookie) {
      proxyResponse.headers.set("set-cookie", setCookie);
    }

    return proxyResponse;
  } catch (error) {
    console.error("BFF Proxy Error:", error);
    return NextResponse.json({ detail: "Failed to connect to backend service" }, { status: 502 });
  }
}
