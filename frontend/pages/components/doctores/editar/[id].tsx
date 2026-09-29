import { useLanguage } from "@/shared/i18n/LanguageContext";
import { resolveApiError } from "@/shared/api/errors";
import { doctorService } from "@/shared/api/serivices/doctor.service";
import { specialityService } from "@/shared/api/serivices/speciality.service";
import type { SpecialityOutput } from "@/shared/api/types/specialty";
import Seo from "@/shared/layouts-components/seo/seo";
import Link from "next/link";
import { useRouter } from "next/router";
import { Fragment, useEffect, useState } from "react";
import { notifier } from "@/utils/notifier";

// Campos del formulario que también existen en el payload del backend, para
// poder colgar el error del input correcto en vez de mostrarlo genérico.
const EDITABLE_FIELDS = ["dni", "email"];

const EditarDoctor = () => {
	const router = useRouter();
	const { t } = useLanguage();
	const id = router.query.id as string | undefined;

	const [data, setData] = useState({
		firstName: "",
		lastName: "",
		dni: "",
		email: "",
		licenseNumber: "",
		specialityId: "",
		// "true" / "false" como texto porque el select devuelve string; el
		// payload lo convierte a boolean.
		active: "",
	});

	const [errors, setErrors] = useState<Record<string, string>>({});
	const [loading, setLoading] = useState(true);
	const [saving, setSaving] = useState(false);

	const [especialidades, setEspecialidades] = useState<SpecialityOutput[]>([]);
	const [cargandoEspecialidades, setCargandoEspecialidades] = useState(true);
	const [errorEspecialidades, setErrorEspecialidades] = useState("");

	useEffect(() => {
		specialityService
			.list()
			.then((response) => {
				setEspecialidades(response.data);
				setErrorEspecialidades("");
			})
			.catch(() => setErrorEspecialidades(t("doctores.errorCargarEspecialidades")))
			.finally(() => setCargandoEspecialidades(false));
		// `t` no va en las dependencias a propósito: el provider recrea la
		// función en cada render y eso dispararía la petición otra vez.
	}, []);

	useEffect(() => {
		if (!id) return;
		setLoading(true);
		doctorService
			.get(Number(id))
			.then((response) => {
				setData({
					firstName: response.data.first_name,
					lastName: response.data.last_name,
					dni: response.data.dni,
					email: response.data.email,
					licenseNumber: response.data.license_number,
					specialityId: String(response.data.speciality_id ?? ""),
					active: String(response.data.active),
				});
				setErrors({});
			})
			.catch((error) => {
				const apiError = resolveApiError(error, t, "doctores.errorCargar");
				setErrors({ form: apiError.message });
			})
			.finally(() => setLoading(false));
		// `t` no va en las dependencias a propósito: el provider recrea la
		// función en cada render y eso dispararía la petición otra vez.
	}, [id]);

	const changeHandler = (
		e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
	) => {
		setData({ ...data, [e.target.name]: e.target.value });
		setErrors({ ...errors, [e.target.name]: "", form: "" });
	};

	const validate = () => {
		const newErrors: Record<string, string> = {};
		if (!data.firstName.trim()) newErrors.firstName = t("formulario.errorNombre");
		if (!data.lastName.trim()) newErrors.lastName = t("formulario.errorApellido");
		if (!data.dni.trim()) {
			newErrors.dni = t("formulario.errorDni");
		} else if (!/^\d{8}$/.test(data.dni.trim())) {
			newErrors.dni = t("formulario.errorDniInvalido");
		}
		if (!data.email.trim()) newErrors.email = t("formulario.errorCorreo");
		if (!data.licenseNumber.trim())
			newErrors.licenseNumber = t("formulario.errorLicencia");
		if (!data.specialityId.trim() || Number.isNaN(Number(data.specialityId)))
			newErrors.specialityId = t("doctores.errorEspecialidad");
		// También protege de guardar antes de que termine el GET: sin esto un
		// "active" todavía vacío se mandaría como false y dada de baja al
		// doctor por accidente.
		if (!data.active) newErrors.active = t("formulario.errorEstado");
		setErrors(newErrors);
		return Object.keys(newErrors).length === 0;
	};

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		if (!id || !validate()) return;
		setSaving(true);
		try {
			// PATCH: solo se envían los campos del formulario, el rol y la
			// contraseña se conservan porque no llegan en el payload.
			await doctorService.update(Number(id), {
				first_name: data.firstName.trim(),
				last_name: data.lastName.trim(),
				dni: data.dni.trim(),
				email: data.email.trim(),
				license_number: data.licenseNumber.trim(),
				speciality_id: Number(data.specialityId),
				// El backend sincroniza con este campo la ficha del doctor y su
				// cuenta de usuario.
				active: data.active === "true",
			});
			notifier.success(t("doctores.actualizado"));
			router.push("/components/doctores");
		} catch (error) {
			const apiError = resolveApiError(
				error,
				t,
				"doctores.errorActualizar",
			);
			setErrors(
				apiError.field && EDITABLE_FIELDS.includes(apiError.field)
					? { [apiError.field]: apiError.message }
					: { form: apiError.message },
			);
		} finally {
			setSaving(false);
		}
	};

	return (
		<Fragment>
			<Seo title={t("doctores.formularioEditarTitulo")} />
			<div className="flex items-center justify-between page-header-breadcrumb flex-wrap gap-2">
				<div>
					<ol className="breadcrumb mb-0">
						<li className="breadcrumb-item">
							<Link scroll={false} href="#!">
								{t("doctores.breadcrumbApp")}
							</Link>
						</li>
						<li className="breadcrumb-item">
							<Link scroll={false} href="/components/doctores">
								{t("doctores.breadcrumbDoctores")}
							</Link>
						</li>
						<li className="breadcrumb-item active" aria-current="page">
							{t("doctores.formularioEditarBreadcrumb")}
						</li>
					</ol>
					<h1 className="page-title font-medium text-lg mb-0">
						{t("doctores.formularioEditarTitulo")}
					</h1>
				</div>
			</div>
			<div className="grid grid-cols-12 gap-x-6">
				<div className="xl:col-span-12 col-span-12">
					<div className="box my-4">
						<div className="box-header">
							<div className="box-title">{t("doctores.boxTitle")}</div>
						</div>
						<div className="box-body !p-[2rem]">
							{loading ? (
								<div className="flex items-center justify-center py-4">
									<span
										className="inline-block w-6 h-6 border-2 border-primary border-t-transparent rounded-full animate-spin"
										role="status"
										aria-label={t("doctores.cargando")}
									></span>
								</div>
							) : (
								<form onSubmit={handleSubmit}>
									<div className="grid grid-cols-12 gap-y-3">
										<div className="xl:col-span-12 col-span-12">
											<label htmlFor="firstName" className="form-label text-defaulttextcolor">
												{t("formulario.nombre")}
												<sup className="text-xs text-danger">*</sup>
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
												{t("formulario.apellido")}
												<sup className="text-xs text-danger">*</sup>
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
												{t("formulario.dni")}
												<sup className="text-xs text-danger">*</sup>
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
												{t("formulario.correo")}
												<sup className="text-xs text-danger">*</sup>
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
											<label
												htmlFor="licenseNumber"
												className="form-label text-defaulttextcolor"
											>
												{t("formulario.licencia")}
												<sup className="text-xs text-danger">*</sup>
											</label>
											<input
												type="text"
												name="licenseNumber"
												className="form-control"
												id="licenseNumber"
												placeholder={t("formulario.placeholderLicencia")}
												value={data.licenseNumber}
												onChange={changeHandler}
											/>
											{errors.licenseNumber && (
												<span className="text-danger text-xs mt-1">{errors.licenseNumber}</span>
											)}
										</div>
										<div className="xl:col-span-12 col-span-12">
											<label
												htmlFor="specialityId"
												className="form-label text-defaulttextcolor"
											>
												{t("doctores.especialidad")}
												<sup className="text-xs text-danger">*</sup>
											</label>
											<select
												name="specialityId"
												className="form-control"
												id="specialityId"
												value={data.specialityId}
												onChange={changeHandler}
												disabled={cargandoEspecialidades || !!errorEspecialidades}
											>
												<option value="">
													{t("doctores.placeholderEspecialidad")}
												</option>
												{especialidades.map((especialidad) => (
													<option key={especialidad.id} value={especialidad.id}>
														{especialidad.name}
													</option>
												))}
											</select>
											{cargandoEspecialidades && (
												<span className="text-textmuted dark:text-textmuted/50 text-xs mt-1">
													{t("doctores.cargandoEspecialidades")}
												</span>
											)}
											{errorEspecialidades && (
												<span className="text-danger text-xs mt-1">
													{errorEspecialidades}{" "}
													<Link
														scroll={false}
														href="/components/especialidades/nuevo"
														className="text-primary"
													>
														{t("doctores.crearEspecialidad")}
													</Link>
												</span>
											)}
											{!cargandoEspecialidades &&
												!errorEspecialidades &&
												especialidades.length === 0 && (
													<span className="text-danger text-xs mt-1">
														{t("doctores.sinEspecialidades")}
													</span>
												)}
											{errors.specialityId && (
												<span className="text-danger text-xs mt-1">{errors.specialityId}</span>
											)}
										</div>
										<div className="xl:col-span-12 col-span-12">
											<label
												htmlFor="active"
												className="form-label text-defaulttextcolor"
											>
												{t("formulario.estado")}
												<sup className="text-xs text-danger">*</sup>
											</label>
											<select
												name="active"
												className="form-control"
												id="active"
												value={data.active}
												onChange={changeHandler}
											>
												<option value="">
													{t("formulario.placeholderEstado")}
												</option>
												<option value="true">{t("formulario.activo")}</option>
												<option value="false">{t("formulario.inactivo")}</option>
											</select>
											<span className="text-textmuted dark:text-textmuted/50 text-xs mt-1">
												{t("formulario.estadoAyuda")}
											</span>
											{errors.active && (
												<span className="text-danger text-xs mt-1">{errors.active}</span>
											)}
										</div>
									</div>
									{errors.form && (
										<span className="text-danger text-xs mt-2">{errors.form}</span>
									)}
									<div className="flex gap-2 mt-4">
										<button
											type="submit"
											className="ti-btn ti-btn-primary btn-wave"
											disabled={saving}
										>
											<i className="ri-save-line align-middle me-1"></i>
											{t("formulario.guardar")}
										</button>
										<Link
											scroll={false}
											href="/components/doctores"
											className="ti-btn ti-btn-light btn-wave"
										>
											{t("formulario.cancelar")}
										</Link>
									</div>
								</form>
							)}
						</div>
					</div>
				</div>
			</div>
		</Fragment>
	);
};

EditarDoctor.layout = "Contentlayout";
export default EditarDoctor;
