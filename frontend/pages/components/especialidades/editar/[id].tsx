import { resolveApiError } from "@/shared/api/errors";
import { specialityService } from "@/shared/api/serivices/speciality.service";
import { useLanguage } from "@/shared/i18n/LanguageContext";
import Seo from "@/shared/layouts-components/seo/seo";
import Link from "next/link";
import { useRouter } from "next/router";
import { Fragment, useEffect, useState } from "react";
import { notifier } from "@/utils/notifier";

const EditarEspecialidad = () => {
	const router = useRouter();
	const { t } = useLanguage();
	const id = router.query.id as string | undefined;

	const [data, setData] = useState({
		name: "",
		description: "",
	});

	const [errors, setErrors] = useState<Record<string, string>>({});
	const [loading, setLoading] = useState(true);
	const [saving, setSaving] = useState(false);

	useEffect(() => {
		if (!id) return;
		setLoading(true);
		specialityService
			.get(Number(id))
			.then((response) => {
				setData({
					name: response.data.name,
					description: response.data.description,
				});
				setErrors({});
			})
			.catch((error) => {
				const apiError = resolveApiError(error, t, "especialidades.errorCargar");
				setErrors({ form: apiError.message });
			})
			.finally(() => setLoading(false));
		// `t` no va en las dependencias a propósito: el provider recrea la
		// función en cada render y eso dispararía la petición otra vez.
	}, [id]);

	const changeHandler = (
		e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
	) => {
		setData({ ...data, [e.target.name]: e.target.value });
		setErrors({ ...errors, [e.target.name]: "", form: "" });
	};

	const validate = () => {
		const newErrors: Record<string, string> = {};
		if (!data.name.trim()) newErrors.name = t("especialidades.errorNombre");
		if (!data.description.trim())
			newErrors.description = t("especialidades.errorDescripcion");
		setErrors(newErrors);
		return Object.keys(newErrors).length === 0;
	};

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		if (!id || !validate()) return;
		setSaving(true);
		try {
			// PATCH: el id va en la ruta, el handler ignora el `id` del body.
			await specialityService.update(Number(id), {
				name: data.name.trim(),
				description: data.description.trim(),
			});
			notifier.success(t("especialidades.actualizada"));
			router.push("/components/especialidades");
		} catch (error) {
			const apiError = resolveApiError(
				error,
				t,
				"especialidades.errorActualizar",
			);
			setErrors({ form: apiError.message });
		} finally {
			setSaving(false);
		}
	};

	return (
		<Fragment>
			<Seo title={t("especialidades.formularioEditarTitulo")} />
			<div className="flex items-center justify-between page-header-breadcrumb flex-wrap gap-2">
				<div>
					<ol className="breadcrumb mb-0">
						<li className="breadcrumb-item">
							<Link scroll={false} href="#!">
								{t("especialidades.breadcrumbApp")}
							</Link>
						</li>
						<li className="breadcrumb-item">
							<Link scroll={false} href="/components/especialidades">
								{t("especialidades.titulo")}
							</Link>
						</li>
						<li className="breadcrumb-item active" aria-current="page">
							{t("especialidades.breadcrumbEditar")}
						</li>
					</ol>
					<h1 className="page-title font-medium text-lg mb-0">
						{t("especialidades.formularioEditarTitulo")}
					</h1>
				</div>
			</div>
			<div className="grid grid-cols-12 gap-x-6">
				<div className="xl:col-span-12 col-span-12">
					<div className="box my-4">
						<div className="box-header">
							<div className="box-title">{t("especialidades.boxTitle")}</div>
						</div>
						<div className="box-body !p-[2rem]">
							{loading ? (
								<div className="flex items-center justify-center py-4">
									<span
										className="inline-block w-6 h-6 border-2 border-primary border-t-transparent rounded-full animate-spin"
										role="status"
										aria-label={t("especialidades.cargando")}
									></span>
								</div>
							) : (
								<form onSubmit={handleSubmit}>
									<div className="grid grid-cols-12 gap-y-3">
										<div className="xl:col-span-12 col-span-12">
											<label htmlFor="name" className="form-label text-defaulttextcolor">
												{t("especialidades.nombre")}
												<sup className="text-xs text-danger">*</sup>
											</label>
											<input
												type="text"
												name="name"
												className="form-control"
												id="name"
												placeholder={t("especialidades.placeholderNombre")}
												value={data.name}
												onChange={changeHandler}
											/>
											{errors.name && (
												<span className="text-danger text-xs mt-1">{errors.name}</span>
											)}
										</div>
										<div className="xl:col-span-12 col-span-12">
											<label
												htmlFor="description"
												className="form-label text-defaulttextcolor"
											>
												{t("especialidades.descripcion")}
												<sup className="text-xs text-danger">*</sup>
											</label>
											<textarea
												name="description"
												className="form-control"
												id="description"
												rows={4}
												placeholder={t("especialidades.placeholderDescripcion")}
												value={data.description}
												onChange={changeHandler}
											></textarea>
											{errors.description && (
												<span className="text-danger text-xs mt-1">
													{errors.description}
												</span>
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
											{t("especialidades.guardar")}
										</button>
										<Link
											scroll={false}
											href="/components/especialidades"
											className="ti-btn ti-btn-light btn-wave"
										>
											{t("especialidades.cancelar")}
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

EditarEspecialidad.layout = "Contentlayout";
export default EditarEspecialidad;
