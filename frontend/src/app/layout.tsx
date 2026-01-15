import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Header } from "@/ui/header";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Infomedia VM Hub",
  description: "PT Infomedia Nusantara - VM Rental Management & Pricing System",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="min-h-screen">
          <Header />
          <div className="flex flex-col m-6 min-h-[90vh]">{children}</div>
        </div>
      </body>
    </html>
  );
}
