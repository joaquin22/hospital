import Spktables from "@/shared/@spk-reusable-components/tables/spk-tables";
import { useLanguage } from "@/shared/i18n/LanguageContext";
import Seo from "@/shared/layouts-components/seo/seo";
import Link from "next/link";
import { Fragment, useEffect, useState } from "react";
import type { DoctorUserOutput } from "@/shared/api/types/doctor";
import { doctorService } from "@/shared/api/serivices/doctor.service";

const Doctores = () => {
	const { t } = useLanguage();
	const [doctores, setDoctores] = useState<DoctorUserOutput[]>([]);

	useEffect(() => {
		doctorService
			.list()
			.then((response) => setDoctores(response.data))
			.catch(() => setDoctores([]));
	}, []);

	return (
		<Fragment>
			<Seo title={t("doctores.titulo")} />
			<div className="flex items-center justify-between page-header-breadcrumb flex-wrap gap-2">
				<div>
					<ol className="breadcrumb mb-0">
						<li className="breadcrumb-item">
							<Link scroll={false} href="#!">
								{t("doctores.breadcrumbApp")}
							</Link>
						</li>
						<li className="breadcrumb-item active" aria-current="page">{t("doctores.titulo")}</li>
					</ol>
					<h1 className="page-title font-medium text-lg mb-0">{t("doctores.titulo")}</h1>
				</div>
			</div>
			<div className="grid grid-cols-12 gap-x-6">
				<div className="xl:col-span-12 col-span-12">
					<div className="box">
						<div className="box-header justify-between">
							<div className="box-title">{t("doctores.listado")}</div>
							<Link scroll={false} href="/components/doctores/nuevo" className="ti-btn ti-btn-primary btn-wave">
								<i className="ri-add-line align-middle me-1"></i>{t("doctores.nuevoDoctor")}
							</Link>
						</div>
						<div className="box-body">
							<div className="table-responsive">
								<Spktables
									tableClass="ti-custom-table text-nowrap"
									header={[
										{ title: t("doctores.columnaNombre") },
										{ title: t("doctores.columnaApellido") },
										{ title: t("doctores.columnaDni") },
										{ title: t("doctores.columnaCorreo") },
										{ title: t("doctores.columnaLicencia") },
									]}
								>
									{doctores.map((idx) => (
										<tr key={idx.id}>
											<td><span className="font-medium">{idx.first_name}</span></td>
											<td>{idx.last_name}</td>
											<td>{idx.dni}</td>
											<td className="text-textmuted dark:text-textmuted/50">{idx.email}</td>
											<td><span className="badge bg-primary/10 text-primary">{idx.license_number}</span></td>
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

Doctores.layout = "Contentlayout";
export default Doctores;
