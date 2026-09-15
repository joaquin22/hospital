import Link from "next/link";
import { Fragment, useEffect, useState } from "react";
import { Helmet } from "react-helmet";
import { useRouter } from "next/router";
import PrelineScript from "./PrelineScript";

export default function Home() {

    const [err, setError] = useState("");
    const [data, setData] = useState({
        "email": "adminnextjs@gmail.com",
        "password": "1234567890",
    });
    const { email, password } = data;
    const changeHandler = (e: any) => {
        setData({ ...data, [e.target.name]: e.target.value });
        setError("");
    };

    const router = useRouter();
    const RouteChange = () => {
        let path = "/components/dashboard/sales";
        router.push(path);
    };
    const Login1 = (_e: any) => {
        if (data.email == "adminnextjs@gmail.com" && data.password == "1234567890") {
            RouteChange();
        }
        else {
            setError("The Auction details did not Match");
            setData({
                "email": "adminnextjs@gmail.com",
                "password": "1234567890",
            });
        }
    };

      const location = useRouter();

     const [backgroundClass, setBackgroundClass] = useState("");

     useEffect(() => {
       if (location.pathname === "/") {
         setBackgroundClass("authentication-background");
       } else {
         setBackgroundClass("");
       }
     }, [location.pathname]);

    return (
        <Fragment>
            <PrelineScript />
            <Helmet>
                    <body className={`${backgroundClass}`}/>
            </Helmet>
            <div className="container">
                <div className="grid grid-cols-12 justify-center items-center authentication authentication-basic h-full">
                    <div className="xxl:col-span-4 xl:col-span-4 lg:col-span-3 md:col-span-3 sm:col-span-2 col-span-12"></div>
                    <div className="xxl:col-span-4 xl:col-span-4 lg:col-span-6 md:col-span-6 sm:col-span-8 col-span-12">
                        <div className="box my-4">
                            <div className="box-body !p-[3rem]">
                                <p className="h5 mb-2 text-center">Iniciar Sesion</p>
                                <div className="grid grid-cols-12 gap-y-3">
                                    <div className="xl:col-span-12 col-span-12">
                                        {err && <div className="alert-danger px-4 py-3 shadow-md mb-2" role="alert">
                                            <div className="flex">
                                                <div className="py-1">
                                                </div>
                                                <div>{err}</div>
                                            </div>
                                        </div>}
                                    </div>
                                    <div className="xl:col-span-12 col-span-12">
                                        <label htmlFor="signin-username" className="form-label text-defaulttextcolor">Usuario<sup className="text-xs text-danger">*</sup></label>
                                        <input type="text" name="email" className="form-control" id="signin-username" placeholder="usuario" defaultValue={email} onChange={changeHandler} />
                                    </div>
                                    <div className="xl:col-span-12 col-span-12 mb-2">
                                        <label htmlFor="signin-password" className="form-label text-defaulttextcolor block">Contrasena<sup className="text-xs text-danger">*</sup></label>
                                        <div className="relative">
                                            <input type="password" name="password" defaultValue={password} onChange={changeHandler} className="form-control create-password-input" id="signin-password" placeholder="contrasena" />
                                        </div>
                                    </div>
                                </div>
                                <div className="grid mt-4">
                                    <Link href="/components/dashboard/sales" className="ti-btn ti-btn-primary" onClick={Login1}>Iniciar Sesion</Link>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="xxl:col-span-4 xl:col-span-4 lg:col-span-3 md:col-span-3 sm:col-span-2 col-span-12"></div>
                </div>
            </div>
        </Fragment>
    );
}
