import "../styles/globals.scss";
import { Provider } from "react-redux";
import store from "@/shared/redux/store";
import {  Suspense, useState } from "react";
import dynamic from "next/dynamic";
const RootWrapper = dynamic(() => import("@/shared/layouts-components/layouts/Rootwrapper"), { ssr: false });
const Authenticationlayout = dynamic(() => import("@/shared/layouts-components/layouts/authenticationlayout"), { ssr: false });
const Contentlayout = dynamic(() => import("@/shared/layouts-components/layouts/contentlayout"), { ssr: false });
import { Initialload } from "@/shared/layouts-components/contextapi";
import Loader from "@/shared/layouts-components/loader/loader";

const layouts:any = {
	Contentlayout: Contentlayout,
	Authenticationlayout: Authenticationlayout,
};

function App({ Component, pageProps }:any) {
	const [pageloading, setpageloading] = useState(false)
	 
	const Layout: any = layouts[Component.layout] || ((props: any) => <Component {...props} />);
	
	return (
		<Provider store={store}>
				<Initialload.Provider value={{ pageloading, setpageloading }}>
						<RootWrapper>
								<Suspense fallback={<Loader/>}>
									<Layout>
										<Component {...pageProps} />
									</Layout>
								</Suspense>
						</RootWrapper>
				</Initialload.Provider>
		</Provider>
	);
}
export default App;
