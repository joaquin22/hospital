import Spktables from "@/shared/@spk-reusable-components/tables/spk-tables";
import { specialityService } from "@/shared/api/serivices/speciality.service";
import type { SpecialityOutput } from "@/shared/api/types/specialty";
import { useLanguage } from "@/shared/i18n/LanguageContext";
import Seo from "@/shared/layouts-components/seo/seo";
import Link from "next/link";
import { Fragment, useEffect, useState } from "react";

const Especialidades = () => {
	const { t } = useLanguage();
	const [especialidades, setEspecialidades] = useState<SpecialityOutput[]>([]);
	const [error, setError] = useState("");

	useEffect(() => {
		specialityService
			.list()
			.then((response) => {
				setEspecialidades(response.data);
				setError("");
			})
			.catch(() => setError(t("especialidades.errorListar")));
		// `t` no va en las dependencias a propósito: el provider recrea la
		// función en cada render y eso dispararía la petición otra vez.
	}, []);

	return (
		<Fragment>
			<Seo title={t("especialidades.titulo")} />
			<div className="flex items-center justify-between page-header-breadcrumb flex-wrap gap-2">
				<div>
					<ol className="breadcrumb mb-0">
						<li className="breadcrumb-item">
							<Link scroll={false} href="#!">
								{t("especialidades.breadcrumbApp")}
							</Link>
						</li>
						<li className="breadcrumb-item active" aria-current="page">
							{t("especialidades.titulo")}
						</li>
					</ol>
					<h1 className="page-title font-medium text-lg mb-0">
						{t("especialidades.titulo")}
					</h1>
				</div>
			</div>
			<div className="grid grid-cols-12 gap-x-6">
				<div className="xl:col-span-12 col-span-12">
					<div className="box">
						<div className="box-header justify-between">
							<div className="box-title">{t("especialidades.listado")}</div>
							<Link
								scroll={false}
								href="/components/especialidades/nuevo"
								className="ti-btn ti-btn-primary btn-wave"
							>
								<i className="ri-add-line align-middle me-1"></i>
								{t("especialidades.nuevaEspecialidad")}
							</Link>
						</div>
						<div className="box-body">
							{error && <span className="text-danger text-xs">{error}</span>}
							{!error && (
								<div className="table-responsive">
									<Spktables
										tableClass="ti-custom-table text-nowrap"
										header={[
											{ title: t("especialidades.columnaNombre") },
											{ title: t("especialidades.columnaDescripcion") },
											{ title: t("especialidades.columnaEstado") },
											{ title: t("especialidades.columnaAcciones") },
										]}
									>
										{especialidades.map((idx) => (
											<tr key={idx.id}>
												<td>
													<span className="font-medium">{idx.name}</span>
												</td>
												<td className="text-textmuted dark:text-textmuted/50">
													{idx.description}
												</td>
												<td>
													{idx.active ? (
														<span className="badge bg-success/10 text-success">
															{t("especialidades.activo")}
														</span>
													) : (
														<span className="badge bg-danger/10 text-danger">
															{t("especialidades.inactivo")}
														</span>
													)}
												</td>
												<td>
													<Link
														scroll={false}
														href={`/components/especialidades/editar/${idx.id}`}
														className="ti-btn ti-btn-sm ti-btn-primary"
													>
														<i className="ri-edit-line me-1"></i>
														{t("especialidades.editar")}
													</Link>
												</td>
											</tr>
										))}
									</Spktables>
									{especialidades.length === 0 && (
										<div className="text-center text-textmuted dark:text-textmuted/50 py-4">
											{t("especialidades.sinRegistros")}
										</div>
									)}
								</div>
							)}
						</div>
					</div>
				</div>
			</div>
		</Fragment>
	);
};

Especialidades.layout = "Contentlayout";
export default Especialidades;
