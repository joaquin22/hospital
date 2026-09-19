import Link from "next/link";
import SpkButton from "@/shared/@spk-reusable-components/uielements/spk-button";
import { Fragment, useEffect, useState } from "react";
import { Helmet } from "react-helmet";
import { useRouter } from "next/router";
import PrelineScript from "./PrelineScript";
import { authService } from "../shared/api/serivices/auth.service";
import { setToken } from "../shared/api/client";

export default function Home() {
  const [passwordshow1, setpasswordshow1] = useState(false);

  const [err, setError] = useState("");
  const [data, setData] = useState({
    email: "admin@example.com",
    password: "password",
  });

  const [loading, setLoading] = useState(false);

  const { email, password } = data;
  const changeHandler = (e: any) => {
    const { name, value } = e.target;
    setData((prevData) => ({
      ...prevData,
      [name]: value,
    }));

    setError("");
  };

  const router = useRouter();
  const RouteChange = () => {
    let path = "/components/dashboard/sales";
    router.push(path);
  };

  const Login = async (e: { preventDefault: () => void }) => {
    e.preventDefault();
    if (loading) return;

    if (!email) {
      setError("El correo electrónico es requerido.");
      return;
    }
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      setError("Por favor, ingresa un correo electrónico válido.");
      return;
    }

    if (!password) {
      setError("La contraseña es requerida.");
      return;
    }

    setLoading(true);
    setError("");

	authService
      .login({ email, password })
      .then((response) => {
			console.log(response);
	      setToken(response.data.token);
        RouteChange();
      })
      .catch((err: any) => {
        if (err.status === 401) {
          setError("Correo electrónico o contraseña incorrectos.");
        } else if (err.status === 500) {
          setError("Error en el servidor. Por favor, intenta de nuevo más tarde.");
        } else if (!err.status) {
          setError("Error de red o conexión. Verifica tu conexión a internet.");
        } else {
          setError(err.errors?.[0] || err.message || "Ha ocurrido un error inesperado.");
        }
      })
      .finally(() => {
        setLoading(false);
      });

    // try {
    //   const response = await authService.login({ email, password });
    //   setToken(response.data.token);
    //   RouteChange();
    // } catch (error: any) {
    //   setError(error.response?.data?.message ?? "Credenciales incorrectas");
    // }
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
        <body className={`${backgroundClass}`} />
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
                    {err && (
                      <div
                        className="alert-danger px-4 py-3 shadow-md mb-2"
                        role="alert"
                      >
                        <div className="flex">
                          <div className="py-1"></div>
                          <div>{err}</div>
                        </div>
                      </div>
                    )}
                  </div>
                  <div className="xl:col-span-12 col-span-12">
                    <label
                      htmlFor="signin-username"
                      className="form-label text-defaulttextcolor"
                    >
                      Usuario<sup className="text-xs text-danger">*</sup>
                    </label>
                    <input
                      type="text"
                      name="email"
                      className="form-control"
                      id="signin-username"
                      placeholder="usuario"
                      defaultValue={email}
                      onChange={changeHandler}
                      disabled={loading}
                    />
                  </div>
                  <div className="xl:col-span-12 col-span-12 mb-2">
                    <label
                      htmlFor="signin-password"
                      className="form-label text-defaulttextcolor block"
                    >
                      Contrasena<sup className="text-xs text-danger">*</sup>
                    </label>
                    <div className="relative">
                      <input
                        type="password"
                        name="password"
                        defaultValue={password}
                        onChange={changeHandler}
                        className="form-control create-password-input"
                        id="signin-password"
                        placeholder="contrasena"
                        disabled={loading}
                      />
                      <SpkButton
                        onclickfunc={() => setpasswordshow1(!passwordshow1)}
                        customClass="show-password-button text-textmuted dark:text-textmuted/50"
                        disabled={loading}
                      >
                        <i
                          className={`${passwordshow1 ? "ri-eye-line" : "ri-eye-off-line"} align-middle`}
                        ></i>
                      </SpkButton>
                    </div>
                  </div>
                </div>
                <div className="grid mt-4">
                    <button onClick={Login} className="ti-btn ti-btn-primary" disabled={loading}>
                        {loading ? "Iniciando Sesión..." : "Iniciar Sesión"}
                    </button>
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
