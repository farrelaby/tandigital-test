import "@/styles/globals.css";
import type { AppProps } from "next/app";

import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { atom, Provider as JotaiProvider } from "jotai";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

export default function App({ Component, pageProps }: AppProps) {
  const [queryClient] = useState(() => new QueryClient());
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <JotaiProvider>
          <Component {...pageProps} />
          <ReactQueryDevtools initialIsOpen={false} />
        </JotaiProvider>
      </QueryClientProvider>
    </>
  );
}
