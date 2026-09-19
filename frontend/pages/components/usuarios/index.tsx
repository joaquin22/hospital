import Spktables from "@/shared/@spk-reusable-components/tables/spk-tables";
import { useLanguage } from "@/shared/i18n/LanguageContext";
import Seo from "@/shared/layouts-components/seo/seo";
import Link from "next/link";
import { Fragment, useEffect, useState } from "react";
import type { UserOutput } from "@/shared/api/types/user";
import { userService } from "@/shared/api/serivices/user.service";

const Usuarios = () => {
	const { t } = useLanguage();
	const [usuarios, setUsuarios] = useState<UserOutput[]>([]);

	useEffect(() => {
		userService
			.list()
			.then((response) => setUsuarios(response.data))
			.catch(() => setUsuarios([]));
	}, []);

	return (
		<Fragment>
			<Seo title={t("usuarios.titulo")} />
			<div className="flex items-center justify-between page-header-breadcrumb flex-wrap gap-2">
				<div>
					<ol className="breadcrumb mb-0">
						<li className="breadcrumb-item">
							<Link scroll={false} href="#!">
								{t("usuarios.breadcrumbApp")}
							</Link>
						</li>
						<li className="breadcrumb-item active" aria-current="page">{t("usuarios.titulo")}</li>
					</ol>
					<h1 className="page-title font-medium text-lg mb-0">{t("usuarios.titulo")}</h1>
				</div>
			</div>
			<div className="grid grid-cols-12 gap-x-6">
				<div className="xl:col-span-12 col-span-12">
					<div className="box">
						<div className="box-header justify-between">
							<div className="box-title">{t("usuarios.listado")}</div>
							<Link scroll={false} href="/components/usuarios/nuevo" className="ti-btn ti-btn-primary btn-wave">
								<i className="ri-add-line align-middle me-1"></i>{t("usuarios.nuevoUsuario")}
							</Link>
						</div>
						<div className="box-body">
							<div className="table-responsive">
								<Spktables
									tableClass="ti-custom-table text-nowrap"
									header={[
										{ title: t("usuarios.columnaNombre") },
										{ title: t("usuarios.columnaApellido") },
										{ title: t("usuarios.columnaDni") },
										{ title: t("usuarios.columnaCorreo") },
										{ title: t("usuarios.columnaRol") },
										{ title: t("usuarios.columnaEstado") },
										{ title: t("usuarios.columnaAcciones") },
									]}
								>
									{usuarios.map((idx) => (
										<tr key={idx.id}>
											<td><span className="font-medium">{idx.first_name}</span></td>
											<td>{idx.last_name}</td>
											<td>{idx.dni}</td>
											<td className="text-textmuted dark:text-textmuted/50">{idx.email}</td>
											<td><span className="badge bg-primary/10 text-primary">{idx.role}</span></td>
											<td>
												{idx.active ? (
													<span className="badge bg-success/10 text-success">{t("usuarios.activo")}</span>
												) : (
													<span className="badge bg-danger/10 text-danger">{t("usuarios.inactivo")}</span>
												)}
											</td>
											<td>
												<Link scroll={false} href={`/components/usuarios/editar/${idx.id}`} className="ti-btn ti-btn-sm ti-btn-primary">
													<i className="ri-edit-line me-1"></i>{t("usuarios.editar")}
												</Link>
											</td>
										</tr>
									))}
								</Spktables>
							</div>
						</div>
					</div>
				</div>
			</div>
		</Fragment>
	);
};

Usuarios.layout = "Contentlayout";
export default Usuarios;