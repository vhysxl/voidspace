import type { Metadata } from "next";
import { Space_Grotesk, Manrope } from "next/font/google";
import ClientThemeProvider from "@/components/layout/ClientThemeProvider";
import "./globals.css";

const spaceGrotesk = Space_Grotesk({
  variable: "--font-space-grotesk",
  subsets: ["latin"],
});

const manrope = Manrope({
  variable: "--font-manrope",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "VOIDSPACE",
  description: "Join the celestial conversation",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html
      lang="en"
      className={`${spaceGrotesk.variable} ${manrope.variable} h-full antialiased`}
      suppressHydrationWarning
    >
      <body className="min-h-full flex flex-col font-manrope">
        <ClientThemeProvider>{children}</ClientThemeProvider>
      </body>
    </html>
  );
}
