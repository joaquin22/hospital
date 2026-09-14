
import { useEffect } from "react";

import { IStaticMethods } from "preline/preline";
import { useRouter } from "next/router";
declare global {
  interface Window {
    HSStaticMethods: IStaticMethods;
  }
}

export default function PrelineScript() {
  const location = useRouter();

  useEffect(() => {
    const loadPreline = async () => {
      await import("preline/preline");

      window.HSStaticMethods.autoInit();
    };

    loadPreline();
  }, [location.pathname]);

  return null;
}