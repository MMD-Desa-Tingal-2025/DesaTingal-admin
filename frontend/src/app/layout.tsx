import type { Metadata } from "next";
import { lora } from "@/libs/fonts";
import "./globals.css";

export const metadata: Metadata = {
  title: "Desa Tingal",
  description: "Data Masyarakat Desa Tingal",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${lora.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
