import { useLanguage } from "@/shared/i18n/LanguageContext";
import Seo from "@/shared/layouts-components/seo/seo";
import { userService } from "@/shared/api/serivices/user.service";
import Link from "next/link";
import { useRouter } from "next/router";
import { Fragment, useState } from "react";

const NuevoUsuario = () => {
	const router = useRouter();
	const { t } = useLanguage();

	const [data, setData] = useState({
		firstName: "",
		lastName: "",
		dni: "",
		email: "",
		password: "",
		role: "",
	});

	const [errors, setErrors] = useState<Record<string, string>>({});

	const changeHandler = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
		setData({ ...data, [e.target.name]: e.target.value });
		setErrors({ ...errors, [e.target.name]: "" });
	};

	const validate = () => {
		const newErrors: Record<string, string> = {};
		if (!data.firstName.trim()) newErrors.firstName = t("formulario.errorNombre");
		if (!data.lastName.trim()) newErrors.lastName = t("formulario.errorApellido");
		if (!data.dni.trim()) newErrors.dni = t("formulario.errorDni");
		if (!data.email.trim()) newErrors.email = t("formulario.errorCorreo");
		if (!data.password.trim()) newErrors.password = t("usuarios.errorPassword");
		if (!data.role.trim()) newErrors.role = t("usuarios.errorRol");
		setErrors(newErrors);
		return Object.keys(newErrors).length === 0;
	};

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		if (!validate()) return;
		try {
			await userService.create({
				first_name: data.firstName,
				last_name: data.lastName,
				dni: data.dni,
				email: data.email,
				password: data.password,
				role: data.role,
			});
			router.push("/components/usuarios");
		} catch {
			setErrors({ form: "error al crear el usuario" });
		}
	};

	return (
		<Fragment>
			<Seo title={t("usuarios.formularioTitulo")} />
			<div className="flex items-center justify-between page-header-breadcrumb flex-wrap gap-2">
				<div>
					<ol className="breadcrumb mb-0">
						<li className="breadcrumb-item">
							<Link scroll={false} href="#!">
								{t("usuarios.breadcrumbApp")}
							</Link>
						</li>
						<li className="breadcrumb-item">
							<Link scroll={false} href="/components/usuarios">
								{t("usuarios.formularioBreadcrumb")}
							</Link>
						</li>
						<li className="breadcrumb-item active" aria-current="page">{t("usuarios.formularioNuevo")}</li>
					</ol>
					<h1 className="page-title font-medium text-lg mb-0">{t("usuarios.formularioTitulo")}</h1>
				</div>
			</div>
			<div className="grid grid-cols-12 gap-x-6">
				<div className="xl:col-span-12 col-span-12">
					<div className="box my-4">
						<div className="box-header">
							<div className="box-title">{t("usuarios.boxTitle")}</div>
						</div>
						<div className="box-body !p-[2rem]">
							<form onSubmit={handleSubmit}>
								<div className="grid grid-cols-12 gap-y-3">
									<div className="xl:col-span-12 col-span-12">
										<label htmlFor="firstName" className="form-label text-defaulttextcolor">
											{t("formulario.nombre")}<sup className="text-xs text-danger">*</sup>
										</label>
										<input
											type="text"
											name="firstName"
											className="form-control"
											id="firstName"
											placeholder={t("formulario.placeholderNombre")}
											value={data.firstName}
											onChange={changeHandler}
										/>
										{errors.firstName && (
											<span className="text-danger text-xs mt-1">{errors.firstName}</span>
										)}
									</div>
									<div className="xl:col-span-12 col-span-12">
										<label htmlFor="lastName" className="form-label text-defaulttextcolor">
											{t("formulario.apellido")}<sup className="text-xs text-danger">*</sup>
										</label>
										<input
											type="text"
											name="lastName"
											className="form-control"
											id="lastName"
											placeholder={t("formulario.placeholderApellido")}
											value={data.lastName}
											onChange={changeHandler}
										/>
										{errors.lastName && (
											<span className="text-danger text-xs mt-1">{errors.lastName}</span>
										)}
									</div>
									<div className="xl:col-span-12 col-span-12">
										<label htmlFor="dni" className="form-label text-defaulttextcolor">
											{t("formulario.dni")}<sup className="text-xs text-danger">*</sup>
										</label>
										<input
											type="text"
											name="dni"
											className="form-control"
											id="dni"
											placeholder={t("formulario.placeholderDni")}
											value={data.dni}
											onChange={changeHandler}
										/>
										{errors.dni && (
											<span className="text-danger text-xs mt-1">{errors.dni}</span>
										)}
									</div>
									<div className="xl:col-span-12 col-span-12">
										<label htmlFor="email" className="form-label text-defaulttextcolor">
											{t("formulario.correo")}<sup className="text-xs text-danger">*</sup>
										</label>
										<input
											type="email"
											name="email"
											className="form-control"
											id="email"
											placeholder={t("formulario.placeholderCorreo")}
											value={data.email}
											onChange={changeHandler}
										/>
										{errors.email && (
											<span className="text-danger text-xs mt-1">{errors.email}</span>
										)}
									</div>
									<div className="xl:col-span-12 col-span-12">
										<label htmlFor="password" className="form-label text-defaulttextcolor">
											{t("usuarios.password")}<sup className="text-xs text-danger">*</sup>
										</label>
										<input
											type="password"
											name="password"
											className="form-control"
											id="password"
											placeholder={t("usuarios.placeholderPassword")}
											value={data.password}
											onChange={changeHandler}
										/>
										{errors.password && (
											<span className="text-danger text-xs mt-1">{errors.password}</span>
										)}
									</div>
									<div className="xl:col-span-12 col-span-12">
										<label htmlFor="role" className="form-label text-defaulttextcolor">
											{t("usuarios.rol")}<sup className="text-xs text-danger">*</sup>
										</label>
										<select
											name="role"
											className="form-control"
											id="role"
											value={data.role}
											onChange={changeHandler}
										>
											<option value="">{t("usuarios.placeholderRol")}</option>
											<option value="admin">{t("usuarios.rolAdmin")}</option>
											<option value="staff">{t("usuarios.rolRecepcionista")}</option>
										</select>
										{errors.role && (
											<span className="text-danger text-xs mt-1">{errors.role}</span>
										)}
									</div>
								</div>
								{errors.form && (
									<span className="text-danger text-xs mt-2">{errors.form}</span>
								)}
								<div className="flex gap-2 mt-4">
									<button type="submit" className="ti-btn ti-btn-primary btn-wave">
										<i className="ri-save-line align-middle me-1"></i>{t("usuarios.guardar")}
									</button>
									<Link scroll={false} href="/components/usuarios" className="ti-btn ti-btn-light btn-wave">
										{t("usuarios.cancelar")}
									</Link>
								</div>
							</form>
						</div>
					</div>
				</div>
			</div>
		</Fragment>
	);
};

NuevoUsuario.layout = "Contentlayout";
export default NuevoUsuario;